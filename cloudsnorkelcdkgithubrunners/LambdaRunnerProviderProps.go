package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

type LambdaRunnerProviderProps struct {
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
	// The size of the function’s /tmp directory in MiB.
	// Default: 10 GiB.
	//
	EphemeralStorageSize awscdk.Size `field:"optional" json:"ephemeralStorageSize" yaml:"ephemeralStorageSize"`
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
	// Runner image builder used to build Docker images containing GitHub Runner and all requirements.
	//
	// The image builder must contain the {@link RunnerImageComponent.lambdaEntrypoint} component.
	//
	// The image builder determines the OS and architecture of the runner.
	// Default: LambdaRunnerProvider.imageBuilder()
	//
	ImageBuilder IRunnerImageBuilder `field:"optional" json:"imageBuilder" yaml:"imageBuilder"`
	// GitHub Actions labels used for this provider.
	//
	// These labels are used to identify which provider should spawn a new on-demand runner. Every job sends a webhook with the labels it's looking for
	// based on runs-on. We match the labels from the webhook with the labels specified here. If all the labels specified here are present in the
	// job's labels, this provider will be chosen and spawn a new runner.
	// Default: ['lambda'].
	//
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// The amount of memory, in MB, that is allocated to your Lambda function.
	//
	// Lambda uses this value to proportionally allocate the amount of CPU
	// power. For more information, see Resource Model in the AWS Lambda
	// Developer Guide.
	// Default: 2048.
	//
	MemorySize *float64 `field:"optional" json:"memorySize" yaml:"memorySize"`
	// Security groups to assign to this instance.
	// Default: public lambda with no security group.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Where to place the network interfaces within the VPC.
	// Default: no subnet.
	//
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The function execution time (in seconds) after which Lambda terminates the function.
	//
	// Because the execution time affects cost, set this value
	// based on the function's expected execution time.
	// Default: Duration.minutes(15)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// VPC to launch the runners in.
	// Default: no VPC.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

