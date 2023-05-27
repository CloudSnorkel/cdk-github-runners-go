//go:build !no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
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

func (c *jsiiProxy_ContainerImageBuilder) validateCreateImageParameters(infra awsimagebuilder.CfnInfrastructureConfiguration, dist awsimagebuilder.CfnDistributionConfiguration, log awslogs.LogGroup) error {
	if infra == nil {
		return fmt.Errorf("parameter infra is required, but nil was provided")
	}

	if dist == nil {
		return fmt.Errorf("parameter dist is required, but nil was provided")
	}

	if log == nil {
		return fmt.Errorf("parameter log is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ContainerImageBuilder) validateCreateInfrastructureParameters(managedPolicies *[]awsiam.IManagedPolicy) error {
	if managedPolicies == nil {
		return fmt.Errorf("parameter managedPolicies is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ContainerImageBuilder) validateCreateLogParameters(recipeName *string) error {
	if recipeName == nil {
		return fmt.Errorf("parameter recipeName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ContainerImageBuilder) validateCreatePipelineParameters(infra awsimagebuilder.CfnInfrastructureConfiguration, dist awsimagebuilder.CfnDistributionConfiguration, log awslogs.LogGroup) error {
	if infra == nil {
		return fmt.Errorf("parameter infra is required, but nil was provided")
	}

	if dist == nil {
		return fmt.Errorf("parameter dist is required, but nil was provided")
	}

	if log == nil {
		return fmt.Errorf("parameter log is required, but nil was provided")
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

func (j *jsiiProxy_ContainerImageBuilder) validateSetComponentsParameters(val *[]ImageBuilderComponent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
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

