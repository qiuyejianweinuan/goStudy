package main

import (
	"MYTESTPROC_GO/main/service"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		//服务端接收消息
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("无法接收")
			break
		}
		msg, _ := service.Completions(reply)

		if msg == "" {
			msg = "和chatGpt的请求可能出现了错误,请尝试重新发送消息"
		}
		msg = "\n" + msg
		fmt.Println("客户端回复消息: " + msg)
		fmt.Println("发送到客户端: " + msg)

		//服务端向客户端发送消息
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("无法发送")
			break
		}
	}
}

func main() {
	//开启websocket连接
	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
