//go:build no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Compliance) validateEncryptParameters(options *ComplianceEncryptionOptions) error {
	return nil
}

func (c *jsiiProxy_Compliance) validateLimitLambdaConcurrencyParameters(options *ComplianceLambdaConcurrencyOptions) error {
	return nil
}

func validateCompliance_OfParameters(scope constructs.IConstruct) error {
	return nil
}

