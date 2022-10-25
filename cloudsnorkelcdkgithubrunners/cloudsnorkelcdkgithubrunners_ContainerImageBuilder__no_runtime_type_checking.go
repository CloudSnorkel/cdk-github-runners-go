//go:build no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerImageBuilder) validateAddComponentParameters(component ImageBuilderComponent) error {
	return nil
}

func (c *jsiiProxy_ContainerImageBuilder) validateAddExtraCertificatesParameters(path *string) error {
	return nil
}

func (c *jsiiProxy_ContainerImageBuilder) validatePrependComponentParameters(component ImageBuilderComponent) error {
	return nil
}

func validateContainerImageBuilder_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewContainerImageBuilderParameters(scope constructs.Construct, id *string, props *ContainerImageBuilderProps) error {
	return nil
}

