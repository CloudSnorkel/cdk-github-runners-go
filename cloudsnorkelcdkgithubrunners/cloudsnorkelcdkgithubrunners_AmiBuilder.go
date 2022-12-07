// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

// An AMI builder that uses AWS Image Builder to build AMIs pre-baked with all the GitHub Actions runner requirements.
//
// Builders can be used with {@link Ec2Runner}.
//
// Each builder re-runs automatically at a set interval to make sure the AMIs contain the latest versions of everything.
//
// You can create an instance of this construct to customize the AMI used to spin-up runners. Some runner providers may require custom components. Check the runner provider documentation.
//
// For example, to set a specific runner version, rebuild the image every 2 weeks, and add a few packages for the EC2 provider, use:
//
// ```
// const builder = new AmiBuilder(this, 'Builder', {
//      runnerVersion: RunnerVersion.specific('2.293.0'),
//      rebuildInterval: Duration.days(14),
// });
// builder.addComponent(new ImageBuilderComponent(scope, id, {
//    platform: 'Linux',
//    displayName: 'p7zip',
//    description: 'Install some more packages',
//    commands: [
//      'set -ex',
//      'apt-get install p7zip',
//    ],
// }));
// new Ec2Runner(this, 'EC2 provider', {
//      label: 'custom-ec2',
//      amiBuilder: builder,
// });
// ```.
// Experimental.
type AmiBuilder interface {
	constructs.Construct
	IAmiBuilder
	awsec2.IConnectable
	// Experimental.
	Architecture() Architecture
	// Experimental.
	Components() *[]ImageBuilderComponent
	// Experimental.
	SetComponents(val *[]ImageBuilderComponent)
	// The network connections associated with this resource.
	// Experimental.
	Connections() awsec2.Connections
	// Experimental.
	Description() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Os() Os
	// Experimental.
	Platform() *string
	// Experimental.
	RunnerVersion() RunnerVersion
	// Add a component to be installed.
	// Experimental.
	AddComponent(component ImageBuilderComponent)
	// Add extra trusted certificates.
	//
	// This helps deal with self-signed certificates for GitHub Enterprise Server.
	// Experimental.
	AddExtraCertificates(path *string)
	// Called by IRunnerProvider to finalize settings and create the AMI builder.
	// Experimental.
	Bind() *RunnerAmi
	// Experimental.
	CreateImage(infra awsimagebuilder.CfnInfrastructureConfiguration, dist awsimagebuilder.CfnDistributionConfiguration, log awslogs.LogGroup, imageRecipeArn *string, containerRecipeArn *string) awsimagebuilder.CfnImage
	// Experimental.
	CreateInfrastructure(managedPolicies *[]awsiam.IManagedPolicy) awsimagebuilder.CfnInfrastructureConfiguration
	// Experimental.
	CreateLog(recipeName *string) awslogs.LogGroup
	// Experimental.
	CreatePipeline(infra awsimagebuilder.CfnInfrastructureConfiguration, dist awsimagebuilder.CfnDistributionConfiguration, log awslogs.LogGroup, imageRecipeArn *string, containerRecipeArn *string) awsimagebuilder.CfnImagePipeline
	// Add a component to be installed before any other components.
	//
	// Useful for required system settings like certificates or proxy settings.
	// Experimental.
	PrependComponent(component ImageBuilderComponent)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AmiBuilder
type jsiiProxy_AmiBuilder struct {
	internal.Type__constructsConstruct
	jsiiProxy_IAmiBuilder
	internal.Type__awsec2IConnectable
}

func (j *jsiiProxy_AmiBuilder) Architecture() Architecture {
	var returns Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmiBuilder) Components() *[]ImageBuilderComponent {
	var returns *[]ImageBuilderComponent
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmiBuilder) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmiBuilder) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmiBuilder) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmiBuilder) Os() Os {
	var returns Os
	_jsii_.Get(
		j,
		"os",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmiBuilder) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmiBuilder) RunnerVersion() RunnerVersion {
	var returns RunnerVersion
	_jsii_.Get(
		j,
		"runnerVersion",
		&returns,
	)
	return returns
}


// Experimental.
func NewAmiBuilder(scope constructs.Construct, id *string, props *AmiBuilderProps) AmiBuilder {
	_init_.Initialize()

	if err := validateNewAmiBuilderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AmiBuilder{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.AmiBuilder",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAmiBuilder_Override(a AmiBuilder, scope constructs.Construct, id *string, props *AmiBuilderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.AmiBuilder",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_AmiBuilder)SetComponents(val *[]ImageBuilderComponent) {
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
func AmiBuilder_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAmiBuilder_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.AmiBuilder",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmiBuilder) AddComponent(component ImageBuilderComponent) {
	if err := a.validateAddComponentParameters(component); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addComponent",
		[]interface{}{component},
	)
}

func (a *jsiiProxy_AmiBuilder) AddExtraCertificates(path *string) {
	if err := a.validateAddExtraCertificatesParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addExtraCertificates",
		[]interface{}{path},
	)
}

func (a *jsiiProxy_AmiBuilder) Bind() *RunnerAmi {
	var returns *RunnerAmi

	_jsii_.Invoke(
		a,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmiBuilder) CreateImage(infra awsimagebuilder.CfnInfrastructureConfiguration, dist awsimagebuilder.CfnDistributionConfiguration, log awslogs.LogGroup, imageRecipeArn *string, containerRecipeArn *string) awsimagebuilder.CfnImage {
	if err := a.validateCreateImageParameters(infra, dist, log); err != nil {
		panic(err)
	}
	var returns awsimagebuilder.CfnImage

	_jsii_.Invoke(
		a,
		"createImage",
		[]interface{}{infra, dist, log, imageRecipeArn, containerRecipeArn},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmiBuilder) CreateInfrastructure(managedPolicies *[]awsiam.IManagedPolicy) awsimagebuilder.CfnInfrastructureConfiguration {
	if err := a.validateCreateInfrastructureParameters(managedPolicies); err != nil {
		panic(err)
	}
	var returns awsimagebuilder.CfnInfrastructureConfiguration

	_jsii_.Invoke(
		a,
		"createInfrastructure",
		[]interface{}{managedPolicies},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmiBuilder) CreateLog(recipeName *string) awslogs.LogGroup {
	if err := a.validateCreateLogParameters(recipeName); err != nil {
		panic(err)
	}
	var returns awslogs.LogGroup

	_jsii_.Invoke(
		a,
		"createLog",
		[]interface{}{recipeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmiBuilder) CreatePipeline(infra awsimagebuilder.CfnInfrastructureConfiguration, dist awsimagebuilder.CfnDistributionConfiguration, log awslogs.LogGroup, imageRecipeArn *string, containerRecipeArn *string) awsimagebuilder.CfnImagePipeline {
	if err := a.validateCreatePipelineParameters(infra, dist, log); err != nil {
		panic(err)
	}
	var returns awsimagebuilder.CfnImagePipeline

	_jsii_.Invoke(
		a,
		"createPipeline",
		[]interface{}{infra, dist, log, imageRecipeArn, containerRecipeArn},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmiBuilder) PrependComponent(component ImageBuilderComponent) {
	if err := a.validatePrependComponentParameters(component); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"prependComponent",
		[]interface{}{component},
	)
}

func (a *jsiiProxy_AmiBuilder) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
