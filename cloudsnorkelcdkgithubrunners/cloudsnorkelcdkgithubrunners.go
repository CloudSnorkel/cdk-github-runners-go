// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

// CPU architecture enum for an image.
// Experimental.
type Architecture interface {
	// Experimental.
	Name() *string
	// Checks if the given architecture is the same as this one.
	// Experimental.
	Is(arch Architecture) *bool
}

// The jsii proxy struct for Architecture
type jsiiProxy_Architecture struct {
	_ byte // padding
}

func (j *jsiiProxy_Architecture) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func Architecture_ARM64() Architecture {
	_init_.Initialize()
	var returns Architecture
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.Architecture",
		"ARM64",
		&returns,
	)
	return returns
}

func Architecture_X86_64() Architecture {
	_init_.Initialize()
	var returns Architecture
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.Architecture",
		"X86_64",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_Architecture) Is(arch Architecture) *bool {
	var returns *bool

	_jsii_.Invoke(
		a,
		"is",
		[]interface{}{arch},
		&returns,
	)

	return returns
}

// An image builder that uses CodeBuild to build Docker images pre-baked with all the GitHub Actions runner requirements.
//
// Builders can be used with runner providers.
//
// Each builder re-runs automatically at a set interval to make sure the images contain the latest versions of everything.
//
// You can create an instance of this construct to customize the image used to spin-up runners. Each provider has its own requirements for what an image should do. That's why they each provide their own Dockerfile.
//
// For example, to set a specific runner version, rebuild the image every 2 weeks, and add a few packages for the Fargate provider, use:
//
// ```
// const builder = new CodeBuildImageBuilder(this, 'Builder', {
//      dockerfilePath: FargateProvider.LINUX_X64_DOCKERFILE_PATH,
//      runnerVersion: RunnerVersion.specific('2.293.0'),
//      rebuildInterval: Duration.days(14),
// });
// builder.setBuildArg('EXTRA_PACKAGES', 'nginx xz-utils');
// new FargateProvider(this, 'Fargate provider', {
//      label: 'customized-fargate',
//      imageBuilder: builder,
// });
// ```.
// Experimental.
type CodeBuildImageBuilder interface {
	constructs.Construct
	IImageBuilder
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Props() *CodeBuildImageBuilderProps
	// Add extra trusted certificates. This helps deal with self-signed certificates for GitHub Enterprise Server.
	//
	// All first party Dockerfiles support this. Others may not.
	// Experimental.
	AddExtraCertificates(path *string)
	// Uploads a folder to the build server at a given folder name.
	// Experimental.
	AddFiles(sourcePath *string, destName *string)
	// Add a policy statement to the builder to access resources required to the image build.
	// Experimental.
	AddPolicyStatement(statement awsiam.PolicyStatement)
	// Adds a command that runs after `docker build` and `docker push`.
	// Experimental.
	AddPostBuildCommand(command *string)
	// Adds a command that runs before `docker build`.
	// Experimental.
	AddPreBuildCommand(command *string)
	// Called by IRunnerProvider to finalize settings and create the image builder.
	// Experimental.
	Bind() *RunnerImage
	// Adds a build argument for Docker.
	//
	// See the documentation for the Dockerfile you're using for a list of supported build arguments.
	// Experimental.
	SetBuildArg(name *string, value *string)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CodeBuildImageBuilder
type jsiiProxy_CodeBuildImageBuilder struct {
	internal.Type__constructsConstruct
	jsiiProxy_IImageBuilder
}

func (j *jsiiProxy_CodeBuildImageBuilder) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildImageBuilder) Props() *CodeBuildImageBuilderProps {
	var returns *CodeBuildImageBuilderProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Experimental.
func NewCodeBuildImageBuilder(scope constructs.Construct, id *string, props *CodeBuildImageBuilderProps) CodeBuildImageBuilder {
	_init_.Initialize()

	j := jsiiProxy_CodeBuildImageBuilder{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.CodeBuildImageBuilder",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCodeBuildImageBuilder_Override(c CodeBuildImageBuilder, scope constructs.Construct, id *string, props *CodeBuildImageBuilderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.CodeBuildImageBuilder",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CodeBuildImageBuilder_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.CodeBuildImageBuilder",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildImageBuilder) AddExtraCertificates(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addExtraCertificates",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CodeBuildImageBuilder) AddFiles(sourcePath *string, destName *string) {
	_jsii_.InvokeVoid(
		c,
		"addFiles",
		[]interface{}{sourcePath, destName},
	)
}

func (c *jsiiProxy_CodeBuildImageBuilder) AddPolicyStatement(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		c,
		"addPolicyStatement",
		[]interface{}{statement},
	)
}

func (c *jsiiProxy_CodeBuildImageBuilder) AddPostBuildCommand(command *string) {
	_jsii_.InvokeVoid(
		c,
		"addPostBuildCommand",
		[]interface{}{command},
	)
}

func (c *jsiiProxy_CodeBuildImageBuilder) AddPreBuildCommand(command *string) {
	_jsii_.InvokeVoid(
		c,
		"addPreBuildCommand",
		[]interface{}{command},
	)
}

func (c *jsiiProxy_CodeBuildImageBuilder) Bind() *RunnerImage {
	var returns *RunnerImage

	_jsii_.Invoke(
		c,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildImageBuilder) SetBuildArg(name *string, value *string) {
	_jsii_.InvokeVoid(
		c,
		"setBuildArg",
		[]interface{}{name, value},
	)
}

func (c *jsiiProxy_CodeBuildImageBuilder) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for CodeBuildImageBuilder construct.
// Experimental.
type CodeBuildImageBuilderProps struct {
	// Path to Dockerfile to be built.
	//
	// It can be a path to a Dockerfile, a folder containing a Dockerfile, or a zip file containing a Dockerfile.
	// Experimental.
	DockerfilePath *string `field:"required" json:"dockerfilePath" yaml:"dockerfilePath"`
	// Image architecture.
	// Experimental.
	Architecture Architecture `field:"optional" json:"architecture" yaml:"architecture"`
	// The type of compute to use for this build.
	//
	// See the {@link ComputeType} enum for the possible values.
	// Experimental.
	ComputeType awscodebuild.ComputeType `field:"optional" json:"computeType" yaml:"computeType"`
	// Removal policy for logs of image builds.
	//
	// If deployment fails on the custom resource, try setting this to `RemovalPolicy.RETAIN`. This way the CodeBuild logs can still be viewed, and you can see why the build failed.
	//
	// We try to not leave anything behind when removed. But sometimes a log staying behind is useful.
	// Experimental.
	LogRemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"logRemovalPolicy" yaml:"logRemovalPolicy"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// Image OS.
	// Experimental.
	Os Os `field:"optional" json:"os" yaml:"os"`
	// Schedule the image to be rebuilt every given interval.
	//
	// Useful for keeping the image up-do-date with the latest GitHub runner version and latest OS updates.
	//
	// Set to zero to disable.
	// Experimental.
	RebuildInterval awscdk.Duration `field:"optional" json:"rebuildInterval" yaml:"rebuildInterval"`
	// Version of GitHub Runners to install.
	// Experimental.
	RunnerVersion RunnerVersion `field:"optional" json:"runnerVersion" yaml:"runnerVersion"`
	// Security Group to assign to this instance.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Where to place the network interfaces within the VPC.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// VPC to launch the runners in.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

// GitHub Actions runner provider using CodeBuild to execute the actions.
//
// Creates a project that gets started for each job.
//
// This construct is not meant to be used by itself. It should be passed in the providers property for GitHubRunners.
// Experimental.
type CodeBuildRunner interface {
	constructs.Construct
	IRunnerProvider
	// The network connections associated with this resource.
	// Experimental.
	Connections() awsec2.Connections
	// Grant principal used to add permissions to the runner role.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// Label associated with this provider.
	// Experimental.
	Label() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// CodeBuild project hosting the runner.
	// Experimental.
	Project() awscodebuild.Project
	// Security group attached to the task.
	// Experimental.
	SecurityGroup() awsec2.ISecurityGroup
	// VPC used for hosting the project.
	// Experimental.
	Vpc() awsec2.IVpc
	// Generate step function task(s) to start a new runner.
	//
	// Called by GithubRunners and shouldn't be called manually.
	// Experimental.
	GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CodeBuildRunner
type jsiiProxy_CodeBuildRunner struct {
	internal.Type__constructsConstruct
	jsiiProxy_IRunnerProvider
}

func (j *jsiiProxy_CodeBuildRunner) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildRunner) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildRunner) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildRunner) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildRunner) Project() awscodebuild.Project {
	var returns awscodebuild.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildRunner) SecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildRunner) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Experimental.
func NewCodeBuildRunner(scope constructs.Construct, id *string, props *CodeBuildRunnerProps) CodeBuildRunner {
	_init_.Initialize()

	j := jsiiProxy_CodeBuildRunner{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.CodeBuildRunner",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCodeBuildRunner_Override(c CodeBuildRunner, scope constructs.Construct, id *string, props *CodeBuildRunnerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.CodeBuildRunner",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CodeBuildRunner_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.CodeBuildRunner",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CodeBuildRunner_LINUX_ARM64_DOCKERFILE_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.CodeBuildRunner",
		"LINUX_ARM64_DOCKERFILE_PATH",
		&returns,
	)
	return returns
}

func CodeBuildRunner_LINUX_X64_DOCKERFILE_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.CodeBuildRunner",
		"LINUX_X64_DOCKERFILE_PATH",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CodeBuildRunner) GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable {
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		c,
		"getStepFunctionTask",
		[]interface{}{parameters},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildRunner) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type CodeBuildRunnerProps struct {
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// The type of compute to use for this build.
	//
	// See the {@link ComputeType} enum for the possible values.
	// Experimental.
	ComputeType awscodebuild.ComputeType `field:"optional" json:"computeType" yaml:"computeType"`
	// Provider running an image to run inside CodeBuild with GitHub runner pre-configured.
	//
	// A user named `runner` is expected to exist with access to Docker-in-Docker.
	// Experimental.
	ImageBuilder IImageBuilder `field:"optional" json:"imageBuilder" yaml:"imageBuilder"`
	// GitHub Actions label used for this provider.
	// Experimental.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Security Group to assign to this instance.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Where to place the network interfaces within the VPC.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// VPC to launch the runners in.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

// GitHub Actions runner provider using Fargate to execute the actions.
//
// Creates a task definition with a single container that gets started for each job.
//
// This construct is not meant to be used by itself. It should be passed in the providers property for GitHubRunners.
// Experimental.
type FargateRunner interface {
	constructs.Construct
	IRunnerProvider
	// Whether task will have a public IP.
	// Experimental.
	AssignPublicIp() *bool
	// Cluster hosting the task hosting the runner.
	// Experimental.
	Cluster() awsecs.Cluster
	// The network connections associated with this resource.
	// Experimental.
	Connections() awsec2.Connections
	// Container definition hosting the runner.
	// Experimental.
	Container() awsecs.ContainerDefinition
	// Grant principal used to add permissions to the runner role.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// Label associated with this provider.
	// Experimental.
	Label() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Security group attached to the task.
	// Experimental.
	SecurityGroup() awsec2.ISecurityGroup
	// Use spot pricing for Fargate tasks.
	// Experimental.
	Spot() *bool
	// Fargate task hosting the runner.
	// Experimental.
	Task() awsecs.FargateTaskDefinition
	// VPC used for hosting the task.
	// Experimental.
	Vpc() awsec2.IVpc
	// Generate step function task(s) to start a new runner.
	//
	// Called by GithubRunners and shouldn't be called manually.
	// Experimental.
	GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FargateRunner
type jsiiProxy_FargateRunner struct {
	internal.Type__constructsConstruct
	jsiiProxy_IRunnerProvider
}

func (j *jsiiProxy_FargateRunner) AssignPublicIp() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) Cluster() awsecs.Cluster {
	var returns awsecs.Cluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) Container() awsecs.ContainerDefinition {
	var returns awsecs.ContainerDefinition
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) SecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) Spot() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"spot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) Task() awsecs.FargateTaskDefinition {
	var returns awsecs.FargateTaskDefinition
	_jsii_.Get(
		j,
		"task",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateRunner) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Experimental.
