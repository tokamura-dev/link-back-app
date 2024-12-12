package usecase

import (
	"link-back-app/api"
	"link-back-app/domain/repository"
	"link-back-app/enum"
	usersmodel "link-back-app/models"
	stringutil "link-back-app/utils/string_util"
	usersutil "link-back-app/utils/users_util"
	"net/http"
	"time"
)

/** 社員IDの初期値 */
const INIT_EMPLOYEE_CD = "000001"

type UsersUsecase interface {
	GetOneByEmployeeIdUsersUsecase(employeeId string) ([]interface{}, error)
	GetAllUsersUsecase() ([]interface{}, error)
	RegisterUsersUsecase(requestCreateUsers usersmodel.RequestCreateUsers) error
	UpdateUsersUsecase(users usersmodel.Users) error
	LogicalDeleteUsersUsecase(employeeId string) error
	DeleteUsersUsecase(employeeId string) error
}

func NewUsersUsecase(repository repository.UsersRepository) UsersUsecase {
	return &usersUsecaseImpl{
		repository: repository,
	}
}

type usersUsecaseImpl struct {
	repository repository.UsersRepository
}

/**
 * ユーザー情報1件取得(社員ID指定)
 **/
func (u *usersUsecaseImpl) GetOneByEmployeeIdUsersUsecase(employeeId string) ([]interface{}, error) {
	data, err := u.repository.GetOneByEmployeeIdUsersRepository(employeeId)
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
	datas, err := u.repository.GetAllUsersRepository()
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
	maxEmployeeId, err := u.repository.GetMaxEmployeeIdUsersRepository()
	if err != nil {
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}

	// 最新の社員IDが取得できな場合は"000001"を設定
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
		CreatedDate:       time.Now().UTC(),
	}

	// ユーザー情報テーブルへの登録処理
	err = u.repository.RegisterUsersRepository(users)
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
	count, err := u.repository.GetExistUsersRepository(users.EmployeeId)
	if err != nil {
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}
	// 社員IDに紐づくユーザー情報が存在しない場合は
	// エラーとする
	if count < 1 {
		return api.NewApiError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
	nowTime := time.Now().UTC()
	users.UpdatedAuthor = ""
	users.UpdatedDate = &nowTime

	// ユーザー情報更新処理
	err = u.repository.UpdateUsersRepository(users)
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
	count, err := u.repository.GetExistUsersRepository(employeeId)
	if err != nil {
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}
	// 社員IDに紐づくユーザー情報が存在しない場合は
	// エラーとする
	if count < 1 {
		return api.NewApiError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	nowTime := time.Now().UTC()
	// 社員IDに紐づくユーザー情報を論理削除
	users := usersmodel.Users{
		EmployeeId:    employeeId,
		DeleteFlg:     enum.LogicalDeleteDiv.Deleted.Code,
		UpdatedAuthor: "",
		UpdatedDate:   &nowTime,
	}
	err = u.repository.LogicalDeleteUsersRepository(users)
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
	count, err := u.repository.GetExistUsersRepository(employeeId)
	if err != nil {
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}
	// 社員IDに紐づくユーザー情報が存在しない場合はエラーとする
	if count < 1 {
		return api.NewApiError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
	// 社員IDに紐づくユーザー情報を削除
	err = u.repository.DeleteUsersRepository(employeeId)
	if err != nil {
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
