package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func echo(w http.ResponseWriter, r *http.Request) {
	//https://godoc.org/github.com/gorilla/websocket#Upgrader.Subprotocols
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{"echo-protocol"},
	} // use default options
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {

	var addr = flag.String("addr", "localhost:8080", "http service address")

	flag.Parse()
	http.HandleFunc("/ws", echo)
	http.ListenAndServe(*addr, nil)
}
