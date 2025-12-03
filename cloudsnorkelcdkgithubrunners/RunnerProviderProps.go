package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Common properties for all runner providers.
// Experimental.
type RunnerProviderProps struct {
	// Add default labels based on OS and architecture of the runner.
	//
	// This will tell GitHub Runner to add default labels like `self-hosted`, `linux`, `x64`, and `arm64`.
	// Default: true.
	//
	// Experimental.
	DefaultLabels *bool `field:"optional" json:"defaultLabels" yaml:"defaultLabels"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Default: logs.RetentionDays.ONE_MONTH
	//
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// Deprecated: use {@link retryOptions } on {@link GitHubRunners } instead.
	RetryOptions *ProviderRetryOptions `field:"optional" json:"retryOptions" yaml:"retryOptions"`
}

