package main

import (
	"fmt"
	service2 "singo/service"

	"log"
	"net/http"
	"os"
	"singo/conf"
	"singo/server"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	go func() {
		router := mux.NewRouter()
		router.HandleFunc("/echo", echoHandler)
		if err := http.ListenAndServe(":6000", router); err != nil {
			fmt.Println("err:", err)
		}
	}()
	// 装载路由
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := server.NewRouter()
	r.Run(":3000")

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func echoHandler(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	// 定时发送心跳消息
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := conn.WriteMessage(websocket.TextMessage, []byte("ping from server")); err != nil {
					//log.Println("Ping error:", err)
					return
				}
				// fmt.Println("send ping")
			}
		}
	}()

	// 处理 WebSocket 消息
	var serialID string
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			//log.Println("Read error:", err)
			log.Println("kill sid: ", serialID)
			service := service2.UserLeaveService{SerialUUID: serialID}
			service.OperateV1()
			return
		}
		serialID = string(message)
	}
}
