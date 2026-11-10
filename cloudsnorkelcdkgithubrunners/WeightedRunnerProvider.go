package cloudsnorkelcdkgithubrunners


// Configuration for weighted distribution of runners.
type WeightedRunnerProvider struct {
	// The runner provider to use.
	Provider IRunnerProvider `field:"required" json:"provider" yaml:"provider"`
	// Weight for this provider.
	//
	// Higher weights mean higher probability of selection.
	// Must be a positive number.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

