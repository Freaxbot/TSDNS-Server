package main
/*
* Developed and published by Freaxbot <support@magiccoder.de>
* This program is free software. It comes without any warranty, to
* the extent permitted by applicable law. You can redistribute it
* and/or modify it under the terms of the Do What The Fuck You Want
* To Public License, Version 2, as published by Sam Hocevar. See
* http://www.wtfpl.net/ for more details.
*/

import (
	"net"
	"fmt"
	"os"
	"strconv"
	"log"
	"gopkg.in/ini.v1"
)

func StartTCP(port int)  {
	l, err := net.Listen("tcp","localhost:"+ strconv.Itoa(port));
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on port ::" + strconv.Itoa(port))
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 512)
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Println("Fail to read file: %v", err)
		os.Exit(1)
	}
	rt := cfg.Section("routing").Key(string(buf[:reqLen])).String()
	if (rt == ""){
		conn.Write([]byte("404"));
	} else if (rt == "NORESPONSE"){
		conn.Close()
	}else{
		conn.Write([]byte(rt));
	}
	log.Println("LOOKUP :: " +  string(buf[:reqLen]) + " == " + string(rt));

	conn.Close()
}