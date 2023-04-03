// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners


// Asset to copy into a built image.
// Experimental.
type RunnerImageAsset struct {
	// Path on local system to copy into the image.
	//
	// Can be a file or a directory.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Target path in the built image.
	// Experimental.
	Target *string `field:"required" json:"target" yaml:"target"`
}

