package model

import "time"

type File struct {
	Id         int64      `json:"id,omitempty" bson:"_id,omitempty"`
	Type       string     `json:"type,omitempty" bson:"type,omitempty"`
	TopicId    int64      `json:"topic_id,omitempty" bson:"topic_id,omitempty"`
	FromUserId int64      `json:"from_user_id,omitempty" bson:"from_user_id,omitempty"`
	WriteName  string     `json:"write_name,omitempty" bson:"write_name,omitempty"`
	Name       string     `json:"name,omitempty" bson:"name,omitempty"`
	Path       string     `json:"path,omitempty" bson:"path,omitempty"`
	Mime       string     `json:"mime,omitempty" bson:"mime,omitempty"`
	Size       int64      `json:"size,omitempty" bson:"size,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" bson:"deleted_at,omitempty""`
}