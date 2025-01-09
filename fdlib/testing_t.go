package fdlib

type TestHelper interface {
	Fatalf(format string, args ...interface{})
}
