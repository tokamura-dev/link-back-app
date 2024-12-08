package usecase

import (
	"link-back-app/domain/repository"
	"link-back-app/models"
	stringutil "link-back-app/utils/string_util"
	usersutil "link-back-app/utils/users_util"
	"time"
)

type UsersUsecase interface {
	GetAllUsersUsecase() ([]interface{}, error)
	RegisterUsersUsecase(requestUsers models.RequestUsers) error
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
 * ユーザー情報全件取得処理
 */
func (u *usersUsecaseImpl) GetAllUsersUsecase() ([]interface{}, error) {
	// 全件取得処理
	datas, err := u.repository.GetAllUsersRepository()
	if err != nil {
		return nil, err
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
func (u *usersUsecaseImpl) RegisterUsersUsecase(requestUsers models.RequestUsers) error {
	// ユーザー情報テーブルから最新の社員IDを取得する
	maxEmployeeId, _ := u.repository.GetMaxEmployeeIdUsers()

	// 最新の社員IDが取得できな場合は"000001"を設定
	if stringutil.IsEmpty(maxEmployeeId) {
		maxEmployeeId = "000001"
	} else {
		maxEmployeeId = usersutil.GetNextEmployeeId(maxEmployeeId)
	}

	// ユーザー情報への値設定
	users := models.Users{
		EmployeeId:        maxEmployeeId,
		UserLastName:      requestUsers.UserLastName,
		UserFirstName:     requestUsers.UserFirstName,
		UserLastNameKana:  requestUsers.UserLastNameKana,
		UserFirstNamaKana: requestUsers.UserFirstNamaKana,
		DateOfJoin:        requestUsers.DateOfJoin,
		BirthDay:          requestUsers.BirthDay,
		Age:               requestUsers.Age,
		Gender:            requestUsers.Gender,
		PhoneNumber:       requestUsers.PhoneNumber,
		MailAddress:       requestUsers.MailAddress,
		Zipcode:           requestUsers.Zipcode,
		Prefcode:          requestUsers.Prefcode,
		Prefecture:        requestUsers.Prefecture,
		Municipalities:    requestUsers.Municipalities,
		Building:          requestUsers.Building,
		DeleteFlg:         0,
		CreatedAuthor:     requestUsers.UserLastName + requestUsers.UserFirstName,
		CreatedDate:       time.Now(),
	}

	// ユーザー情報テーブルへの登録処理
	err := u.repository.RegisterUsersRepository(users)
	return err
}
