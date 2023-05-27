package cloudsnorkelcdkgithubrunners


// Experimental.
type RunnerImageComponentCustomProps struct {
	// Assets to copy into the built image.
	// Experimental.
	Assets *[]*RunnerImageAsset `field:"optional" json:"assets" yaml:"assets"`
	// Commands to run in the built image.
	// Experimental.
	Commands *[]*string `field:"optional" json:"commands" yaml:"commands"`
	// Docker commands to run in the built image.
	//
	// For example: `['ENV foo=bar', 'RUN echo $foo']`
	//
	// These commands are ignored when building AMIs.
	// Experimental.
	DockerCommands *[]*string `field:"optional" json:"dockerCommands" yaml:"dockerCommands"`
	// Component name used for (1) image build logging and (2) identifier for {@link ImageRunnerBuilder.removeComponent }.
	//
	// Name must only contain alphanumeric characters and dashes.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

