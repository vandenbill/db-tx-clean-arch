package main

import (
	"log"

	"github.com/vandenbill/db-tx-clean-arch/internal/domain"
	"github.com/vandenbill/db-tx-clean-arch/internal/repo"
	"github.com/vandenbill/db-tx-clean-arch/pkg"
)

func main() {
	db := pkg.PostgreConn()
	pkg.AutoMigrate(db)

	repoReg := repo.NewRepoRegistry(db)

	userRepo := repoReg.NewUserRepo()
	if err := userRepo.Create(domain.User{ID: 1, Username: "user1"}); err != nil {
		log.Println(err)
	}

	repoReg.DoInTx(func(rr repo.RepoRegistry) (interface{}, error) {
		txUserRepo := rr.NewUserRepo()

		if err := txUserRepo.Create(domain.User{ID: 2, Username: "user2"}); err != nil {
			return nil, err
		}

		/*
			should produce an err bcs duplicate id, so the rollback will be triggered
			and transaction came before this is not comitted
		*/
		if err := txUserRepo.Create(domain.User{ID: 2, Username: "user2"}); err != nil {
			return nil, err
		}

		return nil, nil
	})

	if err := userRepo.Create(domain.User{ID: 4, Username: "user4"}); err != nil {
		log.Println(err)
	}
}
