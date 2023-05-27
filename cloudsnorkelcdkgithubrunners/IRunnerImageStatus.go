package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Image status returned from runner providers to be displayed in status.json.
// Experimental.
type IRunnerImageStatus interface {
	// Log group name for the image builder where history of image builds can be analyzed.
	// Experimental.
	ImageBuilderLogGroup() *string
	// Image repository where image builder pushes runner images.
	// Experimental.
	ImageRepository() *string
	// Tag of image that should be used.
	// Experimental.
	ImageTag() *string
}

// The jsii proxy for IRunnerImageStatus
type jsiiProxy_IRunnerImageStatus struct {
	_ byte // padding
}

func (j *jsiiProxy_IRunnerImageStatus) ImageBuilderLogGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageBuilderLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerImageStatus) ImageRepository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerImageStatus) ImageTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTag",
		&returns,
	)
	return returns
}

