package repo

import (
	"log"

	"github.com/vandenbill/db-tx-clean-arch/internal/domain"
	"gorm.io/gorm"
)

type repoRegistry struct {
	db *gorm.DB
}

func NewRepoRegistry(db *gorm.DB) RepoRegistry {
	return &repoRegistry{db}
}

func (rr *repoRegistry) DoInTx(txFunc TxFunc) (interface{}, error) {
	txExecutor := rr.db.Begin()
	if txExecutor.Error != nil {
		return nil, txExecutor.Error
	}

	txRepoRegistry := rr
	rr.db = txExecutor

	out, txFuncErr := txFunc(txRepoRegistry)
	if txFuncErr != nil {
		log.Printf("ROLLBACK err: %s\n", txFuncErr)
		err := txExecutor.Rollback().Error
		if err != nil {
			log.Printf("ERR WHILE ROLLBACK err: %s\n", err)
		}
		return nil, txFuncErr
	}

	if err := txExecutor.Commit().Error; err != nil {
		log.Printf("FAIL COMMIT err: %s\n", err)
	}

	return out, nil
}

func (rr *repoRegistry) NewUserRepo() UserRepo {
	return &userRepo{db: rr.db}
}

type userRepo struct {
	db *gorm.DB
}

func (ur *userRepo) Create(user domain.User) error {
	if err := ur.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
