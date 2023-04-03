//go:build no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodeBuildRunnerProvider) validateAddRetryParameters(task interface{}, errors *[]*string) error {
	return nil
}

func (c *jsiiProxy_CodeBuildRunnerProvider) validateGetStepFunctionTaskParameters(parameters *RunnerRuntimeParameters) error {
	return nil
}

func (c *jsiiProxy_CodeBuildRunnerProvider) validateGrantStateMachineParameters(_arg awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_CodeBuildRunnerProvider) validateLabelsFromPropertiesParameters(defaultLabel *string) error {
	return nil
}

func (c *jsiiProxy_CodeBuildRunnerProvider) validateStatusParameters(statusFunctionRole awsiam.IGrantable) error {
	return nil
}

func validateCodeBuildRunnerProvider_ImageBuilderParameters(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) error {
	return nil
}

func validateCodeBuildRunnerProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCodeBuildRunnerProviderParameters(scope constructs.Construct, id *string, props *CodeBuildRunnerProviderProps) error {
	return nil
}

