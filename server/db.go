package server

import (
	"context"
	"fmt"
	"os"

	zlog "github.com/rs/zerolog/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
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

// 最新のGPS情報を取得する
func fetchGpsByDeviceId(ctx context.Context, deviceID string) (DeviceGPS, error) {
	var resps []DeviceGPS
	var resp DeviceGPS
	err := gdb.Table(deviceGpsTable).Get("device_id", deviceID).Order(dynamo.Descending).Limit(1).All(&resps)
	if err != nil {
		return resp, err
	}
	// 一件もなかった
	if len(resps) == 0 {
		return resp, dynamo.ErrNotFound
	}
	resp = resps[0]

	return resp, nil
}

func insertGpsByDevice(ctx context.Context, gps DeviceGPS) error {
	return gdb.Table(deviceGpsTable).Put(gps).RunWithContext(ctx)
}
