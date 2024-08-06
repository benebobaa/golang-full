package dto

type RegisterRequest struct {
	Name            string `json:"name" valo:"notblank"`
	Email           string `json:"email" valo:"email"`
	Password        string `json:"password" valo:"notblank,sizeMin=6,sizeMax=20"`
	ConfirmPassword string `json:"confirm_password" valo:"notblank,sizeMin=6,sizeMax=20"`
}
