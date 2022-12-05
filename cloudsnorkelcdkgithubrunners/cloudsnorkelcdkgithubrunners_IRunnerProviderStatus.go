// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for runner image status used by status.json.
// Experimental.
type IRunnerProviderStatus interface {
	// Details about AMI used by this runner provider.
	// Experimental.
	Ami() IRunnerAmiStatus
	// Details about Docker image used by this runner provider.
	// Experimental.
	Image() IRunnerImageStatus
	// Labels associated with provider.
	// Experimental.
	Labels() *[]*string
	// Role attached to runners.
	// Experimental.
	RoleArn() *string
	// Security groups attached to runners.
	// Experimental.
	SecurityGroups() *[]*string
	// Runner provider type.
	// Experimental.
	Type() *string
	// VPC where runners will be launched.
	// Experimental.
	VpcArn() *string
}

// The jsii proxy for IRunnerProviderStatus
type jsiiProxy_IRunnerProviderStatus struct {
	_ byte // padding
}

func (j *jsiiProxy_IRunnerProviderStatus) Ami() IRunnerAmiStatus {
	var returns IRunnerAmiStatus
	_jsii_.Get(
		j,
		"ami",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerProviderStatus) Image() IRunnerImageStatus {
	var returns IRunnerImageStatus
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerProviderStatus) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerProviderStatus) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerProviderStatus) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerProviderStatus) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerProviderStatus) VpcArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcArn",
		&returns,
	)
	return returns
}

