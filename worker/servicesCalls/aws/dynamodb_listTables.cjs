// https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/DynamoDB.html#listTables-property
var AWS = require('aws-sdk');

/**
 * Returns aws dynamo DB tables from a given region
 * @param {string} accessKeyId 
 * @param {string} secretAccessKey 
 * @param {string} region 
 * @return {object} 
 */
function listTables(accessKeyId, secretAccessKey, region) {
    AWS.config.update({accessKeyId, secretAccessKey, region});
    var ddb = new AWS.DynamoDB({apiVersion: '2012-08-10'});

    let result
    ddb.listTables({Limit: 10}, function(err, data) {
        if (err) {
            throw new Error("error :", err)
        } else {
            // data : {TableNames: ["Name1", "Name2", ...]}
            result = data.TableNames;
        }
    });
    
    return result;
}

module.exports.listTables = listTables;