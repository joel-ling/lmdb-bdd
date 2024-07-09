package test

import (
	"context"
	"os"

	"github.com/PowerDNS/lmdb-go/lmdb"
	"github.com/cucumber/godog"
)

func AddStepCleanUp(sc *godog.ScenarioContext) {
	sc.After(cleanUp)

	return
}

func cleanUp(ctx0 context.Context, scenario *godog.Scenario, e0 error) (
	ctx context.Context, e error,
) {
	ctx = ctx0

	if ctx.Value(ctxKeyLMDBCur{}) != nil {
		ctx.Value(ctxKeyLMDBCur{}).(*lmdb.Cursor).Txn().Abort()
	}

	e = os.RemoveAll(
		ctx.Value(ctxKeyTempDir{}).(string),
	)
	if e != nil {
		return
	}

	return
}
