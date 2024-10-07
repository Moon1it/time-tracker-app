package dto

type CreateUserPaylaod struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type UserResponse struct {
	ID        string `json:"id" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type GetAllUsersResponse struct {
	Size  int            `json:"size" validate:"required"`
	Users []UserResponse `json:"users" validate:"required"`
}
