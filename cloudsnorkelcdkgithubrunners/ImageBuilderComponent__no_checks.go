//go:build no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ImageBuilderComponent) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_ImageBuilderComponent) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_ImageBuilderComponent) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (i *jsiiProxy_ImageBuilderComponent) validateGrantAssetsReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_ImageBuilderComponent) validatePrefixCommandsWithErrorHandlingParameters(platform *string, commands *[]*string) error {
	return nil
}

func validateImageBuilderComponent_IsConstructParameters(x interface{}) error {
	return nil
}

func validateImageBuilderComponent_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateImageBuilderComponent_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewImageBuilderComponentParameters(scope constructs.Construct, id *string, props *ImageBuilderComponentProperties) error {
	return nil
}

