//go:build !no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CodeBuildImageBuilder) validateAddExtraCertificatesParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CodeBuildImageBuilder) validateAddFilesParameters(sourcePath *string, destName *string) error {
	if sourcePath == nil {
		return fmt.Errorf("parameter sourcePath is required, but nil was provided")
	}

	if destName == nil {
		return fmt.Errorf("parameter destName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CodeBuildImageBuilder) validateAddPolicyStatementParameters(statement awsiam.PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CodeBuildImageBuilder) validateAddPostBuildCommandParameters(command *string) error {
	if command == nil {
		return fmt.Errorf("parameter command is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CodeBuildImageBuilder) validateAddPreBuildCommandParameters(command *string) error {
	if command == nil {
		return fmt.Errorf("parameter command is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CodeBuildImageBuilder) validateSetBuildArgParameters(name *string, value *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCodeBuildImageBuilder_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewCodeBuildImageBuilderParameters(scope constructs.Construct, id *string, props *CodeBuildImageBuilderProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

