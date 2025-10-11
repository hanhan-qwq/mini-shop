package service

import (
	"errors"
	"mini_shop/repository"
	"time"
)

type UserProfileResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	AvatarURL string    `json:"avatar_url"`
	CreatedAt time.Time `json:"created_at"`
}

type UserService struct {
	UserDB *repository.UserDAO
}

func NewUserService() *UserService {
	return &UserService{
		UserDB: repository.NewUserDAO(),
	}
}

func (s *UserService) GetProfileByUserID(userID uint) (*UserProfileResponse, error) {
	//1.查询用户
	user, err := s.UserDB.GetUserByUserID(userID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	//2.用新的response避开敏感字段
	profile := &UserProfileResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      user.Role,
		AvatarURL: user.Avatar,
		CreatedAt: user.CreatedAt,
	}

	return profile, nil
}

func (s *UserService) UpdateUserProfile(userID uint, username, email, phone string) error {
	//1.检查用户是否存在
	user, err := s.UserDB.GetUserByUserID(userID)
	if err != nil {
		return errors.New("用户不存在")
	}
	//2.更新数据
	user.Username = username
	user.Email = email
	user.Phone = phone

	err = s.UserDB.UpdateUserProfile(user)
	if err != nil {
		return errors.New("用户信息更新失败")
	}

	return nil
}
