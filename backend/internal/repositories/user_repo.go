package repository

import (
	"s-store/internal/model/entity"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreaterUser(user entity.UserEntity) (*entity.UserEntity, error)
	GetById(id string) (*entity.UserEntity, error)
	UpdateUser(user *entity.UserEntity) (*entity.UserEntity, error)
	DeleteUser(id string) error
	FindByEmail(email string) (*entity.UserEntity, error)
}

type userRepo struct {
	db *gorm.DB
}

// CreaterUser implements [UserRepo].
func (u *userRepo) CreaterUser(user entity.UserEntity) (*entity.UserEntity, error) {
	var userEntity entity.UserEntity
	err := u.db.Create(&userEntity).Error
	return &userEntity, err
}

// DeleteUser implements [UserRepo].
func (u *userRepo) DeleteUser(id string) error {
	return u.db.Delete(&entity.UserEntity{}, id).Error
}

// FindById implements [UserRepo].
func (u *userRepo) GetById(id string) (*entity.UserEntity, error) {
	var user entity.UserEntity
	err := u.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

// UpdateUser implements [UserRepo].
func (u *userRepo) UpdateUser(user *entity.UserEntity) (*entity.UserEntity, error) {
	var updatedUser entity.UserEntity
	err := u.db.Model(&updatedUser).Where("id = ?", user.ID).Updates(user).Error
	return &updatedUser, err
}

// Find by email
func (u *userRepo) FindByEmail(email string) (*entity.UserEntity, error) {
	var user entity.UserEntity
	err := u.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func UserRepoNew(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
