package test

import (
	"context"

	"github.com/PowerDNS/lmdb-go/lmdb"
	"github.com/cucumber/godog"
)

func AddStepOpenCursor(sc *godog.ScenarioContext) {
	sc.When(`^I open a cursor to DB "([^"]+)" of environment "([^"]+)"$`,
		openCursor,
	)

	return
}

func openCursor(ctx0 context.Context, dbName, envName string) (
	ctx context.Context, e error,
) {
	ctx = ctx0

	var (
		cur *lmdb.Cursor
		dbi lmdb.DBI
		txn *lmdb.Txn

		env *lmdb.Env = ctx.Value(ctxKeyLMDBEnv{envName}).(*lmdb.Env)
	)

	txn, e = env.BeginTxn(nil, lmdb.Readonly)
	if e != nil {
		return
	}

	dbi, e = txn.OpenDBI(dbName, 0)
	if e != nil {
		return
	}

	cur, e = txn.OpenCursor(dbi)
	if e != nil {
		return
	}

	ctx = context.WithValue(ctx, ctxKeyLMDBCur{}, cur)

	return
}
