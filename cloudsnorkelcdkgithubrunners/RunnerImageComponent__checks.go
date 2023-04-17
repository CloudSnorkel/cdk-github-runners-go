//go:build !no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (r *jsiiProxy_RunnerImageComponent) validateGetAssetsParameters(_os Os, _architecture Architecture) error {
	if _os == nil {
		return fmt.Errorf("parameter _os is required, but nil was provided")
	}

	if _architecture == nil {
		return fmt.Errorf("parameter _architecture is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RunnerImageComponent) validateGetCommandsParameters(_os Os, _architecture Architecture) error {
	if _os == nil {
		return fmt.Errorf("parameter _os is required, but nil was provided")
	}

	if _architecture == nil {
		return fmt.Errorf("parameter _architecture is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RunnerImageComponent) validateGetDockerCommandsParameters(_os Os, _architecture Architecture) error {
	if _os == nil {
		return fmt.Errorf("parameter _os is required, but nil was provided")
	}

	if _architecture == nil {
		return fmt.Errorf("parameter _architecture is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RunnerImageComponent) validateShouldRebootParameters(_os Os, _architecture Architecture) error {
	if _os == nil {
		return fmt.Errorf("parameter _os is required, but nil was provided")
	}

	if _architecture == nil {
		return fmt.Errorf("parameter _architecture is required, but nil was provided")
	}

	return nil
}

func validateRunnerImageComponent_CustomParameters(props *RunnerImageComponentCustomProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateRunnerImageComponent_ExtraCertificatesParameters(source *string, name *string) error {
	if source == nil {
		return fmt.Errorf("parameter source is required, but nil was provided")
	}

	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateRunnerImageComponent_GithubRunnerParameters(runnerVersion RunnerVersion) error {
	if runnerVersion == nil {
		return fmt.Errorf("parameter runnerVersion is required, but nil was provided")
	}

	return nil
}

