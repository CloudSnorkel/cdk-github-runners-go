//go:build no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func validateScheduledWarmRunner_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewScheduledWarmRunnerParameters(scope constructs.Construct, id *string, props *ScheduledWarmRunnerProps) error {
	return nil
}

