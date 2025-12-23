package cloudsnorkelcdkgithubrunners

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Properties for GitHubRunners.
// Experimental.
type GitHubRunnersProps struct {
	// Allow management functions to run in public subnets.
	//
	// Lambda Functions in a public subnet can NOT access the internet.
	// Default: false.
	//
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
	// Default: 5 minutes.
	//
	// Experimental.
	IdleTimeout awscdk.Duration `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// Logging options for the state machine that manages the runners.
	// Default: no logs.
	//
	// Experimental.
	LogOptions *LogOptions `field:"optional" json:"logOptions" yaml:"logOptions"`
	// List of runner providers to use.
	//
	// At least one provider is required. Provider will be selected when its label matches the labels requested by the workflow job.
	// Default: CodeBuild, Lambda and Fargate runners with all the defaults (no VPC or default account VPC).
	//
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Optional Lambda function to customize provider selection logic and label assignment.
	//
	// * The function receives the webhook payload along with default provider and its labels as {@link ProviderSelectorInput }
	// * The function returns a selected provider and its labels as {@link ProviderSelectorResult }
	// * You can decline to provision a runner by returning undefined as the provider selector result
	// * You can fully customize the labels for the about-to-be-provisioned runner (add, remove, modify, dynamic labels, etc.)
	// * Labels don't have to match the labels originally configured for the provider, but see warnings below
	// * This function will be called synchronously during webhook processing, so it should be fast and efficient (webhook limit is 30 seconds total)
	//
	// **WARNING: It is your responsibility to ensure the selected provider's labels match the job's required labels. If you return the wrong labels, the runner will be created but GitHub Actions will not assign the job to it.**
	//
	// **WARNING: Provider selection is not a guarantee that a specific provider will be assigned for the job. GitHub Actions may assign the job to any runner with matching labels. The provider selector only determines which provider's runner will be *created*, but GitHub Actions may route the job to any available runner with the required labels.**
	//
	// **For reliable provider assignment based on job characteristics, consider using repo-level runner registration where you can control which runners are available for specific repositories. See {@link SETUP_GITHUB.md } for more details on the different registration levels. This information is also available while using the setup wizard.
	// Experimental.
	ProviderSelector awslambda.IFunction `field:"optional" json:"providerSelector" yaml:"providerSelector"`
	// Whether to require the `self-hosted` label.
	//
	// If `true`, the runner will only start if the workflow job explicitly requests the `self-hosted` label.
	//
	// Be careful when setting this to `false`. Avoid setting up providers with generic label requirements like `linux` as they may match workflows that are not meant to run on self-hosted runners.
	// Default: true.
	//
	// Experimental.
	RequireSelfHostedLabel *bool `field:"optional" json:"requireSelfHostedLabel" yaml:"requireSelfHostedLabel"`
	// Options to retry operation in case of failure like missing capacity, or API quota issues.
	//
	// GitHub jobs time out after not being able to get a runner for 24 hours. You should not retry for more than 24 hours.
	//
	// Total time spent waiting can be calculated with interval * (backoffRate ^ maxAttempts) / (backoffRate - 1).
	// Default: retry 23 times up to about 24 hours.
	//
	// Experimental.
	RetryOptions *ProviderRetryOptions `field:"optional" json:"retryOptions" yaml:"retryOptions"`
	// Security group attached to all management functions.
	//
	// Use this with to provide access to GitHub Enterprise Server hosted inside a VPC.
	// Deprecated: use {@link securityGroups } instead.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Security groups attached to all management functions.
	//
	// Use this with to provide access to GitHub Enterprise Server hosted inside a VPC.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Access configuration for the setup function.
	//
	// Once you finish the setup process, you can set this to `LambdaAccess.noAccess()` to remove access to the setup function. You can also use `LambdaAccess.apiGateway({ allowedIps: ['my-ip/0']})` to limit access to your IP only.
	// Default: LambdaAccess.lambdaUrl()
	//
	// Experimental.
	SetupAccess LambdaAccess `field:"optional" json:"setupAccess" yaml:"setupAccess"`
	// Access configuration for the status function.
	//
	// This function returns a lot of sensitive information about the runner, so you should only allow access to it from trusted IPs, if at all.
	// Default: LambdaAccess.noAccess()
	//
	// Experimental.
	StatusAccess LambdaAccess `field:"optional" json:"statusAccess" yaml:"statusAccess"`
	// VPC used for all management functions. Use this with GitHub Enterprise Server hosted that's inaccessible from outside the VPC.
	//
	// Make sure the selected VPC and subnets have access to the following with either NAT Gateway or VPC Endpoints:
	// * GitHub Enterprise Server
	// * Secrets Manager
	// * SQS
	// * Step Functions
	// * CloudFormation (status function only)
	// * EC2 (status function only)
	// * ECR (status function only).
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
	// You can also set this to `LambdaAccess.apiGateway({allowedVpc: vpc, allowedIps: ['GHES.IP.ADDRESS/32']})` if your GitHub Enterprise Server is hosted in a VPC. This will create an API Gateway endpoint that's only accessible from within the VPC.
	//
	// *WARNING*: changing access type may change the URL. When the URL changes, you must update GitHub as well.
	// Default: LambdaAccess.lambdaUrl()
	//
	// Experimental.
	WebhookAccess LambdaAccess `field:"optional" json:"webhookAccess" yaml:"webhookAccess"`
}

