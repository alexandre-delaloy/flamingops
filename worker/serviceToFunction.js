import awsServicesCalls from "./aws/index.js";
import swServicesCalls from "./scw/index.js";

export default {
  cloudwatch: {
    function: awsServicesCalls.cw_listMetrics,
    cloud: 'aws',
  },
  dynamodb: {
    function: awsServicesCalls.dynamodb_listTables,
    cloud: 'aws',
  },
  iam: {
    function: awsServicesCalls.iam_listUsers,
    cloud: 'aws',
  },
  s3: {
    function: awsServicesCalls.s3_listBuckets,
    cloud: 'aws',
  },
  ses: {
    function: awsServicesCalls.ses_listIdentities,
    cloud: 'aws',
  },
  sns: {
    function: awsServicesCalls.sns_listTopicsAndSubscriptions,
    cloud: 'aws',
  },
  sqs: {
    function: awsServicesCalls.sqs_listQueues,
    cloud: 'aws',
  },
  applesilicon: {
    function: swServicesCalls.listAppleSiliconServers,
    cloud: 'sw',
  },
  baremetal: {
    function: swServicesCalls.listBareMetalServers,
    cloud: 'sw',
  },
  dns: {
    function: swServicesCalls.listDnsZones,
    cloud: 'sw',
  },
  flexibleip: {
    function: swServicesCalls.listFlexibleIps,
    cloud: 'sw',
  },
  instances: {
    function: swServicesCalls.listInstances,
    cloud: 'sw',
  },
  iothubs: {
    function: swServicesCalls.listIotHubs,
    cloud: 'sw',
  },
  k8s: {
    function: swServicesCalls.listK8sClusters,
    cloud: 'sw',
  },
  loadbalancers: {
    function: swServicesCalls.listLoadBalancers,
    cloud: 'sw',
  },
  namespaces: {
    function: swServicesCalls.listNamespaces,
    cloud: 'sw',
  },
  rdb: {
    function: swServicesCalls.listRdbInstances,
    cloud: 'sw',
  },
  vpc: {
    function: swServicesCalls.listVpcNetworks,
    cloud: 'sw',
  },
  serverlesscontainers: {
    function: swServicesCalls.serverlessContainers,
    cloud: 'sw',
  },
  serverlessfunctions: {
    function: swServicesCalls.serverlessFunctions,
    cloud: 'sw',
  },
}