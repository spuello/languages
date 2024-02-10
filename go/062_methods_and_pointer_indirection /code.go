package main

// 1. Methods with pointer receivers take either a value or a pointer as the receiver when they are called

// 2. Functions that take a value argument must take a value of that specific type
//while methods with value receivers take either a value or a pointer as the receiver when they are called

// CHOOSING A VALUE OR A POINTER RECEIVER
// ======================================

//There are two reasons to use a pointer receiver.

// 1. Method can modify the value that its receiver points to.
// 2. Avoid copying the value on each method call.
// This can be more efficient if the receiver is a large struct.

func main() {

}
