package fuzzdecimal

import (
	"reflect"
	"testing"
)

func Test_parseSeedFuncType(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		config config
		want   []reflect.Type
	}{
		{
			name: "no decimals",
			config: config{
				decimals: []decimalConfig{},
			},
			want: []reflect.Type{
				rtTestingT,
			},
		},
		{
			name: "single decimal with no digits",
			config: config{
				decimals: []decimalConfig{
					{
						maxSignificantDigits: 0,
						signed:               false,
						decimalPointPosition: 0,
					},
				},
			},
			want: []reflect.Type{
				rtTestingT,
			},
		},
		{
			name: "two unsigned decimals with same uints amount",
			config: config{
				decimals: []decimalConfig{
					{
						maxSignificantDigits: 30, // 2 uints
						signed:               false,
						decimalPointPosition: 0, // irrelevant
					},
					{
						maxSignificantDigits: 30, // 2 uints
						signed:               false,
						decimalPointPosition: 0, // irrelevant
					},
				},
			},
			want: []reflect.Type{
				rtTestingT,
				rtUint,
				rtUint,
				rtUint,
				rtUint,
			},
		},
		{
			name: "two signed decimals with same uints amount",
			config: config{
				decimals: []decimalConfig{
					{
						maxSignificantDigits: 30, // 2 uints
						signed:               true,
						decimalPointPosition: 0, // irrelevant
					},
					{
						maxSignificantDigits: 30, // 2 uints
						signed:               true,
						decimalPointPosition: 0, // irrelevant
					},
				},
			},
			want: []reflect.Type{
				rtTestingT,
				// signed with 2 uints
				rtBool,
				rtUint,
				rtUint,
				// signed with 2 uints
				rtBool,
				rtUint,
				rtUint,
			},
		},
		{
			name: "three decimals with varying uints amount and signed same uints amount",
			config: config{
				decimals: []decimalConfig{
					{
						maxSignificantDigits: 10, // 1 uints
						signed:               true,
						decimalPointPosition: 0, // irrelevant
					},
					{
						maxSignificantDigits: 2 * maxDigitsPerUint, // 2 uints
						signed:               false,
						decimalPointPosition: 0, // irrelevant
					},
					{
						maxSignificantDigits: 50, // 3 uints
						signed:               true,
						decimalPointPosition: 0, // irrelevant
					},
				},
			},
			want: []reflect.Type{
				rtTestingT,
				// signed with 1 uints
				rtBool,
				rtUint,
				// unsigned with 2 uints
				rtUint,
				rtUint,
				// signed with 3 uints
				rtBool,
				rtUint,
				rtUint,
				rtUint,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			f := &testing.F{}

			got := parseSeedFuncType(f, tc.config)

			if got.Kind() != reflect.Func {
				t.Fatalf("parseSeedFuncType().Kind() = %s, want %s", got.Kind().String(), reflect.Func.String())
			}

			if got.NumOut() != 0 {
				t.Errorf("parseSeedFuncType().NumOut() = %d, want 0", got.NumOut())
			}

			if got.NumIn() != len(tc.want) {
				t.Fatalf("parseSeedFuncType().NumIn() = %d, want %d", got.NumIn(), len(tc.want))
			}

			for i, want := range tc.want {
				if got.In(i).String() != want.String() {
					t.Errorf("parseSeedFuncType().In(%d) = %s, want %s", i, got.In(i).String(), want.String())
				}
			}
		})
	}
}
