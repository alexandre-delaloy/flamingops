const axios = require('axios');

/**
 * Returns a list of sw rdb instances from a given region
 *
 * @param {string} region 
 * @param {string} auth_token
 * @return {object} 
 */
function listRdbInstances(region, auth_token){
    return axios({
        method: "get",
        url: `/rdb/v1/regions/${region}/instances`,
        baseURL: 'https://api.scaleway.com',
        headers: { 'User-Agent': 'node-scaleway-api-client/2.0.0', 'X-Auth-Token': auth_token },
        params: {},
        data: {},
        responseType: 'json',
    });      
}

module.exports.listRdbInstances = listRdbInstances;