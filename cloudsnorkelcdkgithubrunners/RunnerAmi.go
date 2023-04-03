// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Description of a AMI built by {@link RunnerImageBuilder }.
// Experimental.
type RunnerAmi struct {
	// Architecture of the image.
	// Experimental.
	Architecture Architecture `field:"required" json:"architecture" yaml:"architecture"`
	// Launch template pointing to the latest AMI.
	// Experimental.
	LaunchTemplate awsec2.ILaunchTemplate `field:"required" json:"launchTemplate" yaml:"launchTemplate"`
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

