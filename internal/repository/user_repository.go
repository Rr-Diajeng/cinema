package repository

import (
	"cinema/internal/model"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Save(user model.Users) (err error)
		FindUserByEmailOrUsername(email string) (user model.Users, err error)
		FindOneUser(id uint) (user model.Users, err error)
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

func (ur userRepository) Save(user model.Users) (err error) {

	if err = ur.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (ur userRepository) FindUserByEmailOrUsername(data string) (user model.Users, err error) {

	err = ur.db.Where("email = ? OR username = ?", data, data).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur userRepository) FindOneUser(id uint) (model.Users, error) {
    var user model.Users
	
    err := ur.db.Preload("Role").Where("id = ?", id).First(&user).Error
    if err != nil {
        return user, err
    }

    return user, nil
}
