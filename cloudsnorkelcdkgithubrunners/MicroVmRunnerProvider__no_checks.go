//go:build no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func validateMicroVmRunnerProvider_ImageBuilderParameters(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) error {
	return nil
}

func validateMicroVmRunnerProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewMicroVmRunnerProviderParameters(scope constructs.Construct, id *string, props *MicroVmRunnerProviderProps) error {
	return nil
}

