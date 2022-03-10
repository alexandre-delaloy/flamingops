package queue

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	log "github.com/sirupsen/logrus"

	"github.com/blyndusk/flamingops/internal/database"
	"github.com/blyndusk/flamingops/pkg/helpers"
	"github.com/blyndusk/flamingops/pkg/models"
)

func HandleMessageCreation(user *models.User) {
	// stops if the user services have been updated recently
	var awsServicesData models.AwsServicesData
	if err := database.Db.Where("user_id = ?", user.Id).First(&awsServicesData).Error; err != nil {
		log.Error(err)
		return
	}
	var swServicesData models.SwServicesData
	if err := database.Db.Where("user_id = ?", user.Id).First(&swServicesData).Error; err != nil {
		log.Error(err)
		return
	}
	if awsServicesData.UpdatedAt.After(time.Now().Add(-2*time.Minute)) && swServicesData.UpdatedAt.After(time.Now().Add(-2*time.Minute)) {
		return
	}

	var requestedServices []*string
	var activeServices models.ActiveServices
	if err := database.Db.Where("user_id = ?", user.Id).First(&activeServices).Error; err != nil {
		log.Error(err)
		return
	}
	
	if len(activeServices.AwsServices) == 0 && len(activeServices.SwServices) == 0 {
		return
	}
	
	var requestedRegions models.RequestedRegions
	if err := database.Db.Where("user_id = ?", user.Id).First(&requestedRegions).Error; err != nil {
		log.Error(err)
		return
	}
	
	if requestedRegions.AwsRegion == "" && requestedRegions.SwRegion == "" {
		return
	}
	
	fmt.Println("hey")
	// Create a session that gets credential values from ~/.aws/credentials
	// and the default region from ~/.aws/config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Get URL of queue
	queue := helpers.EnvVar("QUEUE_NAME")
	result, err := GetQueueURL(sess, &queue)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		fmt.Println(err)
		return
	}

	queueURL := result.QueueUrl

	for _, service := range activeServices.AwsServices {
		service := string(service)
		requestedServices = append(requestedServices, &service)
	}
	for _, service := range activeServices.SwServices {
		service := string(service)
		requestedServices = append(requestedServices, &service)
	}

	err = SendMsg(sess, queueURL, user.Username, user.Id, requestedServices, requestedRegions.AwsRegion, requestedRegions.SwRegion)
	if err != nil {
		fmt.Println("Got an error sending the message:")
		fmt.Println(err)
		return
	}

	fmt.Println("Sent message to queue ")
}

func GetQueueURL(sess *session.Session, queue *string) (*sqs.GetQueueUrlOutput, error) {
	// Create an SQS service client
	svc := sqs.New(sess)

	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queue,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func SendMsg(sess *session.Session, queueURL *string, clientName string, clientId uint, requestedServices []*string, requestedAwsRegion string, requestedSwRegion string) error {
	// Create an SQS service client
	// snippet-start:[sqs.go.send_message.call]
	svc := sqs.New(sess)

	_, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"ClientName": {
				DataType:    aws.String("String"),
				StringValue: aws.String(clientName),
			},
			"ClientId": {
				DataType:    aws.String("Number"),
				StringValue: aws.String(fmt.Sprintf("%d", clientId)),
			},
			"RequestedServices": {
				DataType:         aws.String("StringList"),
				StringListValues: requestedServices,
			},
			"RequestedAwsRegion": {
				DataType:    aws.String("String"),
				StringValue: aws.String(requestedAwsRegion),
			},
			"RequestedSwRegion": {
				DataType:    aws.String("String"),
				StringValue: aws.String(requestedSwRegion),
			},
		},
		MessageBody: aws.String(""),
		QueueUrl:    queueURL,
	})

	if err != nil {
		return err
	}

	return nil
}
