package handlers

import (
	"context"
	"fmt"
	"golangapi/jwt"
	"golangapi/models"
	"golangapi/routers"

	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) models.ResApi {
	var r models.ResApi
	r.Status = 400

	isOK, statusCode, msg, _ := validateAuth(ctx, request)

	if !isOK {
		r.Status = statusCode
		r.Message = msg
		return r
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "register":
			return routers.Register(ctx)
		}
	}
	r.Message = "method invalid"
	return r
}

func validateAuth(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)

	if path == "register" || path == "login" || path == "getAvatar" || path == "getBanner" {
		return true, 200, "", models.Claim{}
	}
	token := request.Headers["Authorization"]

	if len(token) == 0 {
		return false, 400, "token requierido", models.Claim{}
	}

	claim, allOK, msg, err := jwt.Process(token, ctx.Value(models.Key("jwtSing")).(string))

	if !allOK {
		if err != nil {
			fmt.Println("error en el token" + err.Error())
			return false, 401, err.Error(), models.Claim{}
		}
	} else {
		fmt.Println("error en el token" + err.Error())
		return false, 401, err.Error(), models.Claim{}
	}

	fmt.Println("token OK")
	return true, 200, msg, *claim
}
