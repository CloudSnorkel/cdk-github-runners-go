package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
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
	// List of step functions errors that should be retried.
	// Deprecated: use {@link Ec2RunnerProvider }.
	RetryableErrors() *[]*string
	// Generate step function task(s) to start a new runner.
	//
	// Called by GithubRunners and shouldn't be called manually.
	// Deprecated: use {@link Ec2RunnerProvider }.
	GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable
	// An optional method that modifies the role of the state machine after all the tasks have been generated.
	//
	// This can be used to add additional policy
	// statements to the state machine role that are not automatically added by the task returned from {@link getStepFunctionTask}.
	// Deprecated: use {@link Ec2RunnerProvider }.
	GrantStateMachine(stateMachineRole awsiam.IGrantable)
	// Deprecated: use {@link Ec2RunnerProvider }.
	LabelsFromProperties(defaultLabel *string, propsLabel *string, propsLabels *[]*string) *[]*string
	// Return status of the runner provider to be used in the main status function.
	//
	// Also gives the status function any needed permissions to query the Docker image or AMI.
	// Deprecated: use {@link Ec2RunnerProvider }.
	Status(statusFunctionRole awsiam.IGrantable) IRunnerProviderStatus
	// Returns a string representation of this construct.
	// Deprecated: use {@link Ec2RunnerProvider }.
	ToString() *string
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

func (j *jsiiProxy_Ec2Runner) RetryableErrors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"retryableErrors",
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
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

func (e *jsiiProxy_Ec2Runner) GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable {
	if err := e.validateGetStepFunctionTaskParameters(parameters); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		e,
		"getStepFunctionTask",
		[]interface{}{parameters},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Runner) GrantStateMachine(stateMachineRole awsiam.IGrantable) {
	if err := e.validateGrantStateMachineParameters(stateMachineRole); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"grantStateMachine",
		[]interface{}{stateMachineRole},
	)
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

func (e *jsiiProxy_Ec2Runner) Status(statusFunctionRole awsiam.IGrantable) IRunnerProviderStatus {
	if err := e.validateStatusParameters(statusFunctionRole); err != nil {
		panic(err)
	}
	var returns IRunnerProviderStatus

	_jsii_.Invoke(
		e,
		"status",
		[]interface{}{statusFunctionRole},
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

