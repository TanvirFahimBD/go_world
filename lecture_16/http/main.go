package main

func main() {

	// 	nl, err := net.Listen("tcp", ":8888") // Port Limit Range - (0-65535)/Usable(1-65535)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		os.Exit(1) //1 = Stop with error
	// 	}

	// 	for {
	// 		conn, err := nl.Accept()
	// 		if err != nil {
	// 			fmt.Println(err.Error())
	// 			//Continue
	// 		}

	// 		// scanner := bufio.NewScanner(conn)
	// 		// for scanner.Scan(){

	// 		// line := scanner.Text()
	// 		// fmt.Println(line)
	// 		// if line == "" {
	// 		// 		fmt.Println("Line Ends")
	// 		// 		break
	// 		// }

	// 		// }

	// 		bs := make([]byte, 1024)

	// 		n, e := conn.Read(bs)
	// 		if e != nil {
	// 			fmt.Println(e.Error())
	// 		}

	// 		fmt.Println(n)
	// 		// fmt.Println(bs)
	// 		reqstr := string(bs) //Conversion
	// 		fmt.Println(reqstr)
	// 		// reqSlc := strings.Fields(reqstr)
	// 		// fmt.Println(reqSlc, len(reqSlc))

	// 		body := `<!DOCTYPE html><html><head><title>Page Title</title></head><body>
	// 		<h1>Want To Learn Code!</h1><p>Welcome To Coding Boot Camp</p></body></html>`

	// 		// fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	// 		// fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	// 		// fmt.Fprint(conn, "Content-Type: text/html\r\n")
	// 		// fmt.Fprint(conn, "\r\n")
	// 		// fmt.Fprint(conn, body)

	// 		resp := "HTTP/1.1 200 OK\r\n" +
	// 			"Content-Length: %d\r\n" +
	// 			"\r\n%s"

	// 		msg := fmt.Sprintf(resp, len(body), body)
	// 		fmt.Println(msg)
	// 		conn.Write([]byte(msg))

	// 		//conn.Close()
	// 	}

	// 	//https://developer.mozilla.org/en-US/docs/Web/HTTP/Messages
}
