/*
 * io.go - Reading from stdin
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
import "os"
import "bufio"

func main() {
	var s string
	var i int

	fmt.Println("Scan to string")
	n, err := fmt.Scanln(&s)
	fmt.Println("scanned items:", n, "err:", err, "s:", s)

	fmt.Println("Scan to int")
	n, err = fmt.Scan(&i)
	fmt.Println("scanned items:", n, "err:", err, "i:", i)

	/* read a whole line from stdin */
	stdin := bufio.NewReader(os.Stdin)
	fmt.Println("Read whole line from stdin")
	s, err = stdin.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("scanned items:", n, "err:", err, "msg:", s)
}
