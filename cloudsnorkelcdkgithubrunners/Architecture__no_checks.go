//go:build no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Architecture) validateInstanceTypeMatchParameters(instanceType awsec2.InstanceType) error {
	return nil
}

func (a *jsiiProxy_Architecture) validateIsParameters(arch Architecture) error {
	return nil
}

func (a *jsiiProxy_Architecture) validateIsInParameters(arches *[]Architecture) error {
	return nil
}

