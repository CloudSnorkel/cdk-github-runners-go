//go:build !no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

import (
	"fmt"
)

func (i *jsiiProxy_IConfigurableRunnerImageBuilder) validateAddComponentParameters(component RunnerImageComponent) error {
	if component == nil {
		return fmt.Errorf("parameter component is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IConfigurableRunnerImageBuilder) validateRemoveComponentParameters(component RunnerImageComponent) error {
	if component == nil {
		return fmt.Errorf("parameter component is required, but nil was provided")
	}

	return nil
}

