package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
)

// Represents a base container image that is used to start from in EC2 Image Builder container builds.
//
// This class is adapted from AWS CDK's BaseContainerImage class to support both string and object inputs.
// Experimental.
type BaseContainerImage interface {
	// The ECR repository if this image was created from an ECR repository.
	//
	// This allows automatic permission granting for CodeBuild.
	// Experimental.
	EcrRepository() awsecr.IRepository
	// The rendered base image to use.
	// Experimental.
	Image() *string
}

// The jsii proxy struct for BaseContainerImage
type jsiiProxy_BaseContainerImage struct {
	_ byte // padding
}

func (j *jsiiProxy_BaseContainerImage) EcrRepository() awsecr.IRepository {
	var returns awsecr.IRepository
	_jsii_.Get(
		j,
		"ecrRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseContainerImage) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}


// Experimental.
func NewBaseContainerImage(image *string, ecrRepository awsecr.IRepository) BaseContainerImage {
	_init_.Initialize()

	if err := validateNewBaseContainerImageParameters(image); err != nil {
		panic(err)
	}
	j := jsiiProxy_BaseContainerImage{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.BaseContainerImage",
		[]interface{}{image, ecrRepository},
		&j,
	)

	return &j
}

// Experimental.
func NewBaseContainerImage_Override(b BaseContainerImage, image *string, ecrRepository awsecr.IRepository) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.BaseContainerImage",
		[]interface{}{image, ecrRepository},
		b,
	)
}

// The DockerHub image to use as the base image in a container recipe.
// Experimental.
func BaseContainerImage_FromDockerHub(repository *string, tag *string) BaseContainerImage {
	_init_.Initialize()

	if err := validateBaseContainerImage_FromDockerHubParameters(repository, tag); err != nil {
		panic(err)
	}
	var returns BaseContainerImage

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.BaseContainerImage",
		"fromDockerHub",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

// The ECR container image to use as the base image in a container recipe.
// Experimental.
func BaseContainerImage_FromEcr(repository awsecr.IRepository, tag *string) BaseContainerImage {
	_init_.Initialize()

	if err := validateBaseContainerImage_FromEcrParameters(repository, tag); err != nil {
		panic(err)
	}
	var returns BaseContainerImage

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.BaseContainerImage",
		"fromEcr",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

// The ECR public container image to use as the base image in a container recipe.
// Experimental.
func BaseContainerImage_FromEcrPublic(registryAlias *string, repositoryName *string, tag *string) BaseContainerImage {
	_init_.Initialize()

	if err := validateBaseContainerImage_FromEcrPublicParameters(registryAlias, repositoryName, tag); err != nil {
		panic(err)
	}
	var returns BaseContainerImage

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.BaseContainerImage",
		"fromEcrPublic",
		[]interface{}{registryAlias, repositoryName, tag},
		&returns,
	)

	return returns
}

// The string value of the base image to use in a container recipe.
//
// This can be an EC2 Image Builder image ARN,
// an ECR or ECR public image, or a container URI sourced from a third-party container registry such as DockerHub.
// Experimental.
func BaseContainerImage_FromString(baseContainerImageString *string) BaseContainerImage {
	_init_.Initialize()

	if err := validateBaseContainerImage_FromStringParameters(baseContainerImageString); err != nil {
		panic(err)
	}
	var returns BaseContainerImage

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.BaseContainerImage",
		"fromString",
		[]interface{}{baseContainerImageString},
		&returns,
	)

	return returns
}

