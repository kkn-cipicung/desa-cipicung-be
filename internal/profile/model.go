package profile

import (
	"time"

	"github.com/lib/pq"
)

type AddProfilePayload struct {
	ID                uint                      `db:"id" json:"id,omitempty"`
	Name              string                    `db:"name" json:"name" binding:"required"`
	Province          string                    `db:"province" json:"province" binding:"required"`
	Regency           string                    `db:"regency" json:"regency" binding:"required"`
	District          string                    `db:"district" json:"district" binding:"required"`
	PostalCode        string                    `db:"postal_code" json:"postal_code"`
	Address           string                    `db:"address" json:"address" binding:"required"`
	Phone             string                    `db:"phone" json:"phone"`
	Email             string                    `db:"email" json:"email"`
	Website           string                    `db:"website" json:"website"`
	Latitude          float64                   `db:"latitude" json:"latitude"`
	Longitude         float64                   `db:"longitude" json:"longitude"`
	Vision            string                    `db:"vision" json:"vision"`
	Mission           []string                  `db:"mission" json:"mission"`
	History           string                    `db:"history" json:"history"`
	Description       string                    `db:"description" json:"description"`
	Region            string                    `db:"region" json:"region"`
	HamletOne         int64                     `db:"hamlet_one" json:"hamlet_one"`
	HamletTwo         int64                     `db:"hamlet_two" json:"hamlet_two"`
	NorthBorder       string                    `db:"north_border" json:"north_border"`
	EastBorder        string                    `db:"east_border" json:"east_border"`
	SouthBorder       string                    `db:"south_border" json:"south_border"`
	WestBorder        string                    `db:"west_border" json:"west_border"`
	Area              string                    `db:"area" json:"area"`
	Population        string                    `db:"population" json:"population"`
	Headman           *ProfileOfficialInput     `json:"headman"`
	Headmen           []ProfileOfficialInput    `json:"headmen"`
	Officials         []GovernmentOfficialInput `json:"officials"`
	ResourcePotential *ResourcePotentialInput   `json:"resource_potential"`
}

type EditProfilePayload struct {
	AddProfilePayload
}

type ProfilePayload struct {
	ID uint `db:"id" json:"id" binding:"required"`
}

type ProfileResponse struct {
	ID                 uint           `db:"id" json:"id"`
	Name               string         `db:"name" json:"name"`
	Province           string         `db:"province" json:"province"`
	Regency            string         `db:"regency" json:"regency"`
	District           string         `db:"district" json:"district"`
	PostalCode         string         `db:"postal_code" json:"postal_code"`
	Address            string         `db:"address" json:"address"`
	Phone              string         `db:"phone" json:"phone"`
	Email              string         `db:"email" json:"email"`
	Website            string         `db:"website" json:"website"`
	Latitude           float64        `db:"latitude" json:"latitude"`
	Longitude          float64        `db:"longitude" json:"longitude"`
	Vision             string         `db:"vision" json:"vision"`
	Mission            pq.StringArray `db:"mission" json:"mission"`
	History            string         `db:"history" json:"history"`
	Description        string         `db:"description" json:"description"`
	Region             string         `db:"region" json:"region"`
	HamletOne          int64          `db:"hamlet_one" json:"hamlet_one"`
	HamletTwo          int64          `db:"hamlet_two" json:"hamlet_two"`
	NorthBorder        string         `db:"north_border" json:"north_border"`
	EastBorder         string         `db:"east_border" json:"east_border"`
	SouthBorder        string         `db:"south_border" json:"south_border"`
	WestBorder         string         `db:"west_border" json:"west_border"`
	Area               string         `db:"area" json:"area"`
	Population         string         `db:"population" json:"population"`
	CreatedAt          time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt          time.Time      `db:"updated_at" json:"updated_at"`
	HeadmanID          uint           `db:"headman_id" json:"headman_id"`
	HeadmanName        string         `db:"headman_name" json:"headman_name"`
	HeadmanPosition    string         `db:"headman_position" json:"headman_position"`
	HeadmanPhone       string         `db:"headman_phone" json:"headman_phone"`
	HeadmanEmail       string         `db:"headman_email" json:"headman_email"`
	HeadmanDescription string         `db:"headman_description" json:"headman_description"`
	HeadmanOrderNumber int            `db:"headman_order_number" json:"headman_order_number"`
	HeadmanIsActive    bool           `db:"headman_is_active" json:"headman_is_active"`
	HeadmanStartDate   *time.Time     `db:"headman_start_date" json:"headman_start_date"`
	HeadmanFinishDate  *time.Time     `db:"headman_finish_date" json:"headman_finish_date"`
}

