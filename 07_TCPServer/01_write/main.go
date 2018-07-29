package main

import (
	"fmt"
	"io"
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

	// 無限ループにて、接続を待機する
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// メッセージ伝送
		io.WriteString(conn, "\n HELLO TCP Server")
		fmt.Fprintln(conn, "\nHello there")
		fmt.Fprintf(conn, "%s", "Hello !\n")

		// 必ず接続を切ること
		conn.Close()
	}
}
