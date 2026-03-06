package repo

import (
	"context"
	"database/sql"

	"github.com/night-sword/kratos-kit/errors"
	"github.com/night-sword/kratos-kit/log"

	"github.com/night-sword/kratos-layout/internal/conf"
	"github.com/night-sword/kratos-layout/internal/dao"
)

type Database struct {
	db    *sql.DB
	query *dao.Queries
}

func NewDatabase(cfg *conf.Bootstrap) (inst *Database, cleanup func(), err error) {
	return newDatabase(cfg.GetData().GetDatabase())
}

func newDatabase(cfg *conf.Data_Database) (inst *Database, cleanup func(), err error) {
	if cfg == nil || cfg.GetDriver() == "" || cfg.GetSource() == "" {
		err = errors.InternalServer(errors.RsnInternal, "database config is invalid")
		return
	}

	db, err := sql.Open(cfg.GetDriver(), cfg.GetSource())
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		_ = db.Close()
		return
	}

	cleanup = func() { closeResource(db, "database") }
	inst = &Database{
		db:    db,
		query: dao.New(db),
	}
	return
}

func (inst *Database) Query() (queries *dao.Queries) {
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
		queries: inst.query.WithTx(tx),
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

func (inst *Database) Dao(ctx context.Context) (queries *dao.Queries) {
	txCtx, ok := ctx.(*TxContext)
	if !ok {
		return inst.query
	}

	if queries = txCtx.GetQueries(); queries == nil {
		queries = inst.query
	}

	return queries
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

	tx      *sql.Tx
	queries *dao.Queries
}

func (inst *TxContext) GetTx() (tx *sql.Tx) {
	return inst.tx
}

func (inst *TxContext) GetQueries() (queries *dao.Queries) {
	return inst.queries
}

func (inst *TxContext) Rollback() error {
	err := inst.tx.Rollback()
	if err != nil && !errors.Is(err, sql.ErrTxDone) {
		return err
	}
	return nil
}
