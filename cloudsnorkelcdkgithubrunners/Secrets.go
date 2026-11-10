package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

// Secrets required for GitHub runners operation.
type Secrets interface {
	constructs.Construct
	// Authentication secret for GitHub containing either app details or personal access token.
	//
	// This secret is used to register runners and
	// cancel jobs when the runner fails to start.
	//
	// This secret is meant to be edited by the user after being created.
	Github() awssecretsmanager.Secret
	// GitHub app private key. Not needed when using personal access tokens.
	//
	// This secret is meant to be edited by the user after being created. It is separate than the main GitHub secret because inserting private keys into JSON is hard.
	GithubPrivateKey() awssecretsmanager.Secret
	// The tree node.
	Node() constructs.Node
	// Setup secret used to authenticate user for our setup wizard.
	//
	// Should be empty after setup has been completed.
	Setup() awssecretsmanager.Secret
	// Webhook secret used to confirm events are coming from GitHub and nowhere else.
	Webhook() awssecretsmanager.Secret
	// Returns a string representation of this construct.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
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

func (s *jsiiProxy_Secrets) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		s,
		"with",
		args,
		&returns,
	)

	return returns
}

