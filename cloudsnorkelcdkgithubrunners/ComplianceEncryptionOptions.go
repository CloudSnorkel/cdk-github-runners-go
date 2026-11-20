package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Keys used by {@link Compliance.encrypt}. Each type of resource uses its own key if set, or `key` otherwise. Types with neither are left as they are.
type ComplianceEncryptionOptions struct {
	// Key for ECR repositories holding runner images, and for AMIs built for EC2 runners.
	//
	// EC2 runner volumes are encrypted with the key of their AMI.
	//
	// Adding a key to an existing repository replaces it. The deployment waits for the image to be rebuilt in the new repository before runners use it.
	// Default: `key`.
	//
	Images awskms.IKey `field:"optional" json:"images" yaml:"images"`
	// Key for every type of resource that doesn't have its own key below.
	// Default: only types with their own key are encrypted.
	//
	Key awskms.IKey `field:"optional" json:"key" yaml:"key"`
	// Key for CloudWatch log groups, including runner logs.
	// Default: `key`.
	//
	Logs awskms.IKey `field:"optional" json:"logs" yaml:"logs"`
	// Key for the Step Functions state machine, SQS queues, SNS topics and the DynamoDB table used to manage runners.
	// Default: `key`.
	//
	Messages awskms.IKey `field:"optional" json:"messages" yaml:"messages"`
	// Key for Secrets Manager secrets holding GitHub credentials.
	// Default: `key`.
	//
	Secrets awskms.IKey `field:"optional" json:"secrets" yaml:"secrets"`
	// Key for runner storage: EBS volumes of ECS instances, Fargate ephemeral storage, CodeBuild build output, and buckets used to build images.
	//
	// ECS clusters and capacity providers you pass in are not changed.
	// Default: `key`.
	//
	Storage awskms.IKey `field:"optional" json:"storage" yaml:"storage"`
}

