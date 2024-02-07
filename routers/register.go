package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"golangapi/db"
	"golangapi/models"
)

func Register(ctx context.Context) models.ResApi {
	var user models.User
	var response models.ResApi
	response.Status = 400

	fmt.Println("entrando al registro")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &user)

	if err != nil {
		response.Message = err.Error()
		fmt.Println(response.Message)
		return response
	}

	if len(user.Email) == 0 {
		response.Message = "El email es necesario"
		fmt.Println(response.Message)
		return response
	}

	if len(user.Password) < 6 {
		response.Message = "contraseña de almenos seis caracteres"
		fmt.Println(response.Message)
		return response
	}

	_, exist, _ := db.ValidateIfExist(user.Email)

	if exist {
		response.Message = "ya viste un usuario con ese email"
		fmt.Println(response.Message)
		return response
	}

	_, status, err := db.InsertUser(user)

	if err != nil {
		response.Message = "error insertando registro"
		fmt.Println(response.Message)
		return response
	}

	if !status {
		response.Message = "status de error"
		fmt.Println(response.Message)
		return response
	}
	return response
}
