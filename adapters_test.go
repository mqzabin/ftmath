package fuzzdecimal

import (
	"errors"
	"testing"

	shopspring "github.com/shopspring/decimal"
)

func Test_convertParseDecimalFunc(t *testing.T) {
	t.Parallel()

	testError := errors.New("test error")

	mustDecimal := func(t *testing.T, s string) shopspring.Decimal {
		t.Helper()

		dec, err := shopspring.NewFromString(s)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		return dec
	}

	testCases := []struct {
		name             string
		parseDecimalFunc func(t *T, s string) (shopspring.Decimal, error)
		decimalStr       string
		wantDecimal      shopspring.Decimal
		wantErr          error
	}{
		{
			name: "succeed to parse",
			parseDecimalFunc: func(t *T, s string) (shopspring.Decimal, error) {
				return shopspring.NewFromString(s)
			},
			decimalStr:  "0.001",
			wantDecimal: mustDecimal(t, "0.001"),
			wantErr:     nil,
		},
		{
			name: "fail to parse",
			parseDecimalFunc: func(t *T, s string) (shopspring.Decimal, error) {
				return shopspring.Decimal{}, testError
			},
			decimalStr:  "0.001",
			wantDecimal: shopspring.Decimal{},
			wantErr:     testError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			nT := &T{
				T:     t,
				seeds: nil,
			}

			newFunc := convertParseDecimalFunc(nT, tc.parseDecimalFunc)

			gotDec, err := newFunc(t, tc.decimalStr)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("unexpected error, want: %v, got: %v", tc.wantErr, err)
			}

			if !gotDec.Equal(tc.wantDecimal) {
				t.Fatalf("unexpected decimal, want: %s, got: %s", gotDec.String(), tc.wantDecimal.String())
			}
		})
	}
}
