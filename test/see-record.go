package test

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

func AddStepSeeRecord(sc *godog.ScenarioContext) {
	sc.Then(`^I should see a record "([^"]+)" "([^"]+)"$`,
		seeRecord,
	)

	return
}

func seeRecord(ctx0 context.Context, key, val string) (
	ctx context.Context, e error,
) {
	ctx = ctx0

	assert.NotNil(
		godog.T(ctx),
		ctx.Value(ctxKeyLMDBKey{}),
	)

	assert.Equal(
		godog.T(ctx),
		key,
		string(
			ctx.Value(ctxKeyLMDBKey{}).([]byte),
		),
	)

	assert.NotNil(
		godog.T(ctx),
		ctx.Value(ctxKeyLMDBVal{}),
	)

	assert.Equal(
		godog.T(ctx),
		val,
		string(
			ctx.Value(ctxKeyLMDBVal{}).([]byte),
		),
	)

	return
}
