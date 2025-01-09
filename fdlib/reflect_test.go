package fdlib

import (
	"reflect"
	"testing"
)

func TestParseSeedFuncType(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		config Config
		want   []reflect.Type
	}{
		{
			name: "no decimals",
			config: Config{
				Decimals: []DecimalConfig{},
			},
			want: []reflect.Type{
				RTTestingT,
			},
		},
		{
			name: "single decimal with no digits",
			config: Config{
				Decimals: []DecimalConfig{
					{
						MaxSignificantDigits: 0,
						Signed:               false,
						DecimalPointPosition: 0,
					},
				},
			},
			want: []reflect.Type{
				RTTestingT,
			},
		},
		{
			name: "two unsigned decimals with same uints amount",
			config: Config{
				Decimals: []DecimalConfig{
					{
						MaxSignificantDigits: 30, // 2 uints
						Signed:               false,
						DecimalPointPosition: 0, // irrelevant
					},
					{
						MaxSignificantDigits: 30, // 2 uints
						Signed:               false,
						DecimalPointPosition: 0, // irrelevant
					},
				},
			},
			want: []reflect.Type{
				RTTestingT,
				RTUint,
				RTUint,
				RTUint,
				RTUint,
			},
		},
		{
			name: "two signed decimals with same uints amount",
			config: Config{
				Decimals: []DecimalConfig{
					{
						MaxSignificantDigits: 30, // 2 uints
						Signed:               true,
						DecimalPointPosition: 0, // irrelevant
					},
					{
						MaxSignificantDigits: 30, // 2 uints
						Signed:               true,
						DecimalPointPosition: 0, // irrelevant
					},
				},
			},
			want: []reflect.Type{
				RTTestingT,
				// signed with 2 uints
				RTBool,
				RTUint,
				RTUint,
				// signed with 2 uints
				RTBool,
				RTUint,
				RTUint,
			},
		},
		{
			name: "three decimals with varying uints amount and signed same uints amount",
			config: Config{
				Decimals: []DecimalConfig{
					{
						MaxSignificantDigits: 10, // 1 uints
						Signed:               true,
						DecimalPointPosition: 0, // irrelevant
					},
					{
						MaxSignificantDigits: 2 * MaxDigitsPerUint, // 2 uints
						Signed:               false,
						DecimalPointPosition: 0, // irrelevant
					},
					{
						MaxSignificantDigits: 50, // 3 uints
						Signed:               true,
						DecimalPointPosition: 0, // irrelevant
					},
				},
			},
			want: []reflect.Type{
				RTTestingT,
				// signed with 1 uints
				RTBool,
				RTUint,
				// unsigned with 2 uints
				RTUint,
				RTUint,
				// signed with 3 uints
				RTBool,
				RTUint,
				RTUint,
				RTUint,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			f := &testing.F{}

			got := ParseSeedFuncType(f, tc.config)

			if got.Kind() != reflect.Func {
				t.Fatalf("ParseSeedFuncType().Kind() = %s, want %s", got.Kind().String(), reflect.Func.String())
			}

			if got.NumOut() != 0 {
				t.Errorf("ParseSeedFuncType().NumOut() = %d, want 0", got.NumOut())
			}

			if got.NumIn() != len(tc.want) {
				t.Fatalf("ParseSeedFuncType().NumIn() = %d, want %d", got.NumIn(), len(tc.want))
			}

			for i, want := range tc.want {
				if got.In(i).String() != want.String() {
					t.Errorf("ParseSeedFuncType().In(%d) = %s, want %s", i, got.In(i).String(), want.String())
				}
			}
		})
	}
}
