//go:build no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func validateAlwaysOnWarmRunner_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAlwaysOnWarmRunnerParameters(scope constructs.Construct, id *string, props *AlwaysOnWarmRunnerProps) error {
	return nil
}

