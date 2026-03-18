package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

// GitHub Runner image builder. Builds a Docker image or AMI with GitHub Runner and other requirements installed.
//
// Images can be customized before passed into the provider by adding or removing components to be installed.
//
// Images are rebuilt every week by default to ensure that the latest security patches are applied.
// Experimental.
type RunnerImageBuilder interface {
	constructs.Construct
	IConfigurableRunnerImageBuilder
	// Experimental.
	Components() *[]RunnerImageComponent
	// Experimental.
	SetComponents(val *[]RunnerImageComponent)
	// The network connections associated with this resource.
	// Experimental.
	Connections() awsec2.Connections
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Add a component to the image builder.
	//
	// The component will be added to the end of the list of components.
	// Experimental.
	AddComponent(component RunnerImageComponent)
	// Build and return an AMI with GitHub Runner installed in it.
	//
	// Anything that ends up with a launch template pointing to an AMI that runs GitHub self-hosted runners can be used. A simple implementation could even point to an existing AMI and nothing else.
	//
	// The AMI can be further updated over time manually or using a schedule as long as it is always written to the same launch template.
	// Experimental.
	BindAmi() *RunnerAmi
	// Build and return a Docker image with GitHub Runner installed in it.
	//
	// Anything that ends up with an ECR repository containing a Docker image that runs GitHub self-hosted runners can be used. A simple implementation could even point to an existing image and nothing else.
	//
	// It's important that the specified image tag be available at the time the repository is available. Providers usually assume the image is ready and will fail if it's not.
	//
	// The image can be further updated over time manually or using a schedule as long as it is always written to the same tag.
	// Experimental.
	BindDockerImage() *RunnerImage
	// Remove a component from the image builder.
	//
	// Removal is done by component name. Multiple components with the same name will all be removed.
	// Experimental.
	RemoveComponent(component RunnerImageComponent)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for RunnerImageBuilder
type jsiiProxy_RunnerImageBuilder struct {
	internal.Type__constructsConstruct
	jsiiProxy_IConfigurableRunnerImageBuilder
}

func (j *jsiiProxy_RunnerImageBuilder) Components() *[]RunnerImageComponent {
	var returns *[]RunnerImageComponent
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerImageBuilder) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerImageBuilder) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerImageBuilder) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewRunnerImageBuilder_Override(r RunnerImageBuilder, scope constructs.Construct, id *string, props *RunnerImageBuilderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.RunnerImageBuilder",
		[]interface{}{scope, id, props},
		r,
	)
}

func (j *jsiiProxy_RunnerImageBuilder)SetComponents(val *[]RunnerImageComponent) {
	if err := j.validateSetComponentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"components",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func RunnerImageBuilder_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRunnerImageBuilder_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.RunnerImageBuilder",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Create a new image builder based on the provided properties.
//
// The implementation will differ based on the OS, architecture, and requested builder type.
// Experimental.
func RunnerImageBuilder_New(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) IConfigurableRunnerImageBuilder {
	_init_.Initialize()

	if err := validateRunnerImageBuilder_NewParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns IConfigurableRunnerImageBuilder

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.RunnerImageBuilder",
		"new",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerImageBuilder) AddComponent(component RunnerImageComponent) {
	if err := r.validateAddComponentParameters(component); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addComponent",
		[]interface{}{component},
	)
}

func (r *jsiiProxy_RunnerImageBuilder) BindAmi() *RunnerAmi {
	var returns *RunnerAmi

	_jsii_.Invoke(
		r,
		"bindAmi",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerImageBuilder) BindDockerImage() *RunnerImage {
	var returns *RunnerImage

	_jsii_.Invoke(
		r,
		"bindDockerImage",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerImageBuilder) RemoveComponent(component RunnerImageComponent) {
	if err := r.validateRemoveComponentParameters(component); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"removeComponent",
		[]interface{}{component},
	)
}

func (r *jsiiProxy_RunnerImageBuilder) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerImageBuilder) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		r,
		"with",
		args,
		&returns,
	)

	return returns
}

