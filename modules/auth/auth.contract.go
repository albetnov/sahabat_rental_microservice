package auth

type Auth struct {
	AppKey string `json:"app_key" binding:"required" validate:"required"`
}
