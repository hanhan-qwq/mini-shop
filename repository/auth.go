package repository

import (
	"gorm.io/gorm"
	"mini_shop/global"
	"mini_shop/model"
)

type AuthDAO struct {
	db *gorm.DB
}

func NewAuthDAO() *AuthDAO {
	return &AuthDAO{
		db: global.GetDB(),
	}
}

func (d *AuthDAO) CreateUser(user *model.User) error {
	return d.db.Create(user).Error
}
func (d *AuthDAO) CheckUserExist(username, phone, email string) (bool, error) {
	var count int64

	err := d.db.Model(&model.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, err
	}
	err = d.db.Model(&model.User{}).Where("phone = ?", phone).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, err
	}
	err = d.db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, err
	}

	return false, nil
}

func (d *AuthDAO) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := d.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
