package repository

import (
	usersmodel "link-back-app/models/users_model"

	"gorm.io/gorm"
)

type UsersRepository interface {
	GetMaxEmployeeIdUsersRepository() (string, error)
	GetExistUsersRepository(employeeId string) (int, error)
	GetOneByKeyUsersRepository(employeeId string) (usersmodel.Users, error)
	GetAllUsersRepository() ([]usersmodel.Users, error)
	RegisterUsersRepository(tx *gorm.DB, users usersmodel.Users) error
	UpdateUsersRepository(tx *gorm.DB, users usersmodel.Users) error
	LogicalDeleteUsersRepository(tx *gorm.DB, users usersmodel.Users) error
	DeleteUsersRepository(tx *gorm.DB, employeeId string) error
}

func NewUsersRepository(database *gorm.DB) UsersRepository {
	return &usersRepositoryImpl{
		database: database,
	}
}

type usersRepositoryImpl struct {
	database *gorm.DB
}

/**
 * 最新の社員IDを取得する処理
 */
func (u *usersRepositoryImpl) GetMaxEmployeeIdUsersRepository() (string, error) {
	var users usersmodel.Users
	var maxEmployeeId string
	err := u.database.Model(&users).Select("COALESCE(MAX(employee_id), '') AS max_employee_id").Scan(&maxEmployeeId).Error
	return maxEmployeeId, err
}

/**
 * 社員IDに紐づくユーザー情報が存在するかを確認する処理
 **/
func (u *usersRepositoryImpl) GetExistUsersRepository(employeeId string) (int, error) {
	var users usersmodel.Users
	var count int
	err := u.database.Model(&users).Where("employee_id = ?", &employeeId).Select("COUNT(*) AS count").Scan(&count).Error
	return count, err
}

/**
 * ユーザー情報1件取得処理(社員ID指定)
 **/
func (u *usersRepositoryImpl) GetOneByKeyUsersRepository(employeeId string) (usersmodel.Users, error) {
	users := usersmodel.Users{}
	err := u.database.Debug().Where("employee_id = ?", employeeId).Take(&users).Error
	return users, err
}

/**
 * ユーザー情報全件取得処理
 */
func (u *usersRepositoryImpl) GetAllUsersRepository() ([]usersmodel.Users, error) {
	users := []usersmodel.Users{}
	err := u.database.Debug().Find(&users).Error
	return users, err
}

/**
 * ユーザー情報登録処理
 */
func (u *usersRepositoryImpl) RegisterUsersRepository(tx *gorm.DB, users usersmodel.Users) error {
	return tx.Debug().Create(&users).Error
}

/**
 * ユーザー情報更新処理
 **/
func (u *usersRepositoryImpl) UpdateUsersRepository(tx *gorm.DB, users usersmodel.Users) error {
	return tx.Model(users).Where("employee_id = ?", &users.EmployeeId).Updates(&users).Error
}

/**
 * ユーザー情報論理削除処理
 **/
func (u *usersRepositoryImpl) LogicalDeleteUsersRepository(tx *gorm.DB, users usersmodel.Users) error {
	return tx.Model(users).Where("employee_id = ?", users.EmployeeId).
		Updates(usersmodel.Users{
			DeleteFlg:       users.DeleteFlg,
			UpdatedAuthor:   users.UpdatedAuthor,
			UpdatedDatetime: users.UpdatedDatetime,
		}).Error
}

/**
 * ユーザー情報削除処理
 **/
func (u *usersRepositoryImpl) DeleteUsersRepository(tx *gorm.DB, employeeId string) error {
	var users usersmodel.Users
	return tx.Where("employee_id = ?", employeeId).Delete(&users).Error
}
