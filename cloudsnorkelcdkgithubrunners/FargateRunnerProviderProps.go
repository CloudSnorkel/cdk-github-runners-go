package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Properties for FargateRunnerProvider.
// Experimental.
type FargateRunnerProviderProps struct {
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
	// Assign public IP to the runner task.
	//
	// Make sure the task will have access to GitHub. A public IP might be required unless you have NAT gateway.
	// Default: true.
	//
	// Experimental.
	AssignPublicIp *bool `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Existing Fargate cluster to use.
	// Default: a new cluster.
	//
	// Experimental.
	Cluster awsecs.Cluster `field:"optional" json:"cluster" yaml:"cluster"`
	// The number of cpu units used by the task.
	//
	// For tasks using the Fargate launch type,
	// this field is required and you must use one of the following values,
	// which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB)
	//
	// 512 (.5 vCPU) - Available memory values: 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB)
	//
	// 1024 (1 vCPU) - Available memory values: 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB)
	//
	// 2048 (2 vCPU) - Available memory values: Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB)
	//
	// 4096 (4 vCPU) - Available memory values: Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB).
	// Default: 1024.
	//
	// Experimental.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// The amount (in GiB) of ephemeral storage to be allocated to the task.
	//
	// The maximum supported value is 200 GiB.
	//
	// NOTE: This parameter is only supported for tasks hosted on AWS Fargate using platform version 1.4.0 or later.
	// Default: 20.
	//
	// Experimental.
	EphemeralStorageGiB *float64 `field:"optional" json:"ephemeralStorageGiB" yaml:"ephemeralStorageGiB"`
	// GitHub Actions runner group name.
	//
	// If specified, the runner will be registered with this group name. Setting a runner group can help managing access to self-hosted runners. It
	// requires a paid GitHub account.
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
	// The image builder determines the OS and architecture of the runner.
	// Default: FargateRunnerProvider.imageBuilder()
	//
	// Experimental.
	ImageBuilder IRunnerImageBuilder `field:"optional" json:"imageBuilder" yaml:"imageBuilder"`
	// GitHub Actions label used for this provider.
	// Default: undefined.
	//
	// Deprecated: use {@link labels } instead.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// GitHub Actions labels used for this provider.
	//
	// These labels are used to identify which provider should spawn a new on-demand runner. Every job sends a webhook with the labels it's looking for
	// based on runs-on. We match the labels from the webhook with the labels specified here. If all the labels specified here are present in the
	// job's labels, this provider will be chosen and spawn a new runner.
	// Default: ['fargate'].
	//
	// Experimental.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// The amount (in MiB) of memory used by the task.
	//
	// For tasks using the Fargate launch type,
	// this field is required and you must use one of the following values, which determines your range of valid values for the cpu parameter:
	//
	// 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available cpu values: 256 (.25 vCPU)
	//
	// 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available cpu values: 512 (.5 vCPU)
	//
	// 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB) - Available cpu values: 1024 (1 vCPU)
	//
	// Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) - Available cpu values: 2048 (2 vCPU)
	//
	// Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available cpu values: 4096 (4 vCPU).
	// Default: 2048.
	//
	// Experimental.
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// Security group to assign to the task.
	// Default: a new security group.
	//
	// Deprecated: use {@link securityGroups }.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Security groups to assign to the task.
	// Default: a new security group.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Use Fargate spot capacity provider to save money.
	//
	// * Runners may fail to start due to missing capacity.
	// * Runners might be stopped prematurely with spot pricing.
	// Default: false.
	//
	// Experimental.
	Spot *bool `field:"optional" json:"spot" yaml:"spot"`
	// Subnets to run the runners in.
	// Default: Fargate default.
	//
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// VPC to launch the runners in.
	// Default: default account VPC.
	//
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

