package outbound

type Meta struct {
	UUID    string               `json:"uuid,omitempty"`
	Topic   *ResponseMetaTopic   `json:"topic,omitempty"`
	User    *ResponseMetaUser    `json:"user,omitempty"`
	Message *ResponseMetaMessage `json:"message,omitempty"`
	Notify  *ResponseMetaNotify  `json:"notify,omitempty"`
	Profile *ResponseMetaProfile `json:"profile,omitempty"`
}

type ResponseMetaTopic struct {
	Topics []*Topic `json:"topics,omitempty"`
}

type ResponseMetaUser struct {
	Users []*User `json:"users,omitempty"`
}

type ResponseMetaMessage struct {
	Messages []*Message `json:"messages,omitempty"`
}

type ResponseMetaNotify struct {
}

type ResponseMetaProfile struct {
}
