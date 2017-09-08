package webSocket

import "testing"

func TestWebSocket(t *testing.T){
	NerverStop := make(chan struct{})
	go ServerHttp()
	go WebSocketClient()
	<- NerverStop
}
