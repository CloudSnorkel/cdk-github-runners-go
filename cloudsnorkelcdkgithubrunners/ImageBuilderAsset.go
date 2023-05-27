package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
)

// An asset including file or directory to place inside the built image.
// Experimental.
type ImageBuilderAsset struct {
	// Asset to place in the image.
	// Experimental.
	Asset awss3assets.Asset `field:"required" json:"asset" yaml:"asset"`
	// Path to place asset in the image.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
}

