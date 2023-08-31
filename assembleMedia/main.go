package main

import (
	"context"
	"log/slog"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	MediaRequests []MediaRequest `json:"mediaRequests"`
	Bucket        BucketInfo     `json:"bucket"`
}

type MediaRequest struct {
	MediaId     string `json:"mediaId,required"`
	BibleId     string `json:"bibleId"`
	Format      string `json:"format"`
	Id3v2       string `json:"id3v2"`
	ProductCode string `json:"productCode"`
}

type BucketInfo struct {
	Bucket string `json:"bucket,omitempty"`
	Key    string `json:"key,omitempty"`
}

type Response struct {
	Complete    bool       `json:"assembleComplete"`
	NextPayload any        `json:"nextPayload"`
	Bucket      BucketInfo `json:"bucket,omitempty"`
}

func Handler(ctx context.Context, event Event) (response Response, e error) {

	currentMediaRequest := event.MediaRequests[0]
	slog.Info("Media Requested",
		"mediaId", currentMediaRequest.MediaId,
		"bibleId", currentMediaRequest.BibleId,
		"format", currentMediaRequest.Format,
		"id3v2", currentMediaRequest.Id3v2,
		"productCode", currentMediaRequest.ProductCode)

	assembleComplete := len(event.MediaRequests[1:]) == 0
	if assembleComplete {
		return Response{
			Complete:    true,
			NextPayload: BucketInfo{Bucket: "BUCKETNAMEGOESHERE", Key: "ASSEMBLEDFOLDERGOESHERE"},
		}, nil
	} else {
		return Response{
			Complete: false,
			NextPayload: Event{
				MediaRequests: event.MediaRequests[1:],
				Bucket:        BucketInfo{Bucket: "BUCKETNAMEGOESHERE", Key: "ASSEMBLEDFOLDERGOESHERE"},
			},
		}, nil
	}
}

func main() {
	lambda.Start(Handler)
}
