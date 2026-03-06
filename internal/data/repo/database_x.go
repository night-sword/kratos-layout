package repo

import "github.com/night-sword/kratos-layout/internal/conf"

// ------ multi database demo ----- //
//type OtherDatabase struct{ *Database }
//
//func OtherDatabase(cfg *conf.Bootstrap) (inst *OtherDatabase, cleanup func(), err error) {
//	return newDatabaseT(cfg.GetData().OtherDatabase(), func(db *Database) *OtherDatabase {
//		return &OtherDatabase{db}
//	})
//}

// ----------- //

func newDatabaseT[T any](cfg *conf.Data_Database, constructor func(*Database) *T) (inst *T, cleanup func(), err error) {
	db, cleanup, err := newDatabase(cfg)
	if err != nil {
		return
	}

	inst = constructor(db)
	return
}
