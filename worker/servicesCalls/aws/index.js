import cw_listMetrics from "./cw_listMetrics.js";
import dynamodb_listTables from "./dynamodb_listTables.js";
import iam_listUsers from "./iam_listUsers.js";
import s3_listBuckets from "./s3_listBuckets.js";
import ses_listIdentities from "./ses_listIdentities.js";
import sns_listTopicsAndSubscriptions from "./sns_listTopicsAndSubscriptions.js";
import sqs_listQueues from "./sqs_listQueues.js";

export default {
    cw_listMetrics,
    dynamodb_listTables,
    iam_listUsers,
    s3_listBuckets,
    ses_listIdentities,
    sns_listTopicsAndSubscriptions,
    sqs_listQueues,
}