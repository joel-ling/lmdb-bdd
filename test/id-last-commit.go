package test

import (
	"context"

	"github.com/PowerDNS/lmdb-go/lmdb"
	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

func AddStepIDLastCommit(sc *godog.ScenarioContext) {
	sc.Then(`^the ID of last committed transaction in environment "([^"]+)" `+
		`should be (\d+)$`,
		idLastCommit,
	)

	return
}

func idLastCommit(ctx0 context.Context, envName string, txnID int64) (
	ctx context.Context, e error,
) {
	ctx = ctx0

	var (
		env *lmdb.Env = ctx.Value(ctxKeyLMDBEnv{envName}).(*lmdb.Env)

		envInfo *lmdb.EnvInfo
	)

	envInfo, e = env.Info()
	if e != nil {
		return
	}

	assert.Equal(
		godog.T(ctx),
		txnID,
		envInfo.LastTxnID,
	)

	return
}
