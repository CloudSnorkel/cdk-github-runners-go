// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners


// Workflow job parameters as parsed from the webhook event. Pass these into your runner executor and run something like:.
//
// ```sh
// ./config.sh --unattended --url "https://${GITHUB_DOMAIN}/${OWNER}/${REPO}" --token "${RUNNER_TOKEN}" --ephemeral --work _work --labels "${RUNNER_LABEL}" --name "${RUNNER_NAME}" --disableupdate
// ```
//
// All parameters are specified as step function paths and therefore must be used only in step function task parameters.
// Experimental.
type RunnerRuntimeParameters struct {
	// Path to GitHub domain.
	//
	// Most of the time this will be github.com but for self-hosted GitHub instances, this will be different.
	// Experimental.
	GithubDomainPath *string `field:"required" json:"githubDomainPath" yaml:"githubDomainPath"`
	// Path to repository owner name.
	// Experimental.
	OwnerPath *string `field:"required" json:"ownerPath" yaml:"ownerPath"`
	// Path to repository name.
	// Experimental.
	RepoPath *string `field:"required" json:"repoPath" yaml:"repoPath"`
	// Path to desired runner name.
	//
	// We specifically set the name to make troubleshooting easier.
	// Experimental.
	RunnerNamePath *string `field:"required" json:"runnerNamePath" yaml:"runnerNamePath"`
	// Path to runner token used to register token.
	// Experimental.
	RunnerTokenPath *string `field:"required" json:"runnerTokenPath" yaml:"runnerTokenPath"`
}

