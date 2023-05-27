package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

// An image builder that uses AWS Image Builder to build Docker images pre-baked with all the GitHub Actions runner requirements.
//
// Builders can be used with runner providers.
//
// The CodeBuild builder is better and faster. Only use this one if you have no choice. For example, if you need Windows containers.
//
// Each builder re-runs automatically at a set interval to make sure the images contain the latest versions of everything.
//
// You can create an instance of this construct to customize the image used to spin-up runners. Some runner providers may require custom components. Check the runner provider documentation. The default components work with CodeBuild and Fargate.
//
// For example, to set a specific runner version, rebuild the image every 2 weeks, and add a few packages for the Fargate provider, use:
//
// ```
// const builder = new ContainerImageBuilder(this, 'Builder', {
//     runnerVersion: RunnerVersion.specific('2.293.0'),
//     rebuildInterval: Duration.days(14),
// });
// new CodeBuildRunnerProvider(this, 'CodeBuild provider', {
//     labels: ['custom-codebuild'],
//     imageBuilder: builder,
// });
// ```.
// Deprecated: use RunnerImageBuilder.
type ContainerImageBuilder interface {
	constructs.Construct
	IRunnerImageBuilder
	// Deprecated: use RunnerImageBuilder.
	Architecture() Architecture
	// Deprecated: use RunnerImageBuilder.
	Components() *[]ImageBuilderComponent
	// Deprecated: use RunnerImageBuilder.
	SetComponents(val *[]ImageBuilderComponent)
	// The network connections associated with this resource.
	// Deprecated: use RunnerImageBuilder.
	Connections() awsec2.Connections
	// Deprecated: use RunnerImageBuilder.
	Description() *string
	// The tree node.
	// Deprecated: use RunnerImageBuilder.
	Node() constructs.Node
	// Deprecated: use RunnerImageBuilder.
	Os() Os
	// Deprecated: use RunnerImageBuilder.
	Platform() *string
	// Deprecated: use RunnerImageBuilder.
	Repository() awsecr.IRepository
	// Deprecated: use RunnerImageBuilder.
	RunnerVersion() RunnerVersion
	// Add a component to be installed.
	// Deprecated: use RunnerImageBuilder.
	AddComponent(component ImageBuilderComponent)
	// Add extra trusted certificates. This helps deal with self-signed certificates for GitHub Enterprise Server.
	//
	// All first party Dockerfiles support this. Others may not.
	// Deprecated: use RunnerImageBuilder.
	AddExtraCertificates(path *string)
	// Build and return an AMI with GitHub Runner installed in it.
	//
	// Anything that ends up with a launch template pointing to an AMI that runs GitHub self-hosted runners can be used. A simple implementation could even point to an existing AMI and nothing else.
	//
	// The AMI can be further updated over time manually or using a schedule as long as it is always written to the same launch template.
	// Deprecated: use RunnerImageBuilder.
	BindAmi() *RunnerAmi
	// Called by IRunnerProvider to finalize settings and create the image builder.
	// Deprecated: use RunnerImageBuilder.
	BindDockerImage() *RunnerImage
	// Deprecated: use RunnerImageBuilder.
	CreateImage(infra awsimagebuilder.CfnInfrastructureConfiguration, dist awsimagebuilder.CfnDistributionConfiguration, log awslogs.LogGroup, imageRecipeArn *string, containerRecipeArn *string) awsimagebuilder.CfnImage
	// Deprecated: use RunnerImageBuilder.
	CreateInfrastructure(managedPolicies *[]awsiam.IManagedPolicy) awsimagebuilder.CfnInfrastructureConfiguration
	// Deprecated: use RunnerImageBuilder.
	CreateLog(recipeName *string) awslogs.LogGroup
	// Deprecated: use RunnerImageBuilder.
	CreatePipeline(infra awsimagebuilder.CfnInfrastructureConfiguration, dist awsimagebuilder.CfnDistributionConfiguration, log awslogs.LogGroup, imageRecipeArn *string, containerRecipeArn *string) awsimagebuilder.CfnImagePipeline
	// Add a component to be installed before any other components.
	//
	// Useful for required system settings like certificates or proxy settings.
	// Deprecated: use RunnerImageBuilder.
	PrependComponent(component ImageBuilderComponent)
	// Returns a string representation of this construct.
	// Deprecated: use RunnerImageBuilder.
	ToString() *string
}

