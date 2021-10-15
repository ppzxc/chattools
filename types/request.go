package types

// domain object type
const (
	WhatTopic   = "topic"
	WhatMessage = "message"
	WhatSession = "session"
	WhatToken   = "token"
	WhatUser    = "user"
	WhatProfile = "profile"

	WhatMention = "mention"
	WhatReply   = "reply"
	WhatNotify  = "notify"
	//WhatStabbing = "stabbing"
)

const (
	HowUpload = "upload"

	//UsingImage = "image"
)

// for notify
const (
	HowDelete = "delete"
)

// use what = session
const (
	HowLogin    = "login"
	HowLogout   = "logout"
	HowRegister = "register"
)

// use what = message
const (
	HowCtrl = "ctrl"
	HowSend = "send"
	HowAck  = "ack"
	HowRead = "read"
)

// use what = topic
const (
	HowJoin   = "join"
	HowLeave  = "leave"
	HowCreate = "create"
	HowSelect = "select"
	HowInvite = "invite"
)

// use what = session, how = login
const (
	UsingId        = "id"
	UsingToken     = "token"
	UsingAnonymous = "anonymous"
	UsingRotary    = "rotary"
	UsingMe        = "me"
	UsingFile      = "file"
	UsingTopic     = "topic"
)

const (
	WhoAlone  = "alone"
	WhoInvite = "invite"
)
