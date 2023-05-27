//go:build no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func validateSecrets_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSecretsParameters(scope constructs.Construct, id *string) error {
	return nil
}

