package types

// A catch-all for numeric types that can be compared during sorting.
type Number interface {
	int | int64 | float64
}