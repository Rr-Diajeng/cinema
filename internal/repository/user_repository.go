package repository

import (
	"cinema/internal/model"
	"context"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Save(c context.Context, user model.Users) (err error)
		FindUserByEmailOrUsername(c context.Context, email string) (user model.Users, err error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

// Jadi disini, ketika manggil UserRepository akan
// menggunakan struktur dari userRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		db: db,
	}
}

func (ur userRepository) Save(c context.Context, user model.Users) (err error) {

	if err = ur.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (ur userRepository) FindUserByEmailOrUsername(c context.Context, data string) (user model.Users, err error) {

    err = ur.db.Where("email = ? OR username = ?", data, data).First(&user).Error
    if err != nil {
        return user, err
    }
    return user, nil
}
