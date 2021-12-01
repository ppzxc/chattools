package model

import "time"

type Paging struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Id        int64      `json:"id,omitempty"`
	Offset    int64      `json:"offset,omitempty"`
	Limit     int64      `json:"limit,omitempty"`
}
