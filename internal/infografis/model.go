package infografis

type InfografisOutput struct {
	Summary    InfografisSummary     `json:"summary"`
	Hamlets    []InfografisSlice     `json:"hamlets"`
	Religions  []InfografisSlice     `json:"religions"`
	ReligionRT []InfografisBar       `json:"religion_rt"`
	Education  []InfografisGenderBar `json:"education"`
	Occupation []InfografisGenderBar `json:"occupation"`
	Ages       []InfografisSingleBar `json:"ages"`
	Source     string                `json:"source"`
}

type InfografisSummary struct {
	Population int64 `json:"population"`
	Family     int64 `json:"family"`
	Male       int64 `json:"male"`
	Female     int64 `json:"female"`
}

type InfografisSlice struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
	Color string `json:"color,omitempty"`
}

type InfografisBar struct {
	Label  string `json:"label"`
	Value  int64  `json:"value"`
	Value2 int64  `json:"value2,omitempty"`
}

type InfografisGenderBar struct {
	Label  string `json:"label"`
	Male   int64  `json:"male"`
	Female int64  `json:"female"`
}

type InfografisSingleBar struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
}

type VillageStats struct {
	Population     int64  `db:"population"`
	Family         int64  `db:"family"`
	Male           int64  `db:"male"`
	Female         int64  `db:"female"`
	HamletOne      int64  `db:"hamlet_one"`
	HamletTwo      int64  `db:"hamlet_two"`
	ReligionsJSON  string `db:"religions"`
	ReligionRTJSON string `db:"religion_rt"`
	EducationJSON  string `db:"education"`
	OccupationJSON string `db:"occupation"`
	AgesJSON       string `db:"ages"`
}
