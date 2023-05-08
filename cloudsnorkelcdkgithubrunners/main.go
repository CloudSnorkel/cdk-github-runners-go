package cloudsnorkelcdkgithubrunners

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.AmiBuilder",
		reflect.TypeOf((*AmiBuilder)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addComponent", GoMethod: "AddComponent"},
			_jsii_.MemberMethod{JsiiMethod: "addExtraCertificates", GoMethod: "AddExtraCertificates"},
			_jsii_.MemberProperty{JsiiProperty: "architecture", GoGetter: "Architecture"},
			_jsii_.MemberMethod{JsiiMethod: "bindAmi", GoMethod: "BindAmi"},
			_jsii_.MemberMethod{JsiiMethod: "bindDockerImage", GoMethod: "BindDockerImage"},
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
			_jsii_.MemberProperty{JsiiProperty: "runnerVersion", GoGetter: "RunnerVersion"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AmiBuilder{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRunnerImageBuilder)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.AmiBuilderProps",
		reflect.TypeOf((*AmiBuilderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.ApiGatewayAccessProps",
		reflect.TypeOf((*ApiGatewayAccessProps)(nil)).Elem(),
	)
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
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.AwsImageBuilderRunnerImageBuilderProps",
		reflect.TypeOf((*AwsImageBuilderRunnerImageBuilderProps)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "bindAmi", GoMethod: "BindAmi"},
			_jsii_.MemberMethod{JsiiMethod: "bindDockerImage", GoMethod: "BindDockerImage"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberMethod{JsiiMethod: "setBuildArg", GoMethod: "SetBuildArg"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CodeBuildImageBuilder{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRunnerImageBuilder)
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
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "retryableErrors", GoGetter: "RetryableErrors"},
			_jsii_.MemberMethod{JsiiMethod: "status", GoMethod: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CodeBuildRunner{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CodeBuildRunnerProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.CodeBuildRunnerImageBuilderProps",
		reflect.TypeOf((*CodeBuildRunnerImageBuilderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.CodeBuildRunnerProvider",
		reflect.TypeOf((*CodeBuildRunnerProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberMethod{JsiiMethod: "getStepFunctionTask", GoMethod: "GetStepFunctionTask"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantStateMachine", GoMethod: "GrantStateMachine"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberMethod{JsiiMethod: "labelsFromProperties", GoMethod: "LabelsFromProperties"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "retryableErrors", GoGetter: "RetryableErrors"},
			_jsii_.MemberMethod{JsiiMethod: "status", GoMethod: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CodeBuildRunnerProvider{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRunnerProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.CodeBuildRunnerProviderProps",
		reflect.TypeOf((*CodeBuildRunnerProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.ContainerImageBuilder",
		reflect.TypeOf((*ContainerImageBuilder)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addComponent", GoMethod: "AddComponent"},
			_jsii_.MemberMethod{JsiiMethod: "addExtraCertificates", GoMethod: "AddExtraCertificates"},
			_jsii_.MemberProperty{JsiiProperty: "architecture", GoGetter: "Architecture"},
			_jsii_.MemberMethod{JsiiMethod: "bindAmi", GoMethod: "BindAmi"},
			_jsii_.MemberMethod{JsiiMethod: "bindDockerImage", GoMethod: "BindDockerImage"},
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
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRunnerImageBuilder)
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
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "retryableErrors", GoGetter: "RetryableErrors"},
			_jsii_.MemberMethod{JsiiMethod: "status", GoMethod: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2Runner{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Ec2RunnerProvider)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.Ec2RunnerProvider",
		reflect.TypeOf((*Ec2RunnerProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberMethod{JsiiMethod: "getStepFunctionTask", GoMethod: "GetStepFunctionTask"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantStateMachine", GoMethod: "GrantStateMachine"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberMethod{JsiiMethod: "labelsFromProperties", GoMethod: "LabelsFromProperties"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "retryableErrors", GoGetter: "RetryableErrors"},
			_jsii_.MemberMethod{JsiiMethod: "status", GoMethod: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2RunnerProvider{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRunnerProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.Ec2RunnerProviderProps",
		reflect.TypeOf((*Ec2RunnerProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.EcsRunnerProvider",
		reflect.TypeOf((*EcsRunnerProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberMethod{JsiiMethod: "getStepFunctionTask", GoMethod: "GetStepFunctionTask"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantStateMachine", GoMethod: "GrantStateMachine"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberMethod{JsiiMethod: "labelsFromProperties", GoMethod: "LabelsFromProperties"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "retryableErrors", GoGetter: "RetryableErrors"},
			_jsii_.MemberMethod{JsiiMethod: "status", GoMethod: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EcsRunnerProvider{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRunnerProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.EcsRunnerProviderProps",
		reflect.TypeOf((*EcsRunnerProviderProps)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "retryableErrors", GoGetter: "RetryableErrors"},
			_jsii_.MemberProperty{JsiiProperty: "spot", GoGetter: "Spot"},
			_jsii_.MemberMethod{JsiiMethod: "status", GoMethod: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "subnetSelection", GoGetter: "SubnetSelection"},
			_jsii_.MemberProperty{JsiiProperty: "task", GoGetter: "Task"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_FargateRunner{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FargateRunnerProvider)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.FargateRunnerProvider",
		reflect.TypeOf((*FargateRunnerProvider)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "retryableErrors", GoGetter: "RetryableErrors"},
			_jsii_.MemberProperty{JsiiProperty: "spot", GoGetter: "Spot"},
			_jsii_.MemberMethod{JsiiMethod: "status", GoMethod: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "subnetSelection", GoGetter: "SubnetSelection"},
			_jsii_.MemberProperty{JsiiProperty: "task", GoGetter: "Task"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_FargateRunnerProvider{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRunnerProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.FargateRunnerProviderProps",
		reflect.TypeOf((*FargateRunnerProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.GitHubRunners",
		reflect.TypeOf((*GitHubRunners)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "failedImageBuildsTopic", GoMethod: "FailedImageBuildsTopic"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricJobCompleted", GoMethod: "MetricJobCompleted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
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
		"@cloudsnorkel/cdk-github-runners.IRunnerImageBuilder",
		reflect.TypeOf((*IRunnerImageBuilder)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bindAmi", GoMethod: "BindAmi"},
			_jsii_.MemberMethod{JsiiMethod: "bindDockerImage", GoMethod: "BindDockerImage"},
		},
		func() interface{} {
			return &jsiiProxy_IRunnerImageBuilder{}
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
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "retryableErrors", GoGetter: "RetryableErrors"},
			_jsii_.MemberMethod{JsiiMethod: "status", GoMethod: "Status"},
		},
		func() interface{} {
			j := jsiiProxy_IRunnerProvider{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
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
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
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
			_jsii_.MemberMethod{JsiiMethod: "prefixCommandsWithErrorHandling", GoMethod: "PrefixCommandsWithErrorHandling"},
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
		"@cloudsnorkel/cdk-github-runners.LambdaAccess",
		reflect.TypeOf((*LambdaAccess)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_LambdaAccess{}
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
			_jsii_.MemberMethod{JsiiMethod: "grantStateMachine", GoMethod: "GrantStateMachine"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberMethod{JsiiMethod: "labelsFromProperties", GoMethod: "LabelsFromProperties"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "retryableErrors", GoGetter: "RetryableErrors"},
			_jsii_.MemberMethod{JsiiMethod: "status", GoMethod: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaRunner{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_LambdaRunnerProvider)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.LambdaRunnerProvider",
		reflect.TypeOf((*LambdaRunnerProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "function", GoGetter: "Function"},
			_jsii_.MemberMethod{JsiiMethod: "getStepFunctionTask", GoMethod: "GetStepFunctionTask"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantStateMachine", GoMethod: "GrantStateMachine"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberMethod{JsiiMethod: "labelsFromProperties", GoMethod: "LabelsFromProperties"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "retryableErrors", GoGetter: "RetryableErrors"},
			_jsii_.MemberMethod{JsiiMethod: "status", GoMethod: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaRunnerProvider{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRunnerProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.LambdaRunnerProviderProps",
		reflect.TypeOf((*LambdaRunnerProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.LinuxUbuntuComponents",
		reflect.TypeOf((*LinuxUbuntuComponents)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_LinuxUbuntuComponents{}
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.LogOptions",
		reflect.TypeOf((*LogOptions)(nil)).Elem(),
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
		"@cloudsnorkel/cdk-github-runners.ProviderRetryOptions",
		reflect.TypeOf((*ProviderRetryOptions)(nil)).Elem(),
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
		"@cloudsnorkel/cdk-github-runners.RunnerImageAsset",
		reflect.TypeOf((*RunnerImageAsset)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.RunnerImageBuilder",
		reflect.TypeOf((*RunnerImageBuilder)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addComponent", GoMethod: "AddComponent"},
			_jsii_.MemberMethod{JsiiMethod: "bindAmi", GoMethod: "BindAmi"},
			_jsii_.MemberMethod{JsiiMethod: "bindDockerImage", GoMethod: "BindDockerImage"},
			_jsii_.MemberProperty{JsiiProperty: "components", GoGetter: "Components"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "removeComponent", GoMethod: "RemoveComponent"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RunnerImageBuilder{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRunnerImageBuilder)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.RunnerImageBuilderProps",
		reflect.TypeOf((*RunnerImageBuilderProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cloudsnorkel/cdk-github-runners.RunnerImageBuilderType",
		reflect.TypeOf((*RunnerImageBuilderType)(nil)).Elem(),
		map[string]interface{}{
			"CODE_BUILD": RunnerImageBuilderType_CODE_BUILD,
			"AWS_IMAGE_BUILDER": RunnerImageBuilderType_AWS_IMAGE_BUILDER,
		},
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-github-runners.RunnerImageComponent",
		reflect.TypeOf((*RunnerImageComponent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAssets", GoMethod: "GetAssets"},
			_jsii_.MemberMethod{JsiiMethod: "getCommands", GoMethod: "GetCommands"},
			_jsii_.MemberMethod{JsiiMethod: "getDockerCommands", GoMethod: "GetDockerCommands"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "shouldReboot", GoMethod: "ShouldReboot"},
		},
		func() interface{} {
			return &jsiiProxy_RunnerImageComponent{}
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-github-runners.RunnerImageComponentCustomProps",
		reflect.TypeOf((*RunnerImageComponentCustomProps)(nil)).Elem(),
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
