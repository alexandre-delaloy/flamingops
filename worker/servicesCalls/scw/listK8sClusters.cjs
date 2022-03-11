const axios = require('axios');

/**
 * Returns a list of sw k8s clusters from a given region
 *
 * @param {string} region 
 * @param {string} auth_token
 * @return {object} 
 */
function listK8sClusters(region, auth_token){
    return axios({
        method: "get",
        url: `/k8s/v1/regions/${region}/clusters`,
        baseURL: 'https://api.scaleway.com',
        headers: { 'User-Agent': 'node-scaleway-api-client/2.0.0', 'X-Auth-Token': auth_token },
        params: {},
        data: {},
        responseType: 'json',
    });      
}

module.exports.listK8sClusters = listK8sClusters;