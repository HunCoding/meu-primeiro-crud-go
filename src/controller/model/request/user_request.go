package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*"`
	Name     string `json:"name" binding:"required,min=4,max=100"`
	Age      int8   `json:"age" binding:"required,min=1,max=140"`
}
