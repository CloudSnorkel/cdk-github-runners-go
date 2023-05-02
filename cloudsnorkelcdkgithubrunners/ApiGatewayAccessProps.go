// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Experimental.
type ApiGatewayAccessProps struct {
	// List of IP addresses in CIDR notation that are allowed to access the API Gateway.
	//
	// If not specified on public API Gateway, all IP addresses are allowed.
	//
	// If not specified on private API Gateway, no IP addresses are allowed (but specified security groups are).
	// Experimental.
	AllowedIps *[]*string `field:"optional" json:"allowedIps" yaml:"allowedIps"`
	// List of security groups that are allowed to access the API Gateway.
	//
	// Only works for private API Gateways with {@link allowedVpc}.
	// Experimental.
	AllowedSecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"allowedSecurityGroups" yaml:"allowedSecurityGroups"`
	// Create a private API Gateway and allow access from the specified VPC.
	// Experimental.
	AllowedVpc awsec2.IVpc `field:"optional" json:"allowedVpc" yaml:"allowedVpc"`
	// Create a private API Gateway and allow access from the specified VPC endpoints.
	//
	// Use this to make use of existing VPC endpoints. The VPC endpoint must point to `ec2.InterfaceVpcEndpointAwsService.APIGATEWAY`.
	//
	// No other settings are supported when using this option.
	// Experimental.
	AllowedVpcEndpoints *[]awsec2.IVpcEndpoint `field:"optional" json:"allowedVpcEndpoints" yaml:"allowedVpcEndpoints"`
}

