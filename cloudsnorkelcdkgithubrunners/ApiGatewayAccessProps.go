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
	// Creates a private API Gateway and allows access from the specified VPC.
	// Experimental.
	AllowedVpc awsec2.IVpc `field:"optional" json:"allowedVpc" yaml:"allowedVpc"`
}

