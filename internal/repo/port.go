package repo

import (
	"github.com/vandenbill/db-tx-clean-arch/internal/domain"
)

type TxFunc func(RepoRegistry) (interface{}, error)

type RepoRegistry interface {
	DoInTx(TxFunc) (interface{}, error)
	NewUserRepo() UserRepo
}

type UserRepo interface {
	Create(domain.User) error
}
