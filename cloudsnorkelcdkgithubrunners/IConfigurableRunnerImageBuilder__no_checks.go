//go:build no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IConfigurableRunnerImageBuilder) validateAddComponentParameters(component RunnerImageComponent) error {
	return nil
}

func (i *jsiiProxy_IConfigurableRunnerImageBuilder) validateRemoveComponentParameters(component RunnerImageComponent) error {
	return nil
}

