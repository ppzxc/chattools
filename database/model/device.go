package model

import "time"

type Device struct {
	Id              int64      `json:"id,omitempty" bson:"_id,omitempty"`
	UserId          int64      `json:"user_id,omitempty" bson:"user_id,omitempty"`
	DeviceId        string     `json:"device_id,omitempty" bson:"device_id,omitempty"`
	BrowserId       string     `json:"browser_id,omitempty" bson:"browser_id,omitempty"`
	UserAgent       string     `json:"user_agent,omitempty" bson:"user_agent,omitempty"`
	OperationSystem string     `json:"operation_system,omitempty" bson:"operation_system,omitempty"`
	Platform        string     `json:"platform,omitempty" bson:"platform,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}
