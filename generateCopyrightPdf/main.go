package main

import (
	"context"
	"log/slog"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jung-kurt/gofpdf"
	"github.com/workspaces/sls-step-functions/common"
)

type Event struct {
	MediaRequests []MediaRequest    `json:"mediaRequests"`
	Bucket        common.BucketInfo `json:"bucket"`
}

type MediaRequest struct {
	MediaId     string `json:"mediaId,required"`
	BibleId     string `json:"bibleId"`
	Format      string `json:"format"`
	Id3v2       string `json:"id3v2"`
	ProductCode string `json:"productCode"`
}

type Response struct {
	Complete    bool              `json:"copyrightPdfComplete"`
	NextPayload any               `json:"nextPayload"`
	Bucket      common.BucketInfo `json:"bucket,omitempty"`
}

func Handler(ctx context.Context, event Event) (response Response, e error) {

	currentMediaRequest := event.MediaRequests[0]
	slog.Info("Copyright PDF Requested",
		"mediaId", currentMediaRequest.MediaId,
		"bibleId", currentMediaRequest.BibleId,
		"format", currentMediaRequest.Format,
		"id3v2", currentMediaRequest.Id3v2,
		"productCode", currentMediaRequest.ProductCode)

	// Create a new PDF document
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add a page to the PDF
	pdf.AddPage()

	// Set font and size
	pdf.SetFont("Arial", "B", 16)

	// Add some text to the PDF
	pdf.Cell(40, 10, "Hello, PDF!")

	// Output the PDF to a file
	err := pdf.OutputFileAndClose("example.pdf")
	if err != nil {
		return Response{
			Complete: false,
			NextPayload: Event{
				MediaRequests: event.MediaRequests[1:],
				Bucket:        common.BucketInfo{Bucket: "BUCKETNAMEGOESHERE", Key: "ASSEMBLEDFOLDERGOESHERE"},
			},
		}, nil
	} else {
		return Response{
			Complete:    true,
			NextPayload: common.BucketInfo{Bucket: "BUCKETNAMEGOESHERE", Key: "ASSEMBLEDFOLDERGOESHERE"},
		}, nil
	}

}

func main() {
	lambda.Start(Handler)
}