// The jsii proxy struct for ContainerImageBuilder
type jsiiProxy_ContainerImageBuilder struct {
	internal.Type__constructsConstruct
	jsiiProxy_IRunnerImageBuilder
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

func (j *jsiiProxy_ContainerImageBuilder) Components() *[]ImageBuilderComponent {
	var returns *[]ImageBuilderComponent
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerImageBuilder) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
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


// Deprecated: use RunnerImageBuilder.
func NewContainerImageBuilder(scope constructs.Construct, id *string, props *ContainerImageBuilderProps) ContainerImageBuilder {
	_init_.Initialize()

	if err := validateNewContainerImageBuilderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerImageBuilder{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.ContainerImageBuilder",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Deprecated: use RunnerImageBuilder.
func NewContainerImageBuilder_Override(c ContainerImageBuilder, scope constructs.Construct, id *string, props *ContainerImageBuilderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.ContainerImageBuilder",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_ContainerImageBuilder)SetComponents(val *[]ImageBuilderComponent) {
	if err := j.validateSetComponentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"components",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func ContainerImageBuilder_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerImageBuilder_IsConstructParameters(x); err != nil {
		panic(err)
	}
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
	if err := c.validateAddComponentParameters(component); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addComponent",
		[]interface{}{component},
	)
}

func (c *jsiiProxy_ContainerImageBuilder) AddExtraCertificates(path *string) {
	if err := c.validateAddExtraCertificatesParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addExtraCertificates",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_ContainerImageBuilder) BindAmi() *RunnerAmi {
	var returns *RunnerAmi

	_jsii_.Invoke(
		c,
		"bindAmi",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerImageBuilder) BindDockerImage() *RunnerImage {
	var returns *RunnerImage

	_jsii_.Invoke(
		c,
		"bindDockerImage",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerImageBuilder) CreateImage(infra awsimagebuilder.CfnInfrastructureConfiguration, dist awsimagebuilder.CfnDistributionConfiguration, log awslogs.LogGroup, imageRecipeArn *string, containerRecipeArn *string) awsimagebuilder.CfnImage {
	if err := c.validateCreateImageParameters(infra, dist, log); err != nil {
		panic(err)
	}
	var returns awsimagebuilder.CfnImage

	_jsii_.Invoke(
		c,
		"createImage",
		[]interface{}{infra, dist, log, imageRecipeArn, containerRecipeArn},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerImageBuilder) CreateInfrastructure(managedPolicies *[]awsiam.IManagedPolicy) awsimagebuilder.CfnInfrastructureConfiguration {
	if err := c.validateCreateInfrastructureParameters(managedPolicies); err != nil {
		panic(err)
	}
	var returns awsimagebuilder.CfnInfrastructureConfiguration

	_jsii_.Invoke(
		c,
		"createInfrastructure",
		[]interface{}{managedPolicies},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerImageBuilder) CreateLog(recipeName *string) awslogs.LogGroup {
	if err := c.validateCreateLogParameters(recipeName); err != nil {
		panic(err)
	}
	var returns awslogs.LogGroup

	_jsii_.Invoke(
		c,
		"createLog",
		[]interface{}{recipeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerImageBuilder) CreatePipeline(infra awsimagebuilder.CfnInfrastructureConfiguration, dist awsimagebuilder.CfnDistributionConfiguration, log awslogs.LogGroup, imageRecipeArn *string, containerRecipeArn *string) awsimagebuilder.CfnImagePipeline {
	if err := c.validateCreatePipelineParameters(infra, dist, log); err != nil {
		panic(err)
	}
	var returns awsimagebuilder.CfnImagePipeline

	_jsii_.Invoke(
		c,
		"createPipeline",
		[]interface{}{infra, dist, log, imageRecipeArn, containerRecipeArn},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerImageBuilder) PrependComponent(component ImageBuilderComponent) {
	if err := c.validatePrependComponentParameters(component); err != nil {
		panic(err)
	}
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

