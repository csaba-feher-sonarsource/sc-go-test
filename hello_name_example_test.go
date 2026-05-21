// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func Example_namedGreeting() {
	greeting := "Hello"
	name := "Alice"
	fmt.Printf("%s, %s!\n", greeting, name)
	// Output: Hello, Alice!
}

func Example_namedGreetingReversed() {
	greeting := "Hello"
	name := "Alice"
	fmt.Printf("%s, %s!\n", reverse.String(greeting), reverse.String(name))
	// Output: olleH, ecilA!
}
