package test

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

func AddStepSeeValue(sc *godog.ScenarioContext) {
	sc.Then(`^I should see a value "([^"]+)"$`,
		seeValue,
	)

	return
}

func seeValue(ctx0 context.Context, val string) (ctx context.Context, e error) {
	ctx = ctx0

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
