package secretmanager

import (
	"encoding/json"
	"fmt"
	"golangapi/awsgo"
	"golangapi/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecrets(secretName string) (models.Secret, error) {
	var secretData models.Secret

	fmt.Println("> pido secreto" + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		fmt.Println("hubo un error" + err.Error())
		return secretData, err
	}

	// con & le dice a go que grabe en la dirección de memoria de la variable que se declaro arriba
	json.Unmarshal([]byte(*key.SecretString), &secretData)

	fmt.Println(">secretdata " + secretName)
	return secretData, nil
}
