//go:build no_runtime_type_checking

package cloudsnorkelcdkgithubrunners

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaAccess) validateBindParameters(scope constructs.Construct, id *string, lambdaFunction awslambda.Function) error {
	return nil
}

func validateLambdaAccess_ApiGatewayParameters(props *ApiGatewayAccessProps) error {
	return nil
}

