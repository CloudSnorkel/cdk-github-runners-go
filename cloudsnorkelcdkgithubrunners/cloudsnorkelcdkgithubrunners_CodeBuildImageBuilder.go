// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

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
//      dockerfilePath: FargateProvider.LINUX_X64_DOCKERFILE_PATH,
//      runnerVersion: RunnerVersion.specific('2.293.0'),
//      rebuildInterval: Duration.days(14),
// });
// builder.setBuildArg('EXTRA_PACKAGES', 'nginx xz-utils');
// new FargateRunner(this, 'Fargate provider', {
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

