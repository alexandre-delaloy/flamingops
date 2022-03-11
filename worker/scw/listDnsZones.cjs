const axios = require('axios');

/**
 * Returns a list of sw dns zones
 * @param {string} region 
 * @param {string} auth_token
 * @return {object} 
 */
function listDnsZones(region, auth_token){
    return axios({
        method: "get",
        url: `/domain/v2beta1/dns-zones`,
        baseURL: 'https://api.scaleway.com',
        headers: { 'User-Agent': 'node-scaleway-api-client/2.0.0', 'X-Auth-Token': auth_token },
        params: {},
        data: {},
        responseType: 'json',
    });      
}

module.exports.listDnsZones = listDnsZones;