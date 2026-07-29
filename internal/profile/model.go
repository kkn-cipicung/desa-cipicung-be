package profile

import (
	"encoding/json"
	"time"

	"cipicung.id/be/pkg/types"
)

type AddProfilePayload struct {
	ID                    uint                      `db:"id" json:"id,omitempty"`
	Name                  string                    `db:"name" json:"name" binding:"required"`
	Province              string                    `db:"province" json:"province" binding:"required"`
	Regency               string                    `db:"regency" json:"regency" binding:"required"`
	District              string                    `db:"district" json:"district" binding:"required"`
	PostalCode            string                    `db:"postal_code" json:"postal_code"`
	Address               string                    `db:"address" json:"address" binding:"required"`
	Phone                 string                    `db:"phone" json:"phone"`
	Email                 string                    `db:"email" json:"email"`
	Website               string                    `db:"website" json:"website"`
	Latitude              float64                   `db:"latitude" json:"latitude"`
	Longitude             float64                   `db:"longitude" json:"longitude"`
	Vision                string                    `db:"vision" json:"vision"`
	Mission               []string                  `db:"mission" json:"mission"`
	History               string                    `db:"history" json:"history"`
	Description           string                    `db:"description" json:"description"`
	Region                string                    `db:"region" json:"region"`
	TotalFamily           int64                     `db:"total_family" json:"total_family"`
	HamletOne             int64                     `db:"hamlet_one" json:"hamlet_one"`
	HamletTwo             int64                     `db:"hamlet_two" json:"hamlet_two"`
	TotalRT               int64                     `db:"total_rt" json:"total_rt"`
	TotalRW               int64                     `db:"total_rw" json:"total_rw"`
	RTHamletOne           int64                     `db:"rt_hamlet_one" json:"rt_hamlet_one"`
	RTHamletTwo           int64                     `db:"rt_hamlet_two" json:"rt_hamlet_two"`
	RWHamletOne           int64                     `db:"rw_hamlet_one" json:"rw_hamlet_one"`
	RWHamletTwo           int64                     `db:"rw_hamlet_two" json:"rw_hamlet_two"`
	NorthBorder           string                    `db:"north_border" json:"north_border"`
	EastBorder            string                    `db:"east_border" json:"east_border"`
	SouthBorder           string                    `db:"south_border" json:"south_border"`
	WestBorder            string                    `db:"west_border" json:"west_border"`
	Area                  string                    `db:"area" json:"area"`
	Population            string                    `db:"population" json:"population"`
	TotalMale             int64                     `db:"total_male" json:"total_male"`
	TotalFemale           int64                     `db:"total_female" json:"total_female"`
	DemographicReligions  []DemographicSlice        `json:"demographic_religions"`
	DemographicReligionRT []DemographicBar          `json:"demographic_religion_rt"`
	DemographicEducation  []DemographicGenderBar    `json:"demographic_education"`
	DemographicOccupation []DemographicGenderBar    `json:"demographic_occupation"`
	DemographicAges       []DemographicSingleBar    `json:"demographic_ages"`
	Headman               *ProfileOfficialInput     `json:"headman"`
	Headmen               []ProfileOfficialInput    `json:"headmen"`
	Officials             []GovernmentOfficialInput `json:"officials"`
	ResourcePotential     *ResourcePotentialInput   `json:"resource_potential"`
}

type EditProfilePayload struct {
	AddProfilePayload
}

type ProfilePayload struct {
	ID uint `db:"id" json:"id"`
}

type ProfileResponse struct {
	ID                        uint                  `db:"id" json:"id"`
	Name                      string                `db:"name" json:"name"`
	Province                  string                `db:"province" json:"province"`
	Regency                   string                `db:"regency" json:"regency"`
	District                  string                `db:"district" json:"district"`
	PostalCode                string                `db:"postal_code" json:"postal_code"`
	Address                   string                `db:"address" json:"address"`
	Phone                     string                `db:"phone" json:"phone"`
	Email                     string                `db:"email" json:"email"`
	Website                   string                `db:"website" json:"website"`
	Latitude                  float64               `db:"latitude" json:"latitude"`
	Longitude                 float64               `db:"longitude" json:"longitude"`
	Vision                    string                `db:"vision" json:"vision"`
	Mission                   types.JSONStringArray `db:"mission" json:"mission"`
	History                   string                `db:"history" json:"history"`
	Description               string                `db:"description" json:"description"`
	Region                    string                `db:"region" json:"region"`
	HamletOne                 int64                 `db:"hamlet_one" json:"hamlet_one"`
	HamletTwo                 int64                 `db:"hamlet_two" json:"hamlet_two"`
	TotalFamily               int64                 `db:"total_family" json:"total_family"`
	TotalRT                   int64                 `db:"total_rt" json:"total_rt"`
	TotalRW                   int64                 `db:"total_rw" json:"total_rw"`
	RTHamletOne               int64                 `db:"rt_hamlet_one" json:"rt_hamlet_one"`
	RTHamletTwo               int64                 `db:"rt_hamlet_two" json:"rt_hamlet_two"`
	RWHamletOne               int64                 `db:"rw_hamlet_one" json:"rw_hamlet_one"`
	RWHamletTwo               int64                 `db:"rw_hamlet_two" json:"rw_hamlet_two"`
	NorthBorder               string                `db:"north_border" json:"north_border"`
	EastBorder                string                `db:"east_border" json:"east_border"`
	SouthBorder               string                `db:"south_border" json:"south_border"`
	WestBorder                string                `db:"west_border" json:"west_border"`
	Area                      string                `db:"area" json:"area"`
	Population                string                `db:"population" json:"population"`
	TotalMale                 int64                 `db:"total_male" json:"total_male"`
	TotalFemale               int64                 `db:"total_female" json:"total_female"`
	DemographicReligionsJSON  string                `db:"demographic_religions" json:"-"`
	DemographicReligionRTJSON string                `db:"demographic_religion_rt" json:"-"`
	DemographicEducationJSON  string                `db:"demographic_education" json:"-"`
	DemographicOccupationJSON string                `db:"demographic_occupation" json:"-"`
	DemographicAgesJSON       string                `db:"demographic_ages" json:"-"`
	IsActive                  bool                  `db:"is_active" json:"is_active"`
	CreatedAt                 time.Time             `db:"created_at" json:"created_at"`
	UpdatedAt                 time.Time             `db:"updated_at" json:"updated_at"`
	HeadmanID                 uint                  `db:"headman_id" json:"headman_id"`
	HeadmanName               string                `db:"headman_name" json:"headman_name"`
	HeadmanPosition           string                `db:"headman_position" json:"headman_position"`
	HeadmanPhone              string                `db:"headman_phone" json:"headman_phone"`
	HeadmanEmail              string                `db:"headman_email" json:"headman_email"`
	HeadmanDescription        string                `db:"headman_description" json:"headman_description"`
	HeadmanOrderNumber        int                   `db:"headman_order_number" json:"headman_order_number"`
	HeadmanIsActive           bool                  `db:"headman_is_active" json:"headman_is_active"`
	HeadmanStartDate          *time.Time            `db:"headman_start_date" json:"headman_start_date"`
	HeadmanFinishDate         *time.Time            `db:"headman_finish_date" json:"headman_finish_date"`
}

type ProfileOutput struct {
	ID                    uint                    `json:"id"`
	Name                  string                  `json:"name"`
	Province              string                  `json:"province"`
	Regency               string                  `json:"regency"`
	District              string                  `json:"district"`
	PostalCode            string                  `json:"postal_code"`
	Address               string                  `json:"address"`
	Phone                 string                  `json:"phone"`
	Email                 string                  `json:"email"`
	Website               string                  `json:"website"`
	Latitude              float64                 `json:"latitude"`
	Longitude             float64                 `json:"longitude"`
	Vision                string                  `json:"vision"`
	Mission               []string                `json:"mission"`
	History               string                  `json:"history"`
	Description           string                  `json:"description"`
	Region                string                  `json:"region"`
	HamletOne             int64                   `json:"hamlet_one"`
	HamletTwo             int64                   `json:"hamlet_two"`
	TotalFamily           int64                   `json:"total_family"`
	TotalRT               int64                   `json:"total_rt"`
	TotalRW               int64                   `json:"total_rw"`
	RTHamletOne           int64                   `json:"rt_hamlet_one"`
	RTHamletTwo           int64                   `json:"rt_hamlet_two"`
	RWHamletOne           int64                   `json:"rw_hamlet_one"`
	RWHamletTwo           int64                   `json:"rw_hamlet_two"`
	NorthBorder           string                  `json:"north_border"`
	EastBorder            string                  `json:"east_border"`
	SouthBorder           string                  `json:"south_border"`
	WestBorder            string                  `json:"west_border"`
	Area                  string                  `json:"area"`
	Population            string                  `json:"population"`
	TotalMale             int64                   `json:"total_male"`
	TotalFemale           int64                   `json:"total_female"`
	DemographicReligions  []DemographicSlice      `json:"demographic_religions"`
	DemographicReligionRT []DemographicBar        `json:"demographic_religion_rt"`
	DemographicEducation  []DemographicGenderBar  `json:"demographic_education"`
	DemographicOccupation []DemographicGenderBar  `json:"demographic_occupation"`
	DemographicAges       []DemographicSingleBar  `json:"demographic_ages"`
	Headman               *ProfileOfficialOutput  `json:"headman"`
	Headmen               []ProfileOfficialOutput `json:"headmen"`
	IsActive              bool                    `json:"is_active"`
	CreatedAt             string                  `json:"created_at"`
	UpdatedAt             string                  `json:"updated_at"`
}

type DemographicSlice struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
	Color string `json:"color,omitempty"`
}

type DemographicBar struct {
	Label  string `json:"label"`
	Value  int64  `json:"value"`
	Value2 int64  `json:"value2,omitempty"`
}

type DemographicGenderBar struct {
	Label  string `json:"label"`
	Male   int64  `json:"male"`
	Female int64  `json:"female"`
}

