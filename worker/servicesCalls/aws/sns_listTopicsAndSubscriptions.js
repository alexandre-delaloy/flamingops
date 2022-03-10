// https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/SNS.html#listTopics-property
var AWS = require('aws-sdk');

/**
 * Returns a list of aws sns topics from a given region
 * @param {string} accessKeyId 
 * @param {string} secretAccessKey 
 * @param {string} region 
 * @return {object} 
 */
 export default function(accessKeyId, secretAccessKey, region) {
    AWS.config.update({accessKeyId, secretAccessKey, region});
    var sns = new AWS.SNS({apiVersion: '2010-03-31'})

    var params = {};

    let result
    sns.listTopics(params, function(err, data) {
        if (err) {
            throw new Error("error :", err)
        } else {
            // data : {Topics: ["topicARN1", "topicARN1", ...], NextToken}
            result.Topics = data.Topics; 
        }
    });
    sns.listSubscriptions(params, function(err, data) {
        if (err) {
            throw new Error("error :", err)
        } else {
            // data : {Subscriptions: [{SubscriptionArn, Owner, Protocol}, ...], NextToken}
            result.Subscriptions = data.Subscriptions; 
        }
    });
    
    return result;
}