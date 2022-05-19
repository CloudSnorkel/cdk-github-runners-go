// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

// GitHub Actions runner provider using CodeBuild to execute the actions.
//
// Creates a project that gets started for each job.
// Experimental.
type CodeBuildRunner interface {
	constructs.Construct
	IRunnerProvider
	// The network connections associated with this resource.
	// Experimental.
	Connections() awsec2.Connections
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// GitHub Actions label associated with this runner provider.
	// Experimental.
	Label() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() awscodebuild.Project
	// Security group associated with runners.
	// Experimental.
	SecurityGroup() awsec2.ISecurityGroup
	// VPC network in which runners will be placed.
	// Experimental.
	Vpc() awsec2.IVpc
	// Generate step function tasks that execute the runner.
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
	// Version of GitHub Runners to install.
	// Experimental.
	RunnerVersion RunnerVersion `field:"optional" json:"runnerVersion" yaml:"runnerVersion"`
	// The type of compute to use for this build.
	//
	// See the {@link ComputeType} enum for the possible values.
	// Experimental.
	ComputeType awscodebuild.ComputeType `field:"optional" json:"computeType" yaml:"computeType"`
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
// Experimental.
type FargateRunner interface {
	constructs.Construct
	IRunnerProvider
	// Experimental.
	AssignPublicIp() *bool
	// Experimental.
	Cluster() awsecs.Cluster
	// The network connections associated with this resource.
	// Experimental.
	Connections() awsec2.Connections
	// Experimental.
	Container() awsecs.ContainerDefinition
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// GitHub Actions label associated with this runner provider.
	// Experimental.
	Label() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Security group associated with runners.
	// Experimental.
	SecurityGroup() awsec2.ISecurityGroup
	// Experimental.
	Task() awsecs.FargateTaskDefinition
	// VPC network in which runners will be placed.
	// Experimental.
	Vpc() awsec2.IVpc
	// Generate step function tasks that execute the runner.
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

// Experimental.
type FargateRunnerProps struct {
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// Version of GitHub Runners to install.
	// Experimental.
	RunnerVersion RunnerVersion `field:"optional" json:"runnerVersion" yaml:"runnerVersion"`
	// Assign public IP to the runner task.
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
	// VPC to launch the runners in.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

// Experimental.
type GitHubRunners interface {
	constructs.Construct
	// Experimental.
	DefaultProvider() IRunnerProvider
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Props() *GitHubRunnersProps
	// Experimental.
	Providers() *[]IRunnerProvider
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

func (j *jsiiProxy_GitHubRunners) DefaultProvider() IRunnerProvider {
	var returns IRunnerProvider
	_jsii_.Get(
		j,
		"defaultProvider",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_GitHubRunners) Props() *GitHubRunnersProps {
	var returns *GitHubRunnersProps
	_jsii_.Get(
		j,
		"props",
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

// Properties of the GitHubRunners.
// Experimental.
type GitHubRunnersProps struct {
	// Experimental.
	DefaultProviderLabel *string `field:"optional" json:"defaultProviderLabel" yaml:"defaultProviderLabel"`
	// Experimental.
	Providers *[]IRunnerProvider `field:"optional" json:"providers" yaml:"providers"`
}

// Experimental.
type IRunnerProvider interface {
	awsec2.IConnectable
	awsiam.IGrantable
	// Generate step function tasks that execute the runner.
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
// Experimental.
type LambdaRunner interface {
	constructs.Construct
	IRunnerProvider
	// The network connections associated with this resource.
	// Experimental.
	Connections() awsec2.Connections
	// Experimental.
	Function() awslambda.Function
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// GitHub Actions label associated with this runner provider.
	// Experimental.
	Label() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Security group associated with runners.
	// Experimental.
	SecurityGroup() awsec2.ISecurityGroup
	// VPC network in which runners will be placed.
	// Experimental.
	Vpc() awsec2.IVpc
	// Generate step function tasks that execute the runner.
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
	// Version of GitHub Runners to install.
	// Experimental.
	RunnerVersion RunnerVersion `field:"optional" json:"runnerVersion" yaml:"runnerVersion"`
	// The size of the function’s /tmp directory in MiB.
	// Experimental.
	EphemeralStorageSize awscdk.Size `field:"optional" json:"ephemeralStorageSize" yaml:"ephemeralStorageSize"`
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

// Experimental.
type RunnerProviderProps struct {
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// Version of GitHub Runners to install.
	// Experimental.
	RunnerVersion RunnerVersion `field:"optional" json:"runnerVersion" yaml:"runnerVersion"`
}

// Experimental.
type RunnerRuntimeParameters struct {
	// Experimental.
	GithubDomainPath *string `field:"required" json:"githubDomainPath" yaml:"githubDomainPath"`
	// Experimental.
	OwnerPath *string `field:"required" json:"ownerPath" yaml:"ownerPath"`
	// Experimental.
	RepoPath *string `field:"required" json:"repoPath" yaml:"repoPath"`
	// Experimental.
	RunnerNamePath *string `field:"required" json:"runnerNamePath" yaml:"runnerNamePath"`
	// Experimental.
	RunnerTokenPath *string `field:"required" json:"runnerTokenPath" yaml:"runnerTokenPath"`
}

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
	// Experimental.
	Github() awssecretsmanager.Secret
	// Experimental.
	GithubPrivateKey() awssecretsmanager.Secret
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
