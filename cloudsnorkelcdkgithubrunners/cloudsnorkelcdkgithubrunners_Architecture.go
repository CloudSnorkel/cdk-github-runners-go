// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// CPU architecture enum for an image.
// Experimental.
type Architecture interface {
	// Experimental.
	Name() *string
	// Checks if a given EC2 instance type matches this architecture.
	// Experimental.
	InstanceTypeMatch(instanceType awsec2.InstanceType) *bool
	// Checks if the given architecture is the same as this one.
	// Experimental.
	Is(arch Architecture) *bool
	// Checks if this architecture is in a given list.
	// Experimental.
	IsIn(arches *[]Architecture) *bool
}

// The jsii proxy struct for Architecture
type jsiiProxy_Architecture struct {
	_ byte // padding
}

func (j *jsiiProxy_Architecture) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func Architecture_ARM64() Architecture {
	_init_.Initialize()
	var returns Architecture
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.Architecture",
		"ARM64",
		&returns,
	)
	return returns
}

func Architecture_X86_64() Architecture {
	_init_.Initialize()
	var returns Architecture
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.Architecture",
		"X86_64",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_Architecture) InstanceTypeMatch(instanceType awsec2.InstanceType) *bool {
	if err := a.validateInstanceTypeMatchParameters(instanceType); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		a,
		"instanceTypeMatch",
		[]interface{}{instanceType},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Architecture) Is(arch Architecture) *bool {
	if err := a.validateIsParameters(arch); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		a,
		"is",
		[]interface{}{arch},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Architecture) IsIn(arches *[]Architecture) *bool {
	if err := a.validateIsInParameters(arches); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		a,
		"isIn",
		[]interface{}{arches},
		&returns,
	)

	return returns
}

