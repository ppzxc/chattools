package domain

import "github.com/ppzxc/chattools/domain/outbound"

type Command int

const (
	TopicCreate Command = iota
	TopicLeave
	TopicJoin
	TopicInvite

	MsgAck
	MsgRead

	NotifyCreate
	NotifyCrud
)

type SyncProtocol struct {
	Command       Command       `json:"command"`
	TransactionId string        `json:"transaction_id"`
	UserId        int64         `json:"user_id"`
	TopicId       int64         `json:"topic_id"`
	Payload       outbound.Root `json:"payload"`
}
