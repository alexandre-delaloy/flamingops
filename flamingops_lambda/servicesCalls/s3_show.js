// snippet-start:[s3.JavaScript.buckets.listObjects]
// Load the AWS SDK for Node.js
var AWS = require('aws-sdk');

/**
 * Returns an aws bucket from a given region
 *
 * @param {string} region 
 * @param {string} bucketName 
 * @return {object} 
 */
export default function(region, bucketName) {
    // Set the region 
    AWS.config.update({region: region});
    // Create S3 service object
    s3 = new AWS.S3({apiVersion: '2006-03-01'});

    // Create the parameters for calling listObjects
    var bucketParams = {
        Bucket : bucketName,
    };

    let result
    // Call S3 to obtain a list of the objects in the bucket
    s3.listObjects(bucketParams, function(err, data) {
        if (err) {
            throw new Error("error :", err)
        } else {
            result = data;
        }
    });

    return result;
}