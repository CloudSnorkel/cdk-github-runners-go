//go:build no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2RunnerProvider) validateGetStepFunctionTaskParameters(parameters *RunnerRuntimeParameters) error {
	return nil
}

func (e *jsiiProxy_Ec2RunnerProvider) validateGrantStateMachineParameters(stateMachineRole awsiam.IGrantable) error {
	return nil
}

func (e *jsiiProxy_Ec2RunnerProvider) validateLabelsFromPropertiesParameters(defaultLabel *string) error {
	return nil
}

func (e *jsiiProxy_Ec2RunnerProvider) validateStatusParameters(statusFunctionRole awsiam.IGrantable) error {
	return nil
}

func validateEc2RunnerProvider_ImageBuilderParameters(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) error {
	return nil
}

func validateEc2RunnerProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewEc2RunnerProviderParameters(scope constructs.Construct, id *string, props *Ec2RunnerProviderProps) error {
	return nil
}

