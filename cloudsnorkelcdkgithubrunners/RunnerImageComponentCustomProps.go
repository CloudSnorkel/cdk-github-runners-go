package cloudsnorkelcdkgithubrunners


type RunnerImageComponentCustomProps struct {
	// Assets to copy into the built image.
	Assets *[]*RunnerImageAsset `field:"optional" json:"assets" yaml:"assets"`
	// Commands to run in the built image.
	Commands *[]*string `field:"optional" json:"commands" yaml:"commands"`
	// Docker commands to run in the built image.
	//
	// For example: `['ENV foo=bar', 'RUN echo $foo']`
	//
	// These commands are ignored when building AMIs.
	DockerCommands *[]*string `field:"optional" json:"dockerCommands" yaml:"dockerCommands"`
	// Component name used for (1) image build logging and (2) identifier for {@link IConfigurableRunnerImageBuilder.removeComponent }.
	//
	// Name must only contain alphanumeric characters and dashes.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

