package dto

type CreateUserPaylaod struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type GetAllUsersResponse struct {
	Size  int            `json:"size" validate:"required"`
	Users []UserResponse `json:"users" validate:"required"`
}
