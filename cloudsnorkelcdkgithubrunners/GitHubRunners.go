package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

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
// const myProvider = new CodeBuildRunnerProvider(
//   this, 'codebuild runner',
//   {
//      labels: ['my-codebuild'],
//      vpc: vpc,
//      securityGroups: [runnerSg],
//   },
// );
// // grant some permissions to the provider
// bucket.grantReadWrite(myProvider);
// dbSg.connections.allowFrom(runnerSg, ec2.Port.tcp(3306), 'allow runners to connect to MySQL database');
//
// // create the runner infrastructure
// new GitHubRunners(
//   this,
//   'runners',
//   {
//     providers: [myProvider],
//   }
// );
// ```.
// Experimental.
type GitHubRunners interface {
	constructs.Construct
	awsec2.IConnectable
	// Manage the connections of all management functions.
	//
	// Use this to enable connections to your GitHub Enterprise Server in a VPC.
	//
	// This cannot be used to manage connections of the runners. Use the `connections` property of each runner provider to manage runner connections.
	// Experimental.
	Connections() awsec2.Connections
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Props() *GitHubRunnersProps
	// Configured runner providers.
	// Experimental.
	Providers() *[]interface{}
	// Secrets for GitHub communication including webhook secret and runner authentication.
	// Experimental.
	Secrets() Secrets
	// Creates CloudWatch Logs Insights saved queries that can be used to debug issues with the runners.
	//
	// * "Webhook errors" helps diagnose configuration issues with GitHub integration
	// * "Ignored webhook" helps understand why runners aren't started
	// * "Ignored jobs based on labels" helps debug label matching issues
	// * "Webhook started runners" helps understand which runners were started.
	// Experimental.
	CreateLogsInsightsQueries()
	// Creates a topic for notifications when a runner image build fails.
	//
	// Runner images are rebuilt every week by default. This provides the latest GitHub Runner version and software updates.
	//
	// If you want to be sure you are using the latest runner version, you can use this topic to be notified when a build fails.
	// Experimental.
	FailedImageBuildsTopic() awssns.Topic
	// Metric for failed runner executions.
	//
	// A failed runner usually means the runner failed to start and so a job was never executed. It doesn't necessarily mean the job was executed and failed. For that, see {@link metricJobCompleted}.
	// Experimental.
	MetricFailed(props *awscloudwatch.MetricProps) awscloudwatch.Metric
	// Metric for the number of GitHub Actions jobs completed.
	//
	// It has `ProviderLabels` and `Status` dimensions. The status can be one of "Succeeded", "SucceededWithIssues", "Failed", "Canceled", "Skipped", or "Abandoned".
	//
	// **WARNING:** this method creates a metric filter for each provider. Each metric has a status dimension with six possible values. These resources may incur cost.
	// Experimental.
	MetricJobCompleted(props *awscloudwatch.MetricProps) awscloudwatch.Metric
	// Metric for successful executions.
	//
	// A successful execution doesn't always mean a runner was started. It can be successful even without any label matches.
	//
	// A successful runner doesn't mean the job it executed was successful. For that, see {@link metricJobCompleted}.
	// Experimental.
	MetricSucceeded(props *awscloudwatch.MetricProps) awscloudwatch.Metric
	// Metric for the interval, in milliseconds, between the time the execution starts and the time it closes.
	//
	// This time may be longer than the time the runner took.
	// Experimental.
	MetricTime(props *awscloudwatch.MetricProps) awscloudwatch.Metric
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GitHubRunners
type jsiiProxy_GitHubRunners struct {
	internal.Type__constructsConstruct
	internal.Type__awsec2IConnectable
}

func (j *jsiiProxy_GitHubRunners) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
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

func (j *jsiiProxy_GitHubRunners) Providers() *[]interface{} {
	var returns *[]interface{}
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

	if err := validateNewGitHubRunnersParameters(scope, id, props); err != nil {
		panic(err)
	}
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

	if err := validateGitHubRunners_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.GitHubRunners",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubRunners) CreateLogsInsightsQueries() {
	_jsii_.InvokeVoid(
		g,
		"createLogsInsightsQueries",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitHubRunners) FailedImageBuildsTopic() awssns.Topic {
	var returns awssns.Topic

	_jsii_.Invoke(
		g,
		"failedImageBuildsTopic",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubRunners) MetricFailed(props *awscloudwatch.MetricProps) awscloudwatch.Metric {
	if err := g.validateMetricFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubRunners) MetricJobCompleted(props *awscloudwatch.MetricProps) awscloudwatch.Metric {
	if err := g.validateMetricJobCompletedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricJobCompleted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubRunners) MetricSucceeded(props *awscloudwatch.MetricProps) awscloudwatch.Metric {
	if err := g.validateMetricSucceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubRunners) MetricTime(props *awscloudwatch.MetricProps) awscloudwatch.Metric {
	if err := g.validateMetricTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricTime",
		[]interface{}{props},
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