func NewFargateRunner(scope constructs.Construct, id *string, props *FargateRunnerProps) FargateRunner {
	_init_.Initialize()

	j := jsiiProxy_FargateRunner{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.FargateRunner",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewFargateRunner_Override(f FargateRunner, scope constructs.Construct, id *string, props *FargateRunnerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.FargateRunner",
		[]interface{}{scope, id, props},
		f,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func FargateRunner_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.FargateRunner",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FargateRunner_LINUX_ARM64_DOCKERFILE_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.FargateRunner",
		"LINUX_ARM64_DOCKERFILE_PATH",
		&returns,
	)
	return returns
}

func FargateRunner_LINUX_X64_DOCKERFILE_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.FargateRunner",
		"LINUX_X64_DOCKERFILE_PATH",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FargateRunner) GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable {
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		f,
		"getStepFunctionTask",
		[]interface{}{parameters},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateRunner) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for FargateRunner.
// Experimental.
type FargateRunnerProps struct {
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// Assign public IP to the runner task.
	//
	// Make sure the task will have access to GitHub. A public IP might be required unless you have NAT gateway.
	// Experimental.
	AssignPublicIp *bool `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Existing Fargate cluster to use.
	// Experimental.
	Cluster awsecs.Cluster `field:"optional" json:"cluster" yaml:"cluster"`
	// The number of cpu units used by the task.
	//
	// For tasks using the Fargate launch type,
	// this field is required and you must use one of the following values,
	// which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB)
	//
	// 512 (.5 vCPU) - Available memory values: 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB)
	//
	// 1024 (1 vCPU) - Available memory values: 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB)
	//
	// 2048 (2 vCPU) - Available memory values: Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB)
	//
	// 4096 (4 vCPU) - Available memory values: Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB).
	// Experimental.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// The amount (in GiB) of ephemeral storage to be allocated to the task.
	//
	// The maximum supported value is 200 GiB.
	//
	// NOTE: This parameter is only supported for tasks hosted on AWS Fargate using platform version 1.4.0 or later.
	// Experimental.
	EphemeralStorageGiB *float64 `field:"optional" json:"ephemeralStorageGiB" yaml:"ephemeralStorageGiB"`
	// Provider running an image to run inside CodeBuild with GitHub runner pre-configured.
	//
	// A user named `runner` is expected to exist.
	//
	// The entry point should start GitHub runner. For example:
	//
	// ```
	// #!/bin/bash
	// set -e -u -o pipefail
	//
	// /home/runner/config.sh --unattended --url "https://${GITHUB_DOMAIN}/${OWNER}/${REPO}" --token "${RUNNER_TOKEN}" --ephemeral --work _work --labels "${RUNNER_LABEL}" --disableupdate --name "${RUNNER_NAME}"
	// /home/runner/run.sh
	// ```.
	// Experimental.
	ImageBuilder IImageBuilder `field:"optional" json:"imageBuilder" yaml:"imageBuilder"`
	// GitHub Actions label used for this provider.
	// Experimental.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The amount (in MiB) of memory used by the task.
	//
	// For tasks using the Fargate launch type,
	// this field is required and you must use one of the following values, which determines your range of valid values for the cpu parameter:
	//
	// 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available cpu values: 256 (.25 vCPU)
	//
	// 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available cpu values: 512 (.5 vCPU)
	//
	// 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB) - Available cpu values: 1024 (1 vCPU)
	//
	// Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) - Available cpu values: 2048 (2 vCPU)
	//
	// Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available cpu values: 4096 (4 vCPU).
	// Experimental.
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// Security Group to assign to the task.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Use Fargate spot capacity provider to save money.
	//
	// * Runners may fail to start due to missing capacity.
	// * Runners might be stopped prematurely with spot pricing.
	// Experimental.
	Spot *bool `field:"optional" json:"spot" yaml:"spot"`
	// VPC to launch the runners in.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

// Create all the required infrastructure to provide self-hosted GitHub runners.
//
// It creates a webhook, secrets, and a step function to orchestrate all runs. Secrets are not automatically filled. See README.md for instructions on how to setup GitHub integration.
//
// By default, this will create a runner provider of each available type with the defaults. This is good enough for the initial setup stage when you just want to get GitHub integration working.
//
// ```typescript
// new GitHubRunners(this, 'runners');
// ```
//
// Usually you'd want to configure the runner providers so the runners can run in a certain VPC or have certain permissions.
//
// ```typescript
// const vpc = ec2.Vpc.fromLookup(this, 'vpc', { vpcId: 'vpc-1234567' });
// const runnerSg = new ec2.SecurityGroup(this, 'runner security group', { vpc: vpc });
// const dbSg = ec2.SecurityGroup.fromSecurityGroupId(this, 'database security group', 'sg-1234567');
// const bucket = new s3.Bucket(this, 'runner bucket');
//
// // create a custom CodeBuild provider
// const myProvider = new CodeBuildRunner(
//    this, 'codebuild runner',
//    {
//       label: 'my-codebuild',
//       vpc: vpc,
//       securityGroup: runnerSg,
//    },
// );
// // grant some permissions to the provider
// bucket.grantReadWrite(myProvider);
// dbSg.connections.allowFrom(runnerSg, ec2.Port.tcp(3306), 'allow runners to connect to MySQL database');
//
// // create the runner infrastructure
// new GitHubRunners(
//    this,
//    'runners',
//    {
//      providers: [myProvider],
//    }
// );
// ```.
// Experimental.
type GitHubRunners interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Configured runner providers.
	// Experimental.
	Providers() *[]IRunnerProvider
	// Secrets for GitHub communication including webhook secret and runner authentication.
	// Experimental.
	Secrets() Secrets
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GitHubRunners
type jsiiProxy_GitHubRunners struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_GitHubRunners) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubRunners) Providers() *[]IRunnerProvider {
	var returns *[]IRunnerProvider
	_jsii_.Get(
		j,
		"providers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubRunners) Secrets() Secrets {
	var returns Secrets
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}


// Experimental.
func NewGitHubRunners(scope constructs.Construct, id *string, props *GitHubRunnersProps) GitHubRunners {
	_init_.Initialize()

	j := jsiiProxy_GitHubRunners{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.GitHubRunners",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGitHubRunners_Override(g GitHubRunners, scope constructs.Construct, id *string, props *GitHubRunnersProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.GitHubRunners",
		[]interface{}{scope, id, props},
		g,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func GitHubRunners_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.GitHubRunners",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubRunners) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for GitHubRunners.
// Experimental.
type GitHubRunnersProps struct {
	// Allow management functions to run in public subnets.
	//
	// Lambda Functions in a public subnet can NOT access the internet.
	// Experimental.
	AllowPublicSubnet *bool `field:"optional" json:"allowPublicSubnet" yaml:"allowPublicSubnet"`
	// Path to a directory containing a file named certs.pem containing any additional certificates required to trust GitHub Enterprise Server. Use this when GitHub Enterprise Server certificates are self-signed.
	//
	// You may also want to use custom images for your runner providers that contain the same certificates. See {@link CodeBuildImageBuilder.addCertificates}.
	//
	// ```typescript
	// const imageBuilder = new CodeBuildImageBuilder(this, 'Image Builder with Certs', {
	//      dockerfilePath: CodeBuildRunner.LINUX_X64_DOCKERFILE_PATH,
	// });
	// imageBuilder.addExtraCertificates('path-to-my-extra-certs-folder');
	//
	// const provider = new CodeBuildRunner(this, 'CodeBuild', {
	//      imageBuilder: imageBuilder,
	// });
	//
	// new GitHubRunners(
	//    this,
	//    'runners',
	//    {
	//      providers: [provider],
	//      extraCertificates: 'path-to-my-extra-certs-folder',
	//    }
	// );
	// ```.
	// Experimental.
	ExtraCertificates *string `field:"optional" json:"extraCertificates" yaml:"extraCertificates"`
	// List of runner providers to use.
	//
	// At least one provider is required. Provider will be selected when its label matches the labels requested by the workflow job.
	// Experimental.
	Providers *[]IRunnerProvider `field:"optional" json:"providers" yaml:"providers"`
	// Security group attached to all management functions.
	//
	// Use this with to provide access to GitHub Enterprise Server hosted inside a VPC.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// VPC used for all management functions.
	//
	// Use this with GitHub Enterprise Server hosted that's inaccessible from outside the VPC.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// VPC subnets used for all management functions.
	//
	// Use this with GitHub Enterprise Server hosted that's inaccessible from outside the VPC.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

// Interface for constructs that build an image that can be used in {@link IRunnerProvider}.
//
// Anything that ends up with an ECR repository containing a Docker image that runs GitHub self-hosted runners can be used. A simple implementation could even point to an existing image and nothing else.
//
// It's important that the specified image tag be available at the time the repository is available. Providers usually assume the image is ready and will fail if it's not.
//
// The image can be further updated over time manually or using a schedule as long as it is always written to the same tag.
// Experimental.
type IImageBuilder interface {
	// ECR repository containing the image.
	//
	// This method can be called multiple times if the image is bound to multiple providers. Make sure you cache the image when implementing or return an error if this builder doesn't support reusing images.
	//
	// Returns: image.
	// Experimental.
	Bind() *RunnerImage
}

// The jsii proxy for IImageBuilder
type jsiiProxy_IImageBuilder struct {
	_ byte // padding
}

func (i *jsiiProxy_IImageBuilder) Bind() *RunnerImage {
	var returns *RunnerImage

	_jsii_.Invoke(
		i,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Interface for all runner providers.
//
// Implementations create all required resources and return a step function task that starts those resources from {@link getStepFunctionTask}.
// Experimental.
type IRunnerProvider interface {
	awsec2.IConnectable
	awsiam.IGrantable
	// Generate step function tasks that execute the runner.
	//
	// Called by GithubRunners and shouldn't be called manually.
	// Experimental.
	GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable
	// GitHub Actions label associated with this runner provider.
	// Experimental.
	Label() *string
	// Security group associated with runners.
	// Experimental.
	SecurityGroup() awsec2.ISecurityGroup
	// VPC network in which runners will be placed.
	// Experimental.
	Vpc() awsec2.IVpc
}

// The jsii proxy for IRunnerProvider
type jsiiProxy_IRunnerProvider struct {
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
}

func (i *jsiiProxy_IRunnerProvider) GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable {
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		i,
		"getStepFunctionTask",
		[]interface{}{parameters},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IRunnerProvider) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerProvider) SecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerProvider) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerProvider) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunnerProvider) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

// GitHub Actions runner provider using Lambda to execute the actions.
//
// Creates a Docker-based function that gets executed for each job.
//
// This construct is not meant to be used by itself. It should be passed in the providers property for GitHubRunners.
// Experimental.
type LambdaRunner interface {
	constructs.Construct
	IRunnerProvider
	// The network connections associated with this resource.
	// Experimental.
	Connections() awsec2.Connections
	// The function hosting the GitHub runner.
	// Experimental.
	Function() awslambda.Function
	// Grant principal used to add permissions to the runner role.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// Label associated with this provider.
	// Experimental.
	Label() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Security group attached to the function.
	// Experimental.
	SecurityGroup() awsec2.ISecurityGroup
	// VPC used for hosting the function.
	// Experimental.
	Vpc() awsec2.IVpc
	// Generate step function task(s) to start a new runner.
	//
	// Called by GithubRunners and shouldn't be called manually.
	// Experimental.
	GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LambdaRunner
type jsiiProxy_LambdaRunner struct {
	internal.Type__constructsConstruct
	jsiiProxy_IRunnerProvider
}

func (j *jsiiProxy_LambdaRunner) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRunner) Function() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"function",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRunner) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRunner) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRunner) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRunner) SecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRunner) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaRunner(scope constructs.Construct, id *string, props *LambdaRunnerProps) LambdaRunner {
	_init_.Initialize()

	j := jsiiProxy_LambdaRunner{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.LambdaRunner",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaRunner_Override(l LambdaRunner, scope constructs.Construct, id *string, props *LambdaRunnerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.LambdaRunner",
		[]interface{}{scope, id, props},
		l,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func LambdaRunner_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.LambdaRunner",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LambdaRunner_LINUX_ARM64_DOCKERFILE_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.LambdaRunner",
		"LINUX_ARM64_DOCKERFILE_PATH",
		&returns,
	)
	return returns
}

