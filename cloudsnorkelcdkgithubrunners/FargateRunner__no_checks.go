//go:build no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FargateRunner) validateGetStepFunctionTaskParameters(parameters *RunnerRuntimeParameters) error {
	return nil
}

func (f *jsiiProxy_FargateRunner) validateGrantStateMachineParameters(_arg awsiam.IGrantable) error {
	return nil
}

func (f *jsiiProxy_FargateRunner) validateLabelsFromPropertiesParameters(defaultLabel *string) error {
	return nil
}

func (f *jsiiProxy_FargateRunner) validateStatusParameters(statusFunctionRole awsiam.IGrantable) error {
	return nil
}

func validateFargateRunner_ImageBuilderParameters(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) error {
	return nil
}

func validateFargateRunner_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewFargateRunnerParameters(scope constructs.Construct, id *string, props *FargateRunnerProviderProps) error {
	return nil
}

