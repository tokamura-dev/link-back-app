package repository

import (
	"fmt"
	"link-back-app/models"

	"gorm.io/gorm"
)

type UsersRepository interface {
	GetMaxEmployeeIdUsers() (string, error)
	GetAllUsersRepository() ([]models.Users, error)
	RegisterUsersRepository(users models.Users) error
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
 * 最新の社員IDを取得する処理
 */
func (u *usersRepositoryImpl) GetMaxEmployeeIdUsers() (string, error) {
	var users models.Users
	var maxEmployeeId string
	err := u.db.Model(&users).Select("COALESCE(MAX(employee_id), '') AS max_employee_id").Scan(&maxEmployeeId).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return maxEmployeeId, err
}

/**
 * ユーザー情報全件取得処理
 */
func (u *usersRepositoryImpl) GetAllUsersRepository() ([]models.Users, error) {
	users := []models.Users{}
	err := u.db.Debug().Find(&users).Error
	return users, err
}

/**
 * ユーザー情報登録処理
 */
func (u *usersRepositoryImpl) RegisterUsersRepository(users models.Users) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if insert_err := tx.Debug().Create(&users).Error; insert_err != nil {
			return insert_err
		}
		return nil
	})
	return err
}
