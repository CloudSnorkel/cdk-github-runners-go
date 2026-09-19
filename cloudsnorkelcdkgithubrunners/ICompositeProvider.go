package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

// Interface for composite runner providers that combine multiple sub-providers.
//
// Unlike IRunnerProvider, composite providers do not have connections, grant capabilities,
// or log groups as they delegate to their sub-providers.
//
// Note that this interface cannot be implemented by external code. Use {@link CompositeProvider } factory methods.
// Experimental.
type ICompositeProvider interface {
	constructs.IConstruct
	// GitHub Actions labels used for this provider.
	//
	// These labels are used to identify which provider should spawn a new on-demand runner. Every job sends a webhook with the labels it's looking for
	// based on runs-on. We use match the labels from the webhook with the labels specified here. If all the labels specified here are present in the
	// job's labels, this provider will be chosen and spawn a new runner.
	// Experimental.
	Labels() *[]*string
	// All sub-providers contained in this composite provider.
	//
	// This is used to extract providers for metric filters and other operations.
	// Experimental.
	Providers() *[]IRunnerProvider
}

// The jsii proxy for ICompositeProvider
type jsiiProxy_ICompositeProvider struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICompositeProvider) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICompositeProvider) Providers() *[]IRunnerProvider {
	var returns *[]IRunnerProvider
	_jsii_.Get(
		j,
		"providers",
		&returns,
	)
	return returns
}

