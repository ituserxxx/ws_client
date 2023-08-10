package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"time"
)

func main(){

	parasm := os.Args[1:]
	if len(parasm) != 1{
		fmt.Println("必须有一个参数，执行示例：ws_client.exe ws://172.16.9.103:1880/api/xxx ")
		return
	}
	fmt.Println("正在连接："+parasm[0])
	// 建立 WebSocket 连接
	conn, _, err := websocket.DefaultDialer.Dial(parasm[0], nil)
	if err != nil {
		fmt.Println("无法建立 WebSocket 连接：", err.Error())
		return
	}
	defer conn.Close()
	fmt.Println("连接成功~~")
	// 接收消息
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("读取消息出错："+err.Error())
				return
			}
			fmt.Println("接收到消息：", string(message))
		}
	}()
	var i int
	for {
		if i== 100{
			return
		}
		time.Sleep(3 *time.Second)
		// 发送消息
		message := []byte("Hello, WebSocket!")
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			fmt.Println("发送消息出错："+err.Error())
		}else{
			fmt.Printf("\n发送消息成功：%s  时间：%s",string(message),time.Now().Format("2006-01-02 15:04:05"))
		}
		i++
	}

}
