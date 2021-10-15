package types

// for user table
const (
	StateUserCreated  = "CREATED"
	StateUserActive   = "ACTIVE"
	StateUserInactive = "INACTIVE"
	StateUserRemoved  = "REMOVED"
)

// for topic table
const (
	StateTopicCreated  = "CREATED"
	StateTopicActive   = "ACTIVE"
	StateTopicInactive = "INACTIVE"
	StateTopicRemoved  = "REMOVED"
)

// for auth table
const (
	StateAuthLevelAdmin     = "admin"
	StateAuthLevelUser      = "user"
	StateAuthLevelAnonymous = "anonymous"
)

const (
	StateAuthTypeId        = "id"
	StateAuthTypeToken     = "token"
	StateAuthTypeAnonymous = "anonymous"
	StateAuthTypeRotary    = "rotary"
)
