import serviceToFunction from './servicesCalls/serviceToFunction.js';
import SecretsManager from './secretsManager.js';
import { setConnection, updateDb } from './dbConnection.js'

// 1. - Récupérer le payload du message
// 2. - Récupérer les credentials de l'utilisateur sur Secrets Manager
// 3. - Executer les calls demandées
// 4. - Stocker les résultats dans rds

exports.handler = async function(event) {
  // 1. - Récupérer le payload du message
  const clientName = event.clientName;
  const clientId = event.clientId;
  const requestedServices = event.requestedServices;
  const requestedAwsRegion = event.requestedAwsRegion;
  const requestedSwRegion = event.requestedSwRegion;

  // 2. - Récupérer les credentials de l'utilisateur sur Secrets Manager
  const AwsSecretName = `${clientName}_aws_iam`;
  const ScalewaySecretName = `${clientName}_sw`;
  const awsSecret = await SecretsManager.getSecret(AwsSecretName, requestedAwsRegion);
  const swSecret = await SecretsManager.getSecret(ScalewaySecretName, requestedSwRegion);

  //3. - Executer les calls demandés
  let servicesData = {
    aws: {},
    sw: {},
  };
  foreach(requestedServices, function(service) {
    switch (serviceToFunction[service].cloud) {
      case 'aws':
        servicesData.aws[service] = await serviceToFunction[service].function(
          awsSecret.accessKeyId,
          awsSecret.secretAccessKey,
          requestedAwsRegion,
        );
        break;
      case 'sw':
        servicesData.sw[service] = await serviceToFunction[service].function(
          swSecret.authToken,
          requestedSwRegion,
        );
        break;
    }
  });

  // 4. - Stocker les résultats dans rds
  const dbHostname = process.env.DB_HOSTNAME;
  const dbPort = process.env.DB_PORT;
  const dbUsername = process.env.DB_USERNAME;
  const dbName = process.env.DB_NAME;
  const dbPassword = process.env.DB_PASSWORD;
  const dbRegion = process.env.DB_REGION;

  const client = setConnection(dbHostname, dbPort, dbUsername, dbName, dbPassword, dbRegion);
  await client.connect();
  await updateDb(client, 'awsServicesDatas', clientId, servicesData.aws);
  await updateDb(client, 'swServicesDatas', clientId, servicesData.sw);
  client.end();

  return {
    statusCode: 200,
    body: JSON.stringify({
      message: 'Successfully executed all services calls',
      data: servicesData,
    }),
  };
};
