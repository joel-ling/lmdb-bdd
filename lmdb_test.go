package lmdbbdd

import (
	"testing"

	"github.com/cucumber/godog"

	"github.com/joel-ling/lmdb-bdd/test"
)

func TestLMDB(t *testing.T) {
	var (
		scenarioInitializer = func(sc *godog.ScenarioContext) {
			test.AddStepSetUp(sc)
			test.AddStepNewLMDBEnv(sc)
			test.AddStepPutRecord(sc)
			test.AddStepGetRecord(sc)
			test.AddStepOpenCursor(sc)
			test.AddStepOpenCursorRoot(sc)
			test.AddStepGetNextRecord(sc)
			test.AddStepGetNextRecordSet(sc)
			test.AddStepSeeRecord(sc)
			test.AddStepResetCursor(sc)
			test.AddStepRenewCursor(sc)
			test.AddStepSeeKey(sc)
			test.AddStepSeeValue(sc)
			test.AddStepSeeError(sc)
			test.AddStepSeeNoError(sc)
			test.AddStepIDLastCommit(sc)
			test.AddStepCleanUp(sc)

			return
		}

		options = &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		}

		suite = godog.TestSuite{
			ScenarioInitializer: scenarioInitializer,
			Options:             options,
		}
	)

	if suite.Run() != 0 {
		t.Fatal()
	}

	return
}
