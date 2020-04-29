package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/wuqinqiang/chitchat/models"
	"net/http"
)

type Msg struct {
	Message     string    `json:"message"`
	UserName    string    `json:"user_name"`
	Type        int       `json:"type"`
	CreatedAt   string    `json:"created_at"`
	ContentType int       `json:"content_type"`
	To          int       `json:"to"`
}

var client_users = make(map[*websocket.Conn]int) //客户端连接绑定user_id

var user_clients = make(map[int]*websocket.Conn) //user_id 绑定客户端

var messageChannel = make(chan interface{}) //消息通道存储

var upgrader = websocket.Upgrader{
	//HandshakeTimeout: 5,
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	//当前预先任何连接
	CheckOrigin: func(r *http.Request) bool {
		return true },
}

//读取发送的消息
func Reader(conn *websocket.Conn, sess models.Session, r *http.Request) {
	for {
		var msg Msg
		//读取信息
		_, p, err := conn.ReadMessage()
		if err != nil {
			CloseClient(conn)
			break
		}

		err = json.Unmarshal([]byte(string(p)), &msg)
		if err != nil {
			danger(err.Error())
			break
		}
		user, err := sess.User()
		//记录发送消息
		if msg.Type == 5 { //是弹幕发送
			fmt.Println("收到消息了")

			if _, err := user.CreateMessage(RemoteIP(r), msg.Message, msg.Type); err != nil {
				danger(err.Error())
				break
			}
			msg.UserName = user.Name
			messageChannel <- msg
		} else { //单聊或者群聊消息
			 err := user.CreateChatMessage(msg.Message, msg.To, msg.Type, msg.ContentType)
			if err != nil {
				danger(err.Error())
				break
			}
			////找到次客户端发送消息
			client := user_clients[msg.To]

			if client == nil {
				//说明并没有登录，这条就算未读
				err:=models.AddUnreadMessage(user.Id,msg.To)
				if err !=nil{
					danger(err.Error())
					break
				}
			} else {
				err = client.WriteJSON(msg)
				if err != nil {
					fmt.Println("出错了")
					CloseClient(client)
					break
				}
			}
		}
	}
}

//获取消息
func SendClientMessage() {
	for {
		msg := <-messageChannel
		for client := range client_users {
			err := client.WriteJSON(msg)
			if err != nil {
				CloseClient(client)
				break
			}
		}
	}
}

func WsContent(w http.ResponseWriter, r *http.Request) {

	//没有登录的人不让连接
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	}

	//将此连接升级为ws

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		warning(err.Error())
		return
	}
	//双向绑定
	client_users[ws] = sess.UserId
	user_clients[sess.UserId] = ws
	go Reader(ws, sess, r)
	go SendClientMessage()
}

func CloseClient(client *websocket.Conn) {
	client.Close()
	delete(user_clients, client_users[client])
	delete(client_users, client)
}

