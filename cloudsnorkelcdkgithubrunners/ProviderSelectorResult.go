package cloudsnorkelcdkgithubrunners


// Result from the provider selector Lambda function.
type ProviderSelectorResult struct {
	// Labels to use when registering the runner.
	//
	// Must be returned when a provider is selected.
	// Can be used to add, remove, or modify labels.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Node path of the provider to use (e.g., "MyStack/MyProvider"). Must match one of the configured provider node paths from the input. If not provided, the job will be skipped (no runner created).
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
}

