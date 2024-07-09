package test

import (
	"context"

	"github.com/PowerDNS/lmdb-go/lmdb"
	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

func AddStepSeeError(sc *godog.ScenarioContext) {
	sc.Then(`^I should see an error "([^"]+)"$`,
		seeError,
	)

	return
}

func seeError(ctx0 context.Context, errName string) (
	ctx context.Context, e error,
) {
	ctx = ctx0

	var (
		err error = ctx.Value(ctxKeyLMDBErr{}).(error)
	)

	switch errName {
	case "MDB_NOTFOUND":
		assert.True(
			godog.T(ctx),
			lmdb.IsNotFound(err),
		)

	default:
		godog.T(ctx).Fail()
	}

	return
}
