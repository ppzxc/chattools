package model

import "time"

type Message struct {
	Id           int64       `json:"id,omitempty" bson:"_id,omitempty"`
	MessageType  string      `json:"message_type,omitempty" bson:"message_type,omitempty"`
	ClintUUID    string      `json:"clint_uuid,omitempty" bson:"clint_uuid,omitempty"`
	FromUserId   int64       `json:"from_user_id,omitempty" bson:"from_user_id,omitempty"`
	FromUserName string      `json:"from_user_name,omitempty" bson:"from_user_name,omitempty"`
	TopicId      int64       `json:"topic_id,omitempty" bson:"topic_id,omitempty"`
	FileId       int64       `json:"file_id,omitempty" bson:"file_id,omitempty"`
	SequenceId   int64       `json:"sequence_id,omitempty" bson:"sequence_id,omitempty"`
	Content      string      `json:"content,omitempty" bson:"content,omitempty"`
	Custom       interface{} `json:"custom,omitempty" bson:"custom,omitempty"`
	CreatedAt    *time.Time  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    *time.Time  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt    *time.Time  `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}