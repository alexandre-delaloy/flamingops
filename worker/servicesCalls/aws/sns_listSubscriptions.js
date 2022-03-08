// https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/SNS.html#listSubscriptions-property
var AWS = require('aws-sdk');

/**
 * Returns a list of aws sns subscriptions from a given region
 *
 * @param {string} region 
 * @return {object} 
 */
export default function(region) {
    AWS.config.update({region});
    var subslist = new AWS.SNS({apiVersion: '2010-03-31'})

    var params = {};

    let result
    subslist.listSubscriptions(params, function(err, data) {
        if (err) {
            throw new Error("error :", err)
        } else {
            // data : {Subscriptions: [{SubscriptionArn, Owner, Protocol}, ...], NextToken}
            result = data.Subscriptions; 
        }
    });
    
    return result;
}