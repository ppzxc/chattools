package model

import "time"

type User struct {
	Id        int64      `json:"id,omitempty" bson:"_id,omitempty"`
	State     string     `json:"state,omitempty" bson:"state,omitempty"`
	StatedAt  *time.Time `json:"stated_at,omitempty" bson:"stated_at,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`

	Device         []*Device       `json:"device,omitempty" bson:"devices,omitempty"`      // relationship
	Authentication *Authentication `json:"auth,omitempty" bson:"authentication,omitempty"` // relationship
	Profile        *Profile        `json:"profile,omitempty" bson:"profile,omitempty"`     // relationship
}
