package webSocket

import (
	"net/http"
	"golang.org/x/net/websocket"
	"fmt"
	"log"
	"time"
)

func WebSocketServer(ws *websocket.Conn){
	msg := make([]byte, 512)
	n, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", msg[:n])

	send_msg := "[" + string(msg[:n]) + "]"
	_, err = ws.Write([]byte(send_msg))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", send_msg)
}

func HelloWorld(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("hello,world"))
}

func ServerHttp(){
	http.Handle("/server", websocket.Handler(WebSocketServer))
	http.HandleFunc("/",HelloWorld)
	err := http.ListenAndServe(":28080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}



func WebSocketClient(){
	ws, err := websocket.Dial("ws://127.0.0.1:28080/server", "", "http://127.0.0.1:28080/")
	if err != nil {
		log.Fatal(err)
	}
	message := []byte("hello, world!你好")
	_, err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", message)
    time.Sleep(4*time.Second)
	var msg = make([]byte, 512)
	m, err := ws.Read(msg)
	if err != nil {
		log.Fatalln("client read",err)
	}
	fmt.Printf("Receive: %s\n", msg[:m])

	ws.Close()//关闭连接
}
