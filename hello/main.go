package main

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/aws/aws-lambda-go/lambda"
)

type RequestItem struct {
	Name string `uri:"name" binding:"required"`
	Time string `form:"time"`
}

type Request []RequestItem
type Response struct {
	Complete  bool          `json:"processingComplete"`
	Remaining []RequestItem `json:"remainingItems"`
}

func Handler(ctx context.Context, requestItems Request) (Response, error) {

	temp, err := json.Marshal(requestItems)
	slog.Info(string(temp))

	requestToProcess := requestItems[0]
	slog.Info("Hello", "name", requestToProcess.Name, "time", requestToProcess.Time)

	if err != nil {
		return Response{Complete: true}, err
	}

	resp := Response{
		Complete:  len(requestItems[1:]) == 0,
		Remaining: requestItems[1:],
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
