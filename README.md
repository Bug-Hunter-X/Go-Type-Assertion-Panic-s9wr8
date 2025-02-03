# Go Type Assertion Panic
This repository demonstrates a common error in Go programming: panics caused by incorrect type assertions.

The `bug.go` file contains code that attempts to assert an interface{} value as an int without checking the type beforehand. This leads to a runtime panic if the underlying type is not an int.

The `bugSolution.go` file provides a corrected version with proper type checking to avoid the panic.

## How to reproduce the bug
1. Clone this repository.
2. Navigate to the repository directory.
3. Run `go run bug.go`.
4. Observe the panic.

## How to fix the bug
Refer to the `bugSolution.go` file for the corrected code, which demonstrates safe type assertion using a type switch.