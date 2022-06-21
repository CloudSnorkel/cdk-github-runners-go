package cloudsnorkelcdkgithubrunners

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.Architecture",
		reflect.TypeOf((*Architecture)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "is", GoMethod: "Is"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_Architecture{}
		},
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.CodeBuildImageBuilder",
		reflect.TypeOf((*CodeBuildImageBuilder)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addFiles", GoMethod: "AddFiles"},
			_jsii_.MemberMethod{JsiiMethod: "addPolicyStatement", GoMethod: "AddPolicyStatement"},
			_jsii_.MemberMethod{JsiiMethod: "addPostBuildCommand", GoMethod: "AddPostBuildCommand"},
			_jsii_.MemberMethod{JsiiMethod: "addPreBuildCommand", GoMethod: "AddPreBuildCommand"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberMethod{JsiiMethod: "setBuildArg", GoMethod: "SetBuildArg"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CodeBuildImageBuilder{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IImageBuilder)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.CodeBuildImageBuilderProps",
		reflect.TypeOf((*CodeBuildImageBuilderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.CodeBuildRunner",
		reflect.TypeOf((*CodeBuildRunner)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberMethod{JsiiMethod: "getStepFunctionTask", GoMethod: "GetStepFunctionTask"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroup", GoGetter: "SecurityGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_CodeBuildRunner{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRunnerProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.CodeBuildRunnerProps",
		reflect.TypeOf((*CodeBuildRunnerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.FargateRunner",
		reflect.TypeOf((*FargateRunner)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "assignPublicIp", GoGetter: "AssignPublicIp"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberMethod{JsiiMethod: "getStepFunctionTask", GoMethod: "GetStepFunctionTask"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroup", GoGetter: "SecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "spot", GoGetter: "Spot"},
			_jsii_.MemberProperty{JsiiProperty: "task", GoGetter: "Task"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_FargateRunner{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRunnerProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.FargateRunnerProps",
		reflect.TypeOf((*FargateRunnerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.GitHubRunners",
		reflect.TypeOf((*GitHubRunners)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "providers", GoGetter: "Providers"},
			_jsii_.MemberProperty{JsiiProperty: "secrets", GoGetter: "Secrets"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GitHubRunners{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.GitHubRunnersProps",
		reflect.TypeOf((*GitHubRunnersProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@cloudsnorkel/cdk-github-runners.IImageBuilder",
		reflect.TypeOf((*IImageBuilder)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IImageBuilder{}
		},
	)
	_jsii_.RegisterInterface(
		"@cloudsnorkel/cdk-github-runners.IRunnerProvider",
		reflect.TypeOf((*IRunnerProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberMethod{JsiiMethod: "getStepFunctionTask", GoMethod: "GetStepFunctionTask"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroup", GoGetter: "SecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_IRunnerProvider{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.LambdaRunner",
		reflect.TypeOf((*LambdaRunner)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "function", GoGetter: "Function"},
			_jsii_.MemberMethod{JsiiMethod: "getStepFunctionTask", GoMethod: "GetStepFunctionTask"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroup", GoGetter: "SecurityGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaRunner{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRunnerProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.LambdaRunnerProps",
		reflect.TypeOf((*LambdaRunnerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.Os",
		reflect.TypeOf((*Os)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "is", GoMethod: "Is"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_Os{}
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.RunnerImage",
		reflect.TypeOf((*RunnerImage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.RunnerProviderProps",
		reflect.TypeOf((*RunnerProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.RunnerRuntimeParameters",
		reflect.TypeOf((*RunnerRuntimeParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.RunnerVersion",
		reflect.TypeOf((*RunnerVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_RunnerVersion{}
		},
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.Secrets",
		reflect.TypeOf((*Secrets)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "github", GoGetter: "Github"},
			_jsii_.MemberProperty{JsiiProperty: "githubPrivateKey", GoGetter: "GithubPrivateKey"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "setup", GoGetter: "Setup"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "webhook", GoGetter: "Webhook"},
		},
		func() interface{} {
			j := jsiiProxy_Secrets{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.StaticRunnerImage",
		reflect.TypeOf((*StaticRunnerImage)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_StaticRunnerImage{}
		},
	)
}
