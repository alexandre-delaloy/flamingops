const axios = require('axios');

/**
 * Returns a list of sw iot hubs from a given region
 *
 * @param {string} region 
 * @param {string} auth_token
 * @return {object} 
 */
function listIotHubs(region, auth_token){
    return axios({
        method: "get",
        url: `/iot/v1/regions/${region}/hubs`,
        baseURL: 'https://api.scaleway.com',
        headers: { 'User-Agent': 'node-scaleway-api-client/2.0.0', 'X-Auth-Token': auth_token },
        params: {},
        data: {},
        responseType: 'json',
    });      
}

module.exports.listIotHubs = listIotHubs;