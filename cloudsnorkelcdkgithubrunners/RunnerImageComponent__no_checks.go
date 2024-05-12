//go:build no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RunnerImageComponent) validateGetAssetsParameters(_os Os, _architecture Architecture) error {
	return nil
}

func (r *jsiiProxy_RunnerImageComponent) validateGetCommandsParameters(_os Os, _architecture Architecture) error {
	return nil
}

func (r *jsiiProxy_RunnerImageComponent) validateGetDockerCommandsParameters(_os Os, _architecture Architecture) error {
	return nil
}

func (r *jsiiProxy_RunnerImageComponent) validateShouldRebootParameters(_os Os, _architecture Architecture) error {
	return nil
}

func validateRunnerImageComponent_CustomParameters(props *RunnerImageComponentCustomProps) error {
	return nil
}

func validateRunnerImageComponent_EnvironmentVariablesParameters(vars *map[string]*string) error {
	return nil
}

func validateRunnerImageComponent_ExtraCertificatesParameters(source *string, name *string) error {
	return nil
}

func validateRunnerImageComponent_GithubRunnerParameters(runnerVersion RunnerVersion) error {
	return nil
}

