package outbound

import "time"

type Root struct {
	Auth   *Authentication `json:"auth,omitempty"`
	Ctrl   *Control        `json:"ctrl,omitempty"`
	Msg    *Msg            `json:"msg,omitempty"`
	Meta   *Meta           `json:"meta,omitempty"`
	Notify *Notify         `json:"notify,omitempty"`
	File   *File           `json:"file,omitempty"`
	Pong   *Pong           `json:"pong,omitempty"`

	StatusCode int64  `json:"status_code,omitempty"`
	Status     string `json:"status,omitempty"`
	Cause      string `json:"cause,omitempty"`
}

type Notify struct {
}
type File struct {
}
type Pong struct {
}

type User struct {
	Id             int64      `json:"id,omitempty"`
	State          string     `json:"state,omitempty"`
	StatedAt       *time.Time `json:"stated_at,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
	Authentication *Auth      `json:"auth,omitempty"`
	Profile        *Profile   `json:"profile,omitempty"`
}

type Auth struct {
	UserId   int64  `json:"user_id,omitempty" bson:"user_id,omitempty"`
	UserName string `json:"user_name,omitempty" bson:"user_name,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
}

type Profile struct {
	FileId      int64  `json:"file_id,omitempty"`
	Description string `json:"description,omitempty"`
}
