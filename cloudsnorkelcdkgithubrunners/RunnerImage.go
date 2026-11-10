package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Description of a Docker image built by {@link RunnerImageBuilder }.
type RunnerImage struct {
	// Architecture of the image.
	Architecture Architecture `field:"required" json:"architecture" yaml:"architecture"`
	// ECR repository containing the image.
	ImageRepository awsecr.IRepository `field:"required" json:"imageRepository" yaml:"imageRepository"`
	// Static image tag where the image will be pushed.
	ImageTag *string `field:"required" json:"imageTag" yaml:"imageTag"`
	// OS type of the image.
	Os Os `field:"required" json:"os" yaml:"os"`
	// Log group where image builds are logged.
	LogGroup awslogs.LogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
}

