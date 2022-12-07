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
			_jsii_.MemberMethod{JsiiMethod: "instanceTypeMatch", GoMethod: "InstanceTypeMatch"},
			_jsii_.MemberMethod{JsiiMethod: "is", GoMethod: "Is"},
			_jsii_.MemberMethod{JsiiMethod: "isIn", GoMethod: "IsIn"},
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
			_jsii_.MemberMethod{JsiiMethod: "addExtraCertificates", GoMethod: "AddExtraCertificates"},
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
			_jsii_.MemberMethod{JsiiMethod: "grantStateMachine", GoMethod: "GrantStateMachine"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberMethod{JsiiMethod: "labelsFromProperties", GoMethod: "LabelsFromProperties"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "status", GoMethod: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
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
		"@cloudsnorkel/cdk-github-runners.ContainerImageBuilder",
		reflect.TypeOf((*ContainerImageBuilder)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addComponent", GoMethod: "AddComponent"},
			_jsii_.MemberMethod{JsiiMethod: "addExtraCertificates", GoMethod: "AddExtraCertificates"},
			_jsii_.MemberProperty{JsiiProperty: "architecture", GoGetter: "Architecture"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "components", GoGetter: "Components"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberMethod{JsiiMethod: "createImage", GoMethod: "CreateImage"},
			_jsii_.MemberMethod{JsiiMethod: "createInfrastructure", GoMethod: "CreateInfrastructure"},
			_jsii_.MemberMethod{JsiiMethod: "createLog", GoMethod: "CreateLog"},
			_jsii_.MemberMethod{JsiiMethod: "createPipeline", GoMethod: "CreatePipeline"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "os", GoGetter: "Os"},
			_jsii_.MemberProperty{JsiiProperty: "platform", GoGetter: "Platform"},
			_jsii_.MemberMethod{JsiiMethod: "prependComponent", GoMethod: "PrependComponent"},
			_jsii_.MemberProperty{JsiiProperty: "repository", GoGetter: "Repository"},
			_jsii_.MemberProperty{JsiiProperty: "runnerVersion", GoGetter: "RunnerVersion"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ContainerImageBuilder{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IImageBuilder)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.ContainerImageBuilderProps",
		reflect.TypeOf((*ContainerImageBuilderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.Ec2Runner",
		reflect.TypeOf((*Ec2Runner)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberMethod{JsiiMethod: "getStepFunctionTask", GoMethod: "GetStepFunctionTask"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantStateMachine", GoMethod: "GrantStateMachine"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberMethod{JsiiMethod: "labelsFromProperties", GoMethod: "LabelsFromProperties"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "status", GoMethod: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2Runner{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRunnerProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.Ec2RunnerProps",
		reflect.TypeOf((*Ec2RunnerProps)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "grantStateMachine", GoMethod: "GrantStateMachine"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberMethod{JsiiMethod: "labelsFromProperties", GoMethod: "LabelsFromProperties"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "spot", GoGetter: "Spot"},
			_jsii_.MemberMethod{JsiiMethod: "status", GoMethod: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "subnetSelection", GoGetter: "SubnetSelection"},
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
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
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
		"@cloudsnorkel/cdk-github-runners.IAmiBuilder",
		reflect.TypeOf((*IAmiBuilder)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IAmiBuilder{}
		},
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
		"@cloudsnorkel/cdk-github-runners.IRunnerAmiStatus",
		reflect.TypeOf((*IRunnerAmiStatus)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "amiBuilderLogGroup", GoGetter: "AmiBuilderLogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplate", GoGetter: "LaunchTemplate"},
		},
		func() interface{} {
			return &jsiiProxy_IRunnerAmiStatus{}
		},
	)
	_jsii_.RegisterInterface(
		"@cloudsnorkel/cdk-github-runners.IRunnerImageStatus",
		reflect.TypeOf((*IRunnerImageStatus)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "imageBuilderLogGroup", GoGetter: "ImageBuilderLogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "imageRepository", GoGetter: "ImageRepository"},
			_jsii_.MemberProperty{JsiiProperty: "imageTag", GoGetter: "ImageTag"},
		},
		func() interface{} {
			return &jsiiProxy_IRunnerImageStatus{}
		},
	)
	_jsii_.RegisterInterface(
		"@cloudsnorkel/cdk-github-runners.IRunnerProvider",
		reflect.TypeOf((*IRunnerProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberMethod{JsiiMethod: "getStepFunctionTask", GoMethod: "GetStepFunctionTask"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantStateMachine", GoMethod: "GrantStateMachine"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberMethod{JsiiMethod: "status", GoMethod: "Status"},
		},
		func() interface{} {
			j := jsiiProxy_IRunnerProvider{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cloudsnorkel/cdk-github-runners.IRunnerProviderStatus",
		reflect.TypeOf((*IRunnerProviderStatus)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ami", GoGetter: "Ami"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "vpcArn", GoGetter: "VpcArn"},
		},
		func() interface{} {
			return &jsiiProxy_IRunnerProviderStatus{}
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.ImageBuilderAsset",
		reflect.TypeOf((*ImageBuilderAsset)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.ImageBuilderComponent",
		reflect.TypeOf((*ImageBuilderComponent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantAssetsRead", GoMethod: "GrantAssetsRead"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "platform", GoGetter: "Platform"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "version", GoMethod: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_ImageBuilderComponent{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.ImageBuilderComponentProperties",
		reflect.TypeOf((*ImageBuilderComponentProperties)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.LambdaRunner",
		reflect.TypeOf((*LambdaRunner)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "function", GoGetter: "Function"},
			_jsii_.MemberMethod{JsiiMethod: "getStepFunctionTask", GoMethod: "GetStepFunctionTask"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantStateMachine", GoMethod: "GrantStateMachine"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberMethod{JsiiMethod: "labelsFromProperties", GoMethod: "LabelsFromProperties"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "status", GoMethod: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
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
		"@cloudsnorkel/cdk-github-runners.LinuxUbuntuComponents",
		reflect.TypeOf((*LinuxUbuntuComponents)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_LinuxUbuntuComponents{}
		},
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.Os",
		reflect.TypeOf((*Os)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "is", GoMethod: "Is"},
			_jsii_.MemberMethod{JsiiMethod: "isIn", GoMethod: "IsIn"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_Os{}
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.RunnerAmi",
		reflect.TypeOf((*RunnerAmi)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "is", GoMethod: "Is"},
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
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.WindowsComponents",
		reflect.TypeOf((*WindowsComponents)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_WindowsComponents{}
		},
	)
}
