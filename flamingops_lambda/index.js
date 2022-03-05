// 1. - Récupérer le payload du message
// 2. - Récupérer les credentials de l'utilisateur sur Secrets Manager
// 3. - Executer les calls demandées
// 4. - Stocker les résultats dans rds

import servicesCalls from "./servicesCalls";
const SecretsManager = require('./SecretsManager.js');

exports.handler = async function(event) {
  // 1. - Récupérer le payload du message
  const clientName = event.clientName;
  const requestedServices = event.requestedServices;
  const requestedRegion = event.requestedRegion;

  // 2. - Récupérer les credentials de l'utilisateur sur Secrets Manager
  const AwsSecretName = `${clientName}_aws_iam`;
  const ScalewaySecretName = `${clientName}_sw`;
  var awsSecret = await SecretsManager.getSecret(AwsSecretName, requestedRegion);
  var swSecret = await SecretsManager.getSecret(ScalewaySecretName, requestedRegion);

  //3. - Executer les calls demandés
  let servicesData = {
    aws: {},
    sw: {},
  } ;
  foreach(requestedServices, function(service) {
    switch(service) {
      case '':

  }

  // 4. - Stocker les résultats dans rds




  const response = {
      statusCode: 200,
      body: JSON.stringify('Hello from Lambda!'),
  };
  return response;
}




// exports.handler = async (event) => {
//   // TODO implement
//   const response = {
//       statusCode: 200,
//       body: JSON.stringify('Hello from Lambda!'),
//   };
//   return response;
// };
