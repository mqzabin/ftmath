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

func createSeedFunc(f *testing.F, signed bool, numbersCount, uintsPerNumber int, seedHandler func(*testing.T, []Seed)) reflect.Value {
	f.Helper()

	signature := seedFuncSignature(f, signed, numbersCount, uintsPerNumber)

	return reflect.MakeFunc(signature, func(args []reflect.Value) []reflect.Value {
		f.Helper()

		testingT, ok := args[0].Interface().(*testing.T)
		if !ok {
			f.Errorf("first argument must be *testing.T")
		}

		seeds := make([]Seed, 0, uintsPerNumber)

		argIndex := 1 // Skipping zero, since *testing.T is the first argument.
		for range numbersCount {
			// Parsing the number's sign.
			var neg bool
			if signed {
				negV, ok := args[argIndex].Interface().(bool)
				if !ok {
					f.Errorf("expected bool at index %d", argIndex)
				}

				neg = negV
				argIndex++
			}

			// Parsing the number's uints.
			uints := make([]uint64, 0, uintsPerNumber)
			for range uintsPerNumber {
				uintN, ok := args[argIndex].Interface().(uint64)
				if !ok {
					f.Errorf("expected uint64 at index %d", argIndex)
				}

				uints = append(uints, uintN)
				argIndex++
			}

			seeds = append(seeds, Seed{uints: uints, neg: neg})
		}

		// Calling the seed function to transform seed into the desired type.
		seedHandler(testingT, seeds)

		return nil
	})
}

func seedFuncSignature(f *testing.F, signed bool, numbersCount, uintsPerNumber int) reflect.Type {
	f.Helper()

	paramsLen := 1 + numbersCount*uintsPerNumber
	if signed {
		paramsLen += numbersCount
	}

	params := make([]reflect.Type, 0, paramsLen)
	params = append(params, rtTestingT)

	for range numbersCount {
		if signed {
			params = append(params, rtBool)
		}

		for range uintsPerNumber {
			params = append(params, rtUint)
		}
	}

	return reflect.FuncOf(params, []reflect.Type{}, false)
}
