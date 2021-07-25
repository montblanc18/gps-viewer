package server

import (
	"context"
	"fmt"
	"os"

	zlog "github.com/rs/zerolog/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/montblanc18/gps-viewer/server/gen/models"
)

var (
	gdb            *dynamo.DB
	region         string
	deviceGpsTable string
)

func init() {

	region = os.Getenv("AWS_REGION")

	deviceGpsTable = os.Getenv("DYNAMO_TABLE_DEVICE_GPS")
	if deviceGpsTable == "" {
		zlog.Fatal().Msgf("missing env variable: DYNAMO_TABLE_DEVICE_GPS")
	}

	dbEndpoint := os.Getenv("DYNAMO_ENDPOINT")
	if dbEndpoint != "" {
		fmt.Printf("DYNAMO_ENDPOINT is set. %s", dbEndpoint)
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			Profile:           "local",
			SharedConfigState: session.SharedConfigEnable,
			Config: aws.Config{
				Endpoint: aws.String(dbEndpoint),
			},
		}))
		gdb = dynamo.New(sess)
	} else {
		zlog.Info().Msgf("DYNAMO_ENDPOINT is not set.")
		gdb = dynamo.New(session.Must(session.NewSession(&aws.Config{
			Region: aws.String(region),
		})))
	}
}

func fetchGpsById(ctx context.Context, deviceID string) (models.DeviceGPS, error) {
	var resp models.DeviceGPS
	err := gdb.Table(deviceGpsTable).Get("device_id", deviceID).OneWithContext(ctx, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
