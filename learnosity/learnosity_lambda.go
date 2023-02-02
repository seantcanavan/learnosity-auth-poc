package learnosity

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/seantcanavan/lambda_jwt_router/lambda_router"
	"net/http"
)

func AuthLambda(_ context.Context, lambdaReq events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var aReq BuildLearnosityReq
	err := lambda_router.UnmarshalReq(lambdaReq, true, &aReq)
	if err != nil {
		return lambda_router.StatusAndErrorRes(http.StatusInternalServerError, err)
	}

	authRes, httpStatus, err := BuildLearnosityRequest(&aReq)
	if err != nil {
		return lambda_router.StatusAndErrorRes(httpStatus, err)
	}

	return lambda_router.SuccessRes(authRes)
}
