package response

type UserResponse struct {
	ID    string `json:"id" example:"82bdd399-321b-41d8-8b40-9a0116db9e92"`
	Email string `json:"email" example:"test@test.com"`
	Name  string `json:"name" example:"John Doe"`
	Age   int8   `json:"age" example:"30"`
}
