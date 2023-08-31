package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

func Handler(ctx context.Context) (Response, error) {

	resp := Response{
		Message: "Hello World!",
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
