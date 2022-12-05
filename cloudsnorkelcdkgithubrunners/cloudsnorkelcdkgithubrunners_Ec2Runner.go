// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

// GitHub Actions runner provider using EC2 to execute jobs.
//
// This construct is not meant to be used by itself. It should be passed in the providers property for GitHubRunners.
// Experimental.
type Ec2Runner interface {
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

// The jsii proxy struct for Ec2Runner
type jsiiProxy_Ec2Runner struct {
	internal.Type__constructsConstruct
	jsiiProxy_IRunnerProvider
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

func (j *jsiiProxy_Ec2Runner) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewEc2Runner(scope constructs.Construct, id *string, props *Ec2RunnerProps) Ec2Runner {
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

// Experimental.
func NewEc2Runner_Override(e Ec2Runner, scope constructs.Construct, id *string, props *Ec2RunnerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.Ec2Runner",
		[]interface{}{scope, id, props},
		e,
	)
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

