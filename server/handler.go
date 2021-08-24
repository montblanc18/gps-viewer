package server

import (
	"fmt"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/guregu/dynamo"
	"github.com/montblanc18/gps-viewer/server/gen/models"
	"github.com/montblanc18/gps-viewer/server/gen/restapi/gpsviewer/gps"
	"github.com/montblanc18/gps-viewer/server/gen/restapi/gpsviewer/system"
)

func GetGpsByDeviceId(p gps.GetGpsByDeviceIDParams) middleware.Responder {

	ctx := p.HTTPRequest.Context()
	g, err := fetchGpsByDeviceId(ctx, p.DeviceID)
	if err != nil {
		if err == dynamo.ErrNotFound {
			return gps.NewGetGpsByDeviceIDNotFound().WithPayload(&models.Error{
				Code:    404,
				Message: "Not Found",
			})
		}
		return gps.NewGetGpsByDeviceIDInternalServerError().WithPayload(&models.Error{
			Code:    500,
			Message: fmt.Sprintf("%v", err),
		})
	}
	t, err := strfmt.ParseDateTime(g.RecordedAt)
	if err != nil {
		return gps.NewGetGpsByDeviceIDInternalServerError().WithPayload(&models.Error{
			Code:    500,
			Message: fmt.Sprintf("%v", err),
		})
	}
	return gps.NewGetGpsByDeviceIDOK().WithPayload(&models.DeviceGPS{
		DeviceID:   &g.DeviceID,
		DeviceType: &g.DeviceType,
		Lat:        g.Lat,
		Lng:        g.Lng,
		Signal:     &g.Signal,
		RecordedAt: &t,
	})
}

func RegisterGpsByDeviceId(p gps.RegisterGpsByDeviceIDParams) middleware.Responder {
	ctx := p.HTTPRequest.Context()
	m := DeviceGPS{
		DeviceID:   p.DeviceID,
		DeviceType: *p.Body.DeviceType,
		Lat:        p.Body.Lat,
		Lng:        p.Body.Lng,
		Signal:     *p.Body.Signal,
		RecordedAt: (*p.Body.RecordedAt).String(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// TODO: 400のケースを作る/validationチェック

	if err := insertGpsByDevice(ctx, m); err != nil {
		return gps.NewRegisterGpsByDeviceIDInternalServerError().WithPayload(&models.Error{
			Code:    500,
			Message: fmt.Sprintf("%v", err),
		})
	}

	return gps.NewRegisterGpsByDeviceIDCreated().WithPayload(&models.DeviceGPS{})
}

func GetHealthCheck(p system.GetHealthCheckParams) middleware.Responder {
	return system.NewGetHealthCheckOK().WithPayload(&models.OK{
		Message: "Health Check OK",
	})
}
