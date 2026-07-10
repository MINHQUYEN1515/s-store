package repository

import (
	"s-store/internal/model/entity"
	"s-store/internal/model/request"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreaterUser(user request.RegisterRequest) (interface{}, error)
	GetById(id string) (*entity.UserEntity, error)
	UpdateUser(user *entity.UserEntity) (*entity.UserEntity, error)
	DeleteUser(id string) error
	FindByEmail(email string) (*entity.UserEntity, error)
}

type userRepo struct {
	db *gorm.DB
}

// CreaterUser implements [UserRepo].
func (u *userRepo) CreaterUser(user request.RegisterRequest) (interface{}, error) {
	panic("unimplemented")
}

// DeleteUser implements [UserRepo].
func (u *userRepo) DeleteUser(id string) error {
	panic("unimplemented")
}

// FindById implements [UserRepo].
func (u *userRepo) GetById(id string) (*entity.UserEntity, error) {
	panic("unimplemented")
}

// UpdateUser implements [UserRepo].
func (u *userRepo) UpdateUser(user *entity.UserEntity) (*entity.UserEntity, error) {
	panic("unimplemented")
}

// Find by email
func (u *userRepo) FindByEmail(email string) (*entity.UserEntity, error) {
	panic("unimplemented")
}

func UserRepoNew(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
