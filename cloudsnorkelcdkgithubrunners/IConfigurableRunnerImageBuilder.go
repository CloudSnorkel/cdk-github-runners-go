package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

// Interface for constructs that build an image that can be used in {@link IRunnerProvider }.
//
// The image can be configured by adding or removing components. The image builder can be configured by adding grants or allowing connections.
//
// An image can be a Docker image or AMI.
// Experimental.
type IConfigurableRunnerImageBuilder interface {
	awsec2.IConnectable
	awsiam.IGrantable
	IRunnerImageBuilder
	// Add a component to the image builder.
	//
	// The component will be added to the end of the list of components.
	// Experimental.
	AddComponent(component RunnerImageComponent)
	// Remove a component from the image builder.
	//
	// Removal is done by component name. Multiple components with the same name will all be removed.
	// Experimental.
	RemoveComponent(component RunnerImageComponent)
}

// The jsii proxy for IConfigurableRunnerImageBuilder
type jsiiProxy_IConfigurableRunnerImageBuilder struct {
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
	jsiiProxy_IRunnerImageBuilder
}

func (i *jsiiProxy_IConfigurableRunnerImageBuilder) AddComponent(component RunnerImageComponent) {
	if err := i.validateAddComponentParameters(component); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addComponent",
		[]interface{}{component},
	)
}

func (i *jsiiProxy_IConfigurableRunnerImageBuilder) RemoveComponent(component RunnerImageComponent) {
	if err := i.validateRemoveComponentParameters(component); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"removeComponent",
		[]interface{}{component},
	)
}

func (i *jsiiProxy_IConfigurableRunnerImageBuilder) BindAmi() *RunnerAmi {
	var returns *RunnerAmi

	_jsii_.Invoke(
		i,
		"bindAmi",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IConfigurableRunnerImageBuilder) BindDockerImage() *RunnerImage {
	var returns *RunnerImage

	_jsii_.Invoke(
		i,
		"bindDockerImage",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IConfigurableRunnerImageBuilder) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfigurableRunnerImageBuilder) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

