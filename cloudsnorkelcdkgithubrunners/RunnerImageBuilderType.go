package cloudsnorkelcdkgithubrunners


// Experimental.
type RunnerImageBuilderType string

const (
	// Build runner images using AWS CodeBuild.
	//
	// Faster than AWS Image Builder, but can only be used to build Linux Docker images.
	// Experimental.
	RunnerImageBuilderType_CODE_BUILD RunnerImageBuilderType = "CODE_BUILD"
	// Build runner images using AWS Image Builder.
	//
	// Slower than CodeBuild, but can be used to build any type of image including AMIs and Windows images.
	// Experimental.
	RunnerImageBuilderType_AWS_IMAGE_BUILDER RunnerImageBuilderType = "AWS_IMAGE_BUILDER"
)

