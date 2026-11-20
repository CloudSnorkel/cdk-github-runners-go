package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Experimental.
type MicroVmRunnerProviderProps struct {
	// Add default labels based on OS and architecture of the runner.
	//
	// This will tell GitHub Runner to add default labels like `self-hosted`, `linux`, `x64`, and `arm64`.
	// Default: true.
	//
	// Experimental.
	DefaultLabels *bool `field:"optional" json:"defaultLabels" yaml:"defaultLabels"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Default: logs.RetentionDays.ONE_MONTH
	//
	// Experimental.
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
	// Experimental.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Runner image builder used to build Docker images containing GitHub Runner and all requirements.
	//
	// The image builder must contain the {@link RunnerImageComponent.microVmEntrypoint} component.
	// Default: MicroVmRunnerProvider.imageBuilder()
	//
	// Experimental.
	ImageBuilder IRunnerImageBuilder `field:"optional" json:"imageBuilder" yaml:"imageBuilder"`
	// GitHub Actions labels used for this provider.
	//
	// These labels are used to identify which provider should spawn a new on-demand runner. Every job sends a webhook with the labels it's looking for
	// based on runs-on. We match the labels from the webhook with the labels specified here. If all the labels specified here are present in the
	// job's labels, this provider will be chosen and spawn a new runner.
	// Default: ['microvm'].
	//
	// Experimental.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Memory given to the MicroVM.
	//
	// CPU is allocated proportionally, about one vCPU for every 2 GiB.
	// Default: 4 GiB (and therefore 2 vCPU).
	//
	// Experimental.
	MemorySize awscdk.Size `field:"optional" json:"memorySize" yaml:"memorySize"`
	// Security groups to assign to the outbound traffic.
	// Default: a new security group, if VPC is used.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Where to place the network interfaces within the VPC.
	// Default: default subnets of the VPC.
	//
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// Time after which the MicroVM is terminated, whether the job is done or not.
	//
	// Must be 8 hours or less.
	// Default: Duration.hours(1)
	//
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// VPC to send the runner's outbound traffic through.
	//
	// The network interfaces never get a public IP, so the subnets
	// need a NAT gateway (or VPC endpoints) to reach GitHub, ECR and CloudWatch. Public subnets won't work.
	// Default: no VPC, runners get public internet access.
	//
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

