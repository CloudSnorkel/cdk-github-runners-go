// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"
)

// Defines desired GitHub Actions runner version.
// Experimental.
type RunnerVersion interface {
	// Experimental.
	Version() *string
	// Check if two versions are the same.
	// Experimental.
	Is(other RunnerVersion) *bool
}

// The jsii proxy struct for RunnerVersion
type jsiiProxy_RunnerVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_RunnerVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Experimental.
func NewRunnerVersion(version *string) RunnerVersion {
	_init_.Initialize()

	if err := validateNewRunnerVersionParameters(version); err != nil {
		panic(err)
	}
	j := jsiiProxy_RunnerVersion{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.RunnerVersion",
		[]interface{}{version},
		&j,
	)

	return &j
}

// Experimental.
func NewRunnerVersion_Override(r RunnerVersion, version *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.RunnerVersion",
		[]interface{}{version},
		r,
	)
}

// Use the latest version available at the time the runner provider image is built.
// Experimental.
func RunnerVersion_Latest() RunnerVersion {
	_init_.Initialize()

	var returns RunnerVersion

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.RunnerVersion",
		"latest",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Use a specific version.
// See: https://github.com/actions/runner/releases
//
// Experimental.
func RunnerVersion_Specific(version *string) RunnerVersion {
	_init_.Initialize()

	if err := validateRunnerVersion_SpecificParameters(version); err != nil {
		panic(err)
	}
	var returns RunnerVersion

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.RunnerVersion",
		"specific",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerVersion) Is(other RunnerVersion) *bool {
	if err := r.validateIsParameters(other); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		r,
		"is",
		[]interface{}{other},
		&returns,
	)

	return returns
}

