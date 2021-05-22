package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net/url"
	"server/msg"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type websocketClientManager struct {
	conn        *websocket.Conn
	addr        *string
	path        string
	sendMsgChan chan string
	recvMsgChan chan string
	isAlive     bool
	timeout     int
}

// 构造函数
func NewWsClientManager(addrIp, addrPort, path string, timeout int) *websocketClientManager {
	addrString := addrIp + ":" + addrPort
	var sendChan = make(chan string, 100)
	var recvChan = make(chan string, 100)
	var conn *websocket.Conn
	return &websocketClientManager{
		addr:        &addrString,
		path:        path,
		conn:        conn,
		sendMsgChan: sendChan,
		recvMsgChan: recvChan,
		isAlive:     false,
		timeout:     timeout,
	}
}

// 链接服务端
func (wsc *websocketClientManager) dail() {
	var err error
	u := url.URL{Scheme: "ws", Host: *wsc.addr, Path: wsc.path}
	log.Printf("connecting to %s", u.String())
	wsc.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return

	}
	wsc.isAlive = true
	log.Printf("connecting to %s 链接成功！！！", u.String())

}

// 发送消息
func (wsc *websocketClientManager) sendMsgThread() {

	// mess := &msg.Hello{
	// 	Name: "i'm proto client name hello",
	// }
	mess := &msg.Test{
		Id: proto.String("10002000"),
	}

	// mess2 := &msg.Hello{}
	data, err := proto.Marshal(mess)
	if err != nil {
		log.Println("proto marshal:", err)
	}

	// 对应服务端数据id

	var btId = []byte{0, 1}
	fmt.Println("proto marshal id  : ", btId)

	m := make([]byte, len(data)+len(btId))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)+len(btId)))

	var buffer bytes.Buffer
	buffer.Write(btId)
	buffer.Write(data)
	newBt := buffer.Bytes()
	copy(m, newBt)

	fmt.Println("mata : ", m)

	fmt.Println("big : ", binary.BigEndian.Uint16(m))
	fmt.Println("small : ", binary.LittleEndian.Uint16(m))

	err = wsc.conn.WriteMessage(websocket.TextMessage, m)
	if err != nil {
		log.Println("write:", err)
	}
}

// 读取消息
func (wsc *websocketClientManager) readMsgThread() {
	go func() {
		for {
			if wsc.conn != nil {
				_, message, err := wsc.conn.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					wsc.isAlive = false
					// 出现错误，退出读取，尝试重连
					break
				}
				log.Printf("recv: %s", message)
				// 需要读取数据，不然会阻塞
				// wsc.recvMsgChan <- string(message)
			}

		}
	}()
}
func main() {
	wsc := NewWsClientManager("127.0.0.1", "7777", "", 10)
	wsc.dail()
	wsc.readMsgThread()
	wsc.sendMsgThread()

	// var c = make(chan int)
	// <-c
}
