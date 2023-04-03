// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

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
//     dockerfilePath: FargateProvider.LINUX_X64_DOCKERFILE_PATH,
//     runnerVersion: RunnerVersion.specific('2.293.0'),
//     rebuildInterval: Duration.days(14),
// });
// builder.setBuildArg('EXTRA_PACKAGES', 'nginx xz-utils');
// new FargateRunner(this, 'Fargate provider', {
//     label: 'customized-fargate',
//     imageBuilder: builder,
// });
// ```.
// Deprecated: use RunnerImageBuilder.
type CodeBuildImageBuilder interface {
	constructs.Construct
	IRunnerImageBuilder
	// Deprecated: use RunnerImageBuilder.
	Connections() awsec2.Connections
	// The tree node.
	// Deprecated: use RunnerImageBuilder.
	Node() constructs.Node
	// Deprecated: use RunnerImageBuilder.
	Props() *CodeBuildImageBuilderProps
	// Add extra trusted certificates. This helps deal with self-signed certificates for GitHub Enterprise Server.
	//
	// All first party Dockerfiles support this. Others may not.
	// Deprecated: use RunnerImageBuilder.
	AddExtraCertificates(path *string)
	// Uploads a folder to the build server at a given folder name.
	// Deprecated: use RunnerImageBuilder.
	AddFiles(sourcePath *string, destName *string)
	// Add a policy statement to the builder to access resources required to the image build.
	// Deprecated: use RunnerImageBuilder.
	AddPolicyStatement(statement awsiam.PolicyStatement)
	// Adds a command that runs after `docker build` and `docker push`.
	// Deprecated: use RunnerImageBuilder.
	AddPostBuildCommand(command *string)
	// Adds a command that runs before `docker build`.
	// Deprecated: use RunnerImageBuilder.
	AddPreBuildCommand(command *string)
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
	// Adds a build argument for Docker.
	//
	// See the documentation for the Dockerfile you're using for a list of supported build arguments.
	// Deprecated: use RunnerImageBuilder.
	SetBuildArg(name *string, value *string)
	// Returns a string representation of this construct.
	// Deprecated: use RunnerImageBuilder.
	ToString() *string
}

// The jsii proxy struct for CodeBuildImageBuilder
type jsiiProxy_CodeBuildImageBuilder struct {
	internal.Type__constructsConstruct
	jsiiProxy_IRunnerImageBuilder
}

func (j *jsiiProxy_CodeBuildImageBuilder) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
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


// Deprecated: use RunnerImageBuilder.
func NewCodeBuildImageBuilder(scope constructs.Construct, id *string, props *CodeBuildImageBuilderProps) CodeBuildImageBuilder {
	_init_.Initialize()

	if err := validateNewCodeBuildImageBuilderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodeBuildImageBuilder{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.CodeBuildImageBuilder",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Deprecated: use RunnerImageBuilder.
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

	if err := validateCodeBuildImageBuilder_IsConstructParameters(x); err != nil {
		panic(err)
	}
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
	if err := c.validateAddExtraCertificatesParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addExtraCertificates",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CodeBuildImageBuilder) AddFiles(sourcePath *string, destName *string) {
	if err := c.validateAddFilesParameters(sourcePath, destName); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addFiles",
		[]interface{}{sourcePath, destName},
	)
}

func (c *jsiiProxy_CodeBuildImageBuilder) AddPolicyStatement(statement awsiam.PolicyStatement) {
	if err := c.validateAddPolicyStatementParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPolicyStatement",
		[]interface{}{statement},
	)
}

func (c *jsiiProxy_CodeBuildImageBuilder) AddPostBuildCommand(command *string) {
	if err := c.validateAddPostBuildCommandParameters(command); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPostBuildCommand",
		[]interface{}{command},
	)
}

func (c *jsiiProxy_CodeBuildImageBuilder) AddPreBuildCommand(command *string) {
	if err := c.validateAddPreBuildCommandParameters(command); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPreBuildCommand",
		[]interface{}{command},
	)
}

func (c *jsiiProxy_CodeBuildImageBuilder) BindAmi() *RunnerAmi {
	var returns *RunnerAmi

	_jsii_.Invoke(
		c,
		"bindAmi",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildImageBuilder) BindDockerImage() *RunnerImage {
	var returns *RunnerImage

	_jsii_.Invoke(
		c,
		"bindDockerImage",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildImageBuilder) SetBuildArg(name *string, value *string) {
	if err := c.validateSetBuildArgParameters(name, value); err != nil {
		panic(err)
	}
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

