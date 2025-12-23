package cloudsnorkelcdkgithubrunners


// Input to the provider selector Lambda function.
// Experimental.
type ProviderSelectorInput struct {
	// Full GitHub webhook payload (workflow_job event structure with action="queued").
	//
	// * Original labels requested by the workflow job can be found at `payload.workflow_job.labels`.
	// * Repository path (e.g. CloudSnorkel/cdk-github-runners) is at `payload.repository.full_name`.
	// * Commit hash is at `payload.workflow_job.head_sha`.
	// See: https://docs.github.com/en/webhooks/webhook-events-and-payloads?actionType=queued#workflow_job
	//
	// Experimental.
	Payload interface{} `field:"required" json:"payload" yaml:"payload"`
	// Map of available provider node paths to their configured labels.
	//
	// Example: { "MyStack/Small": ["linux", "small"], "MyStack/Large": ["linux", "large"] }.
	// Experimental.
	Providers *map[string]*[]*string `field:"required" json:"providers" yaml:"providers"`
	// Labels that would have been used by default (the selected provider's labels).
	//
	// May be undefined if no provider matched by default.
	// Experimental.
	DefaultLabels *[]*string `field:"optional" json:"defaultLabels" yaml:"defaultLabels"`
	// Provider node path that would have been selected by default label matching.
	//
	// Use this to easily return the default selection: `{ provider: input.defaultProvider, labels: input.defaultLabels }`
	// May be undefined if no provider matched by default.
	// Experimental.
	DefaultProvider *string `field:"optional" json:"defaultProvider" yaml:"defaultProvider"`
}

