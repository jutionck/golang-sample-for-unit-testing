package manager

import "enigmacamp.com/golang-sample/usecase"

type UseCaseManager interface {
	CustomerUseCase() usecase.CustomerUsecase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) CustomerUseCase() usecase.CustomerUsecase {
	return usecase.NewCustomerUseCase(u.repoManager.CustomerRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
