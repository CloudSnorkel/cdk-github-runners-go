//go:build no_runtime_type_checking

// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GitHubRunners) validateMetricFailedParameters(props *awscloudwatch.MetricProps) error {
	return nil
}

func (g *jsiiProxy_GitHubRunners) validateMetricJobCompletedParameters(props *awscloudwatch.MetricProps) error {
	return nil
}

func (g *jsiiProxy_GitHubRunners) validateMetricSucceededParameters(props *awscloudwatch.MetricProps) error {
	return nil
}

func (g *jsiiProxy_GitHubRunners) validateMetricTimeParameters(props *awscloudwatch.MetricProps) error {
	return nil
}

func validateGitHubRunners_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewGitHubRunnersParameters(scope constructs.Construct, id *string, props *GitHubRunnersProps) error {
	return nil
}

