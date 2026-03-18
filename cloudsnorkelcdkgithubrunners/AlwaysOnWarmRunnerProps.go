package cloudsnorkelcdkgithubrunners


// Properties for always on warm runners.
// Experimental.
type AlwaysOnWarmRunnerProps struct {
	// Number of warm runners to maintain.
	// Experimental.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// GitHub owner where runners will be registered (org or user login).
	// Experimental.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// Provider to use.
	//
	// Warm runners bypass the provider selector — they always use
	// this provider, regardless of job characteristics. Labels cannot be modified.
	// Experimental.
	Provider interface{} `field:"required" json:"provider" yaml:"provider"`
	// The GitHubRunners construct that owns the shared warm runner infrastructure.
	// Experimental.
	Runners GitHubRunners `field:"required" json:"runners" yaml:"runners"`
	// Registration level — must match how your runners are set up in GitHub.
	//
	// Choose
	// 'org' for org-wide runners, 'repo' for repo-level. See the setup wizard or
	// {@link SETUP_GITHUB.md } for choosing repo vs org.
	// Default: 'repo'.
	//
	// Experimental.
	RegistrationLevel *string `field:"optional" json:"registrationLevel" yaml:"registrationLevel"`
	// Repository name (without owner) where runners will be registered.
	//
	// Required when `registrationLevel` is 'repo'.
	// Experimental.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

