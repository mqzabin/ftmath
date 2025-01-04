# fuzzdecimal

This package provides an easy way to perform fuzzy tests with arbitrary precision decimals.

The `fuzzdecimal` package provides three diferrente "fuzzers":
- `StringFuzzer`: More flexible. Your test will receive N decimal strings and your test function have to parse them into your desired decimal type, and make comparisons.
- `GenericFuzzer`: Easier to use. You have to provide your decimal type string parse function, and your test will receive N decimals (in your chosen decimal type), and you should only implement the comparisons.
- `GenericComparisonFuzzer`: Is an addition to the `GenericFuzzer`. Could be used when you need to compare your decimal type with another reference decimal type operation result. You have to provide a function to parse yours, and the reference decimal type from a string. Then, the `FuzzN` method will receive a reference result function that should return the operation result, and your test function will receive the reference result as a parameter. 

All of them have the same methods in the `FuzzN()` name format, where N is the number of decimal arguments that your tested function needs, e.g. `Fuzz1` for unary operations, `Fuzz2` for binary operations, and so on.

You can define via options, the maximum number of digits for the fuzzed decimals, and if they can be signed (positive/negative) or unsigned.

## Usage

Install as dependency into your project:

```bash
$ go get github.com/mqzabin/fuzzdecimal
```

Then the used will vary according to the fuzzer you choose.

### StringFuzzer

```go
package test

import (
	"testing"
	
	"github.com/shopspring/decimal"
	"github.com/mqzabin/fuzzdecimal"
)

func FuzzAdd(f *testing.F) {
	fuzzer := fuzzdecimal.NewStringFuzzer(f, fuzzdecimal.WithSignedMaxDigits(30))
	
	fuzzer.Fuzz2(func(t *testing.T, x1, x2 string) {
		a, err := decimal.NewFromString(x1)
		if err != nil {
			t.Errorf("failed to parse x1: %v", err)
		}

		b, err := decimal.NewFromString(x1)
		if err != nil {
			t.Errorf("failed to parse x1: %v", err)
		}
        
		if resAB, resBA := a.Add(b), b.Add(a); !resAB.Equal(resBA) {
			t.Errorf("a + b != b + a, where a='%s', b='%s', a+b='%s' and b+a='%s'", a.String(), b.String(), resAB.String(), resBA.String())
		}
	})
}
```

### GenericFuzzer

```go
package test

import (
	"testing"
	
	"github.com/shopspring/decimal"
	"github.com/mqzabin/fuzzdecimal"
)

func FuzzAdd(f *testing.F) {
	parseFunc := func(f *testing.F, s string) (decimal.Decimal, error) {
		f.Helper()
		
        return decimal.NewFromString(s)
    }
	
	fuzzer := fuzzdecimal.NewGenericFuzzer(f, parseFunc, fuzzdecimal.WithSignedMaxDigits(30))
	
	fuzzer.Fuzz2(func(t *testing.T, x1, x2 decimal.Decimal) {
		if res12, res21 := x1.Add(x2), x2.Add(x1); !res12.Equal(res21) {
			t.Errorf("x1 + x2 != x2 + x1, where x1='%s', x2='%s', x1+x2='%s' and x2+x1='%s'", x1.String(), x2.String(), res12.String(), res21.String())
		}
	})
}
```

### GenericComparisonFuzzer

TODO
