package manager

import "enigmacamp.com/golang-sample/repository"

type RepositoryManager interface {
	CustomerRepo() repository.CustomerRepository
}

type repoistoryManager struct {
	infra InfraManager
}

func (r *repoistoryManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerDbRepository(r.infra.DbConn())
}

func NewRepositoryManager(manager InfraManager) RepositoryManager {
	return &repoistoryManager{infra: manager}
}
