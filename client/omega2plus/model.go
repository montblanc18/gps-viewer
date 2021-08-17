package omega2plus

import (
	"context"
	"strconv"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
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
	DeviceID   string `json:"device_id"`
}

func (u Uplink) conv(ctx context.Context) (*models.RegisterDeviceGPS, error) {
	lng, err := strconv.ParseFloat(u.Payloads.Longitude, 64)
	if err != nil {
		return &models.RegisterDeviceGPS{}, err
	}

	lat, err := strconv.ParseFloat(u.Payloads.Latitude, 64)
	if err != nil {
		return &models.RegisterDeviceGPS{}, err
	}
	recordedAt, err := strfmt.ParseDateTime(u.Payloads.RecordedAt)
	if err != nil {
		return &models.RegisterDeviceGPS{}, err
	}
	deviceType := "omega2plus"
	return &models.RegisterDeviceGPS{
		Lng:        swag.Float64(lng),
		Lat:        swag.Float64(lat),
		DeviceID:   &u.Payloads.DeviceID,
		DeviceType: &deviceType,
		RecordedAt: &recordedAt,
	}, nil
}
