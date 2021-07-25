package server

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/montblanc18/gps-viewer/server/gen/models"
)

func TestFetchGpsById(t *testing.T) {

	ctx := context.Background()

	tests := []struct {
		name          string
		deviceID      string
		wantDeviceGps models.DeviceGPS
		wantErr       error
		cmds          []string
	}{
		{
			name:          "正常系",
			deviceID:      "11111",
			wantDeviceGps: models.DeviceGPS{},
			wantErr:       nil,
			cmds: []string{
				`aws --profile local --endpoint-url http://localhost:4566 dynamodb create-table --cli-input-json file://./testdata/schema/local_device_gps.json`,
				`aws --profile local --endpoint-url http://localhost:4566 dynamodb delete-table --table local_device_gps`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdExec(t, tt.cmds)
			got, err := fetchGpsById(ctx, tt.deviceID)
			if err != tt.wantErr {
				//log.Fatalf(err)
				fmt.Printf("%v", err)
			}
			if got != tt.wantDeviceGps {
				t.Errorf("fetchGpsById() = %v, want %v", got, tt.wantDeviceGps)
			}
		})
	}
}

func cmdExec(t *testing.T, cmds []string) error {
	t.Helper()
	for _, cmd := range cmds {
		t.Logf("[INFO] %s", cmd)
		args := strings.Split(cmd, " ")

		if err := exec.Command(args[0], args[1:]...).Run(); err != nil {
			t.Logf("[WARN] %s %v", cmd, err)
		}
	}
	return nil

}
