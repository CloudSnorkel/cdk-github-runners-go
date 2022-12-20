// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Common properties for all runner providers.
// Experimental.
type RunnerProviderProps struct {
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// Options to retry operation in case of failure like missing capacity, or API quota issues.
	// Experimental.
	RetryOptions *ProviderRetryOptions `field:"optional" json:"retryOptions" yaml:"retryOptions"`
}

