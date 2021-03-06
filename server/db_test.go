package server

import (
	"context"
	"testing"

	"github.com/guregu/dynamo"
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
				Signal:     true,
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
				Signal:     true,
				RecordedAt: "2021-08-07T12:34:56+09:00",
			},
			wantErr: nil,
			cmds: []string{
				`aws --profile local --endpoint-url http://localhost:4566 dynamodb create-table --cli-input-json file://./testdata/schema/local_device_gps.json`,
				`aws --profile local --endpoint-url http://localhost:4566 dynamodb batch-write-item --request-items file://./testdata/db/fetchGpsByDeviceId_1.json`,
			},
		},

		{
			name:     "【正常系】1レコードfalseがあるとき",
			deviceID: "00002",
			wantDeviceGps: DeviceGPS{
				DeviceID:   "00002",
				DeviceType: "microcomputer",
				Signal:     false,
				RecordedAt: "2021-08-08T12:34:56+09:00",
			},
			wantErr: nil,
			cmds: []string{
				`aws --profile local --endpoint-url http://localhost:4566 dynamodb create-table --cli-input-json file://./testdata/schema/local_device_gps.json`,
				`aws --profile local --endpoint-url http://localhost:4566 dynamodb batch-write-item --request-items file://./testdata/db/fetchGpsByDeviceId_2.json`,
			},
		},
		{
			name:          "【異常系】当該データが一つもない",
			deviceID:      "10000",
			wantDeviceGps: DeviceGPS{},
			wantErr:       dynamo.ErrNotFound,
			cmds: []string{
				`aws --profile local --endpoint-url http://localhost:4566 dynamodb create-table --cli-input-json file://./testdata/schema/local_device_gps.json`,
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

			if tt.wantErr != nil {
				t.Logf("fetchGpsByDeviceId() got err = %v correctly", err)
				return
			}

			if got != tt.wantDeviceGps {
				t.Errorf("fetchGpsByDeviceId() = %v, want %v", got, tt.wantDeviceGps)
			}
		})
	}
}

func TestInsertGpsById(t *testing.T) {

	ctx := context.Background()

	tests := []struct {
		name            string
		insertDeviceGps DeviceGPS
		wantErr         error
		cmds            []string
	}{
		{
			name: "【正常系】1レコード差し込む/signal true",
			insertDeviceGps: DeviceGPS{
				DeviceID:   "00000",
				DeviceType: "microcomputer",
				Lat:        12.345,
				Lng:        135.12345,
				Signal:     true,
				RecordedAt: "2021-08-06T12:34:56+09:00",
			},
			wantErr: nil,
			cmds: []string{
				`aws --profile local --endpoint-url http://localhost:4566 dynamodb create-table --cli-input-json file://./testdata/schema/local_device_gps.json`,
			},
		},
		{
			name: "【正常系】1レコード差し込む/signal false",
			insertDeviceGps: DeviceGPS{
				DeviceID:   "00000",
				DeviceType: "microcomputer",
				Signal:     false,
				RecordedAt: "2021-08-06T12:34:56+09:00",
			},
			wantErr: nil,
			cmds: []string{
				`aws --profile local --endpoint-url http://localhost:4566 dynamodb create-table --cli-input-json file://./testdata/schema/local_device_gps.json`,
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
			// 当該レコードがないことの確認
			_, err := fetchGpsByDeviceId(ctx, tt.insertDeviceGps.DeviceID)
			if err != dynamo.ErrNotFound {
				t.Fatalf("挿入対象がすでにいるので不適")
			}
			if err = insertGpsByDevice(ctx, tt.insertDeviceGps); err != nil {
				t.Fatalf("%v", err)
			}
			got, err := fetchGpsByDeviceId(ctx, tt.insertDeviceGps.DeviceID)
			if err != nil {
				t.Fatalf("%v", err)
			}
			if got != tt.insertDeviceGps {
				t.Errorf("insertGpsByDeviceId() = %v, want %v", got, tt.insertDeviceGps)
			}
		})
	}
}
