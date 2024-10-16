package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Properties for EcsRunnerProvider.
// Experimental.
type EcsRunnerProviderProps struct {
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
	// Existing capacity provider to use.
	//
	// Make sure the AMI used by the capacity provider is compatible with ECS.
	// Default: new capacity provider.
	//
	// Experimental.
	CapacityProvider awsecs.AsgCapacityProvider `field:"optional" json:"capacityProvider" yaml:"capacityProvider"`
	// Existing ECS cluster to use.
	// Default: a new cluster.
	//
	// Experimental.
	Cluster awsecs.Cluster `field:"optional" json:"cluster" yaml:"cluster"`
	// The number of cpu units used by the task.
	//
	// 1024 units is 1 vCPU. Fractions of a vCPU are supported.
	// Default: 1024.
	//
	// Experimental.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Support building and running Docker images by enabling Docker-in-Docker (dind) and the required CodeBuild privileged mode.
	//
	// Disabling this can
	// speed up provisioning of CodeBuild runners. If you don't intend on running or building Docker images, disable this for faster start-up times.
	// Default: true.
	//
	// Experimental.
	DockerInDocker *bool `field:"optional" json:"dockerInDocker" yaml:"dockerInDocker"`
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
	// Default: EcsRunnerProvider.imageBuilder()
	//
	// Experimental.
	ImageBuilder IRunnerImageBuilder `field:"optional" json:"imageBuilder" yaml:"imageBuilder"`
	// Instance type of ECS cluster instances.
	//
	// Only used when creating a new cluster.
	// Default: m6i.large or m6g.large
	//
	// Experimental.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// GitHub Actions labels used for this provider.
	//
	// These labels are used to identify which provider should spawn a new on-demand runner. Every job sends a webhook with the labels it's looking for
	// based on runs-on. We match the labels from the webhook with the labels specified here. If all the labels specified here are present in the
	// job's labels, this provider will be chosen and spawn a new runner.
	// Default: ['ecs'].
	//
	// Experimental.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// The maximum number of instances to run in the cluster.
	//
	// Only used when creating a new cluster.
	// Default: 5.
	//
	// Experimental.
	MaxInstances *float64 `field:"optional" json:"maxInstances" yaml:"maxInstances"`
	// The amount (in MiB) of memory used by the task.
	// Default: 3500, unless `memoryReservationMiB` is used and then it's undefined.
	//
	// Experimental.
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// The soft limit (in MiB) of memory to reserve for the container.
	// Default: undefined.
	//
	// Experimental.
	MemoryReservationMiB *float64 `field:"optional" json:"memoryReservationMiB" yaml:"memoryReservationMiB"`
	// The minimum number of instances to run in the cluster.
	//
	// Only used when creating a new cluster.
	// Default: 0.
	//
	// Experimental.
	MinInstances *float64 `field:"optional" json:"minInstances" yaml:"minInstances"`
	// Security groups to assign to the task.
	// Default: a new security group.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Use spot capacity.
	// Default: false (true if spotMaxPrice is specified).
	//
	// Experimental.
	Spot *bool `field:"optional" json:"spot" yaml:"spot"`
	// Maximum price for spot instances.
	// Experimental.
	SpotMaxPrice *string `field:"optional" json:"spotMaxPrice" yaml:"spotMaxPrice"`
	// Options for runner instance storage volume.
	// Experimental.
	StorageOptions *StorageOptions `field:"optional" json:"storageOptions" yaml:"storageOptions"`
	// Size of volume available for launched cluster instances.
	//
	// This modifies the boot volume size and doesn't add any additional volumes.
	//
	// Each instance can be used by multiple runners, so make sure there is enough space for all of them.
	// Default: default size for AMI (usually 30GB for Linux and 50GB for Windows).
	//
	// Experimental.
	StorageSize awscdk.Size `field:"optional" json:"storageSize" yaml:"storageSize"`
	// Subnets to run the runners in.
	// Default: ECS default.
	//
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// VPC to launch the runners in.
	// Default: default account VPC.
	//
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

