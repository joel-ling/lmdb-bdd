package test

import (
	"context"
	"os"

	"github.com/PowerDNS/lmdb-go/lmdb"
	"github.com/cucumber/godog"
)

func AddStepNewLMDBEnv(sc *godog.ScenarioContext) {
	sc.Given(`^there is a new LMDB environment "([^"]+)"$`,
		newLMDBEnv,
	)

	return
}

func newLMDBEnv(ctx0 context.Context, name string) (
	ctx context.Context, e error,
) {
	ctx = ctx0

	var (
		newEnv *lmdb.Env
		path   string
	)

	newEnv, e = lmdb.NewEnv()
	if e != nil {
		return
	}

	e = newEnv.SetMaxDBs(4)
	if e != nil {
		return
	}

	path, e = os.MkdirTemp(
		ctx.Value(ctxKeyTempDir{}).(string),
		name,
	)
	if e != nil {
		return
	}

	e = newEnv.Open(path, 0, 0644)
	if e != nil {
		return
	}

	ctx = context.WithValue(ctx, ctxKeyLMDBEnv{name}, newEnv)

	return
}
