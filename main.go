package main

import (
	"context"
	"golangapi/awsgo"
	"golangapi/handlers"
	"golangapi/models"
	secretmanager "golangapi/secretManager"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(callLambda)
}

func callLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	awsgo.InitAws()
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
	SecretModel, erro := secretmanager.GetSecrets(os.Getenv("SecretName"))

	if erro != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en lalectura del secret" + erro.Error(),
			Headers: map[string]string{
				"Content-type": "application/json",
			},
		}
		return res, nil
	}

	path := strings.Replace(request.PathParameters["awsgolangtuiter"], os.Getenv("UrlPrefix"), "", -1)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtsing"), SecretModel.Jwtsing)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	resApi := handlers.Handler(awsgo.Ctx, request)
	if resApi.CustomResp == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: resApi.Status,
			Body:       resApi.Message,
			Headers: map[string]string{
				"Content-type": "application/json",
			},
		}
		return res, nil
	} else {
		return resApi.CustomResp, nil
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
