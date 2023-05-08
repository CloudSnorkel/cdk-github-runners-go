//go:build no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsRunnerProvider) validateGetStepFunctionTaskParameters(parameters *RunnerRuntimeParameters) error {
	return nil
}

func (e *jsiiProxy_EcsRunnerProvider) validateGrantStateMachineParameters(_arg awsiam.IGrantable) error {
	return nil
}

func (e *jsiiProxy_EcsRunnerProvider) validateLabelsFromPropertiesParameters(defaultLabel *string) error {
	return nil
}

func (e *jsiiProxy_EcsRunnerProvider) validateStatusParameters(statusFunctionRole awsiam.IGrantable) error {
	return nil
}

func validateEcsRunnerProvider_ImageBuilderParameters(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) error {
	return nil
}

func validateEcsRunnerProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewEcsRunnerProviderParameters(scope constructs.Construct, id *string, props *EcsRunnerProviderProps) error {
	return nil
}

