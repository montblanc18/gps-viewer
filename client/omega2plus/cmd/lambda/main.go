package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/montblanc18/gps-viewer/client/omega2plus"
)

func main() {
	lambda.Start(omega2plus.Handle)
}
