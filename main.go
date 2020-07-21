package main

import (
	"context"
	"encoding/base64"
	slack "myproject/go-slackbot/slack"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

var functionName string = os.Getenv("AWS_LAMBDA_FUNCTION_NAME")
var encrypted string = os.Getenv("slackURL")
var decrypted string

func init() {
	kmsClient := kms.New(session.New())
	decodedBytes, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		panic(err)
	}
	input := &kms.DecryptInput{
		CiphertextBlob: decodedBytes,
		EncryptionContext: aws.StringMap(map[string]string{
			"LambdaFunctionName": functionName,
		}),
	}
	response, err := kmsClient.Decrypt(input)
	if err != nil {
		panic(err)
	}

	decrypted = string(response.Plaintext[:])
}

func handleRequest(ctx context.Context) (string, error) {
	var slackURL = decrypted
	var slackChannel = os.Getenv("slackChannel")

	sl := slack.NewSlack(slackURL, "This is a test", "test_imp", ":smiling_imp:", "", slackChannel)
	sl.Send()
	return "", nil
}

func main() {
	lambda.Start(handleRequest)
}
