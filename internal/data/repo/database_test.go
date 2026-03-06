package repo

import (
	"context"
	"database/sql"
	"testing"

	"github.com/night-sword/kratos-layout/internal/dao"
)

func TestDatabase_Dao_normalContext(t *testing.T) {
	defaultQuery := dao.New(nil)
	db := &Database{query: defaultQuery}

	got := db.Dao(context.Background())
	if got != defaultQuery {
		t.Error("Dao() with normal context should return default query")
	}
}

func TestDatabase_Dao_txContext(t *testing.T) {
	defaultQuery := dao.New(nil)
	txQuery := dao.New(nil)
	db := &Database{query: defaultQuery}

	txCtx := &TxContext{
		Context: context.Background(),
		queries:  txQuery,
	}

	got := db.Dao(txCtx)
	if got != txQuery {
		t.Error("Dao() with TxContext should return tx query")
	}
}

func TestDatabase_Dao_txContextNilQuerys(t *testing.T) {
	defaultQuery := dao.New(nil)
	db := &Database{query: defaultQuery}

	txCtx := &TxContext{
		Context: context.Background(),
		queries:  nil,
	}

	got := db.Dao(txCtx)
	if got != defaultQuery {
		t.Error("Dao() with TxContext(nil queries) should fallback to default query")
	}
}

func TestTxContext_Rollback_afterCommit(t *testing.T) {
	// Rollback on already-committed tx should return nil (ErrTxDone is ignored)
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/")
	if err != nil {
		t.Skip("skip: cannot open mysql connection:", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		t.Skip("skip: mysql not reachable:", err)
	}

	tx, err := db.Begin()
	if err != nil {
		t.Fatal("Begin() error:", err)
	}

	if err = tx.Commit(); err != nil {
		t.Fatal("Commit() error:", err)
	}

	txCtx := &TxContext{
		Context: context.Background(),
		tx:      tx,
	}

	if err = txCtx.Rollback(); err != nil {
		t.Errorf("Rollback() after commit should return nil, got: %v", err)
	}
}
