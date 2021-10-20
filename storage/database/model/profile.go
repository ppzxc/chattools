package model

import "time"

type Profile struct {
	Id          int64      `json:"id,omitempty" bson:"_id,omitempty"`
	UserId      int64      `json:"user_id,omitempty" bson:"user_id,omitempty"`
	FileId      int64      `json:"file_id,omitempty" bson:"file_id,omitempty"`
	Description string     `json:"description,omitempty" bson:"description,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}
