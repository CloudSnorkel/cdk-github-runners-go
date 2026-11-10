package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

type RunnerImageBuilderProps struct {
	// Image architecture.
	// Default: Architecture.X86_64
	//
	Architecture Architecture `field:"optional" json:"architecture" yaml:"architecture"`
	// Options specific to AWS Image Builder.
	//
	// Only used when builderType is RunnerImageBuilderType.AWS_IMAGE_BUILDER.
	AwsImageBuilderOptions *AwsImageBuilderRunnerImageBuilderProps `field:"optional" json:"awsImageBuilderOptions" yaml:"awsImageBuilderOptions"`
	// Base AMI from which runner AMIs will be built.
	//
	// This can be a BaseImage instance created using static factory methods:
	//   - `BaseImage.fromAmiId('ami-12345')` - Use an AMI ID
	//   - `BaseImage.fromString('arn:aws:imagebuilder:...')` - Use any string (ARN, AMI ID, etc.)
	//   - `BaseImage.fromSsmParameter(parameter)` - Use an SSM parameter object
	//   - `BaseImage.fromSsmParameterName('/aws/service/ami/...')` - Use an SSM parameter by name
	//   - `BaseImage.fromMarketplaceProductId('product-id')` - Use a Marketplace product ID
	//   - `BaseImage.fromImageBuilder(scope, 'ubuntu-server-22-lts-x86')` - Use an AWS-provided Image Builder image
	//
	// For example `BaseImage.fromImageBuilder(scope, 'ubuntu-server-22-lts-x86')` would always use the latest version of Ubuntu 22.04 in each build. If you want a specific version, you can pass the version as the third parameter.
	// Default: latest Ubuntu 22.04 AMI for Os.LINUX_UBUNTU and Os.LINUX_UBUNTU_2204, Ubuntu 24.04 AMI for Os.LINUX_UBUNTU_2404, latest Amazon Linux 2 AMI for Os.LINUX_AMAZON_2, latest Windows Server 2022 AMI for Os.WINDOWS
	//
	BaseAmi BaseImage `field:"optional" json:"baseAmi" yaml:"baseAmi"`
	// Base image from which Docker runner images will be built.
	//
	// This can be a BaseContainerImage instance created using static factory methods:
	//   - `BaseContainerImage.fromDockerHub('ubuntu', '22.04')` - Use DockerHub
	//   - `BaseContainerImage.fromEcr(repo, 'latest')` - Use ECR (automatically grants permissions with CodeBuild)
	//   - `BaseContainerImage.fromEcrPublic('lts', 'ubuntu', '22.04')` - Use ECR Public
	//   - `BaseContainerImage.fromString('public.ecr.aws/lts/ubuntu:22.04')` - Use any string
	//
	// When using private images from a different account or not on ECR, you may need to include additional setup commands with {@link dockerSetupCommands}.
	// Default: public.ecr.aws/lts/ubuntu:22.04 for Os.LINUX_UBUNTU and Os.LINUX_UBUNTU_2204, public.ecr.aws/lts/ubuntu:24.04 for Os.LINUX_UBUNTU_2404, public.ecr.aws/amazonlinux/amazonlinux:2 for Os.LINUX_AMAZON_2, mcr.microsoft.com/windows/servercore:ltsc2019-amd64 for Os.WINDOWS
	//
	BaseDockerImage BaseContainerImage `field:"optional" json:"baseDockerImage" yaml:"baseDockerImage"`
	// Default: CodeBuild for Linux Docker image, AWS Image Builder for Windows Docker image and any AMI.
	//
	BuilderType RunnerImageBuilderType `field:"optional" json:"builderType" yaml:"builderType"`
	// Options specific to CodeBuild image builder.
	//
	// Only used when builderType is RunnerImageBuilderType.CODE_BUILD.
	CodeBuildOptions *CodeBuildRunnerImageBuilderProps `field:"optional" json:"codeBuildOptions" yaml:"codeBuildOptions"`
	// Components to install on the image.
	// Default: none.
	//
	Components *[]RunnerImageComponent `field:"optional" json:"components" yaml:"components"`
	// Additional commands to run on the build host before starting the Docker runner image build.
	//
	// Use this to execute commands such as `docker login` or `aws ecr get-login-password` to pull private base images.
	// Default: [].
	//
	DockerSetupCommands *[]*string `field:"optional" json:"dockerSetupCommands" yaml:"dockerSetupCommands"`
	// Removal policy for logs of image builds.
	//
	// If deployment fails on the custom resource, try setting this to `RemovalPolicy.RETAIN`. This way the CodeBuild logs can still be viewed, and you can see why the build failed.
	//
	// We try to not leave anything behind when removed. But sometimes a log staying behind is useful.
	// Default: RemovalPolicy.DESTROY
	//
	LogRemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"logRemovalPolicy" yaml:"logRemovalPolicy"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Default: logs.RetentionDays.ONE_MONTH
	//
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// Image OS.
	// Default: OS.LINUX_UBUNTU
	//
	Os Os `field:"optional" json:"os" yaml:"os"`
	// Schedule the image to be rebuilt every given interval.
	//
	// Useful for keeping the image up-do-date with the latest GitHub runner version and latest OS updates.
	//
	// Set to zero to disable.
	// Default: Duration.days(7)
	//
	RebuildInterval awscdk.Duration `field:"optional" json:"rebuildInterval" yaml:"rebuildInterval"`
	// Version of GitHub Runners to install.
	// Default: latest version available.
	//
	RunnerVersion RunnerVersion `field:"optional" json:"runnerVersion" yaml:"runnerVersion"`
	// Security Groups to assign to this instance.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Where to place the network interfaces within the VPC.
	// Default: no subnet.
	//
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// VPC to build the image in.
	// Default: no VPC.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Wait for image to finish building during deployment.
	//
	// It's usually best to leave this enabled to ensure everything is ready once deployment is done. However, it can be disabled to speed up deployment in case where you have a lot of image components that can take a long time to build.
	//
	// Disabling this option means a finished deployment is not ready to be used. You will have to wait for the image to finish building before the system can be used.
	//
	// Disabling this option may also mean any changes to settings or components can take up to a week (default rebuild interval) to take effect.
	// Default: true.
	//
	WaitOnDeploy *bool `field:"optional" json:"waitOnDeploy" yaml:"waitOnDeploy"`
}

