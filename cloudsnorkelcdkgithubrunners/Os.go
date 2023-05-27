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
	// Checks if this OS is in a given list.
	// Experimental.
	IsIn(oses *[]Os) *bool
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

func Os_LINUX_AMAZON_2() Os {
	_init_.Initialize()
	var returns Os
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.Os",
		"LINUX_AMAZON_2",
		&returns,
	)
	return returns
}

func Os_LINUX_UBUNTU() Os {
	_init_.Initialize()
	var returns Os
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.Os",
		"LINUX_UBUNTU",
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
	if err := o.validateIsParameters(os); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		o,
		"is",
		[]interface{}{os},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Os) IsIn(oses *[]Os) *bool {
	if err := o.validateIsInParameters(oses); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		o,
		"isIn",
		[]interface{}{oses},
		&returns,
	)

	return returns
}

