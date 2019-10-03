# takeon-validation-wrangler
## About this repo
Standardised validation wranglers using Go. This is a lambda function that reads the configs from input queue (i.e. Validation Wrangler Queue) and then substitutes the config formulae with the corresponding response value of the question so that the formulae can be run by the 'Generic Validation Solver Lambda'. The output json from this wrangler is sent as a message to output queue (i.e. Validation Runner Queue).

## Input and Output
Input Json format is similar to the 'testJSON' variable defined under 'wrangler_test.go'  
Output Json format is similar to the 'expectedOutputJSON' variable defined under 'wrangler_test.go'

## Environment Variables
INPUT_QUEUE_URL: The URL of the input Validation Wrangler Queue  
OUTPUT_QUEUE_URL: The URL of the output Validation Runner Queue

## Depolyment
This lambda deploymnet is implemented with serverless. 
'make build' command for compilation and build. 
'make deploy' command for 'clean build' and 'deployment' to aws. 
The serverless.yml file reads the values from 'config.dev.json' in the same directory. So variables like ROLE, INPUT_QUEUE_URL, OUTPUT_QUEUE_URL and INPUT_QUEUE_ARN must be correctly set in 'config.dev.json' with corresponding values from AWS before deployment to AWS.  
Lambda Function name - takeon-validation-wrangler-dev-main (It is generated from service name, stage and main function defined in serverless.yml)  
Lambda URL - https://eu-west-2.console.aws.amazon.com/lambda/home?region=eu-west-2#/functions/takeon-validation-wrangler-dev-main?tab=graph