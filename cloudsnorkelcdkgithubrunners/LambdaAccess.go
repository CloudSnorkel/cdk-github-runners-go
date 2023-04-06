// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"
)

// Access configuration options for Lambda functions like setup and webhook function.
//
// Use this to limit access to these functions.
// Experimental.
type LambdaAccess interface {
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

