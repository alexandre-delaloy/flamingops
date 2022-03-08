// https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/SES.html#listIdentities-property
var AWS = require('aws-sdk');

/**
 * Returns a list of aws ses identities as mail adresses from a given region
 *
 * @param {string} region 
 * @return {object} 
 */
export default function(region) {
    AWS.config.update({region});
    var listIDsPromise = new AWS.SES({apiVersion: '2010-12-01'}).promise();

    var params = {
        IdentityType: "EmailAddress",
        MaxItems: 10
    };
       
    let result
    listIDsPromise.listIdentities(params, function(err, data) {
        if (err) {
            throw new Error("error :", err)
        } else {
            // data : {Identities: ["user@example.com"], NextToken}
            result = data.Identities; 
        }
    });
    
    return result;
}