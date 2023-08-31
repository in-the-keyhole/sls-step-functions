package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/workspaces/sls-step-functions/common"
)

type Response struct {
	URL string `json:"url"`
}

func Handler(ctx context.Context, bucketInfo common.BucketInfo) (response Response, e error) {

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
