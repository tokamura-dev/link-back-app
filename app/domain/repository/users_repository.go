package repository

import (
	usersmodel "link-back-app/models"

	"gorm.io/gorm"
)

type UsersRepository interface {
	GetMaxEmployeeIdUsersRepository() (string, error)
	GetExistUsersRepository(employeeId string) (int, error)
	GetOneByEmployeeIdUsersRepository(employeeId string) (usersmodel.Users, error)
	GetAllUsersRepository() ([]usersmodel.Users, error)
	RegisterUsersRepository(users usersmodel.Users) error
	LogicalDeleteUsersRepository(requestUpdateUsers usersmodel.RequestUpdateUsers) error
	DeleteUsersRepository(employeeId string) error
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
func (u *usersRepositoryImpl) GetMaxEmployeeIdUsersRepository() (string, error) {
	var users usersmodel.Users
	var maxEmployeeId string
	err := u.db.Model(&users).Select("COALESCE(MAX(employee_id), '') AS max_employee_id").Scan(&maxEmployeeId).Error
	return maxEmployeeId, err
}

/**
 * 社員IDに紐づくユーザー情報が存在するかを確認する処理
 **/
func (u *usersRepositoryImpl) GetExistUsersRepository(employeeId string) (int, error) {
	var users usersmodel.Users
	var count int
	err := u.db.Model(&users).Where("employee_id = ?", &employeeId).Select("COUNT(*) AS count").Scan(&count).Error
	return count, err
}

/**
 * ユーザー情報1件取得処理(社員ID指定)
 **/
func (u *usersRepositoryImpl) GetOneByEmployeeIdUsersRepository(employeeId string) (usersmodel.Users, error) {
	users := usersmodel.Users{}
	err := u.db.Debug().Where("employee_id = ?", employeeId).Take(&users).Error
	return users, err
}

/**
 * ユーザー情報全件取得処理
 */
func (u *usersRepositoryImpl) GetAllUsersRepository() ([]usersmodel.Users, error) {
	users := []usersmodel.Users{}
	err := u.db.Debug().Find(&users).Error
	return users, err
}

/**
 * ユーザー情報登録処理
 */
func (u *usersRepositoryImpl) RegisterUsersRepository(users usersmodel.Users) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		return tx.Debug().Create(&users).Error
	})
	return err
}

/**
 * ユーザー情報論理削除処理
 **/
func (u *usersRepositoryImpl) LogicalDeleteUsersRepository(requestUpdateUsers usersmodel.RequestUpdateUsers) error {
	var uses usersmodel.Users
	err := u.db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&uses).Where("employee_id = ?", requestUpdateUsers.EmployeeId).
			Updates(usersmodel.Users{
				DeleteFlg:     requestUpdateUsers.DeleteFlg,
				UpdatedAuthor: requestUpdateUsers.UpdatedAuthor,
				UpdatedDate:   requestUpdateUsers.UpdatedDate,
			}).Error
	})
	return err
}

/**
 * ユーザー情報削除処理
 **/
func (u *usersRepositoryImpl) DeleteUsersRepository(employeeId string) error {
	var users usersmodel.Users
	err := u.db.Transaction(func(tx *gorm.DB) error {
		return tx.Where("employee_id = ?", employeeId).Delete(&users).Error
	})
	return err
}
