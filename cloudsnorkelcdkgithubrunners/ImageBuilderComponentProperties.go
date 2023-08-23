package cloudsnorkelcdkgithubrunners


// Properties for ImageBuilderComponent construct.
// Experimental.
type ImageBuilderComponentProperties struct {
	// Shell commands to run when adding this component to the image.
	//
	// On Linux, these are bash commands. On Windows, there are PowerShell commands.
	// Experimental.
	Commands *[]*string `field:"required" json:"commands" yaml:"commands"`
	// Component description.
	// Experimental.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Component display name.
	// Experimental.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Component platform.
	//
	// Must match the builder platform.
	// Experimental.
	Platform *string `field:"required" json:"platform" yaml:"platform"`
	// Optional assets to add to the built image.
	// Experimental.
	Assets *[]*ImageBuilderAsset `field:"optional" json:"assets" yaml:"assets"`
	// Require a reboot after installing this component.
	// Default: false.
	//
	// Experimental.
	Reboot *bool `field:"optional" json:"reboot" yaml:"reboot"`
}

