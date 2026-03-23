package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Workflow job parameters as parsed from the webhook event. Pass these into your runner executor and run something like:.
//
// ```sh
// ./config.sh --unattended --url "{REGISTRATION_URL}" --token "${RUNNER_TOKEN}" --ephemeral --work _work --labels "${RUNNER_LABEL}" --name "${RUNNER_NAME}" --disableupdate
// ```
//
// All parameters are specified as step function paths and therefore must be used only in step function task parameters.
// Experimental.
type IRunnerRuntimeParameters interface {
	// Catches all errors and cleans up the failed runner from GitHub Actions.
	//
	// It is important to fully clean up after any failed runner provisioning. GitHub
	// will fail booting a new runner if the previous one with the same name is not
	// fully cleaned up.
	// Experimental.
	AddCatchAndCleanUp(state interface{}, next awsstepfunctions.IChainable)
	// Path to GitHub domain.
	//
	// Most of the time this will be github.com but for self-hosted GitHub instances, this will be different.
	// Experimental.
	GithubDomainPath() *string
	// Path to comma-separated labels string to use for runner.
	// Experimental.
	LabelsPath() *string
	// Path to repository owner name.
	// Experimental.
	OwnerPath() *string
	// Repository or organization URL to register runner at.
	// Experimental.
	RegistrationUrl() *string
	// Path to repository name.
	// Experimental.
	RepoPath() *string
	// Path to desired runner name.
	//
	// We specifically set the name to make troubleshooting easier.
	// Experimental.
	RunnerNamePath() *string
	// Path to runner token used to register token.
	// Experimental.
	RunnerTokenPath() *string
}

// The jsii proxy for IRunnerRuntimeParameters
type jsiiProxy_IRunnerRuntimeParameters struct {
	_ byte // padding
}

func (i *jsiiProxy_IRunnerRuntimeParameters) AddCatchAndCleanUp(state interface{}, next awsstepfunctions.IChainable) {
	if err := i.validateAddCatchAndCleanUpParameters(state); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addCatchAndCleanUp",
		[]interface{}{state, next},
	)
}

func (j *jsiiProxy_IRunnerRuntimeParameters) GithubDomainPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"githubDomainPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerRuntimeParameters) LabelsPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelsPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerRuntimeParameters) OwnerPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerRuntimeParameters) RegistrationUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registrationUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerRuntimeParameters) RepoPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerRuntimeParameters) RunnerNamePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runnerNamePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerRuntimeParameters) RunnerTokenPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runnerTokenPath",
		&returns,
	)
	return returns
}

