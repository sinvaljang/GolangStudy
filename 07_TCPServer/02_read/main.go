package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Protocol, ip, Port番号を設定し、Network接続を待機する
	// ipはlocalhostに
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	// 必ず接続を切ること
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
		}
		//go rutine
		go dosomething(conn)
	}
}

func dosomething(conn net.Conn) {
	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		line := scan.Text()
		fmt.Println(line)
	}
	// 必ず接続を切ること
	defer conn.Close()
}
