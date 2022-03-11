import cw_listMetrics from "./cw_listMetrics.cjs";
import dynamodb_listTables from "./dynamodb_listTables.cjs";
import iam_listUsers from "./iam_listUsers.cjs";
import s3_listBuckets from "./s3_listBuckets.cjs";
import ses_listIdentities from "./ses_listIdentities.cjs";
import sns_listTopicsAndSubscriptions from "./sns_listTopicsAndSubscriptions.cjs";
import sqs_listQueues from "./sqs_listQueues.cjs";

export default {
    cw_listMetrics,
    dynamodb_listTables,
    iam_listUsers,
    s3_listBuckets,
    ses_listIdentities,
    sns_listTopicsAndSubscriptions,
    sqs_listQueues,
}