//go:build !no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateCompositeProvider_DistributeParameters(scope constructs.Construct, id *string, weightedProviders *[]*WeightedRunnerProvider) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if weightedProviders == nil {
		return fmt.Errorf("parameter weightedProviders is required, but nil was provided")
	}
	for idx_023554, v := range *weightedProviders {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter weightedProviders[%#v]", idx_023554) }); err != nil {
			return err
		}
	}

	return nil
}

func validateCompositeProvider_FallbackParameters(scope constructs.Construct, id *string, providers *[]IRunnerProvider) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if providers == nil {
		return fmt.Errorf("parameter providers is required, but nil was provided")
	}

	return nil
}

