package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	//  net = network
	// ring ring ring
	nl, err := net.Listen("tcp", ":8888") //ip:port
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //0,Cause = 1
	}

	//call receive == green button press
	conn, err := nl.Accept() // Layer-5 session layer
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	remoteAddr := conn.RemoteAddr().String() //Caller
	fmt.Println(remoteAddr)                  //[::1]:1731
	//localhost
	//[::1]
	//127.0.0.1

	bs := make([]byte, 1024) //1500 bytes
	conn.Read(bs)            //listen

	//90 10 20 30 50
	//H e l l o
	fmt.Println(string(bs)) // Human readable formated message

	//byte
	conn.Write([]byte("It's working !!! Tanvir You can do it."))

	conn.Close() // red button press
	nl.Close()   // red button press

}
