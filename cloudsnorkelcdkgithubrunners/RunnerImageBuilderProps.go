package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Experimental.
type RunnerImageBuilderProps struct {
	// Image architecture.
	// Default: Architecture.X86_64
	//
	// Experimental.
	Architecture Architecture `field:"optional" json:"architecture" yaml:"architecture"`
	// Options specific to AWS Image Builder.
	//
	// Only used when builderType is RunnerImageBuilderType.AWS_IMAGE_BUILDER.
	// Experimental.
	AwsImageBuilderOptions *AwsImageBuilderRunnerImageBuilderProps `field:"optional" json:"awsImageBuilderOptions" yaml:"awsImageBuilderOptions"`
	// Base AMI from which runner AMIs will be built.
	//
	// This can be an actual AMI or an AWS Image Builder ARN that points to the latest AMI. For example `arn:aws:imagebuilder:us-east-1:aws:image/ubuntu-server-22-lts-x86/x.x.x` would always use the latest version of Ubuntu 22.04 in each build. If you want a specific version, you can replace `x.x.x` with that version.
	// Default: latest Ubuntu 22.04 AMI for Os.LINUX_UBUNTU, latest Amazon Linux 2 AMI for Os.LINUX_AMAZON_2, latest Windows Server 2022 AMI for Os.WINDOWS
	//
	// Experimental.
	BaseAmi *string `field:"optional" json:"baseAmi" yaml:"baseAmi"`
	// Base image from which Docker runner images will be built.
	// Default: public.ecr.aws/lts/ubuntu:22.04 for Os.LINUX_UBUNTU, public.ecr.aws/amazonlinux/amazonlinux:2 for Os.LINUX_AMAZON_2, mcr.microsoft.com/windows/servercore:ltsc2019-amd64 for Os.WINDOWS
	//
	// Experimental.
	BaseDockerImage *string `field:"optional" json:"baseDockerImage" yaml:"baseDockerImage"`
	// Default: CodeBuild for Linux Docker image, AWS Image Builder for Windows Docker image and any AMI.
	//
	// Experimental.
	BuilderType RunnerImageBuilderType `field:"optional" json:"builderType" yaml:"builderType"`
	// Options specific to CodeBuild image builder.
	//
	// Only used when builderType is RunnerImageBuilderType.CODE_BUILD.
	// Experimental.
	CodeBuildOptions *CodeBuildRunnerImageBuilderProps `field:"optional" json:"codeBuildOptions" yaml:"codeBuildOptions"`
	// Components to install on the image.
	// Default: none.
	//
	// Experimental.
	Components *[]RunnerImageComponent `field:"optional" json:"components" yaml:"components"`
	// Removal policy for logs of image builds.
	//
	// If deployment fails on the custom resource, try setting this to `RemovalPolicy.RETAIN`. This way the CodeBuild logs can still be viewed, and you can see why the build failed.
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
	// Default: OS.LINUX_UBUNTU
	//
	// Experimental.
	Os Os `field:"optional" json:"os" yaml:"os"`
	// Schedule the image to be rebuilt every given interval.
	//
	// Useful for keeping the image up-do-date with the latest GitHub runner version and latest OS updates.
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
	// Security Groups to assign to this instance.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Where to place the network interfaces within the VPC.
	// Default: no subnet.
	//
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// VPC to build the image in.
	// Default: no VPC.
	//
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

