package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(callLambda)
}

func callLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *event.APIGatewayProxyResponse

	if !ValidateParams() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las variables de entorno, deben incluir secret bucket y prefix",
			Headers: map[string]string{
				"Content-type": "application/json",
			},
		}
		return res, nil
	}

}

func ValidateParams() bool {
	_, getParam := os.LookupEnv("SecretName")

	if !getParam {
		return getParam
	}

	_, getParam = os.LookupEnv("BucketName")

	if !getParam {
		return getParam
	}

	_, getParam = os.LookupEnv("UrlPrefix")

	if !getParam {
		return getParam
	}

	return getParam
}
