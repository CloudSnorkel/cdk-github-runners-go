package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for runner image status used by status.json.
type IRunnerProviderStatus interface {
	// Details about AMI used by this runner provider.
	Ami() IRunnerAmiStatus
	// CDK construct node path for this provider.
	ConstructPath() *string
	// Details about Docker image used by this runner provider.
	Image() IRunnerImageStatus
	// Labels associated with provider.
	Labels() *[]*string
	// Log group for runners.
	LogGroup() *string
	// Role attached to runners.
	RoleArn() *string
	// Security groups attached to runners.
	SecurityGroups() *[]*string
	// Runner provider type.
	Type() *string
	// VPC where runners will be launched.
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

func (j *jsiiProxy_IRunnerProviderStatus) ConstructPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constructPath",
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

func (j *jsiiProxy_IRunnerProviderStatus) LogGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroup",
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

