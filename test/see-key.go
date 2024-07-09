package test

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

func AddStepSeeKey(sc *godog.ScenarioContext) {
	sc.Then(`^I should see a key "([^"]+)"$`,
		seeKey,
	)

	return
}

func seeKey(ctx0 context.Context, key string) (ctx context.Context, e error) {
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

	return
}
