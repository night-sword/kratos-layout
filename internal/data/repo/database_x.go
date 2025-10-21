package repo

import "github.com/night-sword/kratos-layout/internal/conf"

// ----------- //

func newDatabaseT[T any](cfg *conf.Data_Database, constructor func(*Database) *T) (inst *T, cleanup func(), err error) {
	db, cleanup, err := newDatabase(cfg)
	if err != nil {
		return
	}

	inst = constructor(db)
	return
}
