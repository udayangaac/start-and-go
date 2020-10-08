package repo

import (
	"context"
	"github.com/jinzhu/gorm"
)

type userMySqlRepo struct {
	DB *gorm.DB
}

func (u *userMySqlRepo) Save(ctx context.Context, users []User) (err error) {
	err = u.DB.Table("user").Create(&users).Error
	return
}

func (u *userMySqlRepo) FindOne(ctx context.Context, ID int) (user User, err error) {
	err = u.DB.Table("user").Where("id = ?", ID).First(&user).Error
	return
}

func (u *userMySqlRepo) FindAll(ctx context.Context, criteria Criteria) (user []User, err error) {
	return
}

func (u *userMySqlRepo) Count(ctx context.Context) (count int) {
	return
}

func (u *userMySqlRepo) Delete(ctx context.Context, ID int) (err error) {
	return
}

func (u *userMySqlRepo) IsExist(ctx context.Context, ID int) (isExist bool) {
	return
}