type ProfileOutput struct {
	ID          uint                    `json:"id"`
	Name        string                  `json:"name"`
	Province    string                  `json:"province"`
	Regency     string                  `json:"regency"`
	District    string                  `json:"district"`
	PostalCode  string                  `json:"postal_code"`
	Address     string                  `json:"address"`
	Phone       string                  `json:"phone"`
	Email       string                  `json:"email"`
	Website     string                  `json:"website"`
	Latitude    float64                 `json:"latitude"`
	Longitude   float64                 `json:"longitude"`
	Vision      string                  `json:"vision"`
	Mission     []string                `json:"mission"`
	History     string                  `json:"history"`
	Description string                  `json:"description"`
	Region      string                  `json:"region"`
	HamletOne   int64                   `json:"hamlet_one"`
	HamletTwo   int64                   `json:"hamlet_two"`
	NorthBorder string                  `json:"north_border"`
	EastBorder  string                  `json:"east_border"`
	SouthBorder string                  `json:"south_border"`
	WestBorder  string                  `json:"west_border"`
	Area        string                  `json:"area"`
	Population  string                  `json:"population"`
	Headman     *ProfileOfficialOutput  `json:"headman"`
	Headmen     []ProfileOfficialOutput `json:"headmen"`
	CreatedAt   string                  `json:"created_at"`
	UpdatedAt   string                  `json:"updated_at"`
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
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Position    string  `json:"position"`
	Phone       string  `json:"phone"`
	Email       string  `json:"email"`
	Description string  `json:"description"`
	OrderNumber int     `json:"order_number"`
	IsActive    bool    `json:"is_active"`
	StartDate   string  `json:"start_date"`
	FinishDate  *string `json:"finish_date"`
}

type ProfileRegionBoundaryResponse struct {
	Region      string `db:"region" json:"region"`
	HamletOne   int64  `db:"hamlet_one" json:"hamlet_one"`
	HamletTwo   int64  `db:"hamlet_two" json:"hamlet_two"`
	NorthBorder string `db:"north_border" json:"north_border"`
	EastBorder  string `db:"east_border" json:"east_border"`
	SouthBorder string `db:"south_border" json:"south_border"`
	WestBorder  string `db:"west_border" json:"west_border"`
	Area        string `db:"area" json:"area"`
	Population  string `db:"population" json:"population"`
}

type ProfileVisionMissionResponse struct {
	Vision  string         `db:"vision" json:"vision"`
	Mission pq.StringArray `db:"mission" json:"mission"`
}

type ProfileVisionMissionOutput struct {
	Vision  string   `json:"vision"`
	Mission []string `json:"mission"`
}

type GovernmentStructureResponse struct {
	ID          uint    `db:"id" json:"id"`
	Name        string  `db:"name" json:"name"`
	Position    string  `db:"position" json:"position"`
	Phone       string  `db:"phone" json:"phone"`
	Email       string  `db:"email" json:"email"`
	Description string  `db:"description" json:"description"`
	OrderNumber int     `db:"order_number" json:"order_number"`
	IsActive    bool    `db:"is_active" json:"is_active"`
	StartDate   *string `db:"start_date" json:"start_date"`
	FinishDate  *string `db:"finish_date" json:"finish_date"`
}

type ResourcePotentialResponse struct {
	Title       string `db:"title" json:"title"`
	Detail      string `db:"detail" json:"detail"`
	Description string `db:"description" json:"description"`
}

type GovernmentOfficialInput struct {
	Name        string  `json:"name" binding:"required"`
	Position    string  `json:"position" binding:"required"`
	Phone       string  `json:"phone"`
	Email       string  `json:"email"`
	Description string  `json:"description"`
	OrderNumber int     `json:"order_number"`
	IsActive    bool    `json:"is_active"`
	StartDate   *string `json:"start_date"`
	FinishDate  *string `json:"finish_date"`
}

type ResourcePotentialInput struct {
	Title       string `json:"title" binding:"required"`
	Detail      string `json:"detail" binding:"required"`
	Description string `json:"description"`
}
