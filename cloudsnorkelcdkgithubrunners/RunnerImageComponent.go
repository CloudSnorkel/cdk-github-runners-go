// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"
)

// Components are used to build runner images.
//
// They can run commands in the image, copy files into the image, and run some Docker commands.
// Experimental.
type RunnerImageComponent interface {
	// Component name.
	//
	// Used to identify component in image build logs, and for {@link RunnerImageBuilder.removeComponent }
	// Experimental.
	Name() *string
	// Returns assets to copy into the built image.
	//
	// Can be used to copy files into the image.
	// Experimental.
	GetAssets(_os Os, _architecture Architecture) *[]*RunnerImageAsset
	// Returns commands to run to in built image.
	//
	// Can be used to install packages, setup build prerequisites, etc.
	// Experimental.
	GetCommands(_os Os, _architecture Architecture) *[]*string
	// Returns Docker commands to run to in built image.
	//
	// Can be used to add commands like `VOLUME`, `ENTRYPOINT`, `CMD`, etc.
	//
	// Docker commands are added after assets and normal commands.
	// Experimental.
	GetDockerCommands(_os Os, _architecture Architecture) *[]*string
	// Returns true if the image builder should be rebooted after this component is installed.
	// Experimental.
	ShouldReboot(_os Os, _architecture Architecture) *bool
}

// The jsii proxy struct for RunnerImageComponent
type jsiiProxy_RunnerImageComponent struct {
	_ byte // padding
}

func (j *jsiiProxy_RunnerImageComponent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewRunnerImageComponent_Override(r RunnerImageComponent) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.RunnerImageComponent",
		nil, // no parameters
		r,
	)
}

// A component to install the AWS CLI.
// Experimental.
func RunnerImageComponent_AwsCli() RunnerImageComponent {
	_init_.Initialize()

	var returns RunnerImageComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.RunnerImageComponent",
		"awsCli",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Define a custom component that can run commands in the image, copy files into the image, and run some Docker commands.
//
// The order of operations is (1) assets (2) commands (3) docker commands.
//
// Use this to customize the image for the runner.
//
// **WARNING:** Docker commands are not guaranteed to be included before the next component.
// Experimental.
func RunnerImageComponent_Custom(props *RunnerImageComponentCustomProps) RunnerImageComponent {
	_init_.Initialize()

	if err := validateRunnerImageComponent_CustomParameters(props); err != nil {
		panic(err)
	}
	var returns RunnerImageComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.RunnerImageComponent",
		"custom",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// A component to install Docker.
//
// On Windows this sets up dockerd for Windows containers without Docker Desktop. If you need Linux containers on Windows, you'll need to install Docker Desktop which doesn't seem to play well with servers (PRs welcome).
// Experimental.
func RunnerImageComponent_Docker() RunnerImageComponent {
	_init_.Initialize()

	var returns RunnerImageComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.RunnerImageComponent",
		"docker",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A component to install Docker-in-Docker.
// Experimental.
func RunnerImageComponent_DockerInDocker() RunnerImageComponent {
	_init_.Initialize()

	var returns RunnerImageComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.RunnerImageComponent",
		"dockerInDocker",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A component to add a trusted certificate authority.
//
// This can be used to support GitHub Enterprise Server with self-signed certificate.
// Experimental.
func RunnerImageComponent_ExtraCertificates(source *string, name *string) RunnerImageComponent {
	_init_.Initialize()

	if err := validateRunnerImageComponent_ExtraCertificatesParameters(source, name); err != nil {
		panic(err)
	}
	var returns RunnerImageComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.RunnerImageComponent",
		"extraCertificates",
		[]interface{}{source, name},
		&returns,
	)

	return returns
}

// A component to install the GitHub CLI.
// Experimental.
func RunnerImageComponent_Git() RunnerImageComponent {
	_init_.Initialize()

	var returns RunnerImageComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.RunnerImageComponent",
		"git",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A component to install the GitHub CLI.
// Experimental.
func RunnerImageComponent_GithubCli() RunnerImageComponent {
	_init_.Initialize()

	var returns RunnerImageComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.RunnerImageComponent",
		"githubCli",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A component to install the GitHub Actions Runner.
//
// This is the actual executable that connects to GitHub to ask for jobs and then execute them.
// Experimental.
func RunnerImageComponent_GithubRunner(runnerVersion RunnerVersion) RunnerImageComponent {
	_init_.Initialize()

	if err := validateRunnerImageComponent_GithubRunnerParameters(runnerVersion); err != nil {
		panic(err)
	}
	var returns RunnerImageComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.RunnerImageComponent",
		"githubRunner",
		[]interface{}{runnerVersion},
		&returns,
	)

	return returns
}

// A component to set up the required Lambda entrypoint for Lambda runners.
// Experimental.
func RunnerImageComponent_LambdaEntrypoint() RunnerImageComponent {
	_init_.Initialize()

	var returns RunnerImageComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.RunnerImageComponent",
		"lambdaEntrypoint",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A component to install the required packages for the runner.
// Experimental.
func RunnerImageComponent_RequiredPackages() RunnerImageComponent {
	_init_.Initialize()

	var returns RunnerImageComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.RunnerImageComponent",
		"requiredPackages",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A component to prepare the required runner user.
// Experimental.
func RunnerImageComponent_RunnerUser() RunnerImageComponent {
	_init_.Initialize()

	var returns RunnerImageComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.RunnerImageComponent",
		"runnerUser",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerImageComponent) GetAssets(_os Os, _architecture Architecture) *[]*RunnerImageAsset {
	if err := r.validateGetAssetsParameters(_os, _architecture); err != nil {
		panic(err)
	}
	var returns *[]*RunnerImageAsset

	_jsii_.Invoke(
		r,
		"getAssets",
		[]interface{}{_os, _architecture},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerImageComponent) GetCommands(_os Os, _architecture Architecture) *[]*string {
	if err := r.validateGetCommandsParameters(_os, _architecture); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getCommands",
		[]interface{}{_os, _architecture},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerImageComponent) GetDockerCommands(_os Os, _architecture Architecture) *[]*string {
	if err := r.validateGetDockerCommandsParameters(_os, _architecture); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getDockerCommands",
		[]interface{}{_os, _architecture},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerImageComponent) ShouldReboot(_os Os, _architecture Architecture) *bool {
	if err := r.validateShouldRebootParameters(_os, _architecture); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		r,
		"shouldReboot",
		[]interface{}{_os, _architecture},
		&returns,
	)

	return returns
}

