package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

// GitHub Actions runner provider using EC2 to execute jobs.
//
// This construct is not meant to be used by itself. It should be passed in the providers property for GitHubRunners.
// Experimental.
type Ec2RunnerProvider interface {
	constructs.Construct
	IRunnerProvider
	// The network connections associated with this resource.
	// Experimental.
	Connections() awsec2.Connections
	// Grant principal used to add permissions to the runner role.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
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
	GrantStateMachine(stateMachineRole awsiam.IGrantable)
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

// The jsii proxy struct for Ec2RunnerProvider
type jsiiProxy_Ec2RunnerProvider struct {
	internal.Type__constructsConstruct
	jsiiProxy_IRunnerProvider
}

func (j *jsiiProxy_Ec2RunnerProvider) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2RunnerProvider) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2RunnerProvider) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2RunnerProvider) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2RunnerProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2RunnerProvider) RetryableErrors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"retryableErrors",
		&returns,
	)
	return returns
}


// Experimental.
func NewEc2RunnerProvider(scope constructs.Construct, id *string, props *Ec2RunnerProviderProps) Ec2RunnerProvider {
	_init_.Initialize()

	if err := validateNewEc2RunnerProviderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2RunnerProvider{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.Ec2RunnerProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEc2RunnerProvider_Override(e Ec2RunnerProvider, scope constructs.Construct, id *string, props *Ec2RunnerProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.Ec2RunnerProvider",
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
// Experimental.
func Ec2RunnerProvider_ImageBuilder(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) IConfigurableRunnerImageBuilder {
	_init_.Initialize()

	if err := validateEc2RunnerProvider_ImageBuilderParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns IConfigurableRunnerImageBuilder

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.Ec2RunnerProvider",
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
func Ec2RunnerProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2RunnerProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.Ec2RunnerProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2RunnerProvider) GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable {
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

func (e *jsiiProxy_Ec2RunnerProvider) GrantStateMachine(stateMachineRole awsiam.IGrantable) {
	if err := e.validateGrantStateMachineParameters(stateMachineRole); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"grantStateMachine",
		[]interface{}{stateMachineRole},
	)
}

func (e *jsiiProxy_Ec2RunnerProvider) LabelsFromProperties(defaultLabel *string, propsLabel *string, propsLabels *[]*string) *[]*string {
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

func (e *jsiiProxy_Ec2RunnerProvider) Status(statusFunctionRole awsiam.IGrantable) IRunnerProviderStatus {
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

func (e *jsiiProxy_Ec2RunnerProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

