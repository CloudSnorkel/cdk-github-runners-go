//go:build !no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

func (a *jsiiProxy_Architecture) validateInstanceTypeMatchParameters(instanceType awsec2.InstanceType) error {
	if instanceType == nil {
		return fmt.Errorf("parameter instanceType is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Architecture) validateIsParameters(arch Architecture) error {
	if arch == nil {
		return fmt.Errorf("parameter arch is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Architecture) validateIsInParameters(arches *[]Architecture) error {
	if arches == nil {
		return fmt.Errorf("parameter arches is required, but nil was provided")
	}

	return nil
}

