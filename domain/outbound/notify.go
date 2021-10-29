package outbound

type Notify struct {
	UUID    string          `json:"uuid,omitempty"`
	Mention *RequestMention `json:"mention,omitempty"`
	Reply   *RequestReply   `json:"reply,omitempty"`
}

type RequestMention struct {
	Create  *Empty `json:"create,omitempty"`
	Receive *Empty `json:"receive,omitempty"`
	Read    *Empty `json:"read,omitempty"`
	Delete  *Empty `json:"delete,omitempty"`
}

type RequestReply struct {
	Create  *Empty `json:"create,omitempty"`
	Receive *Empty `json:"receive,omitempty"`
	Read    *Empty `json:"read,omitempty"`
	Delete  *Empty `json:"delete,omitempty"`
}

type Empty struct {
}
