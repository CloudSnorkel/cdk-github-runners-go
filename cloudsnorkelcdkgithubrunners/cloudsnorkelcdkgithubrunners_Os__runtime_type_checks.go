//go:build !no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	"fmt"
)

func (o *jsiiProxy_Os) validateIsParameters(os Os) error {
	if os == nil {
		return fmt.Errorf("parameter os is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_Os) validateIsInParameters(oses *[]Os) error {
	if oses == nil {
		return fmt.Errorf("parameter oses is required, but nil was provided")
	}

	return nil
}

