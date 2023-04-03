//go:build !no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (r *jsiiProxy_RunnerImageBuilder) validateAddComponentParameters(component RunnerImageComponent) error {
	if component == nil {
		return fmt.Errorf("parameter component is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RunnerImageBuilder) validateRemoveComponentParameters(component RunnerImageComponent) error {
	if component == nil {
		return fmt.Errorf("parameter component is required, but nil was provided")
	}

	return nil
}

func validateRunnerImageBuilder_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateRunnerImageBuilder_NewParameters(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_RunnerImageBuilder) validateSetComponentsParameters(val *[]RunnerImageComponent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewRunnerImageBuilderParameters(scope constructs.Construct, id *string, props *RunnerImageBuilderProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

