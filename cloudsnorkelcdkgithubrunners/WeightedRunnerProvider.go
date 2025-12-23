package cloudsnorkelcdkgithubrunners


// Configuration for weighted distribution of runners.
// Experimental.
type WeightedRunnerProvider struct {
	// The runner provider to use.
	// Experimental.
	Provider IRunnerProvider `field:"required" json:"provider" yaml:"provider"`
	// Weight for this provider.
	//
	// Higher weights mean higher probability of selection.
	// Must be a positive number.
	// Experimental.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