func LambdaRunner_LINUX_X64_DOCKERFILE_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.LambdaRunner",
		"LINUX_X64_DOCKERFILE_PATH",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LambdaRunner) GetStepFunctionTask(parameters *RunnerRuntimeParameters) awsstepfunctions.IChainable {
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		l,
		"getStepFunctionTask",
		[]interface{}{parameters},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRunner) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type LambdaRunnerProps struct {
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// The size of the function’s /tmp directory in MiB.
	// Experimental.
	EphemeralStorageSize awscdk.Size `field:"optional" json:"ephemeralStorageSize" yaml:"ephemeralStorageSize"`
	// Provider running an image to run inside CodeBuild with GitHub runner pre-configured.
	//
	// The default command (`CMD`) should be `["runner.handler"]` which points to an included `runner.js` with a function named `handler`. The function should start the GitHub runner.
	// See: https://github.com/CloudSnorkel/cdk-github-runners/tree/main/src/providers/docker-images/lambda
	//
	// Experimental.
	ImageBuilder IImageBuilder `field:"optional" json:"imageBuilder" yaml:"imageBuilder"`
	// GitHub Actions label used for this provider.
	// Experimental.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The amount of memory, in MB, that is allocated to your Lambda function.
	//
	// Lambda uses this value to proportionally allocate the amount of CPU
	// power. For more information, see Resource Model in the AWS Lambda
	// Developer Guide.
	// Experimental.
	MemorySize *float64 `field:"optional" json:"memorySize" yaml:"memorySize"`
	// Security Group to assign to this instance.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Where to place the network interfaces within the VPC.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The function execution time (in seconds) after which Lambda terminates the function.
	//
	// Because the execution time affects cost, set this value
	// based on the function's expected execution time.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// VPC to launch the runners in.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

// OS enum for an image.
// Experimental.
type Os interface {
	// Experimental.
	Name() *string
	// Checks if the given OS is the same as this one.
	// Experimental.
	Is(os Os) *bool
}

// The jsii proxy struct for Os
type jsiiProxy_Os struct {
	_ byte // padding
}

func (j *jsiiProxy_Os) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func Os_LINUX() Os {
	_init_.Initialize()
	var returns Os
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.Os",
		"LINUX",
		&returns,
	)
	return returns
}

