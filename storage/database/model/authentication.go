package model

import "time"

type Authentication struct {
	Id        int64      `json:"id,omitempty" bson:"_id,omitempty"`
	UserId    int64      `json:"user_id,omitempty" bson:"user_id,omitempty"`
	UserName  string     `json:"name,omitempty" bson:"user_name,omitempty"`
	Email     string     `json:"email,omitempty" bson:"email,omitempty"`
	Password  string     `json:"password,omitempty" bson:"password,omitempty"`
	AuthType  string     `json:"auth_type,omitempty" bson:"auth_type,omitempty"`
	AuthLevel string     `json:"auth_level,omitempty" bson:"auth_level,omitempty"`
	Secret    string     `json:"secret,omitempty" bson:"secret,omitempty"`
	Expires   *time.Time `json:"expires,omitempty" bson:"expires,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}
