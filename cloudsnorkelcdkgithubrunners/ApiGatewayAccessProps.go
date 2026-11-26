package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

type ApiGatewayAccessProps struct {
	// Log group for access logs, with one JSON line for each request.
	//
	// API Gateway can only write logs once the account has a CloudWatch Logs role. See
	// https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html
	// Default: no access logs.
	//
	AccessLogGroup awslogs.ILogGroup `field:"optional" json:"accessLogGroup" yaml:"accessLogGroup"`
	// List of IP addresses in CIDR notation that are allowed to access the API Gateway.
	//
	// If not specified on public API Gateway, all IP addresses are allowed.
	//
	// If not specified on private API Gateway, no IP addresses are allowed (but specified security groups are).
	AllowedIps *[]*string `field:"optional" json:"allowedIps" yaml:"allowedIps"`
	// List of security groups that are allowed to access the API Gateway.
	//
	// Only works for private API Gateways with {@link ApiGatewayAccessProps.allowedVpc}.
	AllowedSecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"allowedSecurityGroups" yaml:"allowedSecurityGroups"`
	// Create a private API Gateway and allow access from the specified VPC.
	AllowedVpc awsec2.IVpc `field:"optional" json:"allowedVpc" yaml:"allowedVpc"`
	// Create a private API Gateway and allow access from the specified VPC endpoints.
	//
	// Use this to make use of existing VPC endpoints or to share an endpoint between multiple functions. The VPC endpoint must point to `ec2.InterfaceVpcEndpointAwsService.APIGATEWAY`.
	//
	// No other settings are supported when using this option.
	//
	// All endpoints will be allowed access, but only the first one will be used as the URL by the runner system for setting up the webhook, and as setup URL.
	AllowedVpcEndpoints *[]awsec2.IVpcEndpoint `field:"optional" json:"allowedVpcEndpoints" yaml:"allowedVpcEndpoints"`
	// Level of execution logs. API Gateway writes them to its own log group named `API-Gateway-Execution-Logs_{rest-api-id}/{stage}`.
	//
	// Like access logs, these need a CloudWatch Logs role in the account.
	// Default: no execution logs.
	//
	LoggingLevel awsapigateway.MethodLoggingLevel `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
}

