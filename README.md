# fuzzdecimal

This package provides an easy way to perform fuzzy tests with arbitrary precision decimals.

## Usage

Install as dependency into your project:

```bash
$ go get github.com/mqzabin/fuzzdecimal
```

Example code:

```go
package test

import (
	"testing"
	
	"github.com/shopspring/decimal"
	"github.com/mqzabin/fuzzdecimal"
)

func FuzzOperations(f *testing.F) {
	fuzzer := fuzzdecimal.NewFuzzer(f, fuzzdecimal.WithSignedMaxDigits(30))
	
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

