package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:7777")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	//go func() {
	//	for {
	//		//result, err := ioutil.ReadAll(conn)
	//		b := make([]byte, 1024)
	//		n, _ := conn.Read(b)
	//		a := make([]byte, n)
	//		a = b[:n]
	//		fmt.Println(string(a))
	//	}
	//}()

	//for {
	//	input := bufio.NewScanner(os.Stdin)
	//	input.Scan()
	//	fmt.Println("你输入的是：", input.Text())
	//	_, err = conn.Write([]byte(input.Text()))
	//	checkError(err)
	//}

	msg1 := Msg{
		GroupId: "group1",
		Type:    "commit",
		Command: "create",
		TxCount: 2,
	}

	msg2 := Msg{
		GroupId: "group1",
		Type:    "rollback",
		Command: "add",
		TxCount: 2,
	}

	msg3 := Msg{
		GroupId: "group1",
		Type:    "commit",
		Command: "add",
		TxCount: 2,
		IsEnd:   true,
	}

	start := 2
	if start == 1 {
		bytes1, err := json.Marshal(&msg1)
		checkError(err)

		_, err = conn.Write(bytes1)
		checkError(err)

		bytes2, err := json.Marshal(&msg2)
		checkError(err)

		_, err = conn.Write(bytes2)
		checkError(err)
	} else {
		bytes3, err := json.Marshal(&msg3)
		checkError(err)

		_, err = conn.Write(bytes3)
		checkError(err)
	}

	for {
		//result, err := ioutil.ReadAll(conn)
		b := make([]byte, 1024)
		n, _ := conn.Read(b)
		a := make([]byte, n)
		a = b[:n]
		fmt.Println(string(a))
	}
}

type Msg struct {
	GroupId string
	Type    string
	Command string
	TxCount int
	IsEnd   bool
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
