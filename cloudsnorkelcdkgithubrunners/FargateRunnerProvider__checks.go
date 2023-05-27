//go:build !no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (f *jsiiProxy_FargateRunnerProvider) validateGetStepFunctionTaskParameters(parameters *RunnerRuntimeParameters) error {
	if parameters == nil {
		return fmt.Errorf("parameter parameters is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(parameters, func() string { return "parameter parameters" }); err != nil {
		return err
	}

	return nil
}

func (f *jsiiProxy_FargateRunnerProvider) validateGrantStateMachineParameters(_arg awsiam.IGrantable) error {
	if _arg == nil {
		return fmt.Errorf("parameter _arg is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FargateRunnerProvider) validateLabelsFromPropertiesParameters(defaultLabel *string) error {
	if defaultLabel == nil {
		return fmt.Errorf("parameter defaultLabel is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FargateRunnerProvider) validateStatusParameters(statusFunctionRole awsiam.IGrantable) error {
	if statusFunctionRole == nil {
		return fmt.Errorf("parameter statusFunctionRole is required, but nil was provided")
	}

	return nil
}

func validateFargateRunnerProvider_ImageBuilderParameters(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) error {
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

func validateFargateRunnerProvider_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewFargateRunnerProviderParameters(scope constructs.Construct, id *string, props *FargateRunnerProviderProps) error {
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

