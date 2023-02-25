//go:build no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func validateLinuxUbuntuComponents_AwsCliParameters(scope constructs.Construct, id *string, architecture Architecture) error {
	return nil
}

func validateLinuxUbuntuComponents_DockerParameters(scope constructs.Construct, id *string, _architecture Architecture) error {
	return nil
}

func validateLinuxUbuntuComponents_ExtraCertificatesParameters(scope constructs.Construct, id *string, path *string) error {
	return nil
}

func validateLinuxUbuntuComponents_GitParameters(scope constructs.Construct, id *string, _architecture Architecture) error {
	return nil
}

func validateLinuxUbuntuComponents_GithubCliParameters(scope constructs.Construct, id *string, _architecture Architecture) error {
	return nil
}

func validateLinuxUbuntuComponents_GithubRunnerParameters(scope constructs.Construct, id *string, runnerVersion RunnerVersion, architecture Architecture) error {
	return nil
}

func validateLinuxUbuntuComponents_RequiredPackagesParameters(scope constructs.Construct, id *string, architecture Architecture) error {
	return nil
}

func validateLinuxUbuntuComponents_RunnerUserParameters(scope constructs.Construct, id *string, _architecture Architecture) error {
	return nil
}

