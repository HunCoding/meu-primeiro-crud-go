package request

// UserLogin represents the data required for user login.
// @Summary User Login Data
// @Description Structure containing the necessary fields for user login.
type UserLogin struct {
	// User's email (required and must be a valid email address).
	Email string `json:"email" binding:"required,email" example:"test@test.com"`

	// User's password (required, minimum of 6 characters, and must contain at least one of the characters: !@#$%*).
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*" example:"password#@#@!2121"`
}
