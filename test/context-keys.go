package test

type (
	ctxKeyTempDir struct{}

	ctxKeyLMDBEnv struct {
		Name string
	}

	ctxKeyLMDBCur struct{}

	ctxKeyLMDBErr struct{}

	ctxKeyLMDBKey struct{}

	ctxKeyLMDBVal struct{}
)
