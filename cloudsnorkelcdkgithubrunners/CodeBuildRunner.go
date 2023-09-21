package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
)

// Deprecated: use {@link CodeBuildRunnerProvider }.
type CodeBuildRunner interface {
	CodeBuildRunnerProvider
	// The network connections associated with this resource.
	// Deprecated: use {@link CodeBuildRunnerProvider }.
	Connections() awsec2.Connections
	// Grant principal used to add permissions to the runner role.
	// Deprecated: use {@link CodeBuildRunnerProvider }.
	GrantPrincipal() awsiam.IPrincipal
	// Docker image loaded with GitHub Actions Runner and its prerequisites.
	//
	// The image is built by an image builder and is specific to CodeBuild.
	// Deprecated: use {@link CodeBuildRunnerProvider }.
	Image() *RunnerImage
	// Labels associated with this provider.
	// Deprecated: use {@link CodeBuildRunnerProvider }.
	Labels() *[]*string
	// Log group where provided runners will save their logs.
	//
	// Note that this is not the job log, but the runner itself. It will not contain output from the GitHub Action but only metadata on its execution.
	// Deprecated: use {@link CodeBuildRunnerProvider }.
	LogGroup() awslogs.ILogGroup
	// The tree node.
	// Deprecated: use {@link CodeBuildRunnerProvider }.
	Node() constructs.Node
	// CodeBuild project hosting the runner.
	// Deprecated: use {@link CodeBuildRunnerProvider }.
	Project() awscodebuild.Project
	// List of step functions errors that should be retried.
	// Deprecated: use {@link CodeBuildRunnerProvider }.
	RetryableErrors() *[]*string
	// Generate step function task(s) to start a new runner.
	//
	// Called by GithubRunners and shouldn't be called manually.
	// Deprecated: use {@link CodeBuildRunnerProvider }.
	GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable
	// An optional method that modifies the role of the state machine after all the tasks have been generated.
	//
	// This can be used to add additional policy
	// statements to the state machine role that are not automatically added by the task returned from {@link getStepFunctionTask}.
	// Deprecated: use {@link CodeBuildRunnerProvider }.
	GrantStateMachine(_arg awsiam.IGrantable)
	// Deprecated: use {@link CodeBuildRunnerProvider }.
	LabelsFromProperties(defaultLabel *string, propsLabel *string, propsLabels *[]*string) *[]*string
	// Return status of the runner provider to be used in the main status function.
	//
	// Also gives the status function any needed permissions to query the Docker image or AMI.
	// Deprecated: use {@link CodeBuildRunnerProvider }.
	Status(statusFunctionRole awsiam.IGrantable) IRunnerProviderStatus
	// Returns a string representation of this construct.
	// Deprecated: use {@link CodeBuildRunnerProvider }.
	ToString() *string
}

// The jsii proxy struct for CodeBuildRunner
type jsiiProxy_CodeBuildRunner struct {
	jsiiProxy_CodeBuildRunnerProvider
}

func (j *jsiiProxy_CodeBuildRunner) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildRunner) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildRunner) Image() *RunnerImage {
	var returns *RunnerImage
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildRunner) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildRunner) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildRunner) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildRunner) Project() awscodebuild.Project {
	var returns awscodebuild.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildRunner) RetryableErrors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"retryableErrors",
		&returns,
	)
	return returns
}


// Deprecated: use {@link CodeBuildRunnerProvider }.
func NewCodeBuildRunner(scope constructs.Construct, id *string, props *CodeBuildRunnerProviderProps) CodeBuildRunner {
	_init_.Initialize()

	if err := validateNewCodeBuildRunnerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodeBuildRunner{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.CodeBuildRunner",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Deprecated: use {@link CodeBuildRunnerProvider }.
func NewCodeBuildRunner_Override(c CodeBuildRunner, scope constructs.Construct, id *string, props *CodeBuildRunnerProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.CodeBuildRunner",
		[]interface{}{scope, id, props},
		c,
	)
}

// Create new image builder that builds CodeBuild specific runner images using Ubuntu.
//
// Included components:
//  * `RunnerImageComponent.requiredPackages()`
//  * `RunnerImageComponent.runnerUser()`
//  * `RunnerImageComponent.git()`
//  * `RunnerImageComponent.githubCli()`
//  * `RunnerImageComponent.awsCli()`
//  * `RunnerImageComponent.docker()`
//  * `RunnerImageComponent.githubRunner()`
// Deprecated: use {@link CodeBuildRunnerProvider }.
func CodeBuildRunner_ImageBuilder(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) RunnerImageBuilder {
	_init_.Initialize()

	if err := validateCodeBuildRunner_ImageBuilderParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns RunnerImageBuilder

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.CodeBuildRunner",
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
func CodeBuildRunner_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCodeBuildRunner_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.CodeBuildRunner",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CodeBuildRunner_LINUX_ARM64_DOCKERFILE_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.CodeBuildRunner",
		"LINUX_ARM64_DOCKERFILE_PATH",
		&returns,
	)
	return returns
}

func CodeBuildRunner_LINUX_X64_DOCKERFILE_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.CodeBuildRunner",
		"LINUX_X64_DOCKERFILE_PATH",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CodeBuildRunner) GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable {
	if err := c.validateGetStepFunctionTaskParameters(parameters); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		c,
		"getStepFunctionTask",
		[]interface{}{parameters},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildRunner) GrantStateMachine(_arg awsiam.IGrantable) {
	if err := c.validateGrantStateMachineParameters(_arg); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"grantStateMachine",
		[]interface{}{_arg},
	)
}

func (c *jsiiProxy_CodeBuildRunner) LabelsFromProperties(defaultLabel *string, propsLabel *string, propsLabels *[]*string) *[]*string {
	if err := c.validateLabelsFromPropertiesParameters(defaultLabel); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"labelsFromProperties",
		[]interface{}{defaultLabel, propsLabel, propsLabels},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildRunner) Status(statusFunctionRole awsiam.IGrantable) IRunnerProviderStatus {
	if err := c.validateStatusParameters(statusFunctionRole); err != nil {
		panic(err)
	}
	var returns IRunnerProviderStatus

	_jsii_.Invoke(
		c,
		"status",
		[]interface{}{statusFunctionRole},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildRunner) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

