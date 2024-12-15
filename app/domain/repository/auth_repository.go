package repository

import (
	authmodel "link-back-app/models/auth_model"

	"gorm.io/gorm"
)

type AuthRepository interface {
	GetMaxEmployeeIdLoginRepository() (string, error)
	GetOneByKeyLoginRepository(employeeId string) (authmodel.Login, error)
	RegisterLoginRepository(tx *gorm.DB, login authmodel.Login) error
	LogicalDeleteLoginRepository(tx *gorm.DB, login authmodel.Login) error
}

func NewAuthRepository(database *gorm.DB) AuthRepository {
	return &authRepositoryImpl{
		database: database,
	}
}

type authRepositoryImpl struct {
	database *gorm.DB
}

/**
 * 最新の社員IDを取得する処理
 */
func (a *authRepositoryImpl) GetMaxEmployeeIdLoginRepository() (string, error) {
	var login authmodel.Login
	var maxEmployeeId string
	err := a.database.Model(&login).Select("COALESCE(MAX(employee_id), '') AS max_employee_id").Scan(&maxEmployeeId).Error
	return maxEmployeeId, err
}

/**
 * ログイン情報1件取得処理(社員ID指定)
 **/
func (a *authRepositoryImpl) GetOneByKeyLoginRepository(employeeId string) (authmodel.Login, error) {
	var login authmodel.Login
	err := a.database.Debug().Where("employee_id = ?", employeeId).Take(&login).Error
	return login, err
}

/**
 * ログイン情報登録処理
 **/
func (a *authRepositoryImpl) RegisterLoginRepository(tx *gorm.DB, login authmodel.Login) error {
	return tx.Debug().Create(&login).Error
}

/**
 * ログイン情報論理削除処理
 **/
func (a *authRepositoryImpl) LogicalDeleteLoginRepository(tx *gorm.DB, login authmodel.Login) error {
	// ログイン情報を論理削除
	return tx.Model(&login).Where("employee_id = ?", login.EmployeeId).Updates(
		authmodel.Login{
			DeleteFlg:       login.DeleteFlg,
			UpdatedAuthor:   login.UpdatedAuthor,
			UpdatedDatetime: login.UpdatedDatetime,
		}).Error
}
