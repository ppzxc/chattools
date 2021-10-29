package outbound

type Notify struct {
	UUID    string           `json:"uuid,omitempty"`
	Mention *ResponseMention `json:"mention,omitempty"`
	Reply   *ResponseReply   `json:"reply,omitempty"`
}

type ResponseMention struct {
	Create  *Empty `json:"create,omitempty"`
	Receive *Empty `json:"receive,omitempty"`
	Read    *Empty `json:"read,omitempty"`
	Delete  *Empty `json:"delete,omitempty"`
}

type ResponseReply struct {
	Create  *Empty `json:"create,omitempty"`
	Receive *Empty `json:"receive,omitempty"`
	Read    *Empty `json:"read,omitempty"`
	Delete  *Empty `json:"delete,omitempty"`
}

type Empty struct {
}
