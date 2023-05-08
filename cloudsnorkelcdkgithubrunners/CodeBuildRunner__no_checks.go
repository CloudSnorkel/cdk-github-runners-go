//go:build no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodeBuildRunner) validateGetStepFunctionTaskParameters(parameters *RunnerRuntimeParameters) error {
	return nil
}

func (c *jsiiProxy_CodeBuildRunner) validateGrantStateMachineParameters(_arg awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_CodeBuildRunner) validateLabelsFromPropertiesParameters(defaultLabel *string) error {
	return nil
}

func (c *jsiiProxy_CodeBuildRunner) validateStatusParameters(statusFunctionRole awsiam.IGrantable) error {
	return nil
}

func validateCodeBuildRunner_ImageBuilderParameters(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) error {
	return nil
}

func validateCodeBuildRunner_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCodeBuildRunnerParameters(scope constructs.Construct, id *string, props *CodeBuildRunnerProviderProps) error {
	return nil
}

