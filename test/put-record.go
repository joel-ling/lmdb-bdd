package test

import (
	"context"

	"github.com/PowerDNS/lmdb-go/lmdb"
	"github.com/cucumber/godog"
)

func AddStepPutRecord(sc *godog.ScenarioContext) {
	sc.Given(`^there is a record "([^"]+)" "([^"]+)" in DB "([^"]+)" `+
		`of environment "([^"]+)"$`,
		putRecord,
	)

	sc.When(`^I put a record "([^"]+)" "([^"]+)" in DB "([^"]+)" `+
		`of environment "([^"]+)"$`,
		putRecord,
	)

	return
}

func putRecord(ctx0 context.Context, key, val, dbName, envName string) (
	ctx context.Context, e error,
) {
	ctx = ctx0

	var (
		env *lmdb.Env = ctx.Value(ctxKeyLMDBEnv{envName}).(*lmdb.Env)

		put = func(txn *lmdb.Txn) (err error) {
			var (
				dbi lmdb.DBI
			)

			dbi, err = txn.OpenDBI(dbName, lmdb.Create)
			if err != nil {
				return
			}

			err = txn.Put(dbi,
				[]byte(key),
				[]byte(val),
				0,
			)
			if err != nil {
				return
			}

			return
		}
	)

	e = env.Update(put)
	if e != nil {
		return
	}

	return
}
