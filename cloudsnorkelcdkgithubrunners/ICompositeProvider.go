package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

// Interface for composite runner providers that interact with multiple sub-providers.
//
// Unlike IRunnerProvider, composite providers do not have connections, grant capabilities,
// log groups, or retryable errors as they delegate to their sub-providers.
// Experimental.
type ICompositeProvider interface {
	constructs.IConstruct
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
	// Return statuses of all sub-providers to be used in the main status function.
	//
	// Also gives the status function any needed permissions to query the Docker images or AMIs.
	// Experimental.
	Status(statusFunctionRole awsiam.IGrantable) *[]IRunnerProviderStatus
	// GitHub Actions labels used for this provider.
	//
	// These labels are used to identify which provider should spawn a new on-demand runner. Every job sends a webhook with the labels it's looking for
	// based on runs-on. We use match the labels from the webhook with the labels specified here. If all the labels specified here are present in the
	// job's labels, this provider will be chosen and spawn a new runner.
	// Experimental.
	Labels() *[]*string
	// All sub-providers contained in this composite provider.
	//
	// This is used to extract providers for metric filters and other operations.
	// Experimental.
	Providers() *[]IRunnerProvider
}

// The jsii proxy for ICompositeProvider
type jsiiProxy_ICompositeProvider struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_ICompositeProvider) GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable {
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

func (i *jsiiProxy_ICompositeProvider) GrantStateMachine(stateMachineRole awsiam.IGrantable) {
	if err := i.validateGrantStateMachineParameters(stateMachineRole); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"grantStateMachine",
		[]interface{}{stateMachineRole},
	)
}

func (i *jsiiProxy_ICompositeProvider) Status(statusFunctionRole awsiam.IGrantable) *[]IRunnerProviderStatus {
	if err := i.validateStatusParameters(statusFunctionRole); err != nil {
		panic(err)
	}
	var returns *[]IRunnerProviderStatus

	_jsii_.Invoke(
		i,
		"status",
		[]interface{}{statusFunctionRole},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ICompositeProvider) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICompositeProvider) Providers() *[]IRunnerProvider {
	var returns *[]IRunnerProvider
	_jsii_.Get(
		j,
		"providers",
		&returns,
	)
	return returns
}

