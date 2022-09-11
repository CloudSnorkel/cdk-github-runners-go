// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

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
	// Image digest for providers that need to know the digest like Lambda.
	//
	// If the digest is not specified, imageTag must always point to a new tag on update. If not, the build may try to use the old image.
	//
	// WARNING: the digest might change when the builder automatically rebuilds the image on a schedule. Do not expect for this digest to stay the same between deploys.
	// Experimental.
	ImageDigest *string `field:"optional" json:"imageDigest" yaml:"imageDigest"`
	// Log group where image builds are logged.
	// Experimental.
	LogGroup awslogs.LogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
}

