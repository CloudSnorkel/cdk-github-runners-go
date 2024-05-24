package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Properties for {@link AmiBuilder} construct.
// Experimental.
type AmiBuilderProps struct {
	// Image architecture.
	// Default: Architecture.X86_64
	//
	// Experimental.
	Architecture Architecture `field:"optional" json:"architecture" yaml:"architecture"`
	// Install Docker inside the image, so it can be used by the runner.
	// Default: true.
	//
	// Experimental.
	InstallDocker *bool `field:"optional" json:"installDocker" yaml:"installDocker"`
	// The instance type used to build the image.
	// Default: m6i.large
	//
	// Experimental.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Removal policy for logs of image builds.
	//
	// If deployment fails on the custom resource, try setting this to `RemovalPolicy.RETAIN`. This way the logs can still be viewed, and you can see why the build failed.
	//
	// We try to not leave anything behind when removed. But sometimes a log staying behind is useful.
	// Default: RemovalPolicy.DESTROY
	//
	// Experimental.
	LogRemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"logRemovalPolicy" yaml:"logRemovalPolicy"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Default: logs.RetentionDays.ONE_MONTH
	//
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// Image OS.
	// Default: OS.LINUX
	//
	// Experimental.
	Os Os `field:"optional" json:"os" yaml:"os"`
	// Schedule the AMI to be rebuilt every given interval.
	//
	// Useful for keeping the AMI up-do-date with the latest GitHub runner version and latest OS updates.
	//
	// Set to zero to disable.
	// Default: Duration.days(7)
	//
	// Experimental.
	RebuildInterval awscdk.Duration `field:"optional" json:"rebuildInterval" yaml:"rebuildInterval"`
	// Version of GitHub Runners to install.
	// Default: latest version available.
	//
	// Experimental.
	RunnerVersion RunnerVersion `field:"optional" json:"runnerVersion" yaml:"runnerVersion"`
	// Security group to assign to launched builder instances.
	// Default: new security group.
	//
	// Deprecated: use {@link securityGroups }.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Security groups to assign to launched builder instances.
	// Default: new security group.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Where to place the network interfaces within the VPC.
	//
	// Only the first matched subnet will be used.
	// Default: default VPC subnet.
	//
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// VPC where builder instances will be launched.
	// Default: default account VPC.
	//
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

