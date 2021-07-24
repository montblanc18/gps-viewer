package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/montblanc18/gps-viewer/server/gen/models"
	"github.com/montblanc18/gps-viewer/server/gen/restapi/gpsviewer"
)

func GetDeviceId(p gpsviewer.GetGpsByIDParams) middleware.Responder {
	var resp *models.DeviceGPS
	return gpsviewer.NewGetGpsByIDOK().WithPayload(resp)
}
