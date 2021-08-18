package omega2plus

import (
	"context"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/google/go-cmp/cmp"
	"github.com/montblanc18/gps-viewer/client/gen/models"
)

func TestUplinkConv(t *testing.T) {
	ctx := context.Background()
	deviceType := "omega2plus"
	tokyoZone := time.FixedZone("Tokyo Time", 3600*9)
	recordedAt := strfmt.DateTime(time.Date(2021, time.August, 6, 12, 34, 56, 0, tokyoZone))
	tests := []struct {
		name string
		in   Uplink
		want models.RegisterDeviceGPS
	}{
		{
			name: "正常系",
			in: Uplink{
				Payloads{
					Age:        0,
					Latitude:   "43.8616",
					Longitude:  "-79.3854",
					Elevation:  "184.0",
					Course:     "",
					Speed:      "N",
					RecordedAt: "2021-08-06T12:34:56+09:00",
				},
			},
			want: models.RegisterDeviceGPS{
				DeviceType: &deviceType,
				Lat:        swag.Float64(43.8616),
				Lng:        swag.Float64(-79.3854),
				RecordedAt: &recordedAt,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.in.conv(ctx)
			if err != nil {
				t.Errorf("err=%v", err)
			}
			if diff := cmp.Diff(*got, tt.want); diff != "" {
				//t.Errorf("case %v want=%v, got=%v", tt.name, tt.want, got)
				t.Errorf("case %v diff=%v", tt.name, diff)
			}
		})
	}
}
