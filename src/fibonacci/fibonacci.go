/*
 * fibonacci.go - Calculate fibonacci series
 * Copyright (C) 2013  Sören Brinkmann <soeren.brinkmann@gmail.com>
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 */

package main

import "fmt"
import "strconv"

var fib_cache = map[uint]uint{0: 0, 1: 1}

func fib(n uint) uint {
	if _, ok := fib_cache[n]; ok == false {
		fib_cache[n] = fib(n-2) + fib(n-1)
		if fib_cache[n] < fib_cache[n-1] {
			fmt.Printf("WARNING: Overflow detected at element #%v\n", n)
			fib_cache[n] = ^uint(0)
		}
	}

	return fib_cache[n]
}

func main() {
	var s string
	var limit uint = 42

	fmt.Printf("Enter the number of elements to calculate (%d): ", limit)
	fmt.Scanln(&s)

	tmp, err := strconv.ParseUint(s, 0, 0)
	if err == nil {
		limit = uint(tmp)
	}
	fmt.Printf("Calculating the first %v Fibonacci numbers\n", limit)

	fib(limit)
	for i := uint(0); i < limit; i++ {
		fmt.Printf("fib[%d] = %d\n", i, fib_cache[i])
	}
}
