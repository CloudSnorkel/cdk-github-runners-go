//go:build !no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

import (
	"fmt"
)

func (r *jsiiProxy_RunnerVersion) validateIsParameters(other RunnerVersion) error {
	if other == nil {
		return fmt.Errorf("parameter other is required, but nil was provided")
	}

	return nil
}

func validateRunnerVersion_SpecificParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

func validateNewRunnerVersionParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

