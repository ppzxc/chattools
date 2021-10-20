package model

import "time"

type Subscription struct {
	Id                int64      `json:"id,omitempty" bson:"_id,omitempty"`
	UserId            int64      `json:"user_id,omitempty" bson:"user_id,omitempty"`
	TopicId           int64      `json:"topic_id,omitempty" bson:"topic_id,omitempty"`
	StartSequenceId   int64      `json:"start_sequence_id,omitempty" bson:"start_sequence_id,omitempty"`
	ReceiveSequenceId int64      `json:"receive_sequence_id,omitempty" bson:"receive_sequence_id,omitempty"`
	ReadSequenceId    int64      `json:"read_sequence_id,omitempty" bson:"read_sequence_id,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}
