// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

// GitHub Actions runner provider using Fargate to execute jobs.
//
// Creates a task definition with a single container that gets started for each job.
//
// This construct is not meant to be used by itself. It should be passed in the providers property for GitHubRunners.
// Experimental.
type FargateRunner interface {
	constructs.Construct
	IRunnerProvider
	// Whether runner task will have a public IP.
	// Experimental.
	AssignPublicIp() *bool
	// Cluster hosting the task hosting the runner.
	// Experimental.
	Cluster() awsecs.Cluster
	// The network connections associated with this resource.
	// Experimental.
	Connections() awsec2.Connections
	// Container definition hosting the runner.
	// Experimental.
	Container() awsecs.ContainerDefinition
	// Grant principal used to add permissions to the runner role.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// Docker image loaded with GitHub Actions Runner and its prerequisites.
	//
	// The image is built by an image builder and is specific to Fargate tasks.
	// Experimental.
	Image() *RunnerImage
	// Labels associated with this provider.
	// Experimental.
	Labels() *[]*string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Use spot pricing for Fargate tasks.
	// Experimental.
	Spot() *bool
	// Subnets used for hosting the runner task.
	// Experimental.
	SubnetSelection() *awsec2.SubnetSelection
	// Fargate task hosting the runner.
	// Experimental.
	Task() awsecs.FargateTaskDefinition
	// VPC used for hosting the runner task.
	// Experimental.
	Vpc() awsec2.IVpc
	// Experimental.
	AddRetry(task awsstepfunctions.TaskStateBase, errors *[]*string)
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

// The jsii proxy struct for FargateRunner
type jsiiProxy_FargateRunner struct {
	internal.Type__constructsConstruct
	jsiiProxy_IRunnerProvider
}

func (j *jsiiProxy_FargateRunner) AssignPublicIp() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) Cluster() awsecs.Cluster {
	var returns awsecs.Cluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) Container() awsecs.ContainerDefinition {
	var returns awsecs.ContainerDefinition
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) Image() *RunnerImage {
	var returns *RunnerImage
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) Spot() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"spot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) SubnetSelection() *awsec2.SubnetSelection {
	var returns *awsec2.SubnetSelection
	_jsii_.Get(
		j,
		"subnetSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) Task() awsecs.FargateTaskDefinition {
	var returns awsecs.FargateTaskDefinition
	_jsii_.Get(
		j,
		"task",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Experimental.
func NewFargateRunner(scope constructs.Construct, id *string, props *FargateRunnerProps) FargateRunner {
	_init_.Initialize()

	if err := validateNewFargateRunnerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FargateRunner{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.FargateRunner",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewFargateRunner_Override(f FargateRunner, scope constructs.Construct, id *string, props *FargateRunnerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.FargateRunner",
		[]interface{}{scope, id, props},
		f,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func FargateRunner_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFargateRunner_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.FargateRunner",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FargateRunner_LINUX_ARM64_DOCKERFILE_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.FargateRunner",
		"LINUX_ARM64_DOCKERFILE_PATH",
		&returns,
	)
	return returns
}

func FargateRunner_LINUX_X64_DOCKERFILE_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.FargateRunner",
		"LINUX_X64_DOCKERFILE_PATH",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FargateRunner) AddRetry(task awsstepfunctions.TaskStateBase, errors *[]*string) {
	if err := f.validateAddRetryParameters(task, errors); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addRetry",
		[]interface{}{task, errors},
	)
}

func (f *jsiiProxy_FargateRunner) GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable {
	if err := f.validateGetStepFunctionTaskParameters(parameters); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		f,
		"getStepFunctionTask",
		[]interface{}{parameters},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateRunner) GrantStateMachine(_arg awsiam.IGrantable) {
	if err := f.validateGrantStateMachineParameters(_arg); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"grantStateMachine",
		[]interface{}{_arg},
	)
}

func (f *jsiiProxy_FargateRunner) LabelsFromProperties(defaultLabel *string, propsLabel *string, propsLabels *[]*string) *[]*string {
	if err := f.validateLabelsFromPropertiesParameters(defaultLabel); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"labelsFromProperties",
		[]interface{}{defaultLabel, propsLabel, propsLabels},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateRunner) Status(statusFunctionRole awsiam.IGrantable) IRunnerProviderStatus {
	if err := f.validateStatusParameters(statusFunctionRole); err != nil {
		panic(err)
	}
	var returns IRunnerProviderStatus

	_jsii_.Invoke(
		f,
		"status",
		[]interface{}{statusFunctionRole},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateRunner) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

