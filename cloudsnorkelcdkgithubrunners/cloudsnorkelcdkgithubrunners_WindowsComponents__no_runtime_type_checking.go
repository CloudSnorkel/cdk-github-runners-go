//go:build no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func validateWindowsComponents_AwsCliParameters(scope constructs.Construct, id *string) error {
	return nil
}

func validateWindowsComponents_CloudwatchAgentParameters(scope constructs.Construct, id *string) error {
	return nil
}

func validateWindowsComponents_DockerParameters(scope constructs.Construct, id *string) error {
	return nil
}

func validateWindowsComponents_ExtraCertificatesParameters(scope constructs.Construct, id *string, path *string) error {
	return nil
}

func validateWindowsComponents_GitParameters(scope constructs.Construct, id *string) error {
	return nil
}

func validateWindowsComponents_GithubCliParameters(scope constructs.Construct, id *string) error {
	return nil
}

func validateWindowsComponents_GithubRunnerParameters(scope constructs.Construct, id *string, runnerVersion RunnerVersion) error {
	return nil
}

