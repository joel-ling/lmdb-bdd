package test

import (
	"context"

	"github.com/PowerDNS/lmdb-go/lmdb"
	"github.com/cucumber/godog"
)

func AddStepGetNextRecord(sc *godog.ScenarioContext) {
	sc.When(`^I get the next record using the cursor$`,
		getNextRecord,
	)

	return
}

func getNextRecord(ctx0 context.Context) (ctx context.Context, e error) {
	ctx = ctx0

	var (
		cursor *lmdb.Cursor = ctx.Value(ctxKeyLMDBCur{}).(*lmdb.Cursor)

		key []byte
		val []byte
	)

	key, val, e = cursor.Get(nil, nil, lmdb.Next)
	if e != nil {
		ctx = context.WithValue(ctx, ctxKeyLMDBErr{}, e)

		e = nil
	}

	ctx = context.WithValue(ctx, ctxKeyLMDBKey{}, key)
	ctx = context.WithValue(ctx, ctxKeyLMDBVal{}, val)

	return
}
