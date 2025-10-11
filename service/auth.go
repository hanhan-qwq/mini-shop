package service

import (
	"encoding/base64"
	"errors"
	"mini_shop/jwt"
	"mini_shop/model"
	"mini_shop/repository"
)

type AuthService struct {
	UserDB *repository.AuthDAO
}

func NewAuthService() *AuthService {
	return &AuthService{
		UserDB: repository.NewAuthDAO(),
	}
}

func (s *AuthService) UserRegister(username, password, phone, email string) error {
	//1.检查用户名/手机号/邮箱唯一性
	flag, err := s.UserDB.CheckUserExist(username, phone, email)
	if err != nil {
		return err
	}
	if flag {
		return errors.New("用户名/邮箱/手机号已存在")
	}
	//2.密码加密并注册
	passwordHash := s.encodePassword(password)
	err = s.creatUser(username, passwordHash, phone, email)
	if err != nil {
		return err
	}
	return nil
}
func (s *AuthService) encodePassword(password string) string {
	return base64.StdEncoding.EncodeToString([]byte(password))
}
func (s *AuthService) creatUser(username, passwordHash, phone, email string) error {
	user := model.User{
		Username: username,
		Password: passwordHash,
		Phone:    phone,
		Email:    email,
	}
	return s.UserDB.CreateUser(&user)
}

func (s *AuthService) Login(username, password string) (string, string, string, error) {
	//1.查询用户是否存在
	user, err := s.UserDB.GetUserByUsername(username)
	if err != nil {
		return "", "", "", errors.New("用户不存在")
	}
	//2.验证密码
	if !s.verifyPassword(password, user.Password) {
		return "", "", "", errors.New("密码错误")
	}
	//3.发放jwt令牌
	accessToken, refreshToken, err := jwt.GenerateTokens(user.ID, user.Username, user.Role)
	if err != nil {
		return "", "", "", errors.New("生成 Token 失败")
	}
	return accessToken, refreshToken, user.Role, nil
}
func (s *AuthService) verifyPassword(password, truePassword string) bool {
	encodePassword := s.encodePassword(password)
	return encodePassword == truePassword
}
