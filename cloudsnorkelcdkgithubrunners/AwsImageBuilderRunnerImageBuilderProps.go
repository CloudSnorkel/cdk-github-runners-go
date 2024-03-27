package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Experimental.
type AwsImageBuilderRunnerImageBuilderProps struct {
	// Options for fast launch.
	//
	// This is only supported for Windows AMIs.
	// Default: disabled.
	//
	// Experimental.
	FastLaunchOptions *FastLaunchOptions `field:"optional" json:"fastLaunchOptions" yaml:"fastLaunchOptions"`
	// The instance type used to build the image.
	// Default: m5.large
	//
	// Experimental.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
}

