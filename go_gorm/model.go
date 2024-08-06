package main

// User represents the user model
// @Description User account information
type User struct {
	// gorm.Model
	// @Description The ID of the user
	ID uint `json:"id" gorm:"primaryKey" binding:"-" swaggerignore:"true"`
	// @Description The name of the user
	Name string `json:"name"`
	// @Description The email of the user (must be unique)
	Email string `json:"email" gorm:"unique"`
	// @Description The password of the user (not included in JSON responses)
	Password string `json:"-"`
}
