package repository

import (
	"link-back-app/models"

	"gorm.io/gorm"
)

type UsersRepository interface {
	GetAllUsersRepository() ([]models.Users, error)
	RegisterUsersRepository(users models.Users)
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &usersRepositoryImpl{
		db: db,
	}
}

type usersRepositoryImpl struct {
	db *gorm.DB
}

/**
 * ユーザー情報全件取得処理
 */
func (u *usersRepositoryImpl) GetAllUsersRepository() ([]models.Users, error) {
	users := []models.Users{}
	err := u.db.Debug().Find(&users).Error
	return users, err
}

func (u *usersRepositoryImpl) RegisterUsersRepository(users models.Users) {
	panic("unimplemented")
}
