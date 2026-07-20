package mapdata

import "testing"

func TestValidateMapPayloadHamlets(t *testing.T) {
	zero := int64(0)
	positive := int64(2500)
	negative := int64(-1)

	tests := []struct {
		name    string
		payload AddMapPayload
		wantErr bool
	}{
		{
			name:    "valid hamlet populations",
			payload: AddMapPayload{Elevation: "120 mdpl", Coordinate: "-6.5,107.4", HamletOne: &positive, HamletTwo: &zero},
		},
		{
			name:    "missing hamlet one",
			payload: AddMapPayload{Elevation: "120 mdpl", Coordinate: "-6.5,107.4", HamletTwo: &positive},
			wantErr: true,
		},
		{
			name:    "negative hamlet population",
			payload: AddMapPayload{Elevation: "120 mdpl", Coordinate: "-6.5,107.4", HamletOne: &negative, HamletTwo: &positive},
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
