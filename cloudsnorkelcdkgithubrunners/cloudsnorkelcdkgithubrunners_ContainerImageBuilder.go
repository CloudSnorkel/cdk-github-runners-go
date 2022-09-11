// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

// An image builder that uses Image Builder to build Docker images pre-baked with all the GitHub Actions runner requirements.
//
// Builders can be used with runner providers.
//
// The CodeBuild builder is better and faster. Only use this one if you have no choice. For example, if you need Windows containers.
//
// Each builder re-runs automatically at a set interval to make sure the images contain the latest versions of everything.
//
// You can create an instance of this construct to customize the image used to spin-up runners. Some runner providers may require custom components. Check the runner provider documentation. The default components work with CodeBuild.
//
// For example, to set a specific runner version, rebuild the image every 2 weeks, and add a few packages for the Fargate provider, use:
//
// ```
// const builder = new ContainerImageBuilder(this, 'Builder', {
//      runnerVersion: RunnerVersion.specific('2.293.0'),
//      rebuildInterval: Duration.days(14),
// });
// new CodeBuildRunner(this, 'Fargate provider', {
//      label: 'windows-codebuild',
//      imageBuilder: builder,
// });
// ```.
// Experimental.
type ContainerImageBuilder interface {
	constructs.Construct
	IImageBuilder
	// Experimental.
	Architecture() Architecture
	// Experimental.
	Description() *string
	// Experimental.
	InstanceTypes() *[]*string
	// Experimental.
	LogRemovalPolicy() awscdk.RemovalPolicy
	// Experimental.
	LogRetention() awslogs.RetentionDays
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Os() Os
	// Experimental.
	Platform() *string
	// Experimental.
	RebuildInterval() awscdk.Duration
	// Experimental.
	Repository() awsecr.IRepository
	// Experimental.
	RunnerVersion() RunnerVersion
	// Experimental.
	SecurityGroupIds() *[]*string
	// Experimental.
	SubnetId() *string
	// Add a component to be installed.
	// Experimental.
	AddComponent(component ImageBuilderComponent)
	// Add extra trusted certificates. This helps deal with self-signed certificates for GitHub Enterprise Server.
	//
	// All first party Dockerfiles support this. Others may not.
	// Experimental.
	AddExtraCertificates(path *string)
	// Called by IRunnerProvider to finalize settings and create the image builder.
	// Experimental.
	Bind() *RunnerImage
	// Add a component to be installed before any other components.
	//
	// Useful for required system settings like certificates or proxy settings.
	// Experimental.
	PrependComponent(component ImageBuilderComponent)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerImageBuilder
type jsiiProxy_ContainerImageBuilder struct {
	internal.Type__constructsConstruct
	jsiiProxy_IImageBuilder
}

func (j *jsiiProxy_ContainerImageBuilder) Architecture() Architecture {
	var returns Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerImageBuilder) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerImageBuilder) InstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerImageBuilder) LogRemovalPolicy() awscdk.RemovalPolicy {
	var returns awscdk.RemovalPolicy
	_jsii_.Get(
		j,
		"logRemovalPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerImageBuilder) LogRetention() awslogs.RetentionDays {
	var returns awslogs.RetentionDays
	_jsii_.Get(
		j,
		"logRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerImageBuilder) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerImageBuilder) Os() Os {
	var returns Os
	_jsii_.Get(
		j,
		"os",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerImageBuilder) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerImageBuilder) RebuildInterval() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"rebuildInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerImageBuilder) Repository() awsecr.IRepository {
	var returns awsecr.IRepository
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerImageBuilder) RunnerVersion() RunnerVersion {
	var returns RunnerVersion
	_jsii_.Get(
		j,
		"runnerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerImageBuilder) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerImageBuilder) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}


// Experimental.
func NewContainerImageBuilder(scope constructs.Construct, id *string, props *ContainerImageBuilderProps) ContainerImageBuilder {
	_init_.Initialize()

	j := jsiiProxy_ContainerImageBuilder{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.ContainerImageBuilder",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewContainerImageBuilder_Override(c ContainerImageBuilder, scope constructs.Construct, id *string, props *ContainerImageBuilderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.ContainerImageBuilder",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func ContainerImageBuilder_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.ContainerImageBuilder",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerImageBuilder) AddComponent(component ImageBuilderComponent) {
	_jsii_.InvokeVoid(
		c,
		"addComponent",
		[]interface{}{component},
	)
}

func (c *jsiiProxy_ContainerImageBuilder) AddExtraCertificates(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addExtraCertificates",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_ContainerImageBuilder) Bind() *RunnerImage {
	var returns *RunnerImage

	_jsii_.Invoke(
		c,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerImageBuilder) PrependComponent(component ImageBuilderComponent) {
	_jsii_.InvokeVoid(
		c,
		"prependComponent",
		[]interface{}{component},
	)
}

func (c *jsiiProxy_ContainerImageBuilder) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

