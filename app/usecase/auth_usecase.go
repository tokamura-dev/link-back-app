package usecase

import (
	"link-back-app/api"
	"link-back-app/domain/repository"
	"link-back-app/enum"
	authmodel "link-back-app/models/auth_model"
	usersmodel "link-back-app/models/users_model"
	cryptoutil "link-back-app/utils/crypto_util"
	stringutil "link-back-app/utils/string_util"
	timeutil "link-back-app/utils/time_util"
	usersutil "link-back-app/utils/users_util"
	"net/http"
	"reflect"

	"gorm.io/gorm"
)

/** 社員IDの初期値 */
const INIT_EMPLOYEE_CD = "000001"

type AuthUsecase interface {
	SignUpUsacase(signup authmodel.RequestSignUp) error
	SignInUsecase(signin authmodel.RequstSignIn) (bool, error)
	LogicalDeleteLoginUsecase(employeeId string) error
}

func NewAuthUsecase(database *gorm.DB, authRepository repository.AuthRepository, usersRepository repository.UsersRepository) AuthUsecase {
	return &authUsecaseImpl{
		database:        database,
		authRepository:  authRepository,
		usersRepository: usersRepository,
	}
}

type authUsecaseImpl struct {
	database        *gorm.DB
	authRepository  repository.AuthRepository
	usersRepository repository.UsersRepository
}

/**
 * サインアップ処理
 **/
func (a *authUsecaseImpl) SignUpUsacase(signup authmodel.RequestSignUp) error {
	// ログイン情報テーブルから最新の社員IDを取得する
	maxEmployeeId, err := a.authRepository.GetMaxEmployeeIdLoginRepository()
	if err != nil {
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}

	// 最新の社員IDが取得できない場合は"000001"を設定
	if stringutil.IsEmpty(maxEmployeeId) {
		maxEmployeeId = INIT_EMPLOYEE_CD
	} else {
		maxEmployeeId = usersutil.GetNextEmployeeId(maxEmployeeId)
	}

	// 平文のパスワードを暗号化
	cryptoedPassword, err := cryptoutil.PasswordEncrypt(signup.Password)
	if err != nil {
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}

	tx := a.database.Begin()
	timeNow := timeutil.GetTimeNow()
	// ログイン情報を設定
	login := authmodel.Login{
		EmployeeId:      maxEmployeeId,
		Password:        cryptoedPassword,
		DeleteFlg:       enum.LogicalDeleteDiv.NotDeleted.Code,
		CreatedAuthor:   "",
		CreatedDatetime: timeNow,
	}
	if err := a.authRepository.RegisterLoginRepository(tx, login); err != nil {
		tx.Rollback()
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}

	// ユーザー情報を設定
	users := usersmodel.Users{
		EmployeeId:        maxEmployeeId,
		UserLastName:      signup.UserLastName,
		UserFirstName:     signup.UserFirstName,
		UserLastNameKana:  signup.UserLastNameKana,
		UserFirstNamaKana: signup.UserFirstNamaKana,
		DateOfJoin:        signup.DateOfJoin,
		BirthDay:          signup.BirthDay,
		Age:               signup.Age,
		Gender:            signup.Gender,
		PhoneNumber:       signup.PhoneNumber,
		MailAddress:       signup.MailAddress,
		Zipcode:           signup.Zipcode,
		Prefcode:          signup.Prefcode,
		Prefecture:        signup.Prefecture,
		Municipalities:    signup.Municipalities,
		Building:          signup.Building,
		DeleteFlg:         enum.LogicalDeleteDiv.NotDeleted.Code,
		CreatedAuthor:     "",
		CreatedDatetime:   timeNow,
	}
	// ログイン情報登録処理
	if err = a.usersRepository.RegisterUsersRepository(tx, users); err != nil {
		tx.Rollback()
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}
	return tx.Commit().Error
}

/**
 * サインイン処理
 **/
func (a *authUsecaseImpl) SignInUsecase(signin authmodel.RequstSignIn) (bool, error) {
	// 社員IDに紐づくログイン情報を取得
	login, err := a.authRepository.GetOneByKeyLoginRepository(signin.EmployeeId)
	if err != nil {
		return false, api.NewApiError(http.StatusInternalServerError, err.Error())
	}
	var loginStruct authmodel.Login
	// ログイン情報がない場合はエラー
	if reflect.DeepEqual(loginStruct, login) {
		return false, api.NewApiError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
	// パスワード検証処理
	err = cryptoutil.CompareHashAndPassword(login.Password, signin.Password)
	if err != nil {
		return false, api.NewApiError(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	}
	return true, nil
}

/**
 * ログイン情報関連の論理削除処理
 **/
func (a *authUsecaseImpl) LogicalDeleteLoginUsecase(employeeId string) error {
	// 社員IDに紐づくログイン情報を取得
	login, err := a.authRepository.GetOneByKeyLoginRepository(employeeId)
	if err != nil {
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}
	var loginStruct authmodel.Login
	// ログイン情報がない場合はエラー
	if reflect.DeepEqual(loginStruct, login) {
		return api.NewApiError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
	tx := a.database.Begin()
	timeNow := timeutil.GetTimeNow()
	login = authmodel.Login{
		EmployeeId:      employeeId,
		DeleteFlg:       enum.LogicalDeleteDiv.Deleted.Code,
		UpdatedAuthor:   "",
		UpdatedDatetime: &timeNow,
	}
	// ログイン情報論理削除処理
	if err := a.authRepository.LogicalDeleteLoginRepository(tx, login); err != nil {
		tx.Rollback()
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}
	// ユーザー情報論理削除処理
	users := usersmodel.Users{
		EmployeeId:      employeeId,
		DeleteFlg:       enum.LogicalDeleteDiv.Deleted.Code,
		UpdatedAuthor:   "",
		UpdatedDatetime: &timeNow,
	}
	if err := a.usersRepository.LogicalDeleteUsersRepository(tx, users); err != nil {
		tx.Rollback()
		return api.NewApiError(http.StatusInternalServerError, err.Error())
	}
	return tx.Commit().Error
}
