/*
 How do I know whether a variable is allocated on the heap or the stack?

From a correctness standpoint, you don't need to know. Each variable in Go exists as long as there are references to it.
The storage location chosen by the implementation is irrelevant to the semantics of the language.

The storage location does have an effect on writing efficient programs. When possible,
the Go compilers will allocate variables that are local to a function in that function's stack frame.
However, if the compiler cannot prove that the variable is not referenced after the function returns,
then the compiler must allocate the variable on the garbage-collected heap to avoid dangling pointer errors.
Also, if a local variable is very large, it might make more sense to store it on the heap rather than the stack.

In the current compilers, if a variable has its address taken, that variable is a candidate for allocation on the heap.
However, a basic escape analysis recognizes some cases when such variables will not live past the return from the function and can reside on the stack.

*/

package main

import "fmt"

func testA() *int {
	A := 5
	var testlocal int
	var test int
	test = 5
	testlocal = 7
	fmt.Println("2&A:", &A)
	fmt.Println("3&testlocal:", &testlocal)
	fmt.Println("4&test:", &test)
	return &test
}

func main() {
	B := 6
	var test *int
	fmt.Println("1&B:", &B)

	fmt.Println("1&test:", test)
	test = testA()
	fmt.Println("5&test:", test)
	fmt.Println(*test)
}
