package fuzzdecimal

const asStringFuncName = "AsString"

func AsStringSlice(t *T, name string, fuzzFunc func(t *T, numbers []string)) {
	t.Helper()

	t.Run(name, func(t *T) {
		fuzzFunc(t, t.seeds)
	})
}

func AsString1(t *T, name string, fuzzFunc func(t *T, x1 string)) {
	t.Helper()

	AsStringSlice(t, name, func(t *T, numbers []string) {
		t.Helper()

		t.assertStaticSeedsCount(1, asStringFuncName)

		t.Run(name, func(t *T) {
			fuzzFunc(t, numbers[0])
		})
	})
}

func AsString2(t *T, name string, fuzzFunc func(t *T, x1, x2 string)) {
	t.Helper()

	AsStringSlice(t, name, func(t *T, numbers []string) {
		t.Helper()

		t.assertStaticSeedsCount(2, asStringFuncName)

		t.Run(name, func(t *T) {
			fuzzFunc(t, numbers[0], numbers[1])
		})
	})
}

func AsString3(t *T, name string, fuzzFunc func(t *T, x1, x2, x3 string)) {
	t.Helper()

	AsStringSlice(t, name, func(t *T, numbers []string) {
		t.Helper()

		t.assertStaticSeedsCount(3, asStringFuncName)

		t.Run(name, func(t *T) {
			fuzzFunc(t, numbers[0], numbers[1], numbers[2])
		})
	})
}

func AsString4(t *T, name string, fuzzFunc func(t *T, x1, x2, x3, x4 string)) {
	t.Helper()

	AsStringSlice(t, name, func(t *T, numbers []string) {
		t.Helper()

		t.assertStaticSeedsCount(4, asStringFuncName)

		t.Run(name, func(t *T) {
			fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3])
		})
	})
}

func AsString5(t *T, name string, fuzzFunc func(t *T, x1, x2, x3, x4, x5 string)) {
	t.Helper()

	AsStringSlice(t, name, func(t *T, numbers []string) {
		t.Helper()

		t.assertStaticSeedsCount(5, asStringFuncName)

		t.Run(name, func(t *T) {
			fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4])
		})
	})
}

func AsString6(t *T, name string, fuzzFunc func(t *T, x1, x2, x3, x4, x5, x6 string)) {
	t.Helper()

	AsStringSlice(t, name, func(t *T, numbers []string) {
		t.Helper()

		t.assertStaticSeedsCount(6, asStringFuncName)

		t.Run(name, func(t *T) {
			fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5])
		})
	})
}

func AsString7(t *T, name string, fuzzFunc func(t *T, x1, x2, x3, x4, x5, x6, x7 string)) {
	t.Helper()

	AsStringSlice(t, name, func(t *T, numbers []string) {
		t.Helper()

		t.assertStaticSeedsCount(7, asStringFuncName)

		t.Run(name, func(t *T) {
			fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6])
		})
	})
}

func AsString8(t *T, name string, fuzzFunc func(t *T, x1, x2, x3, x4, x5, x6, x7, x8 string)) {
	t.Helper()

	AsStringSlice(t, name, func(t *T, numbers []string) {
		t.Helper()

		t.assertStaticSeedsCount(8, asStringFuncName)

		t.Run(name, func(t *T) {
			fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7])
		})
	})
}

func AsString9(t *T, name string, fuzzFunc func(t *T, x1, x2, x3, x4, x5, x6, x7, x8, x9 string)) {
	t.Helper()

	AsStringSlice(t, name, func(t *T, numbers []string) {
		t.Helper()

		t.assertStaticSeedsCount(9, asStringFuncName)

		t.Run(name, func(t *T) {
			fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8])
		})
	})
}

func AsString10(t *T, name string, fuzzFunc func(t *T, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10 string)) {
	t.Helper()

	AsStringSlice(t, name, func(t *T, numbers []string) {
		t.Helper()

		t.assertStaticSeedsCount(10, asStringFuncName)

		t.Run(name, func(t *T) {
			fuzzFunc(t, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
		})
	})
}
