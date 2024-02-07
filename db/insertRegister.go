package db

import (
	"context"
	"golangapi/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(user models.User) (string, bool, error) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("usuarios")

	user.Password, _ = EncryptPassword(user.Password)

	result, err := col.InsertOne(ctx, user)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
