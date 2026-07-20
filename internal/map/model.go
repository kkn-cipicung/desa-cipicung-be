package mapdata

type AddMapPayload struct {
	Elevation  string `db:"elevation" json:"elevation"`
	Coordinate string `db:"coordinate" json:"coordinate"`
	HamletOne  *int64 `db:"hamlet_one" json:"hamlet_one" binding:"required,gte=0"`
	HamletTwo  *int64 `db:"hamlet_two" json:"hamlet_two" binding:"required,gte=0"`
}

type EditMapPayload struct {
	ID uint `db:"id" json:"id" binding:"required"`
	AddMapPayload
}

type MapPayload struct {
	ID uint `db:"id" json:"id" binding:"required"`
}

type MapResponse struct {
	Elevation  string `db:"elevation" json:"elevation"`
	Coordinate string `db:"coordinate" json:"coordinate"`
	HamletOne  int64  `db:"hamlet_one" json:"hamlet_one"`
	HamletTwo  int64  `db:"hamlet_two" json:"hamlet_two"`
}

type MapOutput struct {
	Elevation  string `json:"elevation"`
	Coordinate string `json:"coordinate"`
	HamletOne  int64  `json:"hamlet_one"`
	HamletTwo  int64  `json:"hamlet_two"`
	Population int64  `json:"population"`
}
