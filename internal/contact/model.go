package contact

type AddContactPayload struct {
	Name       string `db:"name" json:"name" binding:"required"`
	Province   string `db:"province" json:"province" binding:"required"`
	Regency    string `db:"regency" json:"regency" binding:"required"`
	District   string `db:"district" json:"district" binding:"required"`
	PostalCode string `db:"postal_code" json:"postal_code"`
	Address    string `db:"address" json:"address" binding:"required"`
	Phone      string `db:"phone" json:"phone"`
	Email      string `db:"email" json:"email"`
	Website    string `db:"website" json:"website"`
	Instagram  string `db:"ig_usn" json:"ig_usn"`
	TikTok     string `db:"tiktok_usn" json:"tiktok_usn"`
	YouTube    string `db:"yt_usn" json:"yt_usn"`
}

type EditContactPayload struct {
	ID uint `db:"id" json:"id" binding:"required"`
	AddContactPayload
}

type ContactPayload struct {
	ID uint `db:"id" json:"id" binding:"required"`
}

type ContactResponse struct {
	ID         uint   `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Province   string `db:"province" json:"province"`
	Regency    string `db:"regency" json:"regency"`
	District   string `db:"district" json:"district"`
	PostalCode string `db:"postal_code" json:"postal_code"`
	Address    string `db:"address" json:"address"`
	Phone      string `db:"phone" json:"phone"`
	Email      string `db:"email" json:"email"`
	Website    string `db:"website" json:"website"`
	Instagram  string `db:"ig_usn" json:"ig_usn"`
	TikTok     string `db:"tiktok_usn" json:"tiktok_usn"`
	YouTube    string `db:"yt_usn" json:"yt_usn"`
	IsActive   bool   `db:"is_active" json:"is_active"`
}

type ContactOutput struct {
	ID          uint                 `json:"id"`
	Office      ContactOffice        `json:"office"`
	Contact     ContactInfo          `json:"contact"`
	SocialMedia []ContactSocialMedia `json:"social_media"`
	ServiceHour []ContactServiceHour `json:"service_hour"`
	IsActive    bool                 `json:"is_active"`
}

type ContactOffice struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	District   string `json:"district"`
	Regency    string `json:"regency"`
	Province   string `json:"province"`
	PostalCode string `json:"postal_code"`
}

type ContactInfo struct {
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
}

type ContactSocialMedia struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

type ContactServiceHour struct {
	Day  string `json:"day"`
	Time string `json:"time"`
}
