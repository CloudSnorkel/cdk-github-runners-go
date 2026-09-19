package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Deprecated: use {@link Ec2RunnerProvider }.
type Ec2Runner interface {
	Ec2RunnerProvider
	// The network connections associated with this resource.
	// Deprecated: use {@link Ec2RunnerProvider }.
	Connections() awsec2.Connections
	// Grant principal used to add permissions to the runner role.
	// Deprecated: use {@link Ec2RunnerProvider }.
	GrantPrincipal() awsiam.IPrincipal
	// Labels associated with this provider.
	// Deprecated: use {@link Ec2RunnerProvider }.
	Labels() *[]*string
	// Log group where provided runners will save their logs.
	//
	// Note that this is not the job log, but the runner itself. It will not contain output from the GitHub Action but only metadata on its execution.
	// Deprecated: use {@link Ec2RunnerProvider }.
	LogGroup() awslogs.ILogGroup
	// The tree node.
	// Deprecated: use {@link Ec2RunnerProvider }.
	Node() constructs.Node
	// Deprecated: use {@link Ec2RunnerProvider }.
	LabelsFromProperties(defaultLabel *string, propsLabel *string, propsLabels *[]*string) *[]*string
	// Returns a string representation of this construct.
	// Deprecated: use {@link Ec2RunnerProvider }.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Deprecated: use {@link Ec2RunnerProvider }.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for Ec2Runner
type jsiiProxy_Ec2Runner struct {
	jsiiProxy_Ec2RunnerProvider
}

func (j *jsiiProxy_Ec2Runner) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Runner) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Runner) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Runner) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Runner) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Deprecated: use {@link Ec2RunnerProvider }.
func NewEc2Runner(scope constructs.Construct, id *string, props *Ec2RunnerProviderProps) Ec2Runner {
	_init_.Initialize()

	if err := validateNewEc2RunnerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2Runner{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.Ec2Runner",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Deprecated: use {@link Ec2RunnerProvider }.
func NewEc2Runner_Override(e Ec2Runner, scope constructs.Construct, id *string, props *Ec2RunnerProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.Ec2Runner",
		[]interface{}{scope, id, props},
		e,
	)
}

// Create new image builder that builds EC2 specific runner images.
//
// You can customize the OS, architecture, VPC, subnet, security groups, etc. by passing in props.
//
// You can add components to the image builder by calling `imageBuilder.addComponent()`.
//
// The default OS is Ubuntu running on x64 architecture.
//
// Included components:
//  * `RunnerImageComponent.requiredPackages()`
//  * `RunnerImageComponent.cloudWatchAgent()`
//  * `RunnerImageComponent.runnerUser()`
//  * `RunnerImageComponent.git()`
//  * `RunnerImageComponent.githubCli()`
//  * `RunnerImageComponent.awsCli()`
//  * `RunnerImageComponent.docker()`
//  * `RunnerImageComponent.githubRunner()`
// Deprecated: use {@link Ec2RunnerProvider }.
func Ec2Runner_ImageBuilder(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) IConfigurableRunnerImageBuilder {
	_init_.Initialize()

	if err := validateEc2Runner_ImageBuilderParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns IConfigurableRunnerImageBuilder

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.Ec2Runner",
		"imageBuilder",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
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
// Deprecated: use {@link Ec2RunnerProvider }.
func Ec2Runner_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2Runner_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.Ec2Runner",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Runner) LabelsFromProperties(defaultLabel *string, propsLabel *string, propsLabels *[]*string) *[]*string {
	if err := e.validateLabelsFromPropertiesParameters(defaultLabel); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"labelsFromProperties",
		[]interface{}{defaultLabel, propsLabel, propsLabels},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Runner) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Runner) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		e,
		"with",
		args,
		&returns,
	)

	return returns
}

