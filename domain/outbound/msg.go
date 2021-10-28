package outbound

type Msg struct {
	UUID string        `json:"uuid,omitempty"`
	Send *ResponseSend `json:"send,omitempty"`
	Ack  *ResponseAck  `json:"ack,omitempty"`
	Read *ResponseRead `json:"read,omitempty"`
	File *ResponseFile `json:"file,omitempty"`
}

type ResponseSend struct {
	Message Message `json:"message,omitempty"`
}

type ResponseAck struct {
	TopicId    int64 `json:"topic_id,omitempty"`
	SequenceId int64 `json:"sequence_id,omitempty"`
}

type ResponseRead struct {
	TopicId    int64 `json:"topic_id,omitempty"`
	SequenceId int64 `json:"sequence_id,omitempty"`
}

type ResponseFile struct {
	TopicId int64 `json:"topic_id,omitempty"`
	FileId  int64 `json:"file_id,omitempty"`
}
