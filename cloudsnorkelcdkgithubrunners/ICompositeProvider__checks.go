//go:build !no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

func (i *jsiiProxy_ICompositeProvider) validateGetStepFunctionTaskParameters(parameters IRunnerRuntimeParameters) error {
	if parameters == nil {
		return fmt.Errorf("parameter parameters is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ICompositeProvider) validateGrantStateMachineParameters(stateMachineRole awsiam.IGrantable) error {
	if stateMachineRole == nil {
		return fmt.Errorf("parameter stateMachineRole is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ICompositeProvider) validateStatusParameters(statusFunctionRole awsiam.IGrantable) error {
	if statusFunctionRole == nil {
		return fmt.Errorf("parameter statusFunctionRole is required, but nil was provided")
	}

	return nil
}

