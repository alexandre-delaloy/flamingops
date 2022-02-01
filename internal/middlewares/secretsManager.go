package middlewares

import (
	"net/http"

	"github.com/blyndusk/flamingops/internal/database"
	"github.com/blyndusk/flamingops/pkg/helpers"
	"github.com/blyndusk/flamingops/pkg/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"encoding/base64"
	"fmt"
)

func GetSecretByName(c *gin.Context) {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.GetSecretValueInput{
			SecretId:     aws.String(c.Params.ByName("secretName")),
			VersionStage: aws.String("AWSPREVIOUS"),
	}
	
	result, err := svc.GetSecretValue(input)
	if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
					switch aerr.Code() {
					case secretsmanager.ErrCodeResourceNotFoundException:
							fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
					case secretsmanager.ErrCodeInvalidParameterException:
							fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
					case secretsmanager.ErrCodeInvalidRequestException:
							fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
					case secretsmanager.ErrCodeDecryptionFailure:
							fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())
					case secretsmanager.ErrCodeInternalServiceError:
							fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
					default:
							fmt.Println(aerr.Error())
					}
			} else {
					// Print the error, cast err to awserr.Error to get the Code and
					// Message from an error.
					fmt.Println(err.Error())
			}
			return
	}

	c.JSON(http.StatusOK, result)
	return
}

func CreateSecret(c *gin.Context) {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.CreateSecretInput{
			ClientRequestToken: aws.String(helpers.RandomString(10)),
			Description:        aws.String(""),
			Name:               aws.String(c.Params.ByName("secretName")),
			SecretString:       aws.String(c.Params.ByName("secretString")),
	}

	result, err := svc.CreateSecret(input)
	if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
					switch aerr.Code() {
					case secretsmanager.ErrCodeInvalidParameterException:
							fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
					case secretsmanager.ErrCodeInvalidRequestException:
							fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
					case secretsmanager.ErrCodeLimitExceededException:
							fmt.Println(secretsmanager.ErrCodeLimitExceededException, aerr.Error())
					case secretsmanager.ErrCodeEncryptionFailure:
							fmt.Println(secretsmanager.ErrCodeEncryptionFailure, aerr.Error())
					case secretsmanager.ErrCodeResourceExistsException:
							fmt.Println(secretsmanager.ErrCodeResourceExistsException, aerr.Error())
					case secretsmanager.ErrCodeResourceNotFoundException:
							fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
					case secretsmanager.ErrCodeMalformedPolicyDocumentException:
							fmt.Println(secretsmanager.ErrCodeMalformedPolicyDocumentException, aerr.Error())
					case secretsmanager.ErrCodeInternalServiceError:
							fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
					case secretsmanager.ErrCodePreconditionNotMetException:
							fmt.Println(secretsmanager.ErrCodePreconditionNotMetException, aerr.Error())
					case secretsmanager.ErrCodeDecryptionFailure:
							fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())
					default:
							fmt.Println(aerr.Error())
					}
			} else {
					// Print the error, cast err to awserr.Error to get the Code and
					// Message from an error.
					fmt.Println(err.Error())
			}
			return
	}

	c.JSON(http.StatusOK, result)
	return
}

func UpdateSecret(c *gin.Context){
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.PutSecretValueInput{
			ClientRequestToken: aws.String(helpers.RandomString(10)),
			SecretId:           aws.String(c.Params.ByName("secretName")),
			SecretString:       aws.String(c.Params.ByName("secretString")),
	}

	result, err := svc.PutSecretValue(input)
	if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
					switch aerr.Code() {
					case secretsmanager.ErrCodeInvalidParameterException:
							fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
					case secretsmanager.ErrCodeInvalidRequestException:
							fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
					case secretsmanager.ErrCodeLimitExceededException:
							fmt.Println(secretsmanager.ErrCodeLimitExceededException, aerr.Error())
					case secretsmanager.ErrCodeEncryptionFailure:
							fmt.Println(secretsmanager.ErrCodeEncryptionFailure, aerr.Error())
					case secretsmanager.ErrCodeResourceExistsException:
							fmt.Println(secretsmanager.ErrCodeResourceExistsException, aerr.Error())
					case secretsmanager.ErrCodeResourceNotFoundException:
							fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
					case secretsmanager.ErrCodeInternalServiceError:
							fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
					case secretsmanager.ErrCodeDecryptionFailure:
							fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())
					default:
							fmt.Println(aerr.Error())
					}
			} else {
					// Print the error, cast err to awserr.Error to get the Code and
					// Message from an error.
					fmt.Println(err.Error())
			}
			return
	}

	c.JSON(http.StatusOK, result)
	return
}

func DeleteSecret(c *gin.Context) {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.DeleteSecretInput{
			RecoveryWindowInDays: aws.Int64(7),
			SecretId:             aws.String(c.Params.ByName("secretName")),
	}

	result, err := svc.DeleteSecret(input)
	if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
					switch aerr.Code() {
					case secretsmanager.ErrCodeResourceNotFoundException:
							fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
					case secretsmanager.ErrCodeInvalidParameterException:
							fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
					case secretsmanager.ErrCodeInvalidRequestException:
							fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
					case secretsmanager.ErrCodeInternalServiceError:
							fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
					default:
							fmt.Println(aerr.Error())
					}
			} else {
					// Print the error, cast err to awserr.Error to get the Code and
					// Message from an error.
					fmt.Println(err.Error())
			}
			return
	}

	c.JSON(http.StatusOK, result)
	return
}