package usecase

import (
	"link-back-app/domain/repository"
	"link-back-app/models"
)

type UsersUsecase interface {
	RegisterUsersUsecase()
	GetAllUsersUsecase() ([]models.Users, error)
}

func NewUsersUsecase(repo repository.UsersRepository) UsersUsecase {
	return &usersUsecaseImpl{
		repo: repo,
	}
}

type usersUsecaseImpl struct {
	repo repository.UsersRepository
}

/**
 * ユーザー情報全件取得処理
 */
func (u *usersUsecaseImpl) GetAllUsersUsecase() ([]models.Users, error) {
	// 全件取得処理
	datas, err := u.repo.GetAllUsersRepository()
	if err != nil {
		return nil, err
	}
	return datas, nil
}

func (u *usersUsecaseImpl) RegisterUsersUsecase() {
	panic("unimplemented")
}
