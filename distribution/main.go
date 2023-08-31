package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/aws/aws-lambda-go/lambda"
)

type BucketInfo struct {
	Bucket string `form:"bucket"`
	Key    string `form:"key"`
}

type Response struct {
	URL string `json:"url"`
}

func Handler(ctx context.Context, bucketInfo BucketInfo) (response Response, e error) {

	temp, err := json.Marshal(bucketInfo)
	slog.Info(string(temp))

	slog.Info("Bucket Info",
		"bucket", bucketInfo.Bucket,
		"key", bucketInfo.Key)

	if err != nil {
		return response, err
	}

	return Response{
		URL: fmt.Sprintf("http://somehost.com/%s.zip", bucketInfo.Key),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
