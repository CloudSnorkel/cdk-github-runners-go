//go:build !no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_ContainerImageBuilder) validateAddComponentParameters(component ImageBuilderComponent) error {
	if component == nil {
		return fmt.Errorf("parameter component is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ContainerImageBuilder) validateAddExtraCertificatesParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ContainerImageBuilder) validatePrependComponentParameters(component ImageBuilderComponent) error {
	if component == nil {
		return fmt.Errorf("parameter component is required, but nil was provided")
	}

	return nil
}

func validateContainerImageBuilder_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewContainerImageBuilderParameters(scope constructs.Construct, id *string, props *ContainerImageBuilderProps) error {
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

