//go:build no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2Runner) validateAddRetryParameters(task interface{}, errors *[]*string) error {
	return nil
}

func (e *jsiiProxy_Ec2Runner) validateGetStepFunctionTaskParameters(parameters *RunnerRuntimeParameters) error {
	return nil
}

func (e *jsiiProxy_Ec2Runner) validateGrantStateMachineParameters(stateMachineRole awsiam.IGrantable) error {
	return nil
}

func (e *jsiiProxy_Ec2Runner) validateLabelsFromPropertiesParameters(defaultLabel *string) error {
	return nil
}

func (e *jsiiProxy_Ec2Runner) validateStatusParameters(statusFunctionRole awsiam.IGrantable) error {
	return nil
}

func validateEc2Runner_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewEc2RunnerParameters(scope constructs.Construct, id *string, props *Ec2RunnerProps) error {
	return nil
}

