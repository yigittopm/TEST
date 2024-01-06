package dtos

type GetAllUsersResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	UserType string `json:"userType"`
}
