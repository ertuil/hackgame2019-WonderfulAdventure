package main

import (
	"WonderfulAdventure/asset"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
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
	defer c.Close()

	for {
		mt, message, _ := c.ReadMessage()
		c.WriteMessage(mt, append([]byte("hello "), message[:]...))
	}

}

func StartHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   1200, // 20 minutes
		HttpOnly: true,
	}

	session.Values["attack"] = 0
	session.Values["money"] = 10
	session.Values["name"] = ""
	session.Values["stage"] = 0
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	j, err := loadFromSessionToJson(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(j)
}

func main() {

	// 释放静态资源
	if err := asset.RestoreAssets(".", "statics"); err != nil {
		log.Panic(err)
	}

	// 路由
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("./statics"))))
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/ws", WebSocketHandler)
	http.HandleFunc("/start", StartHandler)

	// 启动服务
	log.Println("Start listening to", dbugHost)
	err := http.ListenAndServe(dbugHost, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
