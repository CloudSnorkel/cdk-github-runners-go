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
	// Default: logs.RetentionDays.ONE_MONTH
	//
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// Deprecated: use {@link retryOptions } on {@link GitHubRunners } instead.
	RetryOptions *ProviderRetryOptions `field:"optional" json:"retryOptions" yaml:"retryOptions"`
	// Deprecated: use imageBuilder.
	AmiBuilder IRunnerImageBuilder `field:"optional" json:"amiBuilder" yaml:"amiBuilder"`
	// Runner image builder used to build AMI containing GitHub Runner and all requirements.
	//
	// The image builder determines the OS and architecture of the runner.
	// Default: Ec2RunnerProvider.imageBuilder()
	//
	// Experimental.
	ImageBuilder IRunnerImageBuilder `field:"optional" json:"imageBuilder" yaml:"imageBuilder"`
	// Instance type for launched runner instances.
	// Default: m6i.large
	//
	// Experimental.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// GitHub Actions labels used for this provider.
	//
	// These labels are used to identify which provider should spawn a new on-demand runner. Every job sends a webhook with the labels it's looking for
	// based on runs-on. We match the labels from the webhook with the labels specified here. If all the labels specified here are present in the
	// job's labels, this provider will be chosen and spawn a new runner.
	// Default: ['ec2'].
	//
	// Experimental.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Security Group to assign to launched runner instances.
	// Default: a new security group.
	//
	// Deprecated: use {@link securityGroups }.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Security groups to assign to launched runner instances.
	// Default: a new security group.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Use spot instances to save money.
	//
	// Spot instances are cheaper but not always available and can be stopped prematurely.
	// Default: false.
	//
	// Experimental.
	Spot *bool `field:"optional" json:"spot" yaml:"spot"`
	// Set a maximum price for spot instances.
	// Default: no max price (you will pay current spot price).
	//
	// Experimental.
	SpotMaxPrice *string `field:"optional" json:"spotMaxPrice" yaml:"spotMaxPrice"`
	// Options for runner instance storage volume.
	// Experimental.
	StorageOptions *StorageOptions `field:"optional" json:"storageOptions" yaml:"storageOptions"`
	// Size of volume available for launched runner instances.
	//
	// This modifies the boot volume size and doesn't add any additional volumes.
	// Default: 30GB.
	//
	// Experimental.
	StorageSize awscdk.Size `field:"optional" json:"storageSize" yaml:"storageSize"`
	// Subnet where the runner instances will be launched.
	// Default: default subnet of account's default VPC.
	//
	// Deprecated: use {@link vpc } and {@link subnetSelection }.
	Subnet awsec2.ISubnet `field:"optional" json:"subnet" yaml:"subnet"`
	// Where to place the network interfaces within the VPC.
	//
	// Only the first matched subnet will be used.
	// Default: default VPC subnet.
	//
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// VPC where runner instances will be launched.
	// Default: default account VPC.
	//
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

