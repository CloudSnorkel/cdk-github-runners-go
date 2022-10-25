//go:build !no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	"fmt"
)

func (a *jsiiProxy_Architecture) validateIsParameters(arch Architecture) error {
	if arch == nil {
		return fmt.Errorf("parameter arch is required, but nil was provided")
	}

	return nil
}

