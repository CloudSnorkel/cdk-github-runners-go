// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Properties for {@link Ec2RunnerProvider} construct.
// Experimental.
type Ec2RunnerProviderProps struct {
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// Options to retry operation in case of failure like missing capacity, or API quota issues.
	// Experimental.
	RetryOptions *ProviderRetryOptions `field:"optional" json:"retryOptions" yaml:"retryOptions"`
	// Deprecated: use imageBuilder.
	AmiBuilder IRunnerImageBuilder `field:"optional" json:"amiBuilder" yaml:"amiBuilder"`
	// Runner image builder used to build AMI containing GitHub Runner and all requirements.
	//
	// The image builder determines the OS and architecture of the runner.
	// Experimental.
	ImageBuilder IRunnerImageBuilder `field:"optional" json:"imageBuilder" yaml:"imageBuilder"`
	// Instance type for launched runner instances.
	// Experimental.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// GitHub Actions labels used for this provider.
	//
	// These labels are used to identify which provider should spawn a new on-demand runner. Every job sends a webhook with the labels it's looking for
	// based on runs-on. We match the labels from the webhook with the labels specified here. If all the labels specified here are present in the
	// job's labels, this provider will be chosen and spawn a new runner.
	// Experimental.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Security Group to assign to launched runner instances.
	// Deprecated: use {@link securityGroups }.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Security groups to assign to launched runner instances.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Use spot instances to save money.
	//
	// Spot instances are cheaper but not always available and can be stopped prematurely.
	// Experimental.
	Spot *bool `field:"optional" json:"spot" yaml:"spot"`
	// Set a maximum price for spot instances.
	// Experimental.
	SpotMaxPrice *string `field:"optional" json:"spotMaxPrice" yaml:"spotMaxPrice"`
	// Size of volume available for launched runner instances.
	//
	// This modifies the boot volume size and doesn't add any additional volumes.
	// Experimental.
	StorageSize awscdk.Size `field:"optional" json:"storageSize" yaml:"storageSize"`
	// Subnet where the runner instances will be launched.
	// Deprecated: use {@link vpc } and {@link subnetSelection }.
	Subnet awsec2.ISubnet `field:"optional" json:"subnet" yaml:"subnet"`
	// Where to place the network interfaces within the VPC.
	//
	// Only the first matched subnet will be used.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// VPC where runner instances will be launched.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

