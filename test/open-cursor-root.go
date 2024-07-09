package test

import (
	"context"

	"github.com/PowerDNS/lmdb-go/lmdb"
	"github.com/cucumber/godog"
)

func AddStepOpenCursorRoot(sc *godog.ScenarioContext) {
	sc.When(`^I open a cursor to the root database of environment "([^"]+)"$`,
		openCursorRoot,
	)

	return
}

func openCursorRoot(ctx0 context.Context, envName string) (
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

	dbi, e = txn.OpenRoot(0)
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
