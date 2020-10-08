package repo

import (
	"context"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string `column:f_name`
	LastName  string `column:l_name`
}

type Criteria struct {
	ID int
}

type UserRepo interface {
	Save(ctx context.Context, users []User) (err error)
	FindOne(ctx context.Context, ID int) (user User, err error)
	FindAll(ctx context.Context, criteria Criteria) (user []User, err error)
	Count(ctx context.Context) (count int)
	Delete(ctx context.Context, ID int) (err error)
	IsExist(ctx context.Context, ID int) (isExist bool)
}
