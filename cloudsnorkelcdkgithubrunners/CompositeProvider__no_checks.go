//go:build no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func validateCompositeProvider_DistributeParameters(scope constructs.Construct, id *string, weightedProviders *[]*WeightedRunnerProvider) error {
	return nil
}

func validateCompositeProvider_FallbackParameters(scope constructs.Construct, id *string, providers *[]IRunnerProvider) error {
	return nil
}

