package main

import (
	"WonderfulAdventure/asset"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"runtime"
	"time"
)

const (
	Version         = "0.0.1"
	Author          = "ertuil"
	dbugHost        = "127.0.0.1:8080"
	debugSessionKey = "adwaaabbbccc123fads90wn"
	Host            = "[::]:80"
)

var (
	store    = sessions.NewCookieStore([]byte(debugSessionKey))
	upgrader = websocket.Upgrader{}
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "statics/html/index.html")
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {

	c, _ := upgrader.Upgrade(w, r, nil)

	go CoreWsHandle(c)
}

func CoreWsHandle(c *websocket.Conn) {
	defer c.Close()

	se := state{}
	se.setState(0, 0, 10, "")

	var err error
	for {
		switch se.Stage {
		case 0:
			err = Stage_0(c, &se)
		case 1:
			err = Stage_1(c, &se)
		case 2:
			err = Stage_2(c, &se)
		case 3:
			err = Stage_3(c, &se)
		case 10:
			getFlag(c, &se)
		}
		if err != nil {
			j, _ := MsgInitJson("系统", err.Error(), []string{}, se)
			_ = c.WriteMessage(websocket.TextMessage, j)
			break
		}
		if se.Stage == failstate {
			break
		}
		time.Sleep(300 * time.Microsecond)
	}
}

func main() {
	// 争取更多资源
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 释放静态资源
	if err := asset.RestoreAssets(".", "statics"); err != nil {
		log.Panic(err)
	}
	// 初始化状态转移矩阵
	stateTransInit()

	// 路由
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("./statics"))))
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/ws", WebSocketHandler)

	// 启动服务
	log.Println("Start listening to", dbugHost)
	err := http.ListenAndServe(dbugHost, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func getFlag(c *websocket.Conn, se *state) {
	flag := "flag{123}"
	j, _ := MsgInitJson("系统", flag, []string{}, *se)
	_ = c.WriteMessage(websocket.TextMessage, j)
	se.Stage = 11
}
