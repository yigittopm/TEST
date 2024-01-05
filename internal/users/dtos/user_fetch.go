package dtos

type GetAllUsersResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserType string `json:"userType"`
}
