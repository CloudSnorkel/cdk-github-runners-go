package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Retry options for providers.
//
// The default is to retry 23 times for about 24 hours with increasing interval.
// Experimental.
type ProviderRetryOptions struct {
	// Multiplication for how much longer the wait interval gets on every retry.
	// Default: 1.3
	//
	// Experimental.
	BackoffRate *float64 `field:"optional" json:"backoffRate" yaml:"backoffRate"`
	// How much time to wait after first retryable failure.
	//
	// This interval will be multiplied by {@link backoffRate} each retry.
	// Default: 1 minute.
	//
	// Experimental.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// How many times to retry.
	// Default: 23.
	//
	// Experimental.
	MaxAttempts *float64 `field:"optional" json:"maxAttempts" yaml:"maxAttempts"`
	// Set to true to retry provider on supported failures.
	//
	// Which failures generate a retry depends on the specific provider.
	// Default: true.
	//
	// Experimental.
	Retry *bool `field:"optional" json:"retry" yaml:"retry"`
}

