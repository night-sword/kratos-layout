package data

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/night-sword/kratos-kit/errors"
	"github.com/night-sword/kratos-kit/log"

	"github.com/night-sword/kratos-layout/internal/dao"
)

type Database struct {
	db    *sql.DB
	query *dao.Queries
}

func NewDatabase(data *Data) *Database {
	return NewDatabaseWithDB(data.db)
}

func NewDatabaseWithDB(db *sql.DB) *Database {
	return &Database{
		db:    db,
		query: newDao(db),
	}
}

func newDao(db *sql.DB) (querys *dao.Queries) {
	return dao.New(db)
}

func (inst *Database) Query() (querys *dao.Queries) {
	return inst.query
}

func (inst *Database) WithTx(ctx context.Context) (txCtx *TxContext, err error) {
	tx, err := inst.db.Begin()
	if err != nil {
		err = errors.InternalServer(errors.RsnInternal, "begin transaction fail").WithCause(err)
		return
	}

	txCtx = &TxContext{
		Context: ctx,
		tx:      tx,
		querys:  inst.query.WithTx(tx),
	}

	return
}

func (inst *Database) Commit(ctx context.Context) (err error) {
	txCtx, ok := ctx.(*TxContext)
	if !ok {
		log.Warn("not in transaction, useless transaction-commit")
		return
	}

	tx := txCtx.GetTx()
	if tx == nil {
		err = errors.InternalServer(errors.RsnInternal, "context tx is nil")
		return
	}

	err = tx.Commit()
	return
}

func (inst *Database) Dao(ctx context.Context) (querys *dao.Queries) {
	txCtx, ok := ctx.(*TxContext)
	if !ok {
		return inst.query
	}

	if querys = txCtx.GetQuerys(); querys == nil {
		querys = inst.query
	}

	return querys
}

func (inst *Database) TxDo(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	txCtx, err := inst.WithTx(ctx)
	if err != nil {
		return
	}
	defer func() { _ = txCtx.Rollback() }()

	err = fn(txCtx)
	if err != nil {
		return
	}

	err = inst.Commit(txCtx)
	return
}

// --- Transaction --- //

type TxContext struct {
	context.Context

	tx     *sql.Tx
	querys *dao.Queries
}

func (inst *TxContext) GetTx() (tx *sql.Tx) {
	return inst.tx
}

func (inst *TxContext) GetQuerys() (queries *dao.Queries) {
	return inst.querys
}

func (inst *TxContext) Rollback() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.InternalServer(errors.RsnInternal, fmt.Sprintf("%s", r))
		}
	}()
	return inst.tx.Rollback()
}