func Os_WINDOWS() Os {
	_init_.Initialize()
	var returns Os
	_jsii_.StaticGet(
		"@cloudsnorkel/cdk-github-runners.Os",
		"WINDOWS",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_Os) Is(os Os) *bool {
	var returns *bool

	_jsii_.Invoke(
		o,
		"is",
		[]interface{}{os},
		&returns,
	)

	return returns
}

// Experimental.
type RunnerImage struct {
	// Architecture of the image.
	// Experimental.
	Architecture Architecture `field:"required" json:"architecture" yaml:"architecture"`
	// Image digest for providers that need to know the digest like Lambda.
	//
	// WARNING: the digest might change when the builder automatically rebuilds the image on a schedule. Do not expect for this digest to stay the same between deploys.
	// Experimental.
	ImageDigest *string `field:"required" json:"imageDigest" yaml:"imageDigest"`
	// ECR repository containing the image.
	// Experimental.
	ImageRepository awsecr.IRepository `field:"required" json:"imageRepository" yaml:"imageRepository"`
	// Static image tag where the image will be pushed.
	// Experimental.
	ImageTag *string `field:"required" json:"imageTag" yaml:"imageTag"`
	// OS type of the image.
	// Experimental.
	Os Os `field:"required" json:"os" yaml:"os"`
}

// Common properties for all runner providers.
// Experimental.
type RunnerProviderProps struct {
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
}

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
	// Path to repostiroy owner name.
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

// Defines desired GitHub Actions runner version.
// Experimental.
type RunnerVersion interface {
	// Experimental.
	Version() *string
}

// The jsii proxy struct for RunnerVersion
type jsiiProxy_RunnerVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_RunnerVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Experimental.
func NewRunnerVersion(version *string) RunnerVersion {
	_init_.Initialize()

	j := jsiiProxy_RunnerVersion{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.RunnerVersion",
		[]interface{}{version},
		&j,
	)

	return &j
}

// Experimental.
func NewRunnerVersion_Override(r RunnerVersion, version *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.RunnerVersion",
		[]interface{}{version},
		r,
	)
}

