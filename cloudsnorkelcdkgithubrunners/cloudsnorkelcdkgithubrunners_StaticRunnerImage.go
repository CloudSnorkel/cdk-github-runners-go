// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/constructs-go/constructs/v10"
)

// Helper class with methods to use static images that are built outside the context of this project.
// Experimental.
type StaticRunnerImage interface {
}

// The jsii proxy struct for StaticRunnerImage
type jsiiProxy_StaticRunnerImage struct {
	_ byte // padding
}

// Experimental.
func NewStaticRunnerImage() StaticRunnerImage {
	_init_.Initialize()

	j := jsiiProxy_StaticRunnerImage{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.StaticRunnerImage",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewStaticRunnerImage_Override(s StaticRunnerImage) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.StaticRunnerImage",
		nil, // no parameters
		s,
	)
}

// Create a builder from an existing Docker Hub image.
//
// The image must already have GitHub Actions runner installed. You are responsible to update it and remove it when done.
//
// We create a CodeBuild image builder behind the scenes to copy the image over to ECR. This helps avoid Docker Hub rate limits and prevent failures.
// Experimental.
func StaticRunnerImage_FromDockerHub(scope constructs.Construct, id *string, image *string, architecture Architecture, os Os) IImageBuilder {
	_init_.Initialize()

	var returns IImageBuilder

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.StaticRunnerImage",
		"fromDockerHub",
		[]interface{}{scope, id, image, architecture, os},
		&returns,
	)

	return returns
}

// Create a builder (that doesn't actually build anything) from an existing image in an existing repository.
//
// The image must already have GitHub Actions runner installed. You are responsible to update it and remove it when done.
// Experimental.
func StaticRunnerImage_FromEcrRepository(repository awsecr.IRepository, tag *string, architecture Architecture, os Os) IImageBuilder {
	_init_.Initialize()

	var returns IImageBuilder

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.StaticRunnerImage",
		"fromEcrRepository",
		[]interface{}{repository, tag, architecture, os},
		&returns,
	)

	return returns
}

