# fuzzdecimal - Fuzzy tests for arbitrary precision decimals

Go 1.18 introduced [fuzzy testing](https://go.dev/doc/security/fuzz/), a way to test your code with "random" inputs. This is a great way to find bugs in your code.

Writing tests for math-related packages is pretty hard, since there are infinite possible inputs.
In this sense, fuzzy testing could be very nice to ensure that your code edge cases are covered.

The fuzzy test API only supports primitive and relatively static types. So, if your math-related package uses primitive types, you could use the fuzzy API directly.

However, if your package uses arbitrary precision decimals, you will have to write a lot of glue code to convert the primitive type inputs into your desired decimal type.

To solve this problem, this package provides an easy way to perform fuzzy tests with arbitrary precision decimals.

The `fuzzdecimal` package provides three different ways to fuzzy yours decimals:
- `AsString`: More flexible. Your test will receive N decimal strings and your test function have to parse them into your desired decimal type, and make comparisons.
- `AsDecimal`: Easier to use. You have to provide your decimal type string parse function, and your test will receive N decimals (in your chosen decimal type), and you should only implement the comparisons.
- `AsDecimalComparison`: Is an addition to the `AsDecimal`. Could be used when you need to compare your decimal type with another reference decimal type operation result. Check examples. 

All this fuzzy functions should be called inside the `Fuzz()` call.

You can define the maximum number of digits for the fuzzed decimals, and if they can be signed (positive/negative) or unsigned via options.

This package will work for any type of decimal that could be parsed from a string.

## Usage

Install as dependency into your project:

```bash
$ go get github.com/mqzabin/fuzzdecimal
```

Then the used will vary according to the fuzzer you choose.

### Fuzz `AsString`

This fuzzer is more flexible, but you have to parse each decimal string into your desired decimal type.

```go
package app_test

import (
	"testing"
	
	"github.com/shopspring/decimal"
	"github.com/mqzabin/fuzzdecimal"
)

func FuzzCommutativeAdd(f *testing.F) {
	// Fuzzing the commutative property of two decimals addition.
	// The number of generated seeds is defined by the second parameter (2).
	fuzzdecimal.Fuzz(f, 2, func(t *fuzzdecimal.T) {
		// AsString2 is called to generate two decimal strings.
		// If the number of parameters change to 3, AsString3 should be called instead.
		// Also, AsStringSlice could be called to receive a slice of decimal strings.
        fuzzdecimal.AsString2(t, "Commutative Add", func(x1, x2 string) {
			// Parsing the first string.
			a, err := decimal.NewFromString(x1)
			if err != nil {
				t.Errorf("failed to parse x1: %v", err)
			}

			// Parsing the second string.
			b, err := decimal.NewFromString(x1)
			if err != nil {
				t.Errorf("failed to parse x1: %v", err)
			}

			// Making the comparison.
			if resAB, resBA := a.Add(b), b.Add(a); !resAB.Equal(resBA) {
				t.Errorf("a + b != b + a, where a='%s', b='%s', a+b='%s' and b+a='%s'", a.String(), b.String(), resAB.String(), resBA.String())
			}	
		})
    },fuzzdecimal.WithAllDecimals(
        fuzzdecimal.WithSigned(),
        fuzzdecimal.WithMaxSignificantDigits(30),
        fuzzdecimal.WithDecimalPointAt(15),
    ))
}
```

### Fuzz `AsDecimal`

This fuzzer is easier than `AsString`, because you only have to implement the desired comparisons. 

```go
package app_test

import (
	"testing"
	
	"github.com/shopspring/decimal"
	"github.com/mqzabin/fuzzdecimal"
)

func FuzzCommutativeAdd(f *testing.F) {
	// Defining how to parse the decimal from a numeric string.
	parseFunc := func(t *testing.T, s string) (decimal.Decimal, error) {
		t.Helper()

		return decimal.NewFromString(s)
	}

	// Fuzzing the commutative property of two decimals addition.
	// The number of generated seeds is defined by the second parameter (2).
	fuzzdecimal.Fuzz(f, 2, func(t *fuzzdecimal.T) {
		// AsDecimal2 is called to generate two decimal from your desired type (defined by parseFunc).
		// If the number of parameters change to 3, AsDecimal3 should be called instead.
		// Also, AsDecimalSlice could be called to receive a slice of decimal numbers.
		fuzzdecimal.AsDecimal2(t, "Commutative Add", parseFunc, func(x1, x2 decimal.Decimal) {
			if res12, res21 := x1.Add(x2), x2.Add(x1); !res12.Equal(res21) {
				t.Errorf("x1 + x2 != x2 + x1, where x1='%s', x2='%s', x1+x2='%s' and x2+x1='%s'", x1.String(), x2.String(), res12.String(), res21.String())
			}
		})
	},fuzzdecimal.WithAllDecimals(
		fuzzdecimal.WithSigned(),
		fuzzdecimal.WithMaxSignificantDigits(30),
		fuzzdecimal.WithDecimalPointAt(15),
	))
}
```

### Fuzz `AsDecimalComparison`

This fuzzy function should be called when you want to compare an operation result from your package, with another reference decimal type operation result.

It needs a little bit more setup than `AsDecimal`, but the `AsDecimalComparison` call will be clean.

```go
package app_test

import (
	"testing"

	shopspring "github.com/shopspring/decimal"
	"github.com/mqzabin/somedecimal"
	"github.com/mqzabin/fuzzdecimal"
)

func FuzzCommutativeAdd(f *testing.F) {
	// Defining how to parse your decimal number.
	parseMyDecimal := func(t *fuzzdecimal.T, s string) (somedecimal.Decimal, error) {
		t.Helper()

		return somedecimal.NewFromString(s)
	}

	// Defining how to parse the reference decimal.
	// If both decimals are the same type, you can use the same function (parseMyDecimal) for both `AsDecimalComparison2` parameters. 
	parseReferenceDecimal := func(t *testing.T, s string) (shopspring.Decimal, error) {
		t.Helper()

		return shopspring.NewFromString(s)
	}

	// Fuzzing the commutative property of two decimals addition.
	// The number of generated seeds is defined by the second parameter (2).
	fuzzdecimal.Fuzz(f, 2, func(t *fuzzdecimal.T) {
		// AsDecimalComparison3 is called to generate two decimal from your desired type (defined by parseFunc).
		// If the number of parameters change to 3, AsDecimalComparison3 should be called instead.
		// Also, AsDecimalSlice could be called to receive a slice of decimal numbers.
		fuzzdecimal.AsDecimalComparison2(t, "Add", parseMyDecimal, parseReferenceDecimal,
			// The first function defines how to compute the reference result.
			// It will be compared (by equality) with the result of the second function (your package/decimal operation).
			func(t *fuzzdecimal.T, x1, x2 shopspring.Decimal) (string, error) {
				return x1.Add(x2).String(), nil
			},
			// The second functions defines how your package computes the result with your chosen Decimal type.
			func(t *fuzzdecimal.T,x1, x2 somedecimal.Decimal) string {
				return x1.Add(x2).String()
			},
		)
	}, fuzzdecimal.WithAllDecimals(
		fuzzdecimal.WithSigned(),
		fuzzdecimal.WithMaxSignificantDigits(30),
		fuzzdecimal.WithDecimalPointAt(15),
	))
}
```
