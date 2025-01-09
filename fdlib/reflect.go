package fdlib

import (
	"reflect"
	"testing"
)

var (
	// RTBool is the reflect.Type of bool.
	RTBool = reflect.TypeOf(true)
	// RTUint is the reflect.Type of uint64.
	RTUint = reflect.TypeOf(uint64(1))
	// RTTestingT is the reflect.Type of *testing.T.
	RTTestingT = reflect.TypeOf(&testing.T{})
)

// CreateSeedFunc creates a reflect.Value of a function that takes a *testing.T and a series of Seeds, according to the provided Config.
func CreateSeedFunc(f *testing.F, cfg Config, seedHandler func(*testing.T, []Seed)) reflect.Value {
	f.Helper()

	signature := ParseSeedFuncType(f, cfg)

	return reflect.MakeFunc(signature, func(args []reflect.Value) []reflect.Value {
		testingT, ok := args[0].Interface().(*testing.T)
		if !ok {
			f.Errorf("first argument must be *testing.TestingT")
		}

		seeds := make([]Seed, 0, len(cfg.Decimals))

		argIndex := 1 // Skipping zero, since *testing.T is the first argument.

		for _, decCfg := range cfg.Decimals {
			// Parsing the number's sign.
			var isNegative bool

			if decCfg.Signed {
				neg, ok := args[argIndex].Interface().(bool)
				if !ok {
					f.Errorf("expected bool at index %d", argIndex)
				}

				isNegative = neg
				argIndex++
			}

			// Parsing the number's uints.
			uintsCount := UintsPerNumber(decCfg.MaxSignificantDigits)
			uints := make([]uint64, 0, uintsCount)

			for i := range uintsCount {
				uintN, ok := args[argIndex].Interface().(uint64)
				if !ok {
					f.Errorf("expected uint64 at index %d", argIndex)
				}

				switch i {
				case 0:
					uints = append(uints, NormalizeUint(testingT, uintN, decCfg.MaxSignificantDigits%MaxDigitsPerUint))
				default:
					uints = append(uints, NormalizeUint(testingT, uintN, MaxDigitsPerUint))
				}

				argIndex++
			}

			seeds = append(seeds, Seed{Uints: uints, Neg: isNegative})
		}

		// Calling the seed function to transform seed into the desired type.
		seedHandler(testingT, seeds)

		return nil
	})
}

// ParseSeedFuncType creates a reflect.Type of a function that takes a *testing.T and a series of arguments according to the provided Config.
func ParseSeedFuncType(f *testing.F, cfg Config) reflect.Type {
	f.Helper()

	// Starting at 1 to cover the *testing.T argument.
	paramsLen := 1

	for _, decCfg := range cfg.Decimals {
		paramsLen += UintsPerNumber(decCfg.MaxSignificantDigits)

		if decCfg.Signed {
			paramsLen++
		}
	}

	params := make([]reflect.Type, 0, paramsLen)
	params = append(params, RTTestingT)

	for _, decCfg := range cfg.Decimals {
		if decCfg.Signed {
			params = append(params, RTBool)
		}

		for range UintsPerNumber(decCfg.MaxSignificantDigits) {
			params = append(params, RTUint)
		}
	}

	return reflect.FuncOf(params, []reflect.Type{}, false)
}
