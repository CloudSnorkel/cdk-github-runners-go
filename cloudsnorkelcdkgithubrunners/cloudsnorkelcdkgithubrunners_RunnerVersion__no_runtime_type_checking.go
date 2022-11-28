//go:build no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RunnerVersion) validateIsParameters(other RunnerVersion) error {
	return nil
}

func validateRunnerVersion_SpecificParameters(version *string) error {
	return nil
}

func validateNewRunnerVersionParameters(version *string) error {
	return nil
}

