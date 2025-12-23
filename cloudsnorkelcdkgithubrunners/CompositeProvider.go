package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// A composite runner provider that implements fallback and distribution strategies.
// Experimental.
type CompositeProvider interface {
}

// The jsii proxy struct for CompositeProvider
type jsiiProxy_CompositeProvider struct {
	_ byte // padding
}

// Experimental.
func NewCompositeProvider() CompositeProvider {
	_init_.Initialize()

	j := jsiiProxy_CompositeProvider{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.CompositeProvider",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCompositeProvider_Override(c CompositeProvider) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.CompositeProvider",
		nil, // no parameters
		c,
	)
}

// Creates a weighted distribution runner provider that randomly selects a provider based on weights.
//
// For example, given providers A (weight 10), B (weight 20), C (weight 30):
// - Total weight = 60
// - Probability of selecting A = 10/60 = 16.67%
// - Probability of selecting B = 20/60 = 33.33%
// - Probability of selecting C = 30/60 = 50%
//
// You can use this to distribute load across multiple instance types or availability zones.
// Experimental.
func CompositeProvider_Distribute(scope constructs.Construct, id *string, weightedProviders *[]*WeightedRunnerProvider) ICompositeProvider {
	_init_.Initialize()

	if err := validateCompositeProvider_DistributeParameters(scope, id, weightedProviders); err != nil {
		panic(err)
	}
	var returns ICompositeProvider

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.CompositeProvider",
		"distribute",
		[]interface{}{scope, id, weightedProviders},
		&returns,
	)

	return returns
}

// Creates a fallback runner provider that tries each provider in order until one succeeds.
//
// For example, given providers A, B, C:
// - Try A first
// - If A fails, try B
// - If B fails, try C
//
// You can use this to try spot instance first, and switch to on-demand instances if spot is unavailable.
//
// Or you can use this to try different instance types in order of preference.
// Experimental.
func CompositeProvider_Fallback(scope constructs.Construct, id *string, providers *[]IRunnerProvider) ICompositeProvider {
	_init_.Initialize()

	if err := validateCompositeProvider_FallbackParameters(scope, id, providers); err != nil {
		panic(err)
	}
	var returns ICompositeProvider

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.CompositeProvider",
		"fallback",
		[]interface{}{scope, id, providers},
		&returns,
	)

	return returns
}

