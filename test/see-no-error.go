package test

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

func AddStepSeeNoError(sc *godog.ScenarioContext) {
	sc.Then(`^I should see no error$`,
		seeNoError,
	)

	return
}

func seeNoError(ctx0 context.Context) (ctx context.Context, e error) {
	ctx = ctx0

	assert.Nil(
		godog.T(ctx),
		ctx.Value(ctxKeyLMDBErr{}),
	)

	return
}
