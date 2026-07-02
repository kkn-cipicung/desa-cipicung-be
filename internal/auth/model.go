package auth

type RegisterPayload struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
