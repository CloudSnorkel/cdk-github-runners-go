//go:build no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaRunner) validateGetStepFunctionTaskParameters(parameters *RunnerRuntimeParameters) error {
	return nil
}

func (l *jsiiProxy_LambdaRunner) validateGrantStateMachineParameters(_arg awsiam.IGrantable) error {
	return nil
}

func (l *jsiiProxy_LambdaRunner) validateLabelsFromPropertiesParameters(defaultLabel *string) error {
	return nil
}

func (l *jsiiProxy_LambdaRunner) validateStatusParameters(statusFunctionRole awsiam.IGrantable) error {
	return nil
}

func validateLambdaRunner_ImageBuilderParameters(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) error {
	return nil
}

func validateLambdaRunner_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewLambdaRunnerParameters(scope constructs.Construct, id *string, props *LambdaRunnerProviderProps) error {
	return nil
}

