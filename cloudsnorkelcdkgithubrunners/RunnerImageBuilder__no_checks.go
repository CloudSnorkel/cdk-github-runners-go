//go:build no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RunnerImageBuilder) validateAddComponentParameters(component RunnerImageComponent) error {
	return nil
}

func (r *jsiiProxy_RunnerImageBuilder) validateRemoveComponentParameters(component RunnerImageComponent) error {
	return nil
}

func validateRunnerImageBuilder_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRunnerImageBuilder_NewParameters(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) error {
	return nil
}

func (j *jsiiProxy_RunnerImageBuilder) validateSetComponentsParameters(val *[]RunnerImageComponent) error {
	return nil
}

func validateNewRunnerImageBuilderParameters(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) error {
	return nil
}

