package fdlib

import (
	"reflect"
	"testing"
)

func TestCreateSeedFunc(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		configFunc func(t *testing.T, f *testing.F) Config
		callFunc   func(t *testing.T, anyFunc any)
		handler    func(t *testing.T, seeds []Seed)
	}{
		{
			name: "two decimals",
			configFunc: func(t *testing.T, f *testing.F) Config {
				return Config{
					Decimals: []DecimalConfig{
						{
							MaxSignificantDigits: 10,
							Signed:               true,
							MaxDecimalPlaces:     0,
						},
						{
							MaxSignificantDigits: 10,
							Signed:               false,
							MaxDecimalPlaces:     0,
						},
					},
				}
			},
			callFunc: func(t *testing.T, anyFunc any) {
				f, ok := anyFunc.(func(*testing.T, bool, uint64, uint64))
				if !ok {
					t.Fatalf("anyFunc is not a function with the expected signature")
				}

				f(t, true, 123, 42)
			},
			handler: func(t *testing.T, seeds []Seed) {
				if len(seeds) != 2 {
					t.Fatalf("len(seeds) = %d, want 2", len(seeds))
				}

				if seeds[0].Neg != true {
					t.Errorf("seeds[0].Neg = %t, want true", seeds[0].Neg)
				}

				if len(seeds[0].Uints) != 1 {
					t.Fatalf("len(seeds[0].Uints) = %d, want 1", len(seeds[0].Uints))
				}

				if seeds[0].Uints[0] != 123 {
					t.Errorf("seeds[0].Uints[0] = %d, want 123", seeds[0].Uints[0])
				}

				if seeds[1].Neg != false {
					t.Errorf("seeds[1].Neg = %t, want true", seeds[1].Neg)
				}

				if len(seeds[1].Uints) != 1 {
					t.Fatalf("len(seeds[1].Uints) = %d, want 1", len(seeds[1].Uints))
				}

				if seeds[1].Uints[0] != 42 {
					t.Errorf("seeds[1].Uints[0] = %d, want 42", seeds[1].Uints[0])
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			f := &testing.F{}

			anyFunc := CreateSeedFunc(f, tc.configFunc(t, f), tc.handler)

			tc.callFunc(t, anyFunc)
		})
	}
}

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
						MaxDecimalPlaces:     0,
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
						MaxDecimalPlaces:     0, // irrelevant
					},
					{
						MaxSignificantDigits: 30, // 2 uints
						Signed:               false,
						MaxDecimalPlaces:     0, // irrelevant
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
						MaxDecimalPlaces:     0, // irrelevant
					},
					{
						MaxSignificantDigits: 30, // 2 uints
						Signed:               true,
						MaxDecimalPlaces:     0, // irrelevant
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
						MaxDecimalPlaces:     0, // irrelevant
					},
					{
						MaxSignificantDigits: 2 * MaxDigitsPerUint, // 2 uints
						Signed:               false,
						MaxDecimalPlaces:     0, // irrelevant
					},
					{
						MaxSignificantDigits: 50, // 3 uints
						Signed:               true,
						MaxDecimalPlaces:     0, // irrelevant
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