type DemographicSingleBar struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
}

func demographicJSON(value any) string {
	data, err := json.Marshal(value)
	if err != nil {
		return "[]"
	}
	return string(data)
}

type ProfileOfficialInput struct {
	Name        string  `db:"name" json:"name" binding:"required"`
	Position    string  `db:"position" json:"position"`
	Phone       string  `db:"phone" json:"phone"`
	Email       string  `db:"email" json:"email"`
	Description string  `db:"description" json:"description"`
	OrderNumber int     `db:"order_number" json:"order_number"`
	IsActive    bool    `db:"is_active" json:"is_active"`
	StartDate   string  `db:"start_date" json:"start_date" binding:"required"`
	FinishDate  *string `db:"finish_date" json:"finish_date"`
}

type ProfileOfficialOutput struct {
	ID          uint    `db:"id" json:"id"`
	Name        string  `db:"name" json:"name"`
	Position    string  `db:"position" json:"position"`
	Phone       string  `db:"phone" json:"phone"`
	Email       string  `db:"email" json:"email"`
	Description string  `db:"description" json:"description"`
	OrderNumber int     `db:"order_number" json:"order_number"`
	IsActive    bool    `db:"is_active" json:"is_active"`
	StartDate   string  `db:"start_date" json:"start_date"`
	FinishDate  *string `db:"finish_date" json:"finish_date"`
}

type ProfileRegionBoundaryResponse struct {
	Region      string `db:"region" json:"region"`
	HamletOne   int64  `db:"hamlet_one" json:"hamlet_one"`
	HamletTwo   int64  `db:"hamlet_two" json:"hamlet_two"`
	TotalFamily int64  `db:"total_family" json:"total_family"`
	TotalRT     int64  `db:"total_rt" json:"total_rt"`
	TotalRW     int64  `db:"total_rw" json:"total_rw"`
	RTHamletOne int64  `db:"rt_hamlet_one" json:"rt_hamlet_one"`
	RTHamletTwo int64  `db:"rt_hamlet_two" json:"rt_hamlet_two"`
	RWHamletOne int64  `db:"rw_hamlet_one" json:"rw_hamlet_one"`
	RWHamletTwo int64  `db:"rw_hamlet_two" json:"rw_hamlet_two"`
	NorthBorder string `db:"north_border" json:"north_border"`
	EastBorder  string `db:"east_border" json:"east_border"`
	SouthBorder string `db:"south_border" json:"south_border"`
	WestBorder  string `db:"west_border" json:"west_border"`
	Area        string `db:"area" json:"area"`
	Population  string `db:"population" json:"population"`
}

type ProfileVisionMissionResponse struct {
	Vision  string                `db:"vision" json:"vision"`
	Mission types.JSONStringArray `db:"mission" json:"mission"`
}

type ProfileVisionMissionOutput struct {
	Vision  string   `json:"vision"`
	Mission []string `json:"mission"`
}

type GovernmentStructureResponse struct {
	ID          uint   `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Position    string `db:"position" json:"position"`
	OrderNumber int    `db:"order_number" json:"order_number"`
	IsActive    bool   `db:"is_active" json:"is_active"`
}

type ResourcePotentialResponse struct {
	Title       string `db:"title" json:"title"`
	Detail      string `db:"detail" json:"detail"`
	Description string `db:"description" json:"description"`
}

type GovernmentOfficialInput struct {
	Name        string `json:"name" binding:"required"`
	Position    string `json:"position" binding:"required"`
	OrderNumber int    `json:"order_number"`
	IsActive    bool   `json:"is_active"`
}

type ResourcePotentialInput struct {
	Title       string `json:"title" binding:"required"`
	Detail      string `json:"detail" binding:"required"`
	Description string `json:"description"`
}

type AddOfficialPayload struct {
	VillageID   uint   `db:"village_id" json:"village_id"`
	Name        string `db:"name" json:"name" binding:"required"`
	Position    string `db:"position" json:"position" binding:"required"`
	OrderNumber int    `db:"order_number" json:"order_number"`
	IsActive    bool   `db:"is_active" json:"is_active"`
}

type EditOfficialPayload struct {
	ID uint `db:"id" json:"id" binding:"required"`
	AddOfficialPayload
}

type OfficialPayload struct {
	ID uint `db:"id" json:"id" binding:"required"`
}

type ListOfficialPayload struct {
	Limit     int  `form:"limit" json:"limit"`
	Index     int  `form:"index" json:"index"`
	VillageID uint `form:"village_id" json:"village_id"`
}

type OfficialResponse struct {
	ID          uint   `db:"id" json:"id"`
	VillageID   uint   `db:"village_id" json:"village_id"`
	Name        string `db:"name" json:"name"`
	Position    string `db:"position" json:"position"`
	OrderNumber int    `db:"order_number" json:"order_number"`
	IsActive    bool   `db:"is_active" json:"is_active"`
	CreatedAt   string `db:"created_at" json:"created_at"`
	UpdatedAt   string `db:"updated_at" json:"updated_at"`
}
