package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Properties for {@link Ec2RunnerProvider} construct.
type Ec2RunnerProviderProps struct {
	// Add default labels based on OS and architecture of the runner.
	//
	// This will tell GitHub Runner to add default labels like `self-hosted`, `linux`, `x64`, and `arm64`.
	// Default: true.
	//
	DefaultLabels *bool `field:"optional" json:"defaultLabels" yaml:"defaultLabels"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Default: logs.RetentionDays.ONE_MONTH
	//
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// GitHub Actions runner group name.
	//
	// If specified, the runner will be registered with this group name. Setting a runner group can help managing access to self-hosted runners. It
	// requires a paid GitHub account and organization level runner registration.
	//
	// The group must exist or the runner will not start.
	//
	// Users will still be able to trigger this runner with the correct labels. But the runner will only be able to run jobs from repos allowed to use the group.
	// Default: undefined.
	//
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Runner image builder used to build AMI containing GitHub Runner and all requirements.
	//
	// The image builder determines the OS and architecture of the runner.
	// Default: Ec2RunnerProvider.imageBuilder()
	//
	ImageBuilder IRunnerImageBuilder `field:"optional" json:"imageBuilder" yaml:"imageBuilder"`
	// Instance type for launched runner instances.
	//
	// For GPU instance types (g4dn, g5, p3, etc.), we automatically use a GPU base image (AWS Deep Learning AMI)
	// with NVIDIA drivers pre-installed. If you provide your own image builder, use
	// `baseAmi: BaseImage.fromGpuBase(os, architecture)` or another image preloaded with NVIDIA drivers, or use
	// an image component to install NVIDIA drivers.
	// Default: m6i.large
	//
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// GitHub Actions labels used for this provider.
	//
	// These labels are used to identify which provider should spawn a new on-demand runner. Every job sends a webhook with the labels it's looking for
	// based on runs-on. We match the labels from the webhook with the labels specified here. If all the labels specified here are present in the
	// job's labels, this provider will be chosen and spawn a new runner.
	// Default: ['ec2'].
	//
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Security groups to assign to launched runner instances.
	// Default: a new security group.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Use spot instances to save money.
	//
	// Spot instances are cheaper but not always available and can be stopped prematurely.
	// Default: false.
	//
	Spot *bool `field:"optional" json:"spot" yaml:"spot"`
	// Set a maximum price for spot instances.
	// Default: no max price (you will pay current spot price).
	//
	SpotMaxPrice *string `field:"optional" json:"spotMaxPrice" yaml:"spotMaxPrice"`
	// Options for runner instance storage volume.
	StorageOptions *StorageOptions `field:"optional" json:"storageOptions" yaml:"storageOptions"`
	// Size of volume available for launched runner instances.
	//
	// This modifies the boot volume size and doesn't add any additional volumes.
	// Default: 30GB.
	//
	StorageSize awscdk.Size `field:"optional" json:"storageSize" yaml:"storageSize"`
	// Where to place the network interfaces within the VPC.
	//
	// Only the first matched subnet will be used.
	// Default: default VPC subnet.
	//
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// Additional tags to apply to launched runner instances and their volumes.
	//
	// These additional tags are set on top of `Name`, `GitHubRunners:Provider`, `GitHubRunners:Repo`, and `GitHubRunners:Labels`, and
	// `GitHubRunners:Runner`. You may only override `Name`. Tags prefixed `GitHubRunners:` are reserved and rejected. Some of our reserved tags are
	// used for house-keeping purposes.
	// Default: no additional tags.
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// VPC where runner instances will be launched.
	// Default: default account VPC.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

