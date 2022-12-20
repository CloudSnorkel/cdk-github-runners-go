// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Retry options for providers.
//
// The default is to retry 10 times for about 45 minutes with increasing interval.
// Experimental.
type ProviderRetryOptions struct {
	// Multiplication for how much longer the wait interval gets on every retry.
	// Experimental.
	BackoffRate *float64 `field:"optional" json:"backoffRate" yaml:"backoffRate"`
	// How much time to wait after first retryable failure.
	//
	// This interval will be multiplied by {@link backoffRate} each retry.
	// Experimental.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// How many times to retry.
	// Experimental.
	MaxAttempts *float64 `field:"optional" json:"maxAttempts" yaml:"maxAttempts"`
	// Set to true to retry provider on supported failures.
	//
	// Which failures generate a retry depends on the specific provider.
	// Experimental.
	Retry *bool `field:"optional" json:"retry" yaml:"retry"`
}

