package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

//clientr端
var wg sync.WaitGroup

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("找不到客户端")
		return
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
					fmt.Println("无法连接服务端")
					wg.Done()
				}
			}
			fmt.Print("服务端：")
			fmt.Println(string(buf[:n]))
		}
	}()
	wg.Add(1)
	wg.Wait()
}
