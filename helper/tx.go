package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
