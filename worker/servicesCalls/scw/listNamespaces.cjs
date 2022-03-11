const axios = require('axios');

/**
 * Returns a list of sw namespaces from a given region
 *
 * @param {string} region 
 * @param {string} auth_token
 * @return {object} 
 */
function listNamespaces(region, auth_token){
    return axios({
        method: "get",
        url: `/registry/v1/regions/${region}/namespaces`,
        baseURL: 'https://api.scaleway.com',
        headers: { 'User-Agent': 'node-scaleway-api-client/2.0.0', 'X-Auth-Token': auth_token },
        params: {},
        data: {},
        responseType: 'json',
    });      
}

module.exports.listNamespaces = listNamespaces;