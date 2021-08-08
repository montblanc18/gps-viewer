package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/montblanc18/gps-viewer/server/gen/models"
	"github.com/montblanc18/gps-viewer/server/gen/restapi/gpsviewer/gps"
)

func GetGpsByDeviceId(p gps.GetGpsByDeviceIDParams) middleware.Responder {
	var resp *models.DeviceGPS

	//ctx := p.HTTPRequest.Context()
	//g, err := fetchGpsByDeviceId(ctx, p.DeviceID)
	//if err != nil {

	//}
	return gps.NewGetGpsByDeviceIDOK().WithPayload(resp)
}
