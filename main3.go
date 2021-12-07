package main

import (
	"github.com/kataras/neffos"
	"github.com/kataras/neffos/gorilla"
	"log"
	"net/http"
)

//var send int64
//var recv int64

var events = neffos.Events{
	neffos.OnNativeMessage: func(c *neffos.NSConn, msg neffos.Message) error {
		//log.Printf("Got: %s", string(msg.Body))

		if !c.Conn.IsClient() {
			return c.Conn.Socket().WriteText(msg.Body, 0)
		}

		return nil
	},
}

func main() {
	websocketServer := neffos.New(
		gorilla.DefaultUpgrader, events)

	router := http.NewServeMux()
	router.Handle("/endpoint", websocketServer)
	router.Handle("/", http.FileServer(http.Dir("./browser")))

	log.Println("Serving websockets on localhost:8080/endpoint")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", router))
}
