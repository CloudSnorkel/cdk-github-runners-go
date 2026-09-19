package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Retry options for providers.
//
// The default is to retry 210 times for a bit over 24 hours with increasing interval.
//
// Retries use full jitter, so every wait is a random time between zero and the calculated interval. This spreads out
// runners that all failed at the same time, so they don't hit the same missing capacity or API quota together again.
// It also means the average wait is half the calculated interval, and that's what the 24 hours are calculated from.
// Experimental.
type ProviderRetryOptions struct {
	// Multiplication for how much longer the wait interval gets on every retry.
	// Default: 2.
	//
	// Experimental.
	BackoffRate *float64 `field:"optional" json:"backoffRate" yaml:"backoffRate"`
	// How much time to wait after first retryable failure.
	//
	// This interval will be multiplied by {@link backoffRate} each retry, up to {@link maxDelay}.
	// Default: 1 minute.
	//
	// Experimental.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// How many times to retry.
	// Default: 210.
	//
	// Experimental.
	MaxAttempts *float64 `field:"optional" json:"maxAttempts" yaml:"maxAttempts"`
	// Maximum wait between retries.
	//
	// Without it, exponential backoff quickly grows to hours between attempts, so a job
	// can end up waiting hours for a runner even though capacity came back minutes after it failed.
	//
	// Don't go too low either. A lower maximum needs more attempts to cover the same 24 hours, and every attempt adds
	// to the execution history that Step Functions caps at 25,000 events.
	// Default: 15 minutes.
	//
	// Experimental.
	MaxDelay awscdk.Duration `field:"optional" json:"maxDelay" yaml:"maxDelay"`
	// Set to true to retry provider on supported failures.
	//
	// Which failures generate a retry depends on the specific provider.
	// Default: true.
	//
	// Experimental.
	Retry *bool `field:"optional" json:"retry" yaml:"retry"`
}

