package server

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/guregu/dynamo"
	"github.com/montblanc18/gps-viewer/server/gen/models"
	"github.com/montblanc18/gps-viewer/server/gen/restapi/gpsviewer/gps"
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
		Lat:        &g.Lat,
		Lng:        &g.Lng,
		RecordedAt: &t,
	})
}
