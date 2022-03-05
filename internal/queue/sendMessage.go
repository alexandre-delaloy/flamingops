package queue

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	"github.com/blyndusk/flamingops/pkg/helpers"
)

func main() {
	queue := helpers.EnvVar("QUEUE_NAME")

	// Create a session that gets credential values from ~/.aws/credentials
	// and the default region from ~/.aws/config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
	}))

	// Get URL of queue
	result, err := GetQueueURL(sess, queue)
	if err != nil {
			fmt.Println("Got an error getting the queue URL:")
			fmt.Println(err)
			return
	}

	queueURL := result.QueueUrl

	err = SendMsg(sess, queueURL)
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

func SendMsg(sess *session.Session, queueURL *string, clientName *string, requestedServices []string, requestedRegion *string) error {
	// Create an SQS service client
	// snippet-start:[sqs.go.send_message.call]
	svc := sqs.New(sess)

	_, err := svc.SendMessage(&sqs.SendMessageInput{
			DelaySeconds: aws.Int64(10),
			MessageAttributes: map[string]*sqs.MessageAttributeValue{
					"ClientName": &sqs.MessageAttributeValue{
							DataType:    aws.String("String"),
							StringValue: aws.String(clientName),
					},
					"RequestedServices": &sqs.MessageAttributeValue{
						DataType:    aws.String("String"),
						StringValue: aws.String(requestedServices),
					},
					"RequestedRegion": &sqs.MessageAttributeValue{
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