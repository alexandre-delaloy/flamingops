// https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/IAM.html#listUsers-property
var AWS = require('aws-sdk');

/**
 * Returns a list of aws IAM users from a given region
 * @param {string} accessKeyId 
 * @param {string} secretAccessKey 
 * @param {string} region 
 * @return {object} 
 */
function listUsers(accessKeyId, secretAccessKey, region) {
    AWS.config.update({accessKeyId, secretAccessKey, region});
    
    var iam = new AWS.IAM({apiVersion: '2010-05-08'});

    var params = {
        MaxItems: 10
    };

    let result
    iam.listUsers(params, function(err, data) {
        if (err) {
            throw new Error("error :", err)
        } else {
            // data : {Users: [{UserName, UserId, PasswordLastUsed, CreateDate, ...}, ...]}
            result = data.Users || [];
        }
    });
    
    return result;
}

module.exports.listUsers = listUsers;