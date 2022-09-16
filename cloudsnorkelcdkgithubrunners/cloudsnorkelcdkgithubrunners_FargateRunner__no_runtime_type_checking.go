//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FargateRunner) validateGetStepFunctionTaskParameters(parameters *RunnerRuntimeParameters) error {
	return nil
}

func validateFargateRunner_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewFargateRunnerParameters(scope constructs.Construct, id *string, props *FargateRunnerProps) error {
	return nil
}

