package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/workspaces/sls-step-functions/common"
)

type Response struct {
	URL string `json:"url"`
}

func Handler(ctx context.Context, bucketInfo common.BucketInfo) (response Response, e error) {

	slog.Info("Bucket Info",
		"bucket", bucketInfo.Bucket,
		"key", bucketInfo.Key)

	return Response{
		URL: fmt.Sprintf("http://somehost.com/%s.zip", bucketInfo.Key),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
