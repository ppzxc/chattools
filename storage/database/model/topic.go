package model

import "time"

type Topic struct {
	Id        int64      `json:"id,omitempty" bson:"_id,omitempty"`
	State     string     `json:"state,omitempty" bson:"state,omitempty"`
	StatedAt  *time.Time `json:"stated_at,omitempty" bson:"stated_at,omitempty"`
	Name      string     `json:"name,omitempty" bson:"name,omitempty"`
	Owner     int64      `json:"owner,omitempty" bson:"owner,omitempty"`
	Private   bool       `json:"private,omitempty" bson:"private,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`

	Message []*Message `json:"messages,omitempty" bson:"messages,omitempty"`
}
