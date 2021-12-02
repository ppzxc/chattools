package model

import "time"

type Notify struct {
	Id            int64       `json:"id,omitempty" bson:"_id,omitempty"`
	Type          string      `json:"type,omitempty" bson:"type,omitempty"`
	RequestUserId int64       `json:"request_user_id,omitempty" bson:"request_user_id,omitempty"`
	ReceiveUserId int64       `json:"receive_user_id,omitempty" bson:"receive_user_id,omitempty"`
	TopicId       int64       `json:"topic_id,omitempty" bson:"topic_id,omitempty"`
	SequenceId    int64       `json:"sequence_id,omitempty" bson:"sequence_id,omitempty"`
	IsReceived    bool        `json:"is_received,omitempty" bson:"is_received,omitempty"`
	IsRead        bool        `json:"is_read,omitempty" bson:"is_read,omitempty"`
	Custom        interface{} `json:"custom,omitempty" bson:"custom"`
	CreatedAt     *time.Time  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     *time.Time  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt     *time.Time  `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}
