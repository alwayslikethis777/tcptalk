package main

import (
	"fmt"
	"io"
	"net"
)

//server端

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.lisen.err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.accept.err:", err)
			continue
		} else {
			fmt.Println("客户端上线了")
		}
		go func() {
			var str string
			for {

				fmt.Scan(&str)
				conn.Write([]byte(str))
			}
		}()
		go func() {

			for {
				buf := make([]byte, 1024)
				n, err := conn.Read(buf)

				if err != nil {
					if err == io.EOF {
						return
					} else {
						fmt.Println("客户端退出了")
						return
					}
				}
				fmt.Print("客户端：")
				fmt.Println(string(buf[:n]))
			}
		}()

	}

}
