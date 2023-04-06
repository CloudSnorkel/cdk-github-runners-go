// CDK construct to create GitHub Actions self-hosted runners. A webhook listens to events and creates ephemeral runners on the fly.
package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
	// You may also want to use custom images for your runner providers that contain the same certificates. See {@link CodeBuildImageBuilder.addCertificates }.
	//
	// ```typescript
	// const imageBuilder = CodeBuildRunnerProvider.imageBuilder(this, 'Image Builder with Certs');
	// imageBuilder.addComponent(RunnerImageComponent.extraCertificates('path-to-my-extra-certs-folder/certs.pem', 'private-ca');
	//
	// const provider = new CodeBuildRunnerProvider(this, 'CodeBuild', {
	//     imageBuilder: imageBuilder,
	// });
	//
	// new GitHubRunners(
	//   this,
	//   'runners',
	//   {
	//     providers: [provider],
	//     extraCertificates: 'path-to-my-extra-certs-folder',
	//   }
	// );
	// ```.
	// Experimental.
	ExtraCertificates *string `field:"optional" json:"extraCertificates" yaml:"extraCertificates"`
	// Time to wait before stopping a runner that remains idle.
	//
	// If the user cancelled the job, or if another runner stole it, this stops the runner to avoid wasting resources.
	// Experimental.
	IdleTimeout awscdk.Duration `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// Logging options for the state machine that manages the runners.
	// Experimental.
	LogOptions *LogOptions `field:"optional" json:"logOptions" yaml:"logOptions"`
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
	// Access configuration for the setup function.
	//
	// Once you finish the setup process, you can set this to `LambdaAccess.noAccess()` to remove access to the setup function. You can also use `LambdaAccess.apiGateway({ allowedIps: ['my-ip/0']})` to limit access to your IP only.
	// Experimental.
	SetupAccess LambdaAccess `field:"optional" json:"setupAccess" yaml:"setupAccess"`
	// Access configuration for the status function.
	//
	// This function returns a lot of sensitive information about the runner, so you should only allow access to it from trusted IPs, if at all.
	// Experimental.
	StatusAccess LambdaAccess `field:"optional" json:"statusAccess" yaml:"statusAccess"`
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
	// Access configuration for the webhook function.
	//
	// This function is called by GitHub when a new workflow job is scheduled. For an extra layer of security, you can set this to `LambdaAccess.apiGateway({ allowedIps: LambdaAccess.githubWebhookIps() })`.
	//
	// You can also set this to `LambdaAccess.privateApiGateway()` if your GitHub Enterprise Server is hosted in a VPC. This will create an API Gateway endpoint that's only accessible from within the VPC.
	//
	// *WARNING*: changing access type may change the URL. When the URL changes, you must update GitHub as well.
	// Experimental.
	WebhookAccess LambdaAccess `field:"optional" json:"webhookAccess" yaml:"webhookAccess"`
}

