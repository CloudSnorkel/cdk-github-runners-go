package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// Access configuration options for Lambda functions like setup and webhook function. Use this to limit access to these functions.
//
// If you need a custom access point, you can implement this abstract class yourself. Note that the Lambda functions expect API Gateway v1 or v2 input. They also expect every URL under the constructed URL to point to the function.
// Experimental.
type LambdaAccess interface {
	// Creates all required resources and returns access URL or empty string if disabled.
	//
	// Returns: access URL or empty string if disabled.
	// Experimental.
	Bind(scope constructs.Construct, id *string, lambdaFunction awslambda.Function) *string
}

// The jsii proxy struct for LambdaAccess
type jsiiProxy_LambdaAccess struct {
	_ byte // padding
}

// Experimental.
func NewLambdaAccess_Override(l LambdaAccess) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.LambdaAccess",
		nil, // no parameters
		l,
	)
}

// Provide access using API Gateway.
//
// This is the most secure option, but requires additional configuration. It allows you to limit access to specific IP addresses and even to a specific VPC.
//
// To limit access to GitHub.com use:
//
// ```
// LambdaAccess.apiGateway({
//   allowedIps: LambdaAccess.githubWebhookIps(),
// });
// ```
//
// Alternatively, get and manually update the list manually with:
//
// ```
// curl https://api.github.com/meta | jq .hooks
// ```.
// Experimental.
func LambdaAccess_ApiGateway(props *ApiGatewayAccessProps) LambdaAccess {
	_init_.Initialize()

	if err := validateLambdaAccess_ApiGatewayParameters(props); err != nil {
		panic(err)
	}
	var returns LambdaAccess

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.LambdaAccess",
		"apiGateway",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Downloads the list of IP addresses used by GitHub.com for webhooks.
//
// Note that downloading dynamic data during deployment is not recommended in CDK. This is a workaround for the lack of a better solution.
// Experimental.
func LambdaAccess_GithubWebhookIps() *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.LambdaAccess",
		"githubWebhookIps",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Provide access using Lambda URL.
//
// This is the default and simplest option. It puts no limits on the requester, but the Lambda functions themselves authenticate every request.
// Experimental.
func LambdaAccess_LambdaUrl() LambdaAccess {
	_init_.Initialize()

	var returns LambdaAccess

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.LambdaAccess",
		"lambdaUrl",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Disables access to the configured Lambda function.
//
// This is useful for the setup function after setup is done.
// Experimental.
func LambdaAccess_NoAccess() LambdaAccess {
	_init_.Initialize()

	var returns LambdaAccess

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.LambdaAccess",
		"noAccess",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaAccess) Bind(scope constructs.Construct, id *string, lambdaFunction awslambda.Function) *string {
	if err := l.validateBindParameters(scope, id, lambdaFunction); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, id, lambdaFunction},
		&returns,
	)

	return returns
}

