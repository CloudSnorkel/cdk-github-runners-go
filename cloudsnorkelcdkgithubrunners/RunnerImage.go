package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Description of a Docker image built by {@link RunnerImageBuilder }.
// Experimental.
type RunnerImage struct {
	// Architecture of the image.
	// Experimental.
	Architecture Architecture `field:"required" json:"architecture" yaml:"architecture"`
	// ECR repository containing the image.
	// Experimental.
	ImageRepository awsecr.IRepository `field:"required" json:"imageRepository" yaml:"imageRepository"`
	// Static image tag where the image will be pushed.
	// Experimental.
	ImageTag *string `field:"required" json:"imageTag" yaml:"imageTag"`
	// OS type of the image.
	// Experimental.
	Os Os `field:"required" json:"os" yaml:"os"`
	// Installed runner version.
	// Deprecated: open a ticket if you need this.
	RunnerVersion RunnerVersion `field:"required" json:"runnerVersion" yaml:"runnerVersion"`
	// Log group where image builds are logged.
	// Experimental.
	LogGroup awslogs.LogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
}

