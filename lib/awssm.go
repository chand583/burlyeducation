package lib

import (
	"encoding/base64"
	"encoding/json"
	"burlyeducation/log"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/beego/beego/v2/core/config"
)

type AWSSecretManager struct{}

type DBConfig struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Engine     string `json:"engine"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	DBName     string `json:"dbname"`
	DBInstance string `json:"dbInstanceIdentifier"`
}

type RedisConfig struct {
	RedisEndpoint string `json:"Redis_endpoint"`
}

type TllmsConfig struct {
	TllmsSecret string `json:"TLLMS-API-KEY"`
}

func (aws AWSSecretManager) GetDBConString() (string, error) {

	secretId, _ := config.String("aws::secret_id_db")
	dbConfigString, err := getSecretFromAWS(secretId)

	if err != nil {
		return "", err
	}

	dbConfig := DBConfig{}
	json.Unmarshal([]byte(dbConfigString), &dbConfig)

	dbConString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)

	return dbConString, nil
}

func (aws AWSSecretManager) GetRedisConString() (string, error) {

	secretId, _ := config.String("aws::secret_id_redis")
	redisPort, _ := config.String("aws::redis_port")

	redisConfigString, err := getSecretFromAWS(secretId)

	if err != nil {
		return "", err
	}

	redisConfig := RedisConfig{}
	json.Unmarshal([]byte(redisConfigString), &redisConfig)

	redisConString := fmt.Sprintf(`{"key":"default", "conn":"%s:%s", "dbNum":"0"}`, redisConfig.RedisEndpoint, redisPort)

	return redisConString, nil
}

func (aws AWSSecretManager) GetTllmsApiSecret() (string, error) {

	secretIdTllms, _ := config.String("aws::secret_id_tllms")

	tllmsAPiSecret, err := getSecretFromAWS(secretIdTllms)

	if err != nil {
		return "", err
	}

	tllmsConfig := TllmsConfig{}
	json.Unmarshal([]byte(tllmsAPiSecret), &tllmsConfig)
	return tllmsConfig.TllmsSecret, nil
}

func getSecretFromAWS(secretName string) (string, error) {

	region, _ := config.String("aws::region")
	versionStage, _ := config.String("aws::version_stage")

	//Create a Secrets Manager client
	svc := secretsmanager.New(session.New(), aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String(versionStage), // VersionStage defaults to AWSCURRENT if unspecified
	}

	// In this sample we only handle the specific exceptions for the 'GetSecretValue' API.
	// See https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html

	result, err := svc.GetSecretValue(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeDecryptionFailure:
				// Secrets Manager can't decrypt the protected secret text using the provided KMS key.
				log.Error(1101, map[string]interface{}{"error_details": secretsmanager.ErrCodeDecryptionFailure, "error_message": aerr.Error()})

			case secretsmanager.ErrCodeInternalServiceError:
				// An error occurred on the server side.
				log.Error(1101, map[string]interface{}{"error_details": secretsmanager.ErrCodeInternalServiceError, "error_message": aerr.Error()})

			case secretsmanager.ErrCodeInvalidParameterException:
				// You provided an invalid value for a parameter.
				log.Error(1101, map[string]interface{}{"error_details": secretsmanager.ErrCodeInvalidParameterException, "error_message": aerr.Error()})

			case secretsmanager.ErrCodeInvalidRequestException:
				// You provided a parameter value that is not valid for the current state of the resource.
				log.Error(1101, map[string]interface{}{"error_details": secretsmanager.ErrCodeInvalidRequestException, "error_message": aerr.Error()})

			case secretsmanager.ErrCodeResourceNotFoundException:
				// We can't find the resource that you asked for.
				log.Error(1101, map[string]interface{}{"error_details": secretsmanager.ErrCodeResourceNotFoundException, "error_message": aerr.Error()})
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Error(1101, map[string]interface{}{"error_details": aerr, "error_message": aerr.Error()})

		}
		return "", err
	}

	// Decrypts secret using the associated KMS CMK.
	// Depending on whether the secret is a string or binary, one of these fields will be populated.
	var secretString, decodedBinarySecret string
	if result.SecretString != nil {
		secretString = *result.SecretString
	} else {
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
		len, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
		if err != nil {
			//fmt.Println("Base64 Decode Error:", err)
			log.Error(err, map[string]interface{}{"error_details": err, "error_message": err.Error()})
			return "", err
		}
		decodedBinarySecret = string(decodedBinarySecretBytes[:len])
	}

	if len(secretString) == 0 && len(decodedBinarySecret) > 0 {
		secretString = decodedBinarySecret
	}

	return secretString, nil
}
