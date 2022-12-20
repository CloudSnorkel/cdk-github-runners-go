// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Defines what execution history events are logged and where they are logged.
// Experimental.
type LogOptions struct {
	// Determines whether execution data is included in your log.
	// Experimental.
	IncludeExecutionData *bool `field:"optional" json:"includeExecutionData" yaml:"includeExecutionData"`
	// Defines which category of execution history events are logged.
	// Experimental.
	Level awsstepfunctions.LogLevel `field:"optional" json:"level" yaml:"level"`
	// The log group where the execution history events will be logged.
	// Experimental.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
}

