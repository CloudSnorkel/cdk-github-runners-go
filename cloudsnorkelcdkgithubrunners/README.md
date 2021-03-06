# GitHub Self-Hosted Runners CDK Constructs

[![NPM](https://img.shields.io/npm/v/@cloudsnorkel/cdk-github-runners?label=npm&logo=npm)](https://www.npmjs.com/package/@cloudsnorkel/cdk-github-runners)
[![PyPI](https://img.shields.io/pypi/v/cloudsnorkel.cdk-github-runners?label=pypi&logo=pypi)](https://pypi.org/project/cloudsnorkel.cdk-github-runners)
[![Maven Central](https://img.shields.io/maven-central/v/com.cloudsnorkel/cdk.github.runners.svg?label=Maven%20Central&logo=java)](https://search.maven.org/search?q=g:%22com.cloudsnorkel%22%20AND%20a:%22cdk.github.runners%22)
[![Go](https://img.shields.io/github/v/tag/CloudSnorkel/cdk-github-runners?color=red&label=go&logo=go)](https://pkg.go.dev/github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners)
[![Nuget](https://img.shields.io/nuget/v/CloudSnorkel.Cdk.Github.Runners?color=red&&logo=nuget)](https://www.nuget.org/packages/CloudSnorkel.Cdk.Github.Runners/)
[![Release](https://github.com/CloudSnorkel/cdk-github-runners/actions/workflows/release.yml/badge.svg)](https://github.com/CloudSnorkel/cdk-github-runners/actions/workflows/release.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](https://github.com/CloudSnorkel/cdk-github-runners/blob/main/LICENSE)

Use this CDK construct to create ephemeral [self-hosted GitHub runners](https://docs.github.com/en/actions/hosting-your-own-runners/about-self-hosted-runners) on-demand inside your AWS account.

* Easy to configure GitHub integration with a web-based interface
* Customizable runners with decent defaults
* Multiple runner configurations controlled by labels
* Everything fully hosted in your account

Self-hosted runners in AWS are useful when:

* You need easy access to internal resources in your actions
* You want to pre-install some software for your actions
* You want to provide some basic AWS API access (but [aws-actions/configure-aws-credentials](https://github.com/marketplace/actions/configure-aws-credentials-action-for-github-actions) has more security controls)

Ephemeral (or on-demand) runners are the [recommended way by GitHub](https://docs.github.com/en/actions/hosting-your-own-runners/autoscaling-with-self-hosted-runners#using-ephemeral-runners-for-autoscaling) for auto-scaling, and they make sure all jobs run with a clean image. Runners are started on-demand. You don't pay unless a job is running.

## API

The best way to browse API documentation is on [Constructs Hub](https://constructs.dev/packages/@cloudsnorkel/cdk-github-runners/). It is available in all supported programming languages.

## Providers

A runner provider creates compute resources on-demand and uses [actions/runner](https://github.com/actions/runner) to start a runner.

|                  | CodeBuild                | Fargate        | Lambda         |
|------------------|--------------------------|----------------|----------------|
| **Time limit**   | 8 hours                  | Unlimited      | 15 minutes     |
| **vCPUs**        | 2, 4, 8, or 72           | 0.25 to 4      | 1 to 6         |
| **RAM**          | 3gb, 7gb, 15gb, or 145gb | 512mb to 30gb  | 128mb to 10gb  |
| **Storage**      | 50gb to 824gb            | 20gb to 200gb  | Up to 10gb     |
| **Architecture** | x86_64, ARM64            | x86_64, ARM64  | x86_64, ARM64  |
| **sudo**         | ???                        | ???              | ???              |
| **Docker**       | ???                        | ???              | ???              |
| **Spot pricing** | ???                        | ???              | ???              |

The best provider to use mostly depends on your current infrastructure. When in doubt, CodeBuild is always a good choice. Execution history and logs are easy to view, and it has no restrictive limits unless you need to run for more than 8 hours.

You can also create your own provider by implementing `IRunnerProvider`.

## Installation

1. Confirm you're using CDK v2
2. Install the appropriate package

   1. [Python](https://pypi.org/project/cloudsnorkel.cdk-github-runners)

      ```
      pip install cloudsnorkel.cdk-github-runners
      ```
   2. [TypeScript or JavaScript](https://www.npmjs.com/package/@cloudsnorkel/cdk-github-runners)

      ```
      npm i @cloudsnorkel/cdk-github-runners
      ```
   3. [Java](https://search.maven.org/search?q=g:%22com.cloudsnorkel%22%20AND%20a:%22cdk.github.runners%22)

      ```xml
      <dependency>
      <groupId>com.cloudsnorkel</groupId>
      <artifactId>cdk.github.runners</artifactId>
      </dependency>
      ```
   4. [Go](https://pkg.go.dev/github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners)

      ```
      go get github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners
      ```
   5. [.NET](https://www.nuget.org/packages/CloudSnorkel.Cdk.Github.Runners/)

      ```
      dotnet add package CloudSnorkel.Cdk.Github.Runners
      ```
3. Use `GitHubRunners` construct in your code (starting with default arguments is fine)
4. Deploy your stack
5. Look for the status command output similar to `aws --region us-east-1 lambda invoke --function-name status-XYZ123 status.json`
6. Execute the status command (you may need to specify `--profile` too) and open the resulting `status.json` file
7. Open the URL in `github.setup.url` from `status.json` or [manually setup GitHub](SETUP_GITHUB.md) integration as an app or with personal access token
8. Run status command again to confirm `github.auth.status` and `github.webhook.status` are OK
9. Trigger a GitHub action that has a `self-hosted` label with `runs-on: [self-hosted, linux, codebuild]` or similar
10. If the action is not successful, see [troubleshooting](#Troubleshooting)

[![Demo](demo-thumbnail.jpg)](https://youtu.be/wlyv_3V8lIw)

## Customizing

The default providers configured by `GitHubRunners` are useful for testing but probably not too much for actual production work. They run in the default VPC or no VPC and have no added IAM permissions. You would usually want to configure the providers yourself.

For example:

```go
let vpc: ec2.Vpc;
let runnerSg: ec2.SecurityGroup;
let dbSg: ec2.SecurityGroup;
let bucket: s3.Bucket;

// create a custom CodeBuild provider
const myProvider = new CodeBuildRunner(this, 'codebuild runner', {
    label: 'my-codebuild',
    vpc: vpc,
    securityGroup: runnerSg,
});
// grant some permissions to the provider
bucket.grantReadWrite(myProvider);
dbSg.connections.allowFrom(runnerSg, ec2.Port.tcp(3306), 'allow runners to connect to MySQL database');

// create the runner infrastructure
new GitHubRunners(this, 'runners', {
    providers: [myProvider],
});
```

Another way to customize runners is by modifying the image used to spin them up. The image contains the [runner](https://github.com/actions/runner), any required dependencies, and integration code with the provider. You may choose to customize this image by adding more packages, for example.

```go
const myBuilder = new CodeBuildImageBuilder(this, 'image builder', {
   dockerfilePath: FargateProvider.LINUX_X64_DOCKERFILE_PATH,
   runnerVersion: RunnerVersion.specific('2.291.0'),
   rebuildInterval: Duration.days(14),
});
myBuilder.setBuildArg('EXTRA_PACKAGES', 'nginx xz-utils');

const myProvider = new FargateProvider(this, 'fargate runner', {
   label: 'customized-fargate',
   vpc: vpc,
   securityGroup: runnerSg,
});

// create the runner infrastructure
new GitHubRunners(stack, 'runners', {
   providers: [myProvider],
});
```

Your workflow will then look like:

```yaml
name: self-hosted example
on: push
jobs:
  self-hosted:
    runs-on: [self-hosted, customized-fargate]
    steps:
      - run: echo hello world
```

## Architecture

![Architecture diagram](architecture.svg)

## Troubleshooting

1. Always start with the status function, make sure no errors are reported, and confirm all status codes are OK
2. Confirm the webhook Lambda was called by visiting the URL in `troubleshooting.webhookHandlerUrl` from `status.json`

   1. If it's not called or logs errors, confirm the webhook settings on the GitHub side
   2. If you see too many errors, make sure you're only sending `workflow_job` events
3. When using GitHub app, make sure there are active installation in `github.auth.app.installations`
4. Check execution details of the orchestrator step function by visiting the URL in `troubleshooting.stepFunctionUrl` from `status.json`

   1. Use the details tab to find the specific execution of the provider (Lambda, CodeBuild, Fargate, etc.)
   2. Every step function execution should be successful, even if the runner action inside it failed

## Other Options

1. [philips-labs/terraform-aws-github-runner](https://github.com/philips-labs/terraform-aws-github-runner) if you're using Terraform
2. [actions-runner-controller/actions-runner-controller](https://github.com/actions-runner-controller/actions-runner-controller) if you're using Kubernetes
