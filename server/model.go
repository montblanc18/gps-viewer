package server

import "time"

type DeviceGPS struct {
	DeviceID   string    `dynamo:"device_id"`
	DeviceType string    `dynamo:"device_type"`
	Lng        float64   `dynamo:"lng"`
	Lat        float64   `dynamo:"lat"`
	RecordedAt string    `dynamo:"recorded_at"`
	CreatedAt  time.Time `dynamo:"created_at, omitempty"`
	UpdatedAt  time.Time `dynamo:"updated_at, omitempty"`
}
