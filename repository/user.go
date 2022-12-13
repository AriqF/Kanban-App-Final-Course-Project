package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

// *implemented
func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	query := r.db.Model(&entity.User{}).First(&user, id)
	if query.Error != nil {
		return entity.User{}, query.Error
	}
	return user, nil
}

// *implemented
func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	query := r.db.Model(&entity.User{}).First(&user, "email = ?", email)
	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return entity.User{}, nil
		}
		return entity.User{}, query.Error
	}
	return user, nil
}

// *implemented
func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	res := r.db.Create(&user)
	if res.Error != nil {
		return entity.User{}, res.Error
	}
	fmt.Println("res err:", res.Error)
	return user, nil
}

// *implemented
func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	res := r.db.Model(&entity.User{}).Where("id = ?", user.ID).Updates(user)
	if res.Error != nil {
		return entity.User{}, res.Error
	}
	return user, nil
}

// *implemented
func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	res := r.db.Delete(&entity.User{}, id)
	return res.Error
}
