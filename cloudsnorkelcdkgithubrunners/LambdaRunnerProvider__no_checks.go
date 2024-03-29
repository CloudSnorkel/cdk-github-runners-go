//go:build no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaRunnerProvider) validateGetStepFunctionTaskParameters(parameters *RunnerRuntimeParameters) error {
	return nil
}

func (l *jsiiProxy_LambdaRunnerProvider) validateGrantStateMachineParameters(_arg awsiam.IGrantable) error {
	return nil
}

func (l *jsiiProxy_LambdaRunnerProvider) validateLabelsFromPropertiesParameters(defaultLabel *string) error {
	return nil
}

func (l *jsiiProxy_LambdaRunnerProvider) validateStatusParameters(statusFunctionRole awsiam.IGrantable) error {
	return nil
}

func validateLambdaRunnerProvider_ImageBuilderParameters(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) error {
	return nil
}

func validateLambdaRunnerProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewLambdaRunnerProviderParameters(scope constructs.Construct, id *string, props *LambdaRunnerProviderProps) error {
	return nil
}

