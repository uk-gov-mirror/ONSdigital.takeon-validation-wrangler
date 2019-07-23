package main

import "github.com/aws/aws-lambda-go/lambda"

// LambdaResponse - Contain all the information the lambda function returns to the caller
// type LambdaResponse struct {
// 	StatusCode int                     `json:"statusCode"`
// 	Headers    map[string]string       `json:"headers"`
// 	Body       ValidationOutputWrapper `json:"body"`
// }

// HandleLambdaEvent - Main entry point to call the wrangling process
func HandleLambdaEvent(config Config) (ValidationOutputWrapper, error) {

	outputJSON, err := Wrangle(config)

	// return LambdaResponse{
	// 		StatusCode: 1,
	// 		Headers:    map[string]string{"Content-Type": "application/json"},
	// 		Body:       outputJSON,
	// 	},
	// 	err
	return outputJSON, err

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
