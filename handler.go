package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var region = os.Getenv("AWS_REGION")

// LambdaResponse - Contain all the information the lambda function returns to the caller
// type LambdaResponse struct {
// 	StatusCode int                     `json:"statusCode"`
// 	Headers    map[string]string       `json:"headers"`
// 	Body       ValidationOutputWrapper `json:"body"`
// }

// HandleLambdaEvent - Main entry point to call the wrangling process
func HandleLambdaEvent(config Config) {

	outputJSON, err := Wrangle(config)

	// return LambdaResponse{
	// 		StatusCode: 1,
	// 		Headers:    map[string]string{"Content-Type": "application/json"},
	// 		Body:       outputJSON,
	// 	},
	// 	err
	// fmt.Printf("outputJSON %v\n", outputJSON)
	DataToOutput, err := json.Marshal(outputJSON)
	if err != nil {
		fmt.Printf("An error occured while marshaling DataToOutput: %s", err)
	}
	fmt.Printf("DataToOutput %v\n", string(DataToOutput))

	AwsConfig := &aws.Config{
		Region: aws.String(region),
	}

	sess := session.New(AwsConfig)

	svc := sqs.New(sess)
	// send message
	OutputQueueUrl := os.Getenv("OUTPUT_QUEUE_URL")

	sendParams := &sqs.SendMessageInput{
		MessageBody: aws.String(string(DataToOutput)), // Required
		QueueUrl:    aws.String(OutputQueueUrl),
	}

	sendResp, err := svc.SendMessage(sendParams)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("[Send message] \n%v \n\n", sendResp)

}

func main() {

	lambda.Start(handler)
	// lambda.Start(HandleLambdaEvent)

	// fmt.Println("In queue wrangler")
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))

	// svc := sqs.New(sess)

	// // URL to our queue
	// qURL := "https://sqs.eu-west-2.amazonaws.com/014669633018/WranglerQueue"

	// result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
	// 	AttributeNames: []*string{
	// 		aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
	// 	},
	// 	MessageAttributeNames: []*string{
	// 		aws.String(sqs.QueueAttributeNameAll),
	// 	},
	// 	QueueUrl:            &qURL,
	// 	MaxNumberOfMessages: aws.Int64(1),
	// 	VisibilityTimeout:   aws.Int64(20), // 20 seconds
	// 	WaitTimeSeconds:     aws.Int64(0),
	// })

	// if err != nil {
	// 	fmt.Println("Error", err)
	// 	return
	// }

	// if len(result.Messages) == 0 {
	// 	fmt.Println("Received no messages")
	// 	return
	// }

	// resultDelete, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
	// 	QueueUrl:      &qURL,
	// 	ReceiptHandle: result.Messages[0].ReceiptHandle,
	// })

	// if err != nil {
	// 	fmt.Println("Delete Error", err)
	// 	return
	// }

	// fmt.Println("Message Deleted", resultDelete)
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {

	config := Config{}

	if len(sqsEvent.Records) == 0 {
		return errors.New("No SQS message passed to function")
	}

	for _, msg := range sqsEvent.Records {
		fmt.Printf("Got SQS message %q with body %q\n", msg.MessageId, msg.Body)
		err := json.Unmarshal([]byte(msg.Body), &config)
		if err != nil {
			fmt.Println(err.Error())
		}

		// outputJSON, err := Wrangle(config)
		// return outputJSON, err
		HandleLambdaEvent(config)
	}

	return nil
}
