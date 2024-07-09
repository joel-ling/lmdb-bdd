package test

import (
	"context"

	"github.com/PowerDNS/lmdb-go/lmdb"
	"github.com/cucumber/godog"
)

func AddStepGetRecord(sc *godog.ScenarioContext) {
	sc.When(
		`^I get a record "([^"]+)" in DB "([^"]+)" of environment "([^"]+)"$`,
		getRecord,
	)

	return
}

func getRecord(ctx0 context.Context, key, dbName, envName string) (
	ctx context.Context, e error,
) {
	ctx = ctx0

	var (
		val []byte

		env *lmdb.Env = ctx.Value(ctxKeyLMDBEnv{envName}).(*lmdb.Env)

		get = func(txn *lmdb.Txn) (err error) {
			var (
				dbi lmdb.DBI
			)

			dbi, err = txn.OpenDBI(dbName, 0)
			if err != nil {
				return
			}

			val, err = txn.Get(dbi,
				[]byte(key),
			)
			if err != nil {
				return
			}

			return
		}
	)

	e = env.View(get)
	if e != nil {
		ctx = context.WithValue(ctx, ctxKeyLMDBErr{}, e)

		e = nil
	}

	ctx = context.WithValue(ctx, ctxKeyLMDBVal{}, val)

	return
}
