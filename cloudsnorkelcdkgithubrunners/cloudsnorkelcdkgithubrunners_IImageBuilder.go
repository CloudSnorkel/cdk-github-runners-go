// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for constructs that build an image that can be used in {@link IRunnerProvider}.
//
// Anything that ends up with an ECR repository containing a Docker image that runs GitHub self-hosted runners can be used. A simple implementation could even point to an existing image and nothing else.
//
// It's important that the specified image tag be available at the time the repository is available. Providers usually assume the image is ready and will fail if it's not.
//
// The image can be further updated over time manually or using a schedule as long as it is always written to the same tag.
// Experimental.
type IImageBuilder interface {
	// ECR repository containing the image.
	//
	// This method can be called multiple times if the image is bound to multiple providers. Make sure you cache the image when implementing or return an error if this builder doesn't support reusing images.
	//
	// Returns: image.
	// Experimental.
	Bind() *RunnerImage
}

// The jsii proxy for IImageBuilder
type jsiiProxy_IImageBuilder struct {
	_ byte // padding
}

func (i *jsiiProxy_IImageBuilder) Bind() *RunnerImage {
	var returns *RunnerImage

	_jsii_.Invoke(
		i,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

