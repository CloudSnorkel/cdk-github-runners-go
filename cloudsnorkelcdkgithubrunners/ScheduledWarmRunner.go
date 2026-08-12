package cloudsnorkelcdkgithubrunners

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-github-runners-go/cloudsnorkelcdkgithubrunners/internal"
)

// Warm runners active during a time window specified by start time (`schedule`) and duration (`duration`).
//
// Runners will be provisioned using the specified provider and registered in the specified repository or organization.
//
// Registration level must match the one selected during setup.
// See: https://github.com/CloudSnorkel/cdk-github-runners/blob/main/SETUP_GITHUB.md
//
// **Limitations**
//
// - **No deployment-fill**: Unlike `AlwaysOnWarmRunner`, scheduled warm runners do not get an initial
// fill on deploy. The first fill happens at the next schedule occurrence. If you deploy at 1pm for
// a 2pm schedule, runners will not appear until 2pm.
// - Jobs will still trigger provisioning of on-demand runners, even if a warm runner ends up being used.
// - You may briefly see more than `count` runners when changing config or at rotation.
// - To remove: set `count` to 0, deploy, wait for warm runners to stop, then remove and deploy again.
// If you don't follow this procedure, warm runners may linger until they expire.
// - Provider failures or timeouts (like Lambda provider timing out after 15 minutes) will result in a
// gap in coverage until the retry succeeds. Current retry mechanism has built-in back-off rate and
// can be tweaked using `retryOptions`. This will be improved in the future.
//
// ```typescript
// // Cron: fill at 1pm on weekdays
// new ScheduledWarmRunner(stack, 'Business Hours', {
// runners,
// provider: myProvider,
// count: 3,
// owner: 'my-org',
// repo: 'my-repo',
// schedule: events.Schedule.cron({ hour: '13', minute: '0', weekDay: 'MON-FRI' }),
// duration: cdk.Duration.hours(2),
// });
// ```
//
// ```typescript
// // Rate: fill every 12 hours
// new ScheduledWarmRunner(stack, 'Every 12 Hours', {
// runners,
// provider: myProvider,
// count: 2,
// owner: 'my-org',
// repo: 'my-repo',
// schedule: events.Schedule.rate(cdk.Duration.hours(5)),
// duration: cdk.Duration.hours(12),
// });
// ```.
//
// Experimental.
type ScheduledWarmRunner interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for ScheduledWarmRunner
type jsiiProxy_ScheduledWarmRunner struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ScheduledWarmRunner) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewScheduledWarmRunner(scope constructs.Construct, id *string, props *ScheduledWarmRunnerProps) ScheduledWarmRunner {
	_init_.Initialize()

	if err := validateNewScheduledWarmRunnerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ScheduledWarmRunner{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.ScheduledWarmRunner",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewScheduledWarmRunner_Override(s ScheduledWarmRunner, scope constructs.Construct, id *string, props *ScheduledWarmRunnerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-github-runners.ScheduledWarmRunner",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func ScheduledWarmRunner_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateScheduledWarmRunner_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-github-runners.ScheduledWarmRunner",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduledWarmRunner) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduledWarmRunner) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		s,
		"with",
		args,
		&returns,
	)

	return returns
}

