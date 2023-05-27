//go:build no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func validateStaticRunnerImage_FromDockerHubParameters(scope constructs.Construct, id *string, image *string) error {
	return nil
}

func validateStaticRunnerImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	return nil
}

