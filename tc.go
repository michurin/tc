// Package tc provides simple functions for type converting, casting, assigning with default and comparing.
package tc

// P simply returns pointer to any type T.
func P[T any](x T) *T {
	return &x
}

// Pn returns pointer to T or nil pointer for zero value of T.
func Pn[T comparable](x T) *T {
	if x == *new(T) {
		return nil
	}
	return &x
}

// SafePx denotes the pointer's underlying value or return zero values of T for nil pointers.
func SafePx[T any](x *T) T {
	if x == nil {
		return *new(T)
	}
	return *x
}

// DefPx denotes the pointer's underlying value or return def if pointer is nil.
func DefPx[T any](x *T, def T) T {
	if x == nil {
		return def
	}
	return *x
}

// Cmp is sort of ternary operator. Returns t or f as c is true or false.
func Cmp[T any](c bool, t T, f T) T {
	if c {
		return t
	}
	return f
}

// CmpN is sort of ternary operator, however it is considering pointer to bool, including nil value.
// Returns t or f as c is true or false. Or n if nil.
func CmpN[T any](c *bool, t T, f T, n T) T {
	if c == nil {
		return n
	}
	if *c {
		return t
	}
	return f
}

// SafeCast casts any to T or return zero value of T.
func SafeCast[T any](x any) T {
	a, _ := x.(T)
	return a
}

// DefCast casts any to T or return def.
func DefCast[T any](x any, def T) T {
	a, ok := x.(T)
	if ok {
		return a
	}
	return def
}

// Zero turns true if x has zero value.
func Zero[T comparable](x T) bool {
	return x == *new(T)
}

// DefZero returns def if argument has zero value.
func DefZero[T comparable](x T, def T) T {
	if x == *new(T) {
		return def
	}
	return x
}
