package omega2plus

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/montblanc18/gps-viewer/client/gen/client"
	"github.com/montblanc18/gps-viewer/client/gen/client/gps"
	zlog "github.com/rs/zerolog/log"
)

var (
	API     *client.Gpsviewer
	apiHost string
)

func init() {
	apiHost = os.Getenv("API_HOST")
	if apiHost == "" {
		zlog.Fatal().Msgf("missing env variable: API_HOST")
	}

	schema := client.DefaultSchemes
	if strings.ToLower(os.Getenv("CLIENT_TRANSPORT_SCHEMA")) == "https" {
		schema = []string{"https"}
	}

	basePath := os.Getenv("BASE_PATH")

	// Initialize API Client
	hc := &http.Client{}

	transport := httptransport.NewWithClient(apiHost, basePath, schema, hc)
	API = client.New(transport, nil)
}

func Handle(ctx context.Context, e events.KinesisEvent) {
	zlog.Info().Msgf("START lambda: size=%v", len(e.Records))
	for _, r := range e.Records {
		var u Uplink
		if err := json.Unmarshal(r.Kinesis.Data, &u); err != nil {
			zlog.Warn().Msgf("invalid data: err=%v, payloards=%v", err, r.Kinesis.Data)
		}

		body, err := u.conv(ctx)
		if err != nil {
			zlog.Warn().Msgf("failed to conv Uplink: err=%v", err)
		}
		params := gps.NewRegisterGpsByDeviceIDParamsWithContext(ctx).WithBody(body)
		if _, err = API.Gps.RegisterGpsByDeviceID(params); err != nil {
			zlog.Warn().Msgf("failed to register GpsByDeviceID: err=%v", err)
		}

	}
}
