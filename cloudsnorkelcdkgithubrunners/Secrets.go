// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

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

	if err := validateNewSecretsParameters(scope, id); err != nil {
		panic(err)
	}
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

	if err := validateSecrets_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

