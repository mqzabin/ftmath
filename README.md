# fuzzymath

This package provides an easy way to perform math-related fuzzy tests.

## Usage

Install as dependency into your project:

```bash
$ go get github.com/mqzabin/fuzzymath
```

```go
package test

import (
	"testing"
)

func TestAdd(t *testing.T) {
    fuzzymath.NewFuzzer()
}
```

