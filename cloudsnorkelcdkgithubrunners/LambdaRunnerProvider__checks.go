//go:build !no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
)

func (l *jsiiProxy_LambdaRunnerProvider) validateAddRetryParameters(task interface{}, errors *[]*string) error {
	if task == nil {
		return fmt.Errorf("parameter task is required, but nil was provided")
	}
	switch task.(type) {
	case awsstepfunctions.Parallel:
		// ok
	case awsstepfunctions.TaskStateBase:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(task) {
			return fmt.Errorf("parameter task must be one of the allowed types: awsstepfunctions.Parallel, awsstepfunctions.TaskStateBase; received %#v (a %T)", task, task)
		}
	}

	if errors == nil {
		return fmt.Errorf("parameter errors is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LambdaRunnerProvider) validateGetStepFunctionTaskParameters(parameters *RunnerRuntimeParameters) error {
	if parameters == nil {
		return fmt.Errorf("parameter parameters is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(parameters, func() string { return "parameter parameters" }); err != nil {
		return err
	}

	return nil
}

func (l *jsiiProxy_LambdaRunnerProvider) validateGrantStateMachineParameters(_arg awsiam.IGrantable) error {
	if _arg == nil {
		return fmt.Errorf("parameter _arg is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LambdaRunnerProvider) validateLabelsFromPropertiesParameters(defaultLabel *string) error {
	if defaultLabel == nil {
		return fmt.Errorf("parameter defaultLabel is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LambdaRunnerProvider) validateStatusParameters(statusFunctionRole awsiam.IGrantable) error {
	if statusFunctionRole == nil {
		return fmt.Errorf("parameter statusFunctionRole is required, but nil was provided")
	}

	return nil
}

func validateLambdaRunnerProvider_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewLambdaRunnerProviderParameters(scope constructs.Construct, id *string, props *LambdaRunnerProviderProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

