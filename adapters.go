package fuzzdecimal

import "testing"

func convertParseDecimalFunc[Decimal any](t *T, parseDecimalFunc func(t *T, s string) (Decimal, error)) func(t *testing.T, s string) (Decimal, error) {
	t.Helper()

	return func(testingT *testing.T, s string) (Decimal, error) {
		testingT.Helper()

		return parseDecimalFunc(&T{
			T:     testingT,
			seeds: t.seeds,
		}, s)
	}
}
