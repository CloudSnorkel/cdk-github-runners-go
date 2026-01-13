package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
)

// Experimental.
type CodeBuildRunnerImageBuilderProps struct {
	// Build image to use in CodeBuild.
	//
	// This is the image that's going to run the code that builds the runner image.
	//
	// The only action taken in CodeBuild is running `docker build`. You would therefore not need to change this setting often.
	// Default: Amazon Linux 2023.
	//
	// Experimental.
	BuildImage awscodebuild.IBuildImage `field:"optional" json:"buildImage" yaml:"buildImage"`
	// The type of compute to use for this build. See the {@link ComputeType} enum for the possible values.
	//
	// The compute type determines CPU, memory, and disk space:
	// - SMALL: 2 vCPU, 3 GB RAM, 64 GB disk
	// - MEDIUM: 4 vCPU, 7 GB RAM, 128 GB disk
	// - LARGE: 8 vCPU, 15 GB RAM, 128 GB disk
	// - X2_LARGE: 72 vCPU, 145 GB RAM, 256 GB disk (Linux) or 824 GB disk (Windows)
	//
	// Use a larger compute type when you need more disk space for building larger Docker images.
	//
	// For more details, see https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html#environment.types
	// Default: {@link ComputeType#SMALL }.
	//
	// Experimental.
	ComputeType awscodebuild.ComputeType `field:"optional" json:"computeType" yaml:"computeType"`
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	// Default: Duration.hours(1)
	//
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

