package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
	// Default: m6i.large
	//
	// Experimental.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Size of volume available for builder instances. This modifies the boot volume size and doesn't add any additional volumes.
	//
	// Use this if you're building images with big components and need more space.
	// Default: default size for AMI (usually 30GB for Linux and 50GB for Windows).
	//
	// Experimental.
	StorageSize awscdk.Size `field:"optional" json:"storageSize" yaml:"storageSize"`
}

