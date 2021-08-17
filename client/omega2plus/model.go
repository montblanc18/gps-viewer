package omega2plus

import (
	"context"

	"github.com/montblanc18/gps-viewer/client/gen/models"
)

type Uplink struct {
	Payloads Payloads `json:"payloads" validate:"required"`
}

type Payloads struct {
	Age        int    `json:"age"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Elevation  string `json:"elevation"`
	Course     string `json:"course"`
	Speed      string `json:"speed"`
	RecordedAt string `json:"recorded_at"`
}

func (u Uplink) conv(ctx context.Context) *models.RegisterDeviceGPS {
	return &models.RegisterDeviceGPS{}
}
