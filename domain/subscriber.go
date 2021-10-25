package domain

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/kataras/iris/v12/websocket"
	"github.com/sirupsen/logrus"
	"time"
)

type Subscriber interface {
	Serve()
	Close()
	GetPubSub() *redis.PubSub
	GetKey() string
}

func NewSubscriber(root context.Context, key string, redisPubSub *redis.PubSub, conn *websocket.Conn) Subscriber {
	logrus.WithFields(logrus.Fields{
		"key":         key,
		"session.id":  conn.ID(),
		"remote.addr": conn.Socket().NetConn().RemoteAddr(),
		"local.addr":  conn.Socket().NetConn().LocalAddr(),
	}).Debug("new subscriber")

	ctx, cancel := context.WithCancel(root)

	return &subscriber{
		ctx:         ctx,
		cancel:      cancel,
		key:         key,
		redisPubSub: redisPubSub,
		conn:        conn,
	}
}

type subscriber struct {
	key         string
	ctx         context.Context
	cancel      context.CancelFunc
	redisPubSub *redis.PubSub
	conn        *websocket.Conn
}

func (s subscriber) GetPubSub() *redis.PubSub {
	return s.redisPubSub
}

func (s subscriber) GetKey() string {
	return s.key
}

func (s subscriber) Serve() {
	go func() {
		logrus.WithFields(logrus.Fields{
			"key":         s.key,
			"session.id":  s.conn.ID(),
			"remote.addr": s.conn.Socket().NetConn().RemoteAddr(),
			"local.addr":  s.conn.Socket().NetConn().LocalAddr(),
		}).Debug("connection default subscribe receiver start")

		defer logrus.WithFields(logrus.Fields{
			"key":         s.key,
			"session.id":  s.conn.ID(),
			"remote.addr": s.conn.Socket().NetConn().RemoteAddr(),
			"local.addr":  s.conn.Socket().NetConn().LocalAddr(),
		}).Debug("connection default subscribe receiver close")

		ch := s.redisPubSub.Channel()
		for {
			select {
			case <-s.ctx.Done():
				logrus.WithFields(logrus.Fields{
					"key":         s.key,
					"session.id":  s.conn.ID(),
					"remote.addr": s.conn.Socket().NetConn().RemoteAddr(),
					"local.addr":  s.conn.Socket().NetConn().LocalAddr(),
				}).Debug("call context cancel")
				return
			case msg, isClose := <-ch:
				if !isClose {
					logrus.WithFields(logrus.Fields{
						"key":           s.key,
						"session.id":    s.conn.ID(),
						"remote.addr":   s.conn.Socket().NetConn().RemoteAddr(),
						"local.addr":    s.conn.Socket().NetConn().LocalAddr(),
						"channel":       msg.Channel,
						"pattern":       msg.Pattern,
						"payload":       msg.Payload,
						"payload.slice": msg.PayloadSlice,
					}).Debug("Receiver receive")

					if err := s.conn.Socket().WriteText([]byte(msg.Payload), 10*time.Second); err != nil {
						logrus.WithFields(logrus.Fields{
							"key":           s.key,
							"session.id":    s.conn.ID(),
							"remote.addr":   s.conn.Socket().NetConn().RemoteAddr(),
							"local.addr":    s.conn.Socket().NetConn().LocalAddr(),
							"channel":       msg.Channel,
							"pattern":       msg.Pattern,
							"payload":       msg.Payload,
							"payload.slice": msg.PayloadSlice,
						}).WithError(err).Error("Receiver write error")
					}
				} else {
					logrus.WithFields(logrus.Fields{
						"key":         s.key,
						"session.id":  s.conn.ID(),
						"remote.addr": s.conn.Socket().NetConn().RemoteAddr(),
						"local.addr":  s.conn.Socket().NetConn().LocalAddr(),
					}).Debug("channel closed")
				}
			}
		}
	}()
}

func (s subscriber) Close() {
	logrus.WithFields(logrus.Fields{
		"key":         s.key,
		"session.id":  s.conn.ID(),
		"remote.addr": s.conn.Socket().NetConn().RemoteAddr(),
		"local.addr":  s.conn.Socket().NetConn().LocalAddr(),
	}).Debug("subscribe receiver close")
	s.cancel()
}
