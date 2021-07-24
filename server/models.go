package server

import "time"

type deviceGPS struct {
	deviceID   string    `dynamodbav:"device_id"`
	deviceType string    `dynamodbav:"device_tyoe"`
	lng        float64   `dynamodbav:"lng"`
	lat        float64   `dynamodbav:"lat"`
	createdAt  time.Time `dynamodbav:"created_at"`
	updatedAt  time.Time `dynamodbav:"updated_at"`
}
