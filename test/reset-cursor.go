package test

import (
	"context"

	"github.com/PowerDNS/lmdb-go/lmdb"
	"github.com/cucumber/godog"
)

func AddStepResetCursor(sc *godog.ScenarioContext) {
	sc.When(`^I reset the transaction relating to the cursor$`,
		resetCursor,
	)

	return
}

func resetCursor(ctx0 context.Context) (ctx context.Context, e error) {
	ctx = ctx0

	var (
		cur *lmdb.Cursor = ctx.Value(ctxKeyLMDBCur{}).(*lmdb.Cursor)
	)

	cur.Txn().Reset()

	return
}
