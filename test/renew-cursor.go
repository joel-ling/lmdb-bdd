package test

import (
	"context"

	"github.com/PowerDNS/lmdb-go/lmdb"
	"github.com/cucumber/godog"
)

func AddStepRenewCursor(sc *godog.ScenarioContext) {
	sc.When(`^I renew the transaction relating to the cursor, and the cursor$`,
		renewCursor,
	)

	return
}

func renewCursor(ctx0 context.Context) (ctx context.Context, e error) {
	ctx = ctx0

	var (
		cur *lmdb.Cursor = ctx.Value(ctxKeyLMDBCur{}).(*lmdb.Cursor)
	)

	cur.Txn().Renew()

	cur.Renew(
		cur.Txn(),
	)

	return
}
