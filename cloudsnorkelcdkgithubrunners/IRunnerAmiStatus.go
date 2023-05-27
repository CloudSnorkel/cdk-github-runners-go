package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AMI status returned from runner providers to be displayed as output of status function.
// Experimental.
type IRunnerAmiStatus interface {
	// Log group name for the AMI builder where history of builds can be analyzed.
	// Experimental.
	AmiBuilderLogGroup() *string
	// Id of launch template pointing to the latest AMI built by the AMI builder.
	// Experimental.
	LaunchTemplate() *string
}

// The jsii proxy for IRunnerAmiStatus
type jsiiProxy_IRunnerAmiStatus struct {
	_ byte // padding
}

func (j *jsiiProxy_IRunnerAmiStatus) AmiBuilderLogGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiBuilderLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerAmiStatus) LaunchTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

