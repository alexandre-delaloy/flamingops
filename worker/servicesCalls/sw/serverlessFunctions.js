const axios = require('axios');

/**
 * Returns a list of sw serverless functions datas from a given region
 *
 * @param {string} region 
 * @param {string} auth_token
 * @return {object} 
 */
export default function(region, auth_token){
    base_url = 'https://api.scaleway.com'

    const functions = axios({
        method: "get",
        url: `/functions/v1beta1/regions/${region}/functions`,
        baseURL: 'https://api.scaleway.com',
        headers: { 'User-Agent': 'node-scaleway-api-client/2.0.0', 'X-Auth-Token': auth_token },
        params: {},
        data: {},
        responseType: 'json',
    });  

    const crons = axios({
        method: "get",
        url: `/functions/v1beta1/regions/${region}/crons`,
        baseURL: 'https://api.scaleway.com',
        headers: { 'User-Agent': 'node-scaleway-api-client/2.0.0', 'X-Auth-Token': auth_token },
        params: {},
        data: {},
        responseType: 'json',
    });       
    
    const namespaces = axios({
        method: "get",
        url: `/functions/v1beta1/regions/${region}/namespaces`,
        baseURL: 'https://api.scaleway.com',
        headers: { 'User-Agent': 'node-scaleway-api-client/2.0.0', 'X-Auth-Token': auth_token },
        params: {},
        data: {},
        responseType: 'json',
    });   

    const domains = axios({
        method: "get",
        url: `/functions/v1beta1/regions/${region}/domains`,
        baseURL: 'https://api.scaleway.com',
        headers: { 'User-Agent': 'node-scaleway-api-client/2.0.0', 'X-Auth-Token': auth_token },
        params: {},
        data: {},
        responseType: 'json',
    });   

    return Promise.all([functions, crons, namespaces, domains]).then(responses => {
        return {
            functions: responses[0].data,
            crons: responses[1].data,
            namespaces: responses[2].data,
            domains: responses[3].data,
        }
    });
}