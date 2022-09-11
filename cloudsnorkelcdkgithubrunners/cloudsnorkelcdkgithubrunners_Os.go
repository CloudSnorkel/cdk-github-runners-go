// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"
)

// OS enum for an image.
// Experimental.
type Os interface {
	// Experimental.
	Name() *string
	// Checks if the given OS is the same as this one.
	// Experimental.
	Is(os Os) *bool
}

// The jsii proxy struct for Os
type jsiiProxy_Os struct {
	_ byte // padding
}

func (j *jsiiProxy_Os) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func Os_LINUX() Os {
	_init_.Initialize()
	var returns Os
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.Os",
		"LINUX",
		&returns,
	)
	return returns
}

func Os_WINDOWS() Os {
	_init_.Initialize()
	var returns Os
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.Os",
		"WINDOWS",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_Os) Is(os Os) *bool {
	var returns *bool

	_jsii_.Invoke(
		o,
		"is",
		[]interface{}{os},
		&returns,
	)

	return returns
}

