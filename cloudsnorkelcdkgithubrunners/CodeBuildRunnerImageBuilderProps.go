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
	// Experimental.
	BuildImage awscodebuild.IBuildImage `field:"optional" json:"buildImage" yaml:"buildImage"`
	// The type of compute to use for this build.
	//
	// See the {@link ComputeType} enum for the possible values.
	// Experimental.
	ComputeType awscodebuild.ComputeType `field:"optional" json:"computeType" yaml:"computeType"`
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

