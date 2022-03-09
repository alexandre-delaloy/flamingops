package queue

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	"github.com/blyndusk/flamingops/pkg/models"
	"github.com/blyndusk/flamingops/pkg/helpers"
)

func HandleMessageCreation(user *models.User) {
	queue := helpers.EnvVar("QUEUE_NAME")

	// Create a session that gets credential values from ~/.aws/credentials
	// and the default region from ~/.aws/config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Get URL of queue
	result, err := GetQueueURL(sess, &queue)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		fmt.Println(err)
		return
	}

	queueURL := result.QueueUrl

	requestedServices := []*string
	for _, service := range user.ActiveServices.AwsServices {
		requestedServices = append(requestedServices, aws.String(service))
	}
	for _, service := range user.ActiveServices.SwServices {
		requestedServices = append(requestedServices, aws.String(service))
	}

	err = SendMsg(sess, queueURL, user, requestedServices, "test")
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

func SendMsg(sess *session.Session, queueURL *string, clientName string, requestedServices []*string, requestedRegion string) error {
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
			"RequestedServices": {
				DataType:         aws.String("StringList"),
				StringListValues: requestedServices,
			},
			"RequestedRegion": {
				DataType:    aws.String("String"),
				StringValue: aws.String(requestedRegion),
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
