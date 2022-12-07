// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Components for Windows that can be used with AWS Image Builder based builders.
//
// These cannot be used by {@link CodeBuildImageBuilder}.
// Experimental.
type WindowsComponents interface {
}

// The jsii proxy struct for WindowsComponents
type jsiiProxy_WindowsComponents struct {
	_ byte // padding
}

// Experimental.
func NewWindowsComponents() WindowsComponents {
	_init_.Initialize()

	j := jsiiProxy_WindowsComponents{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.WindowsComponents",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewWindowsComponents_Override(w WindowsComponents) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.WindowsComponents",
		nil, // no parameters
		w,
	)
}

// Experimental.
func WindowsComponents_AwsCli(scope constructs.Construct, id *string) ImageBuilderComponent {
	_init_.Initialize()

	if err := validateWindowsComponents_AwsCliParameters(scope, id); err != nil {
		panic(err)
	}
	var returns ImageBuilderComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.WindowsComponents",
		"awsCli",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

// Experimental.
func WindowsComponents_CloudwatchAgent(scope constructs.Construct, id *string) ImageBuilderComponent {
	_init_.Initialize()

	if err := validateWindowsComponents_CloudwatchAgentParameters(scope, id); err != nil {
		panic(err)
	}
	var returns ImageBuilderComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.WindowsComponents",
		"cloudwatchAgent",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

// Experimental.
func WindowsComponents_Docker(scope constructs.Construct, id *string) ImageBuilderComponent {
	_init_.Initialize()

	if err := validateWindowsComponents_DockerParameters(scope, id); err != nil {
		panic(err)
	}
	var returns ImageBuilderComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.WindowsComponents",
		"docker",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

// Experimental.
func WindowsComponents_ExtraCertificates(scope constructs.Construct, id *string, path *string) ImageBuilderComponent {
	_init_.Initialize()

	if err := validateWindowsComponents_ExtraCertificatesParameters(scope, id, path); err != nil {
		panic(err)
	}
	var returns ImageBuilderComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.WindowsComponents",
		"extraCertificates",
		[]interface{}{scope, id, path},
		&returns,
	)

	return returns
}

// Experimental.
func WindowsComponents_Git(scope constructs.Construct, id *string) ImageBuilderComponent {
	_init_.Initialize()

	if err := validateWindowsComponents_GitParameters(scope, id); err != nil {
		panic(err)
	}
	var returns ImageBuilderComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.WindowsComponents",
		"git",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

// Experimental.
func WindowsComponents_GithubCli(scope constructs.Construct, id *string) ImageBuilderComponent {
	_init_.Initialize()

	if err := validateWindowsComponents_GithubCliParameters(scope, id); err != nil {
		panic(err)
	}
	var returns ImageBuilderComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.WindowsComponents",
		"githubCli",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

// Experimental.
func WindowsComponents_GithubRunner(scope constructs.Construct, id *string, runnerVersion RunnerVersion) ImageBuilderComponent {
	_init_.Initialize()

	if err := validateWindowsComponents_GithubRunnerParameters(scope, id, runnerVersion); err != nil {
		panic(err)
	}
	var returns ImageBuilderComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.WindowsComponents",
		"githubRunner",
		[]interface{}{scope, id, runnerVersion},
		&returns,
	)

	return returns
}

