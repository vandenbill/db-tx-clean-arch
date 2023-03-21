this repo provide simple example how to doing database transaction while also maintaining technology
agnostic in clean architecture

basically just make a wrapper func for db transaction mode (manual commit) and switch the executor from auto-commit to manual commit executor and check if there is err in wrapper func we do rollback to maintaining data integration

this example we use the famous orm, gorm library, if we want to manual-commit transaction we can just make concrite gorm.DB struct and call Begin() method from DB struct, it will return db transaction executor(manual commit) that we use to replace the default executor(auto-commit)
this operation is happen in wrapper

```
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
```

if u use package that have different struct for different executor, u can add interface type in registry struct field that both manual and auto commit executor implement and also auto-commit executor, and use that interface in each repo, so while we make new repo, we can switch between that 2 executor

the concept also can be implemented if u add each repo to the service struct for better space complexity, just make wrapper that switch the executor

if u looking from some issue or mistake in the example, just make an issue, it really helps me
