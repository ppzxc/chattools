package main

import (
	"fmt"
	gorillaWs "github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	context2 "github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
	"github.com/kataras/neffos/gorilla"
	"github.com/ppzxc/chattools/utils/mono"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

var send int64
var recv int64

func main() {
	mtx := sync.Mutex{}
	clients := make(map[string]struct{})

	mono.Init()

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-ticker.C:
				fmt.Printf("SEND: %v, RECV: %v\r\n", send, recv)
			}
		}
	}()

	ws := websocket.New(
		gorilla.Upgrader(gorillaWs.Upgrader{
			CheckOrigin: func(*http.Request) bool { return true },
		}),
		websocket.Events{
			websocket.OnNativeMessage: func(conn *neffos.NSConn, message neffos.Message) error {
				mtx.Lock()
				if _, ok := clients[conn.Conn.ID()]; !ok {
					panic(fmt.Sprintf("%v: not contains key", conn.Conn.ID()))
				}
				mtx.Unlock()
				atomic.AddInt64(&recv, 1)

				//fmt.Printf("%v: %v\r\n", conn.Conn.ID(), string(message.Body))
				err := conn.Conn.Socket().WriteText(message.Body, 5*time.Second)
				if err != nil {
					return err
				}
				atomic.AddInt64(&send, 1)
				return nil
			},
		},
	)

	// websocket connect
	ws.OnConnect = func(c *neffos.Conn) error {
		mtx.Lock()
		defer mtx.Unlock()
		fmt.Printf("%v: On Connect\r\n", c.ID())
		if _, ok := clients[c.ID()]; ok {
			panic(fmt.Sprintf("%v: same key", c.ID()))
		} else {
			clients[c.ID()] = struct{}{}
		}
		return nil
	}

	// websocket disconnect
	ws.OnDisconnect = func(c *neffos.Conn) {
		mtx.Lock()
		defer mtx.Unlock()
		fmt.Printf("%v: On Disconnect\r\n", c.ID())
		if _, ok := clients[c.ID()]; ok {
			delete(clients, c.ID())
		} else {
			panic(fmt.Sprintf("%v: no key", c.ID()))
		}
		return
	}

	app := iris.New()
	app.Get("/ws/{token}", websocket.Handler(ws, wrapIDGenerator()))

	fmt.Println(app.Listen("0.0.0.0:3000"))
}

func wrapIDGenerator() websocket.IDGenerator {
	return func(c context2.Context) string {
		//id, err := uuid.NewV4()
		//if err != nil {
		//	return strconv.FormatInt(time.Now().Unix(), 10)
		//}
		//return id.String()
		return mono.GetMONOID()
	}
}
