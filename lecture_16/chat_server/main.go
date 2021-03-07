package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	nl, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer nl.Close() // await - run at last after executing all others function
	log.Printf("server started on :8888")

	for {
		conn, err := nl.Accept()
		if err != nil {
			fmt.Println(err.Error())
			//Continue
		}

		bs := make([]byte, 1024)
		n, e := conn.Read(bs)
		if e != nil {
			fmt.Println(e.Error())
		}

		//0-5
		//Hello
		//01234
		//bs[0:5]

		// fmt.Println(n)
		reqstr := string(bs[:n]) //Conversion
		fmt.Println(reqstr)
		recvTime := time.Now().Format("2006-01-02 15:04:05")
		msg := fmt.Sprintf(`Your message: %s, revived at %s`, reqstr, recvTime)
		conn.Write([]byte(msg))

	}

	//https://developer.mozilla.org/en-US/docs/Web/HTTP/Messages
}
