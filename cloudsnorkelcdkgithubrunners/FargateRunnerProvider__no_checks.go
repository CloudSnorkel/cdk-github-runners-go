//go:build no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FargateRunnerProvider) validateAddRetryParameters(task interface{}, errors *[]*string) error {
	return nil
}

func (f *jsiiProxy_FargateRunnerProvider) validateGetStepFunctionTaskParameters(parameters *RunnerRuntimeParameters) error {
	return nil
}

func (f *jsiiProxy_FargateRunnerProvider) validateGrantStateMachineParameters(_arg awsiam.IGrantable) error {
	return nil
}

func (f *jsiiProxy_FargateRunnerProvider) validateLabelsFromPropertiesParameters(defaultLabel *string) error {
	return nil
}

func (f *jsiiProxy_FargateRunnerProvider) validateStatusParameters(statusFunctionRole awsiam.IGrantable) error {
	return nil
}

func validateFargateRunnerProvider_ImageBuilderParameters(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) error {
	return nil
}

func validateFargateRunnerProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewFargateRunnerProviderParameters(scope constructs.Construct, id *string, props *FargateRunnerProviderProps) error {
	return nil
}

