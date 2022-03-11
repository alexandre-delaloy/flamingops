// https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/S3.html#listBuckets-property
var AWS = require('aws-sdk');

/**
 * Returns a list of aws s3 buckets from a given region
 * @param {string} accessKeyId 
 * @param {string} secretAccessKey 
 * @param {string} region 
 * @return {object} 
 */
function listBuckets(accessKeyId, secretAccessKey, region) {
    AWS.config.update({accessKeyId, secretAccessKey, region});
    var s3 = new AWS.S3({apiVersion: '2006-03-01'});

    var params = {};

    let result
    s3.listBuckets(params, function(err, data) {
        if (err) {
            throw new Error("error :", err)
        } else {
            // data : {Buckets: [{CreationDate, Name}, ...], Owner}
            result = data.Buckets; 
        }
    });
    
    return result;
}

module.exports.listBuckets = listBuckets;