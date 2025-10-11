package repository

import (
	"gorm.io/gorm"
	"mini_shop/global"
	"mini_shop/model"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO() *UserDAO {
	return &UserDAO{
		db: global.GetDB(),
	}
}

func (d *UserDAO) GetUserByUserID(id uint) (*model.User, error) {
	var user model.User
	err := d.db.Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (d *UserDAO) UpdateUserProfile(user *model.User) error {
	return d.db.Model(&user).Updates(user).Error
}
