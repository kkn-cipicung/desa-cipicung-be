package mapdata

import "testing"

func TestValidateMapPayload(t *testing.T) {
	tests := []struct {
		name    string
		payload AddMapPayload
		wantErr bool
	}{
		{
			name:    "valid map payload",
			payload: AddMapPayload{Elevation: "120 mdpl", Coordinate: "-6.5,107.4"},
		},
		{
			name:    "missing elevation",
			payload: AddMapPayload{Coordinate: "-6.5,107.4"},
			wantErr: true,
		},
		{
			name:    "missing coordinate",
			payload: AddMapPayload{Elevation: "120 mdpl"},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateMapPayload(test.payload)
			if (err != nil) != test.wantErr {
				t.Fatalf("validateMapPayload() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func TestMapMapOutputCalculatesPopulation(t *testing.T) {
	output := mapMapOutput(MapResponse{HamletOne: 2500, HamletTwo: 1750})
	if output.Population != 4250 {
		t.Fatalf("population = %d, want 4250", output.Population)
	}
}
