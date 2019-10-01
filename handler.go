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

// HandleLambdaEvent - Main entry point to call the wrangling process
func HandleLambdaEvent(config Config) {

	outputJSON, err := Wrangle(config)

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

	// Get the output Queue URL
	OutputQueueURL := os.Getenv("OUTPUT_QUEUE_URL")

	//Generate params and send msg to the output queue
	sendParams := &sqs.SendMessageInput{
		MessageBody: aws.String(string(DataToOutput)),
		QueueUrl:    aws.String(OutputQueueURL),
	}

	sendResp, err := svc.SendMessage(sendParams)

	//If errors in sending to the queue, log it
	if err != nil {
		fmt.Println(err)
	}
	//Log the response MessageId
	fmt.Printf("[Send message] \n%v \n\n", sendResp)

}

//main entry point from lambda that calls handler
func main() {

	lambda.Start(handler)
}

//This function polls the input queue and extracts the message body to sends it to underlying wrangler process
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

		HandleLambdaEvent(config)
	}

	return nil
}
