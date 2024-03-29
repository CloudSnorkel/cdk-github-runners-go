package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Components for Ubuntu Linux that can be used with AWS Image Builder based builders.
//
// These cannot be used by {@link CodeBuildImageBuilder }.
// Deprecated: Use `RunnerImageComponent` instead.
type LinuxUbuntuComponents interface {
}

// The jsii proxy struct for LinuxUbuntuComponents
type jsiiProxy_LinuxUbuntuComponents struct {
	_ byte // padding
}

// Deprecated: Use `RunnerImageComponent` instead.
func NewLinuxUbuntuComponents() LinuxUbuntuComponents {
	_init_.Initialize()

	j := jsiiProxy_LinuxUbuntuComponents{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.LinuxUbuntuComponents",
		nil, // no parameters
		&j,
	)

	return &j
}

// Deprecated: Use `RunnerImageComponent` instead.
func NewLinuxUbuntuComponents_Override(l LinuxUbuntuComponents) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.LinuxUbuntuComponents",
		nil, // no parameters
		l,
	)
}

// Deprecated: Use `RunnerImageComponent` instead.
func LinuxUbuntuComponents_AwsCli(scope constructs.Construct, id *string, architecture Architecture) ImageBuilderComponent {
	_init_.Initialize()

	if err := validateLinuxUbuntuComponents_AwsCliParameters(scope, id, architecture); err != nil {
		panic(err)
	}
	var returns ImageBuilderComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.LinuxUbuntuComponents",
		"awsCli",
		[]interface{}{scope, id, architecture},
		&returns,
	)

	return returns
}

// Deprecated: Use `RunnerImageComponent` instead.
func LinuxUbuntuComponents_Docker(scope constructs.Construct, id *string, _architecture Architecture) ImageBuilderComponent {
	_init_.Initialize()

	if err := validateLinuxUbuntuComponents_DockerParameters(scope, id, _architecture); err != nil {
		panic(err)
	}
	var returns ImageBuilderComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.LinuxUbuntuComponents",
		"docker",
		[]interface{}{scope, id, _architecture},
		&returns,
	)

	return returns
}

// Deprecated: Use `RunnerImageComponent` instead.
func LinuxUbuntuComponents_ExtraCertificates(scope constructs.Construct, id *string, path *string) ImageBuilderComponent {
	_init_.Initialize()

	if err := validateLinuxUbuntuComponents_ExtraCertificatesParameters(scope, id, path); err != nil {
		panic(err)
	}
	var returns ImageBuilderComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.LinuxUbuntuComponents",
		"extraCertificates",
		[]interface{}{scope, id, path},
		&returns,
	)

	return returns
}

// Deprecated: Use `RunnerImageComponent` instead.
func LinuxUbuntuComponents_Git(scope constructs.Construct, id *string, _architecture Architecture) ImageBuilderComponent {
	_init_.Initialize()

	if err := validateLinuxUbuntuComponents_GitParameters(scope, id, _architecture); err != nil {
		panic(err)
	}
	var returns ImageBuilderComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.LinuxUbuntuComponents",
		"git",
		[]interface{}{scope, id, _architecture},
		&returns,
	)

	return returns
}

// Deprecated: Use `RunnerImageComponent` instead.
func LinuxUbuntuComponents_GithubCli(scope constructs.Construct, id *string, _architecture Architecture) ImageBuilderComponent {
	_init_.Initialize()

	if err := validateLinuxUbuntuComponents_GithubCliParameters(scope, id, _architecture); err != nil {
		panic(err)
	}
	var returns ImageBuilderComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.LinuxUbuntuComponents",
		"githubCli",
		[]interface{}{scope, id, _architecture},
		&returns,
	)

	return returns
}

// Deprecated: Use `RunnerImageComponent` instead.
func LinuxUbuntuComponents_GithubRunner(scope constructs.Construct, id *string, runnerVersion RunnerVersion, architecture Architecture) ImageBuilderComponent {
	_init_.Initialize()

	if err := validateLinuxUbuntuComponents_GithubRunnerParameters(scope, id, runnerVersion, architecture); err != nil {
		panic(err)
	}
	var returns ImageBuilderComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.LinuxUbuntuComponents",
		"githubRunner",
		[]interface{}{scope, id, runnerVersion, architecture},
		&returns,
	)

	return returns
}

// Deprecated: Use `RunnerImageComponent` instead.
func LinuxUbuntuComponents_RequiredPackages(scope constructs.Construct, id *string, architecture Architecture) ImageBuilderComponent {
	_init_.Initialize()

	if err := validateLinuxUbuntuComponents_RequiredPackagesParameters(scope, id, architecture); err != nil {
		panic(err)
	}
	var returns ImageBuilderComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.LinuxUbuntuComponents",
		"requiredPackages",
		[]interface{}{scope, id, architecture},
		&returns,
	)

	return returns
}

// Deprecated: Use `RunnerImageComponent` instead.
func LinuxUbuntuComponents_RunnerUser(scope constructs.Construct, id *string, _architecture Architecture) ImageBuilderComponent {
	_init_.Initialize()

	if err := validateLinuxUbuntuComponents_RunnerUserParameters(scope, id, _architecture); err != nil {
		panic(err)
	}
	var returns ImageBuilderComponent

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.LinuxUbuntuComponents",
		"runnerUser",
		[]interface{}{scope, id, _architecture},
		&returns,
	)

	return returns
}

