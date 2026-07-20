package profile

import "testing"

func TestValidateHeadmanPeriods(t *testing.T) {
	finish2023 := "2023-12-31"
	finish2024 := "2024-06-30"

	tests := []struct {
		name    string
		headmen []ProfileOfficialInput
		wantErr bool
	}{
		{
			name: "different years are valid",
			headmen: []ProfileOfficialInput{
				{Name: "A", StartDate: "2018-01-01", FinishDate: &finish2023},
				{Name: "B", StartDate: "2024-01-01", FinishDate: nil},
			},
		},
		{
			name: "same boundary year is rejected",
			headmen: []ProfileOfficialInput{
				{Name: "A", StartDate: "2018-01-01", FinishDate: &finish2024},
				{Name: "B", StartDate: "2024-07-01", FinishDate: nil},
			},
			wantErr: true,
		},
		{
			name: "nested period is rejected",
			headmen: []ProfileOfficialInput{
				{Name: "A", StartDate: "2018-01-01", FinishDate: &finish2024},
				{Name: "B", StartDate: "2020-01-01", FinishDate: &finish2023},
			},
			wantErr: true,
		},
		{
			name: "two open periods are rejected",
			headmen: []ProfileOfficialInput{
				{Name: "A", StartDate: "2018-01-01", FinishDate: nil},
				{Name: "B", StartDate: "2024-01-01", FinishDate: nil},
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateHeadmanPeriods(test.headmen)
			if (err != nil) != test.wantErr {
				t.Fatalf("validateHeadmanPeriods() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}
