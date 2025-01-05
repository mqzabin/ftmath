package fuzzdecimal

import (
	"reflect"
	"testing"
)

var (
	rtBool     = reflect.TypeOf(true)
	rtUint     = reflect.TypeOf(uint64(1))
	rtTestingT = reflect.TypeOf(&testing.T{})
)

func createSeedFunc(f *testing.F, cfg config, seedHandler func(*testing.T, []seed)) reflect.Value {
	f.Helper()

	signature := parseSeedFuncType(f, cfg)

	return reflect.MakeFunc(signature, func(args []reflect.Value) []reflect.Value {
		testingT, ok := args[0].Interface().(*testing.T)
		if !ok {
			f.Errorf("first argument must be *testing.TestingT")
		}

		seeds := make([]seed, 0, len(cfg.decimals))

		argIndex := 1 // Skipping zero, since *testing.T is the first argument.

		for _, cfg := range cfg.decimals {
			// Parsing the number's sign.
			var isNegative bool

			if cfg.signed {
				neg, ok := args[argIndex].Interface().(bool)
				if !ok {
					f.Errorf("expected bool at index %d", argIndex)
				}

				isNegative = neg
				argIndex++
			}

			// Parsing the number's uints.
			uintsCount := uintsPerNumber(cfg.maxSignificantDigits)
			uints := make([]uint64, 0, uintsCount)

			for i := range uintsCount {
				uintN, ok := args[argIndex].Interface().(uint64)
				if !ok {
					f.Errorf("expected uint64 at index %d", argIndex)
				}

				switch i {
				case 0:
					uints = append(uints, normalizeUint(uintN, cfg.maxSignificantDigits%maxDigitsPerUint))
				default:
					uints = append(uints, normalizeUint(uintN, maxDigitsPerUint))
				}

				argIndex++
			}

			seeds = append(seeds, seed{uints: uints, neg: isNegative})
		}

		// Calling the seed function to transform seed into the desired type.
		seedHandler(testingT, seeds)

		return nil
	})
}

func parseSeedFuncType(f *testing.F, cfg config) reflect.Type {
	f.Helper()

	// Starting at 1 to cover the *testing.T argument.
	paramsLen := 1

	for _, cfg := range cfg.decimals {
		paramsLen += uintsPerNumber(cfg.maxSignificantDigits)

		if cfg.signed {
			paramsLen++
		}
	}

	params := make([]reflect.Type, 0, paramsLen)
	params = append(params, rtTestingT)

	for _, cfg := range cfg.decimals {
		if cfg.signed {
			params = append(params, rtBool)
		}

		for range uintsPerNumber(cfg.maxSignificantDigits) {
			params = append(params, rtUint)
		}
	}

	return reflect.FuncOf(params, []reflect.Type{}, false)
}
