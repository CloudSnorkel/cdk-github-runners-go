//go:build !no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/constructs-go/constructs/v10"
)

func validateBaseImage_FromAmiIdParameters(amiId *string) error {
	if amiId == nil {
		return fmt.Errorf("parameter amiId is required, but nil was provided")
	}

	return nil
}

func validateBaseImage_FromGpuBaseParameters(os Os, architecture Architecture) error {
	if os == nil {
		return fmt.Errorf("parameter os is required, but nil was provided")
	}

	if architecture == nil {
		return fmt.Errorf("parameter architecture is required, but nil was provided")
	}

	return nil
}

func validateBaseImage_FromImageBuilderParameters(scope constructs.Construct, resourceName *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if resourceName == nil {
		return fmt.Errorf("parameter resourceName is required, but nil was provided")
	}

	return nil
}

func validateBaseImage_FromMarketplaceProductIdParameters(productId *string) error {
	if productId == nil {
		return fmt.Errorf("parameter productId is required, but nil was provided")
	}

	return nil
}

func validateBaseImage_FromSsmParameterParameters(parameter awsssm.IParameter) error {
	if parameter == nil {
		return fmt.Errorf("parameter parameter is required, but nil was provided")
	}

	return nil
}

func validateBaseImage_FromSsmParameterNameParameters(parameterName *string) error {
	if parameterName == nil {
		return fmt.Errorf("parameter parameterName is required, but nil was provided")
	}

	return nil
}

func validateBaseImage_FromStringParameters(baseImageString *string) error {
	if baseImageString == nil {
		return fmt.Errorf("parameter baseImageString is required, but nil was provided")
	}

	return nil
}

func validateNewBaseImageParameters(image *string) error {
	if image == nil {
		return fmt.Errorf("parameter image is required, but nil was provided")
	}

	return nil
}

