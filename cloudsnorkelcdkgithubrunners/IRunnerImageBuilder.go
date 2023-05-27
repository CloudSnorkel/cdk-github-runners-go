package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for constructs that build an image that can be used in {@link IRunnerProvider }.
//
// An image can be a Docker image or AMI.
// Experimental.
type IRunnerImageBuilder interface {
	// Build and return an AMI with GitHub Runner installed in it.
	//
	// Anything that ends up with a launch template pointing to an AMI that runs GitHub self-hosted runners can be used. A simple implementation could even point to an existing AMI and nothing else.
	//
	// The AMI can be further updated over time manually or using a schedule as long as it is always written to the same launch template.
	// Experimental.
	BindAmi() *RunnerAmi
	// Build and return a Docker image with GitHub Runner installed in it.
	//
	// Anything that ends up with an ECR repository containing a Docker image that runs GitHub self-hosted runners can be used. A simple implementation could even point to an existing image and nothing else.
	//
	// It's important that the specified image tag be available at the time the repository is available. Providers usually assume the image is ready and will fail if it's not.
	//
	// The image can be further updated over time manually or using a schedule as long as it is always written to the same tag.
	// Experimental.
	BindDockerImage() *RunnerImage
}

// The jsii proxy for IRunnerImageBuilder
type jsiiProxy_IRunnerImageBuilder struct {
	_ byte // padding
}

func (i *jsiiProxy_IRunnerImageBuilder) BindAmi() *RunnerAmi {
	var returns *RunnerAmi

	_jsii_.Invoke(
		i,
		"bindAmi",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRunnerImageBuilder) BindDockerImage() *RunnerImage {
	var returns *RunnerImage

	_jsii_.Invoke(
		i,
		"bindDockerImage",
		nil, // no parameters
		&returns,
	)

	return returns
}

