var AWS = require('aws-sdk');

/**
 * Returns aws cloudwatch metrics from a given region
 * @param {string} accessKeyId 
 * @param {string} secretAccessKey 
 * @param {string} region 
 * @return {object} 
 */
function listMetrics(accessKeyId, secretAccessKey, region) {
    // Set the region 
    AWS.config.update({accessKeyId, secretAccessKey, region});
    // Create CloudWatch service object
    var cw = new AWS.CloudWatch({apiVersion: '2010-08-01'});

    // todo : change parameters for variables
    var params = {
        Dimensions: [
          {
            Name: 'LogGroupName', /* required */
          },
        ],
        MetricName: 'IncomingLogEvents',
        Namespace: 'AWS/Logs'
    };

    let result;
    cw.listMetrics(params, function(err, data) {
        if (err) {
            throw new Error("error :", err);
        } else {
            // data : 
            result = JSON.stringify(data.Metrics);
        }
    })
    
    return result;
}

module.exports.listMetrics = listMetrics;