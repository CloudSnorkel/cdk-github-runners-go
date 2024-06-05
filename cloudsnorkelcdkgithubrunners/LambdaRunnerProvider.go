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
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

// GitHub Actions runner provider using Lambda to execute jobs.
//
// Creates a Docker-based function that gets executed for each job.
//
// This construct is not meant to be used by itself. It should be passed in the providers property for GitHubRunners.
// Experimental.
type LambdaRunnerProvider interface {
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
	// Log group where provided runners will save their logs.
	//
	// Note that this is not the job log, but the runner itself. It will not contain output from the GitHub Action but only metadata on its execution.
	// Experimental.
	LogGroup() awslogs.ILogGroup
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// List of step functions errors that should be retried.
	// Experimental.
	RetryableErrors() *[]*string
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

// The jsii proxy struct for LambdaRunnerProvider
type jsiiProxy_LambdaRunnerProvider struct {
	internal.Type__constructsConstruct
	jsiiProxy_IRunnerProvider
}

func (j *jsiiProxy_LambdaRunnerProvider) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRunnerProvider) Function() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"function",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRunnerProvider) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRunnerProvider) Image() *RunnerImage {
	var returns *RunnerImage
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRunnerProvider) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRunnerProvider) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRunnerProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRunnerProvider) RetryableErrors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"retryableErrors",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaRunnerProvider(scope constructs.Construct, id *string, props *LambdaRunnerProviderProps) LambdaRunnerProvider {
	_init_.Initialize()

	if err := validateNewLambdaRunnerProviderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaRunnerProvider{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.LambdaRunnerProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaRunnerProvider_Override(l LambdaRunnerProvider, scope constructs.Construct, id *string, props *LambdaRunnerProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.LambdaRunnerProvider",
		[]interface{}{scope, id, props},
		l,
	)
}

// Create new image builder that builds Lambda specific runner images.
//
// You can customize the OS, architecture, VPC, subnet, security groups, etc. by passing in props.
//
// You can add components to the image builder by calling `imageBuilder.addComponent()`.
//
// The default OS is Amazon Linux 2023 running on x64 architecture.
//
// Included components:
//  * `RunnerImageComponent.requiredPackages()`
//  * `RunnerImageComponent.runnerUser()`
//  * `RunnerImageComponent.git()`
//  * `RunnerImageComponent.githubCli()`
//  * `RunnerImageComponent.awsCli()`
//  * `RunnerImageComponent.githubRunner()`
//  * `RunnerImageComponent.lambdaEntrypoint()`
// Experimental.
func LambdaRunnerProvider_ImageBuilder(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) IConfigurableRunnerImageBuilder {
	_init_.Initialize()

	if err := validateLambdaRunnerProvider_ImageBuilderParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns IConfigurableRunnerImageBuilder

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.LambdaRunnerProvider",
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
func LambdaRunnerProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaRunnerProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.LambdaRunnerProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LambdaRunnerProvider_LINUX_ARM64_DOCKERFILE_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.LambdaRunnerProvider",
		"LINUX_ARM64_DOCKERFILE_PATH",
		&returns,
	)
	return returns
}

func LambdaRunnerProvider_LINUX_X64_DOCKERFILE_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.LambdaRunnerProvider",
		"LINUX_X64_DOCKERFILE_PATH",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LambdaRunnerProvider) GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable {
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

func (l *jsiiProxy_LambdaRunnerProvider) GrantStateMachine(_arg awsiam.IGrantable) {
	if err := l.validateGrantStateMachineParameters(_arg); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"grantStateMachine",
		[]interface{}{_arg},
	)
}

func (l *jsiiProxy_LambdaRunnerProvider) LabelsFromProperties(defaultLabel *string, propsLabel *string, propsLabels *[]*string) *[]*string {
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

func (l *jsiiProxy_LambdaRunnerProvider) Status(statusFunctionRole awsiam.IGrantable) IRunnerProviderStatus {
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

func (l *jsiiProxy_LambdaRunnerProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

