package fuzzdecimal

import (
	"fmt"
	"testing"
)

// T is a wrapper around testing.T that propagated the fuzzy seeds to the subtest called with Run.
type T struct {
	*testing.T

	seeds []string
}

func (ft *T) Run(name string, f func(t *T)) {
	ft.T.Run(name, func(t *testing.T) {
		t.Helper()

		f(&T{
			seeds: ft.seeds,
			T:     t,
		})
	})
}

func (ft *T) assertStaticSeedsCount(expectedCount int, funcName string) {
	if len(ft.seeds) != expectedCount {
		ft.Fatalf("calling %[1]s%[2]d with %[3]d seeds, please call %[1]s%[3]d instead", funcName, expectedCount, len(ft.seeds))
	}
}

func (ft *T) seedsErrorMessage(linePrefix string) string {
	var errorMessage string
	for index, seed := range ft.seeds {
		errorMessage += fmt.Sprintf("%sindex %d value: '%s'\n", linePrefix, index, seed)
	}

	return errorMessage
}
