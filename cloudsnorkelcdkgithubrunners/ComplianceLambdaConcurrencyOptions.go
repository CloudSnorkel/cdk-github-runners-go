package cloudsnorkelcdkgithubrunners


// Reserved concurrency for Lambda functions used by runners.
type ComplianceLambdaConcurrencyOptions struct {
	// How many jobs you expect to start at the same time.
	//
	// Functions called for every job, like the webhook handler, can run this many at once.
	// If more jobs start at once, they wait a bit longer for a runner.
	//
	// Status, setup, warm runner and deployment helper functions are limited to 5.
	// Default: 20.
	//
	MaxJobStarts *float64 `field:"optional" json:"maxJobStarts" yaml:"maxJobStarts"`
	// How many jobs each Lambda runner provider can run at the same time.
	// Default: 20.
	//
	MaxLambdaJobs *float64 `field:"optional" json:"maxLambdaJobs" yaml:"maxLambdaJobs"`
}

