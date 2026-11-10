package cloudsnorkelcdkgithubrunners


// Asset to copy into a built image.
type RunnerImageAsset struct {
	// Path on local system to copy into the image.
	//
	// Can be a file or a directory.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Target path in the built image.
	Target *string `field:"required" json:"target" yaml:"target"`
}

