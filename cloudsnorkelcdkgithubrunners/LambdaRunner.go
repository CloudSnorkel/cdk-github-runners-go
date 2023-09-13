package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
)

// Deprecated: use {@link LambdaRunnerProvider }.
type LambdaRunner interface {
	LambdaRunnerProvider
	// The network connections associated with this resource.
	// Deprecated: use {@link LambdaRunnerProvider }.
	Connections() awsec2.Connections
	// The function hosting the GitHub runner.
	// Deprecated: use {@link LambdaRunnerProvider }.
	Function() awslambda.Function
	// Grant principal used to add permissions to the runner role.
	// Deprecated: use {@link LambdaRunnerProvider }.
	GrantPrincipal() awsiam.IPrincipal
	// Docker image loaded with GitHub Actions Runner and its prerequisites.
	//
	// The image is built by an image builder and is specific to Lambda.
	// Deprecated: use {@link LambdaRunnerProvider }.
	Image() *RunnerImage
	// Labels associated with this provider.
	// Deprecated: use {@link LambdaRunnerProvider }.
	Labels() *[]*string
	// Log group where provided runners will save their logs.
	//
	// Note that this is not the job log, but the runner itself. It will not contain output from the GitHub Action but only metadata on its execution.
	// Deprecated: use {@link LambdaRunnerProvider }.
	LogGroup() awslogs.ILogGroup
	// The tree node.
	// Deprecated: use {@link LambdaRunnerProvider }.
	Node() constructs.Node
	// List of step functions errors that should be retried.
	// Deprecated: use {@link LambdaRunnerProvider }.
	RetryableErrors() *[]*string
	// Generate step function task(s) to start a new runner.
	//
	// Called by GithubRunners and shouldn't be called manually.
	// Deprecated: use {@link LambdaRunnerProvider }.
	GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable
	// An optional method that modifies the role of the state machine after all the tasks have been generated.
	//
	// This can be used to add additional policy
	// statements to the state machine role that are not automatically added by the task returned from {@link getStepFunctionTask}.
	// Deprecated: use {@link LambdaRunnerProvider }.
	GrantStateMachine(_arg awsiam.IGrantable)
	// Deprecated: use {@link LambdaRunnerProvider }.
	LabelsFromProperties(defaultLabel *string, propsLabel *string, propsLabels *[]*string) *[]*string
	// Return status of the runner provider to be used in the main status function.
	//
	// Also gives the status function any needed permissions to query the Docker image or AMI.
	// Deprecated: use {@link LambdaRunnerProvider }.
	Status(statusFunctionRole awsiam.IGrantable) IRunnerProviderStatus
	// Returns a string representation of this construct.
	// Deprecated: use {@link LambdaRunnerProvider }.
	ToString() *string
}

// The jsii proxy struct for LambdaRunner
type jsiiProxy_LambdaRunner struct {
	jsiiProxy_LambdaRunnerProvider
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

func (j *jsiiProxy_LambdaRunner) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
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

func (j *jsiiProxy_LambdaRunner) RetryableErrors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"retryableErrors",
		&returns,
	)
	return returns
}


// Deprecated: use {@link LambdaRunnerProvider }.
func NewLambdaRunner(scope constructs.Construct, id *string, props *LambdaRunnerProviderProps) LambdaRunner {
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

// Deprecated: use {@link LambdaRunnerProvider }.
func NewLambdaRunner_Override(l LambdaRunner, scope constructs.Construct, id *string, props *LambdaRunnerProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.LambdaRunner",
		[]interface{}{scope, id, props},
		l,
	)
}

// Create new image builder that builds Lambda specific runner images using Amazon Linux 2.
//
// Included components:
//  * `RunnerImageComponent.requiredPackages()`
//  * `RunnerImageComponent.runnerUser()`
//  * `RunnerImageComponent.git()`
//  * `RunnerImageComponent.githubCli()`
//  * `RunnerImageComponent.awsCli()`
//  * `RunnerImageComponent.githubRunner()`
//  * `RunnerImageComponent.lambdaEntrypoint()`
//
//  Base Docker image: `public.ecr.aws/lambda/nodejs:16-x86_64` or `public.ecr.aws/lambda/nodejs:16-arm64`
// Deprecated: use {@link LambdaRunnerProvider }.
func LambdaRunner_ImageBuilder(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) RunnerImageBuilder {
	_init_.Initialize()

	if err := validateLambdaRunner_ImageBuilderParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns RunnerImageBuilder

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.LambdaRunner",
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

