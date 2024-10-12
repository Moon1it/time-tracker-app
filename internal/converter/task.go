package converter

import (
	"time-tracker-app/internal/domain"
	"time-tracker-app/internal/transport/http/dto"
)

func ToTaskInfoFromHandler(info *dto.TaskInfo) *domain.TaskInfo {
	return &domain.TaskInfo{
		UserID: info.UserID,
		Name:   info.Name,
	}
}

// func ToGetTasksByUserIDResponseFromService(users []domain.User) *dto.GetAllUsersResponse {
// 	usersResponse := make([]dto.UserResponse, 0, len(users))

// 	for _, user := range users {
// 		usersResponse = append(usersResponse, dto.UserResponse{
// 			ID:        user.ID,
// 			FirstName: user.Info.FirstName,
// 			LastName:  user.Info.LastName,
// 		})
// 	}

// 	return &dto.GetAllUsersResponse{
// 		Size:  len(users),
// 		Users: usersResponse,
// 	}
// }
