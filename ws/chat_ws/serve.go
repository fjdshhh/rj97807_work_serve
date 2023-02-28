package chat_ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// 所有客户端
var aliveUsers = make(map[*websocket.Conn]bool)

//消息体
type Msg struct {
	Message string `json:"message"`
}

// BroadcastMsg 向所有客户端广播消息
func BroadcastMsg(mt int, message []byte, conn *websocket.Conn) {
	for aliveUser := range aliveUsers {
		if aliveUser == conn {
			break
		}
		aliveUser.WriteMessage(mt, message)
	}
}

func ServeWs(w http.ResponseWriter, r *http.Request, engineWeb *gorm.DB) {
	r.Header.Del("Origin")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logx.Error("upgrader error:" + err.Error())
		return
	}
	aliveUsers[conn] = true
	for {
		mt, readMsg, err := conn.ReadMessage()
		var result = new(Msg)
		json.Unmarshal(readMsg, result)
		if err != nil || result.Message == "" {
			continue
		}
		if result.Message == ".exit" {
			delete(aliveUsers, conn)
			exit := Msg{Message: fmt.Sprintf("您已退出，当前还有%v人数", len(aliveUsers))}
			byteExit, _ := json.Marshal(exit)
			conn.WriteMessage(websocket.TextMessage, byteExit)
			conn.Close()
			break
		}
		if result.Message == "ping" {
			bytePong, _ := json.Marshal(Msg{Message: "pong"})
			conn.WriteMessage(websocket.TextMessage, bytePong)
			//conn.WriteMessage(websocket.PongMessage, []byte("pong"))
		}
		BroadcastMsg(mt, readMsg, conn)
	}
}
