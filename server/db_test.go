package server

import (
	"context"
	"os/exec"
	"strings"
	"testing"
)

func TestFetchGpsById(t *testing.T) {

	ctx := context.Background()

	tests := []struct {
		name          string
		deviceID      string
		wantDeviceGps DeviceGPS
		wantErr       error
		cmds          []string
	}{
		{
			name:     "【正常系】1レコードのみ取得する",
			deviceID: "00000",
			wantDeviceGps: DeviceGPS{
				DeviceID:   "00000",
				DeviceType: "microcomputer",
				Lat:        12.345,
				Lng:        135.12345,
				RecordedAt: "2021-08-06T12:34:56+09:00",
			},
			wantErr: nil,
			cmds: []string{
				`aws --profile local --endpoint-url http://localhost:4566 dynamodb create-table --cli-input-json file://./testdata/schema/local_device_gps.json`,
				`aws --profile local --endpoint-url http://localhost:4566 dynamodb batch-write-item --request-items file://./testdata/db/fetchGpsByDeviceId_0.json`,
			},
		},
		{
			name:     "【正常系】2レコードあるときに最新レコードを取得する",
			deviceID: "00001",
			wantDeviceGps: DeviceGPS{
				DeviceID:   "00001",
				DeviceType: "microcomputer",
				Lat:        13.345,
				Lng:        136.12345,
				RecordedAt: "2021-08-07T12:34:56+09:00",
			},
			wantErr: nil,
			cmds: []string{
				`aws --profile local --endpoint-url http://localhost:4566 dynamodb create-table --cli-input-json file://./testdata/schema/local_device_gps.json`,
				`aws --profile local --endpoint-url http://localhost:4566 dynamodb batch-write-item --request-items file://./testdata/db/fetchGpsByDeviceId_1.json`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				cmds := []string{`aws --profile local --endpoint-url http://localhost:4566 dynamodb delete-table --table local_device_gps`}
				cmdExec(t, cmds)
			})
			cmdExec(t, tt.cmds)
			got, err := fetchGpsByDeviceId(ctx, tt.deviceID)
			if err != tt.wantErr {
				t.Fatalf("%v", err)
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
