// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

// Interface for all runner providers.
//
// Implementations create all required resources and return a step function task that starts those resources from {@link getStepFunctionTask}.
// Experimental.
type IRunnerProvider interface {
	awsec2.IConnectable
	awsiam.IGrantable
	// Generate step function tasks that execute the runner.
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
	// Return status of the runner provider to be used in the main status function.
	//
	// Also gives the status function any needed permissions to query the Docker image or AMI.
	// Experimental.
	Status(statusFunctionRole awsiam.IGrantable) IRunnerProviderStatus
	// GitHub Actions labels used for this provider.
	//
	// These labels are used to identify which provider should spawn a new on-demand runner. Every job sends a webhook with the labels it's looking for
	// based on runs-on. We use match the labels from the webhook with the labels specified here. If all the labels specified here are present in the
	// job's labels, this provider will be chosen and spawn a new runner.
	// Experimental.
	Labels() *[]*string
	// Security group associated with runners.
	// Experimental.
	SecurityGroup() awsec2.ISecurityGroup
	// VPC network in which runners will be placed.
	// Experimental.
	Vpc() awsec2.IVpc
}

// The jsii proxy for IRunnerProvider
type jsiiProxy_IRunnerProvider struct {
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
}

func (i *jsiiProxy_IRunnerProvider) GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable {
	if err := i.validateGetStepFunctionTaskParameters(parameters); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		i,
		"getStepFunctionTask",
		[]interface{}{parameters},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRunnerProvider) GrantStateMachine(stateMachineRole awsiam.IGrantable) {
	if err := i.validateGrantStateMachineParameters(stateMachineRole); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"grantStateMachine",
		[]interface{}{stateMachineRole},
	)
}

func (i *jsiiProxy_IRunnerProvider) Status(statusFunctionRole awsiam.IGrantable) IRunnerProviderStatus {
	if err := i.validateStatusParameters(statusFunctionRole); err != nil {
		panic(err)
	}
	var returns IRunnerProviderStatus

	_jsii_.Invoke(
		i,
		"status",
		[]interface{}{statusFunctionRole},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IRunnerProvider) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerProvider) SecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerProvider) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerProvider) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerProvider) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

