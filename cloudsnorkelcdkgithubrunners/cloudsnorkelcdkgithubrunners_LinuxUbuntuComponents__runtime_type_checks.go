//go:build !no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateLinuxUbuntuComponents_AwsCliParameters(scope constructs.Construct, id *string, architecture Architecture) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if architecture == nil {
		return fmt.Errorf("parameter architecture is required, but nil was provided")
	}

	return nil
}

func validateLinuxUbuntuComponents_DockerParameters(scope constructs.Construct, id *string, _architecture Architecture) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if _architecture == nil {
		return fmt.Errorf("parameter _architecture is required, but nil was provided")
	}

	return nil
}

func validateLinuxUbuntuComponents_GitParameters(scope constructs.Construct, id *string, _architecture Architecture) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if _architecture == nil {
		return fmt.Errorf("parameter _architecture is required, but nil was provided")
	}

	return nil
}

func validateLinuxUbuntuComponents_GithubCliParameters(scope constructs.Construct, id *string, _architecture Architecture) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if _architecture == nil {
		return fmt.Errorf("parameter _architecture is required, but nil was provided")
	}

	return nil
}

func validateLinuxUbuntuComponents_GithubRunnerParameters(scope constructs.Construct, id *string, runnerVersion RunnerVersion, architecture Architecture) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if runnerVersion == nil {
		return fmt.Errorf("parameter runnerVersion is required, but nil was provided")
	}

	if architecture == nil {
		return fmt.Errorf("parameter architecture is required, but nil was provided")
	}

	return nil
}

func validateLinuxUbuntuComponents_RequiredPackagesParameters(scope constructs.Construct, id *string, architecture Architecture) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if architecture == nil {
		return fmt.Errorf("parameter architecture is required, but nil was provided")
	}

	return nil
}

func validateLinuxUbuntuComponents_RunnerUserParameters(scope constructs.Construct, id *string, _architecture Architecture) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if _architecture == nil {
		return fmt.Errorf("parameter _architecture is required, but nil was provided")
	}

	return nil
}

