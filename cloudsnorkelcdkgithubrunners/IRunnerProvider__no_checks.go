//go:build no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IRunnerProvider) validateGetStepFunctionTaskParameters(parameters *RunnerRuntimeParameters) error {
	return nil
}

func (i *jsiiProxy_IRunnerProvider) validateGrantStateMachineParameters(stateMachineRole awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IRunnerProvider) validateStatusParameters(statusFunctionRole awsiam.IGrantable) error {
	return nil
}

