//go:build !no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

func (i *jsiiProxy_IRunnerRuntimeParameters) validateAddCatchAndCleanUpParameters(state interface{}) error {
	if state == nil {
		return fmt.Errorf("parameter state is required, but nil was provided")
	}
	switch state.(type) {
	case awsstepfunctions.TaskStateBase:
		// ok
	case awsstepfunctions.Parallel:
		// ok
	case awsstepfunctions.Map:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(state) {
			return fmt.Errorf("parameter state must be one of the allowed types: awsstepfunctions.TaskStateBase, awsstepfunctions.Parallel, awsstepfunctions.Map; received %#v (a %T)", state, state)
		}
	}

	return nil
}

