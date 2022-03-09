const axios = require('axios');

/**
 * Returns a list of sw dns zones
 * @param {string} region 
 * @param {string} auth_token
 * @return {object} 
 */
export default function(region, auth_token){
    base_url = 'https://api.scaleway.com'

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