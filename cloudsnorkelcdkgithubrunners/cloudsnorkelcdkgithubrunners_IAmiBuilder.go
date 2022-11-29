// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for constructs that build an AMI that can be used in {@link IRunnerProvider}.
//
// Anything that ends up with a launch template pointing to an AMI that runs GitHub self-hosted runners can be used. A simple implementation could even point to an existing AMI and nothing else.
//
// The AMI can be further updated over time manually or using a schedule as long as it is always written to the same launch template.
// Experimental.
type IAmiBuilder interface {
	// Finalize and return all required information about the AMI built by this builder.
	//
	// This method can be called multiple times if the image is bound to multiple providers. Make sure you cache the image when implementing or return an error if this builder doesn't support reusing images.
	//
	// Returns: ami.
	// Experimental.
	Bind() *RunnerAmi
}

// The jsii proxy for IAmiBuilder
type jsiiProxy_IAmiBuilder struct {
	_ byte // padding
}

func (i *jsiiProxy_IAmiBuilder) Bind() *RunnerAmi {
	var returns *RunnerAmi

	_jsii_.Invoke(
		i,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

