package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a base image that is used to start from in EC2 Image Builder image builds.
//
// This class is adapted from AWS CDK's BaseImage class to support both string and object inputs.
// Experimental.
type BaseImage interface {
	// The rendered base image to use.
	// Experimental.
	Image() *string
}

// The jsii proxy struct for BaseImage
type jsiiProxy_BaseImage struct {
	_ byte // padding
}

func (j *jsiiProxy_BaseImage) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}


// Experimental.
func NewBaseImage(image *string) BaseImage {
	_init_.Initialize()

	if err := validateNewBaseImageParameters(image); err != nil {
		panic(err)
	}
	j := jsiiProxy_BaseImage{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.BaseImage",
		[]interface{}{image},
		&j,
	)

	return &j
}

// Experimental.
func NewBaseImage_Override(b BaseImage, image *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.BaseImage",
		[]interface{}{image},
		b,
	)
}

// The AMI ID to use as a base image in an image recipe.
// Experimental.
func BaseImage_FromAmiId(amiId *string) BaseImage {
	_init_.Initialize()

	if err := validateBaseImage_FromAmiIdParameters(amiId); err != nil {
		panic(err)
	}
	var returns BaseImage

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.BaseImage",
		"fromAmiId",
		[]interface{}{amiId},
		&returns,
	)

	return returns
}

// A base AMI with NVIDIA drivers pre-installed for GPU workloads.
//
// Uses AWS Deep Learning AMIs for Linux (Ubuntu, Amazon Linux 2, Amazon Linux 2023).
// For Windows, subscribe to NVIDIA RTX Virtual Workstation in AWS Marketplace, then use
// {@link fromMarketplaceProductId} with the product ID.
// Experimental.
func BaseImage_FromGpuBase(os Os, architecture Architecture) BaseImage {
	_init_.Initialize()

	if err := validateBaseImage_FromGpuBaseParameters(os, architecture); err != nil {
		panic(err)
	}
	var returns BaseImage

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.BaseImage",
		"fromGpuBase",
		[]interface{}{os, architecture},
		&returns,
	)

	return returns
}

// An AWS-provided EC2 Image Builder image to use as a base image in an image recipe.
//
// This constructs an Image Builder ARN for AWS-provided images like `ubuntu-server-22-lts-x86/x.x.x`.
// Experimental.
func BaseImage_FromImageBuilder(scope constructs.Construct, resourceName *string, version *string) BaseImage {
	_init_.Initialize()

	if err := validateBaseImage_FromImageBuilderParameters(scope, resourceName); err != nil {
		panic(err)
	}
	var returns BaseImage

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.BaseImage",
		"fromImageBuilder",
		[]interface{}{scope, resourceName, version},
		&returns,
	)

	return returns
}

// The marketplace product ID for an AMI product to use as the base image in an image recipe.
// Experimental.
func BaseImage_FromMarketplaceProductId(productId *string) BaseImage {
	_init_.Initialize()

	if err := validateBaseImage_FromMarketplaceProductIdParameters(productId); err != nil {
		panic(err)
	}
	var returns BaseImage

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.BaseImage",
		"fromMarketplaceProductId",
		[]interface{}{productId},
		&returns,
	)

	return returns
}

// The SSM parameter to use as the base image in an image recipe.
// Experimental.
func BaseImage_FromSsmParameter(parameter awsssm.IParameter) BaseImage {
	_init_.Initialize()

	if err := validateBaseImage_FromSsmParameterParameters(parameter); err != nil {
		panic(err)
	}
	var returns BaseImage

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.BaseImage",
		"fromSsmParameter",
		[]interface{}{parameter},
		&returns,
	)

	return returns
}

// The parameter name for the SSM parameter to use as the base image in an image recipe.
// Experimental.
func BaseImage_FromSsmParameterName(parameterName *string) BaseImage {
	_init_.Initialize()

	if err := validateBaseImage_FromSsmParameterNameParameters(parameterName); err != nil {
		panic(err)
	}
	var returns BaseImage

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.BaseImage",
		"fromSsmParameterName",
		[]interface{}{parameterName},
		&returns,
	)

	return returns
}

// The direct string value of the base image to use in an image recipe.
//
// This can be an EC2 Image Builder image ARN,
// an SSM parameter, an AWS Marketplace product ID, or an AMI ID.
// Experimental.
func BaseImage_FromString(baseImageString *string) BaseImage {
	_init_.Initialize()

	if err := validateBaseImage_FromStringParameters(baseImageString); err != nil {
		panic(err)
	}
	var returns BaseImage

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.BaseImage",
		"fromString",
		[]interface{}{baseImageString},
		&returns,
	)

	return returns
}

