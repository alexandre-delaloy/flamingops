// https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/S3.html#listObjects-property
var AWS = require('aws-sdk');

/**
 * Returns an aws s3 bucket from a given region
 *
 * @param {string} region 
 * @param {string} bucketName 
 * @return {object} 
 */
export default function(region, bucketName) {
    AWS.config.update({region: region});
    var s3 = new AWS.S3({apiVersion: '2006-03-01'});

    var params = {
        Bucket : bucketName,
    };

    let result
    s3.listObjects(params, function(err, data) {
        if (err) {
            throw new Error("error :", err)
        } else {
            // data : {Contents : [{ETag, Key, LastModified, Owner, ...}, ...], NextMarker}
            result = data.Contents;
        }
    });

    return result;
}