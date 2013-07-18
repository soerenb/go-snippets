/*
 * client.go - Project to demonstrate IPC via Linux sockets, client part
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

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	conn, err := net.Dial("unix", "./sock_srv")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer conn.Close()

	for {
		fmt.Print("Enter message to transmit: ")
		msg, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Print(err)
			return
		}

		msg = msg[:len(msg)-1]
		if (strings.ToLower(msg) == "quit") || (strings.ToLower(msg) == "exit") {
			fmt.Println("bye")
			return
		}

		n, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Printf("CLIENT: sent %v bytes\n", n)

		n, err = conn.Read([]byte(msg))
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Printf("CLIENT: received %v bytes\n", n)

		fmt.Println("Received message:", msg)
	}
}
