// https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/SNS.html#listTopics-property
var AWS = require('aws-sdk');

/**
 * Returns a list of aws sns topics from a given region
 *
 * @param {string} region 
 * @return {object} 
 */
export default function(region) {
    AWS.config.update({region});
    var listTopics = new AWS.SNS({apiVersion: '2010-03-31'})

    var params = {};

    let result
    listTopics.listTopics(params, function(err, data) {
        if (err) {
            throw new Error("error :", err)
        } else {
            // data : {Topics: ["topicARN1", "topicARN1", ...], NextToken}
            result = data.Topics; 
        }
    });
    
    return result;
}