package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Description of a AMI built by {@link RunnerImageBuilder }.
type RunnerAmi struct {
	// Architecture of the image.
	Architecture Architecture `field:"required" json:"architecture" yaml:"architecture"`
	// Launch template pointing to the latest AMI.
	LaunchTemplate awsec2.ILaunchTemplate `field:"required" json:"launchTemplate" yaml:"launchTemplate"`
	// OS type of the image.
	Os Os `field:"required" json:"os" yaml:"os"`
	// Set this to a value that changes whenever the AMI changes (the AMI id or any version string works).
	//
	// It's used to know when the AMI's root device name needs to be looked up again. If left empty, the root
	// device name is looked up once and reused. That's fine as long as the AMI's root device never changes.
	//
	// This value may be used for other things in the future that require knowing when the AMI changed.
	CacheKey *string `field:"optional" json:"cacheKey" yaml:"cacheKey"`
	// Log group where image builds are logged.
	LogGroup awslogs.LogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
}

