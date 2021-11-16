package outbound

type Meta struct {
	UUID    string           `json:"uuid,omitempty"`
	Topic   *ResponseTopic   `json:"topic,omitempty"`
	User    *ResponseUser    `json:"user,omitempty"`
	Message *ResponseMessage `json:"message,omitempty"`
	Notify  *ResponseNotify  `json:"notify,omitempty"`
	Profile *ResponseProfile `json:"profile,omitempty"`
}

type ResponseTopic struct {
	Topics []*Topic `json:"topics,omitempty"`
}

type ResponseUser struct {
	Users []*User `json:"users,omitempty"`
}

type ResponseMessage struct {
	Messages []*Message `json:"messages,omitempty"`
}

type ResponseNotify struct {
	Notifications []*Notification `json:"notifications,omitempty"`
}

type ResponseProfile struct {
	UserId      int64  `json:"user_id,omitempty"`
	FileId      int64  `json:"file_id,omitempty"`
	Description string `json:"description,omitempty"`
}
