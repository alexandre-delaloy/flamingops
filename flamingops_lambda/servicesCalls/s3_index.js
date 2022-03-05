// snippet-start:[s3.JavaScript.buckets.listBuckets]
var AWS = require('aws-sdk');

/**
 * Returns an aws bucket from a given region
 *
 * @param {string} region 
 * @return {object} 
 */
export default function(region) {
    // Set the region 
    AWS.config.update({region});
    // Create S3 service object
    s3 = new AWS.S3({apiVersion: '2006-03-01'});

    let result
    // Call S3 to list the buckets
    s3.listBuckets(function(err, data) {
        if (err) {
            throw new Error("error :", err)
        } else {
            result = data.Buckets; // CreationDate, Name
        }
    });
    
    return result;
}
// snippet-end:[s3.JavaScript.buckets.listBuckets]