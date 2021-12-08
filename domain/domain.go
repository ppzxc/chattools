package domain

import (
	"github.com/ppzxc/chattools/domain/outbound"
	"github.com/sirupsen/logrus"
)

type Command int

const (
	PubSubTopicCreate Command = iota + 1
	PubSubTopicLeave
	PubSubTopicJoin
	PubSubTopicInvite

	PubSubMsgAck
	PubSubMsgRead
	PubSubMsgSend

	PubSubNotifyCreate
	PubSubNotifyCrud

	PubSubDefaultWriteOut
)

type PubSubProtocol struct {
	Command Command       `json:"command"`
	TopicId int64         `json:"topic_id"`
	Payload outbound.Root `json:"payload"`
	Fields  logrus.Fields `json:"fields"`

	//TransactionId string        `json:"transaction_id"`
	//UserId        int64         `json:"user_id"`
}

type PubSubs struct {
	Command Command `json:"command"`
	TopicId int64   `json:"topic_id"`
	Payload []byte  `json:"payload"`
}
