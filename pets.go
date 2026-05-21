// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func greetPet(name, pet string) string {
	return fmt.Sprintf("%s's favourite pet is a %s.", name, pet)
}
