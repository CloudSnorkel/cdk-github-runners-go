//go:build no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodeBuildImageBuilder) validateAddExtraCertificatesParameters(path *string) error {
	return nil
}

func (c *jsiiProxy_CodeBuildImageBuilder) validateAddFilesParameters(sourcePath *string, destName *string) error {
	return nil
}

func (c *jsiiProxy_CodeBuildImageBuilder) validateAddPolicyStatementParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (c *jsiiProxy_CodeBuildImageBuilder) validateAddPostBuildCommandParameters(command *string) error {
	return nil
}

func (c *jsiiProxy_CodeBuildImageBuilder) validateAddPreBuildCommandParameters(command *string) error {
	return nil
}

func (c *jsiiProxy_CodeBuildImageBuilder) validateSetBuildArgParameters(name *string, value *string) error {
	return nil
}

func validateCodeBuildImageBuilder_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCodeBuildImageBuilderParameters(scope constructs.Construct, id *string, props *CodeBuildImageBuilderProps) error {
	return nil
}

