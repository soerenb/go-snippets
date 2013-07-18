/*
 * server.go - Project to demonstrate IPC via Linux sockets, server part
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
	"fmt"
	"net"
	"sync"
)

func echo_srv(c net.Conn, wg *sync.WaitGroup) {
	defer c.Close()
	defer wg.Done()

	for {
		var msg = make([]byte, 1024)

		n, err := c.Read(msg)
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Printf("SERVER: received %v bytes\n", n)

		n, err = c.Write(msg[:n])
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Printf("SERVER: sent %v bytes\n", n)
	}
}

func main() {
	var wg sync.WaitGroup

	ln, err := net.Listen("unix", "./sock_srv")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		fmt.Print(err)
		return
	}
	wg.Add(1)
	go echo_srv(conn, &wg)

	wg.Wait()
}
