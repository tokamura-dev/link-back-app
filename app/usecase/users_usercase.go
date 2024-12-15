package usecase

import (
	"link-back-app/api"
	"link-back-app/domain/repository"
	"link-back-app/enum"
	usersmodel "link-back-app/models/users_model"
	stringutil "link-back-app/utils/string_util"
	timeutil "link-back-app/utils/time_util"
	usersutil "link-back-app/utils/users_util"
	"net/http"

	"gorm.io/gorm"
)

type UsersUsecase interface {
	GetOneByEmployeeIdUsersUsecase(employeeId string) ([]interface{}, error)
	GetAllUsersUsecase() ([]interface{}, error)
	RegisterUsersUsecase(requestCreateUsers usersmodel.RequestCreateUsers) error
	UpdateUsersUsecase(users usersmodel.Users) error
	LogicalDeleteUsersUsecase(employeeId string) error
	DeleteUsersUsecase(employeeId string) error
}

func NewUsersUsecase(database *gorm.DB, usersRepository repository.UsersRepository) UsersUsecase {
	return &usersUsecaseImpl{
		database:        database,
		usersRepository: usersRepository,
	}
}

type usersUsecaseImpl struct {
	database        *gorm.DB
	usersRepository repository.UsersRepository
}

/**
 * ユーザー情報1件取得(社員ID指定)
 **/
func (u *usersUsecaseImpl) GetOneByEmployeeIdUsersUsecase(employeeId string) ([]interface{}, error) {
	data, err := u.usersRepository.GetOneByKeyUsersRepository(employeeId)
	if err != nil {
		return nil, api.NewApiError(http.StatusInternalServerError, err.Error())
	}
	var genericData []interface{}
	genericData = append(genericData, data)
	return genericData, nil
}

/**
 * ユーザー情報全件取得処理
 */
func (u *usersUsecaseImpl) GetAllUsersUsecase() ([]interface{}, error) {
	// 全件取得処理
	datas, err := u.usersRepository.GetAllUsersRepository()
	if err != nil {
		return nil, api.NewApiError(http.StatusInternalServerError, err.Error())
	}
	var genericData []interface{}
	for _, data := range datas {
		genericData = append(genericData, data)
	}
	return genericData, nil
}

/**
 * ユーザー情報登録処理
 **/
func (u *usersUsecaseImpl) RegisterUsersUsecase(requestCreateUsers usersmodel.RequestCreateUsers) error {
	// ユーザー情報テーブルから最新の社員IDを取得する
	maxEmployeeId, err := u.usersRepository.GetMaxEmployeeIdUsersRepository()
	if err != nil {
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}

	// 最新の社員IDが取得できない場合は"000001"を設定
	if stringutil.IsEmpty(maxEmployeeId) {
		maxEmployeeId = INIT_EMPLOYEE_CD
	} else {
		maxEmployeeId = usersutil.GetNextEmployeeId(maxEmployeeId)
	}

	// ユーザー情報への値設定
	users := usersmodel.Users{
		EmployeeId:        maxEmployeeId,
		UserLastName:      requestCreateUsers.UserLastName,
		UserFirstName:     requestCreateUsers.UserFirstName,
		UserLastNameKana:  requestCreateUsers.UserLastNameKana,
		UserFirstNamaKana: requestCreateUsers.UserFirstNamaKana,
		DateOfJoin:        requestCreateUsers.DateOfJoin,
		BirthDay:          requestCreateUsers.BirthDay,
		Age:               requestCreateUsers.Age,
		Gender:            requestCreateUsers.Gender,
		PhoneNumber:       requestCreateUsers.PhoneNumber,
		MailAddress:       requestCreateUsers.MailAddress,
		Zipcode:           requestCreateUsers.Zipcode,
		Prefcode:          requestCreateUsers.Prefcode,
		Prefecture:        requestCreateUsers.Prefecture,
		Municipalities:    requestCreateUsers.Municipalities,
		Building:          requestCreateUsers.Building,
		DeleteFlg:         0,
		CreatedAuthor:     requestCreateUsers.UserLastName + requestCreateUsers.UserFirstName,
		CreatedDatetime:   timeutil.GetTimeNow(),
	}

	tx := u.database.Begin()
	// ユーザー情報テーブルへの登録処理
	err = u.usersRepository.RegisterUsersRepository(tx, users)
	if err != nil {
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

/**
 * ユーザー情報更新処理
 **/
func (u *usersUsecaseImpl) UpdateUsersUsecase(users usersmodel.Users) error {
	// ユーザー情報テーブルから社員IDに紐づくユーザー情報が存在するか確認
	count, err := u.usersRepository.GetExistUsersRepository(users.EmployeeId)
	if err != nil {
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}
	// 社員IDに紐づくユーザー情報が存在しない場合は
	// エラーとする
	if count < 1 {
		return api.NewApiError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
	nowTime := timeutil.GetTimeNow()
	users.UpdatedAuthor = ""
	users.UpdatedDatetime = &nowTime

	tx := u.database.Begin()
	// ユーザー情報更新処理
	err = u.usersRepository.UpdateUsersRepository(tx, users)
	if err != nil {
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

/**
 * ユーザー情報論理削除処理
 **/
func (u *usersUsecaseImpl) LogicalDeleteUsersUsecase(employeeId string) error {
	// ユーザー情報テーブルから社員IDに紐づくユーザー情報が存在するか確認
	count, err := u.usersRepository.GetExistUsersRepository(employeeId)
	if err != nil {
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}
	// 社員IDに紐づくユーザー情報が存在しない場合は
	// エラーとする
	if count < 1 {
		return api.NewApiError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	nowTime := timeutil.GetTimeNow()
	// 社員IDに紐づくユーザー情報を論理削除
	users := usersmodel.Users{
		EmployeeId:      employeeId,
		DeleteFlg:       enum.LogicalDeleteDiv.Deleted.Code,
		UpdatedAuthor:   "",
		UpdatedDatetime: &nowTime,
	}
	tx := u.database.Begin()
	err = u.usersRepository.LogicalDeleteUsersRepository(tx, users)
	if err != nil {
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

/**
 * ユーザー情報削除処理
 **/
func (u *usersUsecaseImpl) DeleteUsersUsecase(employeeId string) error {
	// ユーザー情報テーブルから社員IDに紐づくユーザー情報が存在するか確認
	count, err := u.usersRepository.GetExistUsersRepository(employeeId)
	if err != nil {
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}
	// 社員IDに紐づくユーザー情報が存在しない場合はエラーとする
	if count < 1 {
		return api.NewApiError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
	tx := u.database.Begin()
	// 社員IDに紐づくユーザー情報を削除
	err = u.usersRepository.DeleteUsersRepository(tx, employeeId)
	if err != nil {
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
