//go:build no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ICompositeProvider) validateGetStepFunctionTaskParameters(parameters *RunnerRuntimeParameters) error {
	return nil
}

func (i *jsiiProxy_ICompositeProvider) validateGrantStateMachineParameters(stateMachineRole awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_ICompositeProvider) validateStatusParameters(statusFunctionRole awsiam.IGrantable) error {
	return nil
}

