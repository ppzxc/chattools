package model

type Serial struct {
	ID  string `bson:"_id"`
	Seq int64  `bson:"seq"`
}
