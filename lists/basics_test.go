package lists

import (
	"fmt"
)

func ExampleNew() {
	l := New()

	fmt.Println(l.IsEmpty())
	fmt.Println(l)

	// Output:
	// true
	// []
}

func ExampleList_Append() {
	l := New()

	l.Append(42)
	l.Append(23)

	fmt.Println(l.IsEmpty())
	fmt.Println(l)

	// Output:
	// false
	// [42 -> 23]
}

func ExampleList_Length() {
	l := New()

	fmt.Println(l.Length())

	l.Append(42)
	fmt.Println(l.Length())

	l.Append(23)
	fmt.Println(l.Length())

	// Output:
	// 0
	// 1
	// 2
}

func ExampleList_Contains() {
	l := New()

	fmt.Println(l.Contains(42))

	l.Append(42)
	fmt.Println(l.Contains(42))
	fmt.Println(l.Contains(23))

	l.Append(23)
	fmt.Println(l.Contains(42))
	fmt.Println(l.Contains(23))

	// Output:
	// false
	// true
	// false
	// true
	// true
}
