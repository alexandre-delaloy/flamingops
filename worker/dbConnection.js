const { Client } = require('pg');

function setConnection(){
  // TODO : use a config file
  const signerOptions = {
    credentials: {
      accessKeyId: 'YOUR-ACCESS-KEY',
      secretAccessKey: 'YOUR-SECRET-ACCESS-KEY',
    },
    region: 'eu-west-3',
    hostname: 'example.aslfdewrlk.eu-west-3.rds.amazonaws.com',
    port: 5432,
    username: 'api-user',
  };

  const signer = new RDS.Signer()
  const getPassword = () => signer.getAuthToken(signerOptions)
  const client = new Client({
    host: signerOptions.hostname,
    port: signerOptions.port,
    user: signerOptions.username,
    database: 'my-db',
    password: getPassword,
  });
  return client;
}

async function updateDb(client, table, clientId, data){
  const query = `UPDATE ${table} SET data = '${JSON.stringify(data)}' WHERE user_id = '${clientId}'`;
  return client.query(query);
}

export default {
  setConnection,
  updateDb,
};
