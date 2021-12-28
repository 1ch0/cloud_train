package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://127.0.0.1:8080/")
	defer resp.Body.Close()

	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			//fmt.Println("读取完毕")
			res := string(buf[:n])
			fmt.Println(res)
			break
		}
	}
}