// Use the latest version available at the time the runner provider image is built.
// Experimental.
func RunnerVersion_Latest() RunnerVersion {
	_init_.Initialize()

	var returns RunnerVersion

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.RunnerVersion",
		"latest",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Use a specific version.
// See: https://github.com/actions/runner/releases
//
// Experimental.
func RunnerVersion_Specific(version *string) RunnerVersion {
	_init_.Initialize()

	var returns RunnerVersion

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.RunnerVersion",
		"specific",
		[]interface{}{version},
		&returns,
	)

	return returns
}

// Secrets required for GitHub runners operation.
// Experimental.
type Secrets interface {
	constructs.Construct
	// Authentication secret for GitHub containing either app details or personal authentication token.
	//
	// This secret is used to register runners and
	// cancel jobs when the runner fails to start.
	//
	// This secret is meant to be edited by the user after being created.
	// Experimental.
	Github() awssecretsmanager.Secret
	// GitHub app private key. Not needed when using personal authentication tokens.
	//
	// This secret is meant to be edited by the user after being created. It is separate than the main GitHub secret because inserting private keys into JSON is hard.
	// Experimental.
	GithubPrivateKey() awssecretsmanager.Secret
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Setup secret used to authenticate user for our setup wizard.
	//
	// Should be empty after setup has been completed.
	// Experimental.
	Setup() awssecretsmanager.Secret
	// Webhook secret used to confirm events are coming from GitHub and nowhere else.
	// Experimental.
	Webhook() awssecretsmanager.Secret
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Secrets
type jsiiProxy_Secrets struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Secrets) Github() awssecretsmanager.Secret {
	var returns awssecretsmanager.Secret
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secrets) GithubPrivateKey() awssecretsmanager.Secret {
	var returns awssecretsmanager.Secret
	_jsii_.Get(
		j,
		"githubPrivateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secrets) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secrets) Setup() awssecretsmanager.Secret {
	var returns awssecretsmanager.Secret
	_jsii_.Get(
		j,
		"setup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secrets) Webhook() awssecretsmanager.Secret {
	var returns awssecretsmanager.Secret
	_jsii_.Get(
		j,
		"webhook",
		&returns,
	)
	return returns
}


// Experimental.
func NewSecrets(scope constructs.Construct, id *string) Secrets {
	_init_.Initialize()

	j := jsiiProxy_Secrets{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.Secrets",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Experimental.
func NewSecrets_Override(s Secrets, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.Secrets",
		[]interface{}{scope, id},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Secrets_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.Secrets",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Secrets) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Helper class with methods to use static images that are built outside the context of this project.
// Experimental.
type StaticRunnerImage interface {
}

// The jsii proxy struct for StaticRunnerImage
type jsiiProxy_StaticRunnerImage struct {
	_ byte // padding
}

// Experimental.
func NewStaticRunnerImage() StaticRunnerImage {
	_init_.Initialize()

	j := jsiiProxy_StaticRunnerImage{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.StaticRunnerImage",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewStaticRunnerImage_Override(s StaticRunnerImage) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.StaticRunnerImage",
		nil, // no parameters
		s,
	)
}

// Create a builder from an existing Docker Hub image.
//
// The image must already have GitHub Actions runner installed. You are responsible to update it and remove it when done.
//
// We create a CodeBuild image builder behind the scenes to copy the image over to ECR. This helps avoid Docker Hub rate limits and prevent failures.
// Experimental.
func StaticRunnerImage_FromDockerHub(scope constructs.Construct, id *string, image *string, architecture Architecture, os Os) IImageBuilder {
	_init_.Initialize()

	var returns IImageBuilder

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.StaticRunnerImage",
		"fromDockerHub",
		[]interface{}{scope, id, image, architecture, os},
		&returns,
	)

	return returns
}

// Create a builder (that doesn't actually build anything) from an existing image in an existing repository.
//
// The image must already have GitHub Actions runner installed. You are responsible to update it and remove it when done.
// Experimental.
func StaticRunnerImage_FromEcrRepository(repository awsecr.IRepository, tag *string, architecture Architecture, os Os) IImageBuilder {
	_init_.Initialize()

	var returns IImageBuilder

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.StaticRunnerImage",
		"fromEcrRepository",
		[]interface{}{repository, tag, architecture, os},
		&returns,
	)

	return returns
}

