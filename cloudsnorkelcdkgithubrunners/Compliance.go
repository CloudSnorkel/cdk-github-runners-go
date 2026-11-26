package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Settings that compliance frameworks like HIPAA, PCI DSS, and NIST 800-53 ask for, applied to everything this library creates in a scope.
//
// The scope can be an app, a stage, a stack, or a single construct like a runner provider. Settings of the closest scope win, so you can set a key
// for a whole stack and a stricter one for a single provider. Resources you create yourself are not changed.
//
// Once anything uses {@link Compliance}, synth warns about runner resources in the same stage that no scope covers. Like other aspects, an app
// scope doesn't reach into nested stages, so call {@link Compliance.of} on those stages too.
//
// ```typescript
// Compliance.of(stack).encrypt({ key });
// Compliance.of(gpuProvider).encrypt({ storage: gpuKey });
// Compliance.of(stack).limitLambdaConcurrency();
// ```.
type Compliance interface {
	// Encrypt data at rest with your own KMS keys.
	//
	// Keys created in the same app get the key policy they need. For imported keys, a warning lists the statements to add to the key policy.
	Encrypt(options *ComplianceEncryptionOptions)
	// Set reserved concurrency on Lambda functions.
	//
	// This reserves 5 x `maxJobStarts` for functions called on every job, `maxLambdaJobs` for every Lambda runner provider, and 5 for every other
	// function. With the defaults, that's about 120 for a simple deployment and 190 for one using every feature. Reserved concurrency comes out of
	// your account limit, and the account must keep 100 unreserved.
	LimitLambdaConcurrency(options *ComplianceLambdaConcurrencyOptions)
}

// The jsii proxy struct for Compliance
type jsiiProxy_Compliance struct {
	_ byte // padding
}

// Compliance settings for everything this library creates in the given scope.
func Compliance_Of(scope constructs.IConstruct) Compliance {
	_init_.Initialize()

	if err := validateCompliance_OfParameters(scope); err != nil {
		panic(err)
	}
	var returns Compliance

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.Compliance",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Compliance) Encrypt(options *ComplianceEncryptionOptions) {
	if err := c.validateEncryptParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"encrypt",
		[]interface{}{options},
	)
}

func (c *jsiiProxy_Compliance) LimitLambdaConcurrency(options *ComplianceLambdaConcurrencyOptions) {
	if err := c.validateLimitLambdaConcurrencyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"limitLambdaConcurrency",
		[]interface{}{options},
	)
}

