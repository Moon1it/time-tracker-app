package converter

import (
	"time-tracker-app/internal/domain"
	"time-tracker-app/internal/transport/http/dto"
)

func ToUserInfoFromHandler(info *dto.UserInfo) *domain.UserInfo {
	return &domain.UserInfo{
		FirstName: info.FirstName,
		LastName:  info.LastName,
	}
}

func ToGetAllUsersResponseFromService(users []domain.User) *dto.GetAllUsersResponse {
	usersResponse := make([]dto.UserResponse, 0, len(users))

	for _, user := range users {
		usersResponse = append(usersResponse, dto.UserResponse{
			ID: user.ID,
			Info: dto.UserInfo{
				FirstName: user.Info.FirstName,
				LastName:  user.Info.LastName,
			},
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return &dto.GetAllUsersResponse{
		Size:  len(users),
		Users: usersResponse,
	}
}
