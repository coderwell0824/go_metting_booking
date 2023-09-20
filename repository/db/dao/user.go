package dao

import (
	"context"
	"gorm.io/gorm"
	"meetingBooking/repository/db/model"
)

type UserDao struct {
	*gorm.DB
}

// NewUserDao 创建一个用户的Dto
func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

// FindUserByUsername 根据用户名找到用户
func (dao *UserDao) FindUserByUsername(userName string) (user *model.User, err error) {
	
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).First(&user).Error
	
	return
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Create(user).Error
	return
}