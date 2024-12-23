

package api

// ServiceName returns the SDK's naming of the service. Has
// backwards compatibility built in for services that were
// incorrectly named with the service's endpoint prefix.
func ServiceName(a *API) string {
	if oldName, ok := oldServiceNames[a.PackageName()]; ok {
		return oldName
	}

	return ServiceID(a)
}

var oldServiceNames = map[string]string{
	"migrationhub":                    "mgh",
	"acmpca":                          "acm-pca",
	"acm":                             "acm",
	"alexaforbusiness":                "a4b",
	"apigateway":                      "apigateway",
	"applicationautoscaling":          "autoscaling",
	"appstream":                       "appstream2",
	"appsync":                         "appsync",
	"athena":                          "athena",
	"autoscalingplans":                "autoscaling",
	"autoscaling":                     "autoscaling",
	"batch":                           "batch",
	"budgets":                         "budgets",
	"costexplorer":                    "ce",
	"cloud9":                          "cloud9",
	"clouddirectory":                  "clouddirectory",
	"cloudformation":                  "cloudformation",
	"cloudfront":                      "cloudfront",
	"cloudhsm":                        "cloudhsm",
	"cloudhsmv2":                      "cloudhsmv2",
	"cloudsearch":                     "cloudsearch",
	"cloudsearchdomain":               "cloudsearchdomain",
	"cloudtrail":                      "cloudtrail",
	"codebuild":                       "codebuild",
	"codecommit":                      "codecommit",
	"codedeploy":                      "codedeploy",
	"codepipeline":                    "codepipeline",
	"codestar":                        "codestar",
	"cognitoidentity":                 "cognito-identity",
	"cognitoidentityprovider":         "cognito-idp",
	"cognitosync":                     "cognito-sync",
	"comprehend":                      "comprehend",
	"configservice":                   "config",
	"connect":                         "connect",
	"costandusagereportservice":       "cur",
	"datapipeline":                    "datapipeline",
	"dax":                             "dax",
	"devicefarm":                      "devicefarm",
	"directconnect":                   "directconnect",
	"applicationdiscoveryservice":     "discovery",
	"databasemigrationservice":        "dms",
	"directoryservice":                "ds",
	"dynamodb":                        "dynamodb",
	"ec2":                             "ec2",
	"ecr":                             "ecr",
	"ecs":                             "ecs",
	"eks":                             "eks",
	"elasticache":                     "elasticache",
	"elasticbeanstalk":                "elasticbeanstalk",
	"efs":                             "elasticfilesystem",
	"elb":                             "elasticloadbalancing",
	"elbv2":                           "elasticloadbalancing",
	"emr":                             "elasticmapreduce",
	"elastictranscoder":               "elastictranscoder",
	"ses":                             "email",
	"marketplaceentitlementservice":   "entitlement.marketplace",
	"elasticsearchservice":            "es",
	"cloudwatchevents":                "events",
	"firehose":                        "firehose",
	"fms":                             "fms",
	"gamelift":                        "gamelift",
	"glacier":                         "glacier",
	"glue":                            "glue",
	"greengrass":                      "greengrass",
	"guardduty":                       "guardduty",
	"health":                          "health",
	"iam":                             "iam",
	"inspector":                       "inspector",
	"iotdataplane":                    "data.iot",
	"iotjobsdataplane":                "data.jobs.iot",
	"iot":                             "iot",
	"iot1clickdevicesservice":         "devices.iot1click",
	"iot1clickprojects":               "projects.iot1click",
	"iotanalytics":                    "iotanalytics",
	"kinesisvideoarchivedmedia":       "kinesisvideo",
	"kinesisvideomedia":               "kinesisvideo",
	"kinesis":                         "kinesis",
	"kinesisanalytics":                "kinesisanalytics",
	"kinesisvideo":                    "kinesisvideo",
	"kms":                             "kms",
	"lambda":                          "lambda",
	"lexmodelbuildingservice":         "models.lex",
	"lightsail":                       "lightsail",
	"cloudwatchlogs":                  "logs",
	"machinelearning":                 "machinelearning",
	"marketplacecommerceanalytics":    "marketplacecommerceanalytics",
	"mediaconvert":                    "mediaconvert",
	"medialive":                       "medialive",
	"mediapackage":                    "mediapackage",
	"mediastoredata":                  "data.mediastore",
	"mediastore":                      "mediastore",
	"mediatailor":                     "api.mediatailor",
	"marketplacemetering":             "metering.marketplace",
	"mobile":                          "mobile",
	"mobileanalytics":                 "mobileanalytics",
	"cloudwatch":                      "monitoring",
	"mq":                              "mq",
	"mturk":                           "mturk-requester",
	"neptune":                         "rds",
	"opsworks":                        "opsworks",
	"opsworkscm":                      "opsworks-cm",
	"organizations":                   "organizations",
	"pi":                              "pi",
	"pinpoint":                        "pinpoint",
	"polly":                           "polly",
	"pricing":                         "api.pricing",
	"rds":                             "rds",
	"redshift":                        "redshift",
	"rekognition":                     "rekognition",
	"resourcegroups":                  "resource-groups",
	"resourcegroupstaggingapi":        "tagging",
	"route53":                         "route53",
	"route53domains":                  "route53domains",
	"lexruntimeservice":               "runtime.lex",
	"sagemakerruntime":                "runtime.sagemaker",
	"s3":                              "s3",
	"sagemaker":                       "sagemaker",
	"simpledb":                        "sdb",
	"secretsmanager":                  "secretsmanager",
	"serverlessapplicationrepository": "serverlessrepo",
	"servicecatalog":                  "servicecatalog",
	"servicediscovery":                "servicediscovery",
	"shield":                          "shield",
	"sms":                             "sms",
	"snowball":                        "snowball",
	"sns":                             "sns",
	"sqs":                             "sqs",
	"ssm":                             "ssm",
	"sfn":                             "states",
	"storagegateway":                  "storagegateway",
	"dynamodbstreams":                 "streams.dynamodb",
	"sts":                             "sts",
	"support":                         "support",
	"swf":                             "swf",
	"transcribeservice":               "transcribe",
	"translate":                       "translate",
	"wafregional":                     "waf-regional",
	"waf":                             "waf",
	"workdocs":                        "workdocs",
	"workmail":                        "workmail",
	"workspaces":                      "workspaces",
	"xray":                            "xray",
}
