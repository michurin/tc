package tc_test

import (
	. "github.com/davecgh/go-spew/spew"

	. "github.com/michurin/tc"
)

func ExampleP() {
	Config.DisablePointerAddresses = true
	Dump(P(1)) // (*int)(1) — just take pointer
	Dump(P(0)) // (*int)(0) — zero value as is
	// Output:
	// (*int)(1)
	// (*int)(0)
}

func ExamplePn() {
	Config.DisablePointerAddresses = true
	Dump(Pn(1)) // (*int)(1) — smart pointer with
	Dump(Pn(0)) // (*int)(<nil>) — zero values to nil pointer
	// Output:
	// (*int)(1)
	// (*int)(<nil>)
}

func ExampleSafePx() {
	Config.DisablePointerAddresses = true
	p := P(1)                 // pointer to 1
	Dump(SafePx(p))           // 1 — &1 -> 1
	Dump(SafePx((*int)(nil))) // 0 — &nil -> 0 (zero value)
	// Output:
	// (int) 1
	// (int) 0
}

func ExampleDefPx() {
	Config.DisablePointerAddresses = true
	p := P(1)                   // pointer to 1
	Dump(DefPx(p, 3))           // 1 — &1 -> 1
	Dump(DefPx((*int)(nil), 3)) // 3 — &nil -> def=3
	// Output:
	// (int) 1
	// (int) 3
}

func ExampleCmp() {
	Config.DisablePointerAddresses = true
	Dump(Cmp(true, "T", "F"))  // "T" — sort of ternary operator
	Dump(Cmp(false, "T", "F")) // "F"
	// Output:
	// (string) (len=1) "T"
	// (string) (len=1) "F"
}

func ExampleCmpN() {
	Config.DisablePointerAddresses = true
	t := P(true)                            // pointer to true
	f := P(false)                           // pointer to false
	Dump(CmpN(t, "T", "F", "N"))            // "T" — sort of ternary operator
	Dump(CmpN(f, "T", "F", "N"))            // "F"
	Dump(CmpN((*bool)(nil), "T", "F", "N")) // "N"
	// Output:
	// (string) (len=1) "T"
	// (string) (len=1) "F"
	// (string) (len=1) "N"
}

func ExampleSafeCast() {
	Config.DisablePointerAddresses = true
	Dump(SafeCast[string](any("X"))) // "X" — cast any to type T
	Dump(SafeCast[string](any(1)))   // "" — cast to zero value on error
	// Output:
	// (string) (len=1) "X"
	// (string) ""
}

func ExampleDefCast() {
	Config.DisablePointerAddresses = true
	Dump(DefCast(any("X"), "Y")) // "X" — cast any to type T
	Dump(DefCast(any(1), "Y"))   // "Y" — def
	Dump(DefCast(any(""), "Y"))  // "" — cast zero value to itself
	// Output:
	// (string) (len=1) "X"
	// (string) (len=1) "Y"
	// (string) ""
}

func ExampleDefZero() {
	Config.DisablePointerAddresses = true
	Dump(DefZero("X", "D")) // "X"
	Dump(DefZero("", "D"))  // "D"
	// Output:
	// (string) (len=1) "X"
	// (string) (len=1) "D"
}

func ExampleZero() {
	Config.DisablePointerAddresses = true
	Dump(Zero("X")) // false
	Dump(Zero(""))  // true — zero value
	// Output:
	// (bool) false
	// (bool) true
}
