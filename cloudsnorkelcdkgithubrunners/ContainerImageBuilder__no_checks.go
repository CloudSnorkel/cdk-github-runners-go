//go:build no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerImageBuilder) validateAddComponentParameters(component ImageBuilderComponent) error {
	return nil
}

func (c *jsiiProxy_ContainerImageBuilder) validateAddExtraCertificatesParameters(path *string) error {
	return nil
}

func (c *jsiiProxy_ContainerImageBuilder) validateCreateImageParameters(infra awsimagebuilder.CfnInfrastructureConfiguration, dist awsimagebuilder.CfnDistributionConfiguration, log awslogs.LogGroup) error {
	return nil
}

func (c *jsiiProxy_ContainerImageBuilder) validateCreateInfrastructureParameters(managedPolicies *[]awsiam.IManagedPolicy) error {
	return nil
}

func (c *jsiiProxy_ContainerImageBuilder) validateCreateLogParameters(recipeName *string) error {
	return nil
}

func (c *jsiiProxy_ContainerImageBuilder) validateCreatePipelineParameters(infra awsimagebuilder.CfnInfrastructureConfiguration, dist awsimagebuilder.CfnDistributionConfiguration, log awslogs.LogGroup) error {
	return nil
}

func (c *jsiiProxy_ContainerImageBuilder) validatePrependComponentParameters(component ImageBuilderComponent) error {
	return nil
}

func validateContainerImageBuilder_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerImageBuilder) validateSetComponentsParameters(val *[]ImageBuilderComponent) error {
	return nil
}

func validateNewContainerImageBuilderParameters(scope constructs.Construct, id *string, props *ContainerImageBuilderProps) error {
	return nil
}

