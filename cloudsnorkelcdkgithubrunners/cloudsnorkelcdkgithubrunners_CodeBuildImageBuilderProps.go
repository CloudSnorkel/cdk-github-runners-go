// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Properties for CodeBuildImageBuilder construct.
// Experimental.
type CodeBuildImageBuilderProps struct {
	// Path to Dockerfile to be built.
	//
	// It can be a path to a Dockerfile, a folder containing a Dockerfile, or a zip file containing a Dockerfile.
	// Experimental.
	DockerfilePath *string `field:"required" json:"dockerfilePath" yaml:"dockerfilePath"`
	// Image architecture.
	// Experimental.
	Architecture Architecture `field:"optional" json:"architecture" yaml:"architecture"`
	// The type of compute to use for this build.
	//
	// See the {@link ComputeType} enum for the possible values.
	// Experimental.
	ComputeType awscodebuild.ComputeType `field:"optional" json:"computeType" yaml:"computeType"`
	// Removal policy for logs of image builds.
	//
	// If deployment fails on the custom resource, try setting this to `RemovalPolicy.RETAIN`. This way the CodeBuild logs can still be viewed, and you can see why the build failed.
	//
	// We try to not leave anything behind when removed. But sometimes a log staying behind is useful.
	// Experimental.
	LogRemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"logRemovalPolicy" yaml:"logRemovalPolicy"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// Image OS.
	// Experimental.
	Os Os `field:"optional" json:"os" yaml:"os"`
	// Schedule the image to be rebuilt every given interval.
	//
	// Useful for keeping the image up-do-date with the latest GitHub runner version and latest OS updates.
	//
	// Set to zero to disable.
	// Experimental.
	RebuildInterval awscdk.Duration `field:"optional" json:"rebuildInterval" yaml:"rebuildInterval"`
	// Version of GitHub Runners to install.
	// Experimental.
	RunnerVersion RunnerVersion `field:"optional" json:"runnerVersion" yaml:"runnerVersion"`
	// Security Group to assign to this instance.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Where to place the network interfaces within the VPC.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// VPC to build the image in.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

