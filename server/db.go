package server

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/montblanc18/gps-viewer/server/gen/models"
)

var (
	gdb      *dynamo.DB
	region   string
	gpsTable string
)

func init() {

	region = os.Getenv("AWS_REGION")

	gpsTable = os.Getenv("DYNAMO_TABLE_gps")
	if gpsTable == "" {
		log.Fatal("missing env variable: DYNAMO_TABLE_USERS")
	}

	gdb = dynamo.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	})))
}

func fetchGpsById(ctx context.Context, deviceID string) (models.DeviceGPS, error) {
	var resp models.DeviceGPS
	err := gdb.Table(gpsTable).Get("device_id", deviceID).OneWithContext(ctx, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
