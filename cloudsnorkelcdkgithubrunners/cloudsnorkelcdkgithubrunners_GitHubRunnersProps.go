// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for GitHubRunners.
// Experimental.
type GitHubRunnersProps struct {
	// Allow management functions to run in public subnets.
	//
	// Lambda Functions in a public subnet can NOT access the internet.
	// Experimental.
	AllowPublicSubnet *bool `field:"optional" json:"allowPublicSubnet" yaml:"allowPublicSubnet"`
	// Path to a directory containing a file named certs.pem containing any additional certificates required to trust GitHub Enterprise Server. Use this when GitHub Enterprise Server certificates are self-signed.
	//
	// You may also want to use custom images for your runner providers that contain the same certificates. See {@link CodeBuildImageBuilder.addCertificates}.
	//
	// ```typescript
	// const imageBuilder = new CodeBuildImageBuilder(this, 'Image Builder with Certs', {
	//      dockerfilePath: CodeBuildRunner.LINUX_X64_DOCKERFILE_PATH,
	// });
	// imageBuilder.addExtraCertificates('path-to-my-extra-certs-folder');
	//
	// const provider = new CodeBuildRunner(this, 'CodeBuild', {
	//      imageBuilder: imageBuilder,
	// });
	//
	// new GitHubRunners(
	//    this,
	//    'runners',
	//    {
	//      providers: [provider],
	//      extraCertificates: 'path-to-my-extra-certs-folder',
	//    }
	// );
	// ```.
	// Experimental.
	ExtraCertificates *string `field:"optional" json:"extraCertificates" yaml:"extraCertificates"`
	// List of runner providers to use.
	//
	// At least one provider is required. Provider will be selected when its label matches the labels requested by the workflow job.
	// Experimental.
	Providers *[]IRunnerProvider `field:"optional" json:"providers" yaml:"providers"`
	// Security group attached to all management functions.
	//
	// Use this with to provide access to GitHub Enterprise Server hosted inside a VPC.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// VPC used for all management functions.
	//
	// Use this with GitHub Enterprise Server hosted that's inaccessible from outside the VPC.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// VPC subnets used for all management functions.
	//
	// Use this with GitHub Enterprise Server hosted that's inaccessible from outside the VPC.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

