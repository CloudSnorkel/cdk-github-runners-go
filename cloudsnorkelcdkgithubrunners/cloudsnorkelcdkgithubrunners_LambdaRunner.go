// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

// GitHub Actions runner provider using Lambda to execute jobs.
//
// Creates a Docker-based function that gets executed for each job.
//
// This construct is not meant to be used by itself. It should be passed in the providers property for GitHubRunners.
// Experimental.
type LambdaRunner interface {
	constructs.Construct
	IRunnerProvider
	// The network connections associated with this resource.
	// Experimental.
	Connections() awsec2.Connections
	// The function hosting the GitHub runner.
	// Experimental.
	Function() awslambda.Function
	// Grant principal used to add permissions to the runner role.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// Docker image loaded with GitHub Actions Runner and its prerequisites.
	//
	// The image is built by an image builder and is specific to Lambda.
	// Experimental.
	Image() *RunnerImage
	// Labels associated with this provider.
	// Experimental.
	Labels() *[]*string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Generate step function task(s) to start a new runner.
	//
	// Called by GithubRunners and shouldn't be called manually.
	// Experimental.
	GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable
	// An optional method that modifies the role of the state machine after all the tasks have been generated.
	//
	// This can be used to add additional policy
	// statements to the state machine role that are not automatically added by the task returned from {@link getStepFunctionTask}.
	// Experimental.
	GrantStateMachine(_arg awsiam.IGrantable)
	// Experimental.
	LabelsFromProperties(defaultLabel *string, propsLabel *string, propsLabels *[]*string) *[]*string
	// Return status of the runner provider to be used in the main status function.
	//
	// Also gives the status function any needed permissions to query the Docker image or AMI.
	// Experimental.
	Status(statusFunctionRole awsiam.IGrantable) IRunnerProviderStatus
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LambdaRunner
type jsiiProxy_LambdaRunner struct {
	internal.Type__constructsConstruct
	jsiiProxy_IRunnerProvider
}

func (j *jsiiProxy_LambdaRunner) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRunner) Function() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"function",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRunner) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRunner) Image() *RunnerImage {
	var returns *RunnerImage
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRunner) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRunner) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaRunner(scope constructs.Construct, id *string, props *LambdaRunnerProps) LambdaRunner {
	_init_.Initialize()

	if err := validateNewLambdaRunnerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaRunner{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.LambdaRunner",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaRunner_Override(l LambdaRunner, scope constructs.Construct, id *string, props *LambdaRunnerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.LambdaRunner",
		[]interface{}{scope, id, props},
		l,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func LambdaRunner_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaRunner_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.LambdaRunner",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LambdaRunner_LINUX_ARM64_DOCKERFILE_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.LambdaRunner",
		"LINUX_ARM64_DOCKERFILE_PATH",
		&returns,
	)
	return returns
}

func LambdaRunner_LINUX_X64_DOCKERFILE_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.LambdaRunner",
		"LINUX_X64_DOCKERFILE_PATH",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LambdaRunner) GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable {
	if err := l.validateGetStepFunctionTaskParameters(parameters); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		l,
		"getStepFunctionTask",
		[]interface{}{parameters},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRunner) GrantStateMachine(_arg awsiam.IGrantable) {
	if err := l.validateGrantStateMachineParameters(_arg); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"grantStateMachine",
		[]interface{}{_arg},
	)
}

func (l *jsiiProxy_LambdaRunner) LabelsFromProperties(defaultLabel *string, propsLabel *string, propsLabels *[]*string) *[]*string {
	if err := l.validateLabelsFromPropertiesParameters(defaultLabel); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"labelsFromProperties",
		[]interface{}{defaultLabel, propsLabel, propsLabels},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRunner) Status(statusFunctionRole awsiam.IGrantable) IRunnerProviderStatus {
	if err := l.validateStatusParameters(statusFunctionRole); err != nil {
		panic(err)
	}
	var returns IRunnerProviderStatus

	_jsii_.Invoke(
		l,
		"status",
		[]interface{}{statusFunctionRole},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRunner) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

