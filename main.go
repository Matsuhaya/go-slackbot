package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type myEvent struct {
	Name string `json:"What is your name?"`
}

type myResponse struct {
	Message string `json:"Answer:"`
}

func hello(event myEvent) (myResponse, error) {
	log.Println("Hello in log")
	return myResponse{Message: fmt.Sprintf("Hello %s!!", event.Name)}, nil
}

func main() {
	lambda.Start(hello)
}
