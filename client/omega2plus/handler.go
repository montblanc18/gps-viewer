package omega2plus

import (
	"context"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	zlog "github.com/rs/zerolog/log"
)

var apiHost string

func init() {
	apiHost = os.Getenv("API_HOST")
	if apiHost == "" {
		zlog.Fatal().Msgf("missing env variable: API_HOST")
	}
}

func Handle(ctx context.Context, e events.KinesisEvent) {
	zlog.Info().Msgf("START lambda: size=%v", len(e.Records))
	for _, r := range e.Records {
		var u Uplink
		if err := json.Unmarshal(r.Kinesis.Data, &u); err != nil {
			zlog.Warn().Msgf("invalid data: err=%w, payloards=%v", err, r.Kinesis.Data)
		}
	}
}
