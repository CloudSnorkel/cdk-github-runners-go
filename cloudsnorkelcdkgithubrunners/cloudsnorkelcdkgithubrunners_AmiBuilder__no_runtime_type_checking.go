//go:build no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AmiBuilder) validateAddComponentParameters(component ImageBuilderComponent) error {
	return nil
}

func (a *jsiiProxy_AmiBuilder) validateAddExtraCertificatesParameters(path *string) error {
	return nil
}

func (a *jsiiProxy_AmiBuilder) validateCreateImageParameters(infra awsimagebuilder.CfnInfrastructureConfiguration, dist awsimagebuilder.CfnDistributionConfiguration, log awslogs.LogGroup) error {
	return nil
}

func (a *jsiiProxy_AmiBuilder) validateCreateInfrastructureParameters(managedPolicies *[]awsiam.IManagedPolicy) error {
	return nil
}

func (a *jsiiProxy_AmiBuilder) validateCreateLogParameters(recipeName *string) error {
	return nil
}

func (a *jsiiProxy_AmiBuilder) validateCreatePipelineParameters(infra awsimagebuilder.CfnInfrastructureConfiguration, dist awsimagebuilder.CfnDistributionConfiguration, log awslogs.LogGroup) error {
	return nil
}

func (a *jsiiProxy_AmiBuilder) validatePrependComponentParameters(component ImageBuilderComponent) error {
	return nil
}

func validateAmiBuilder_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AmiBuilder) validateSetComponentsParameters(val *[]ImageBuilderComponent) error {
	return nil
}

func validateNewAmiBuilderParameters(scope constructs.Construct, id *string, props *AmiBuilderProps) error {
	return nil
}

