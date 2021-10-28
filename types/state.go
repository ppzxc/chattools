package types

const (
	StateUserCreated  = "CREATED"
	StateUserActive   = "ACTIVE"
	StateUserInactive = "INACTIVE"
	StateUserRemoved  = "REMOVED"
)

const (
	StateTopicCreated  = "CREATED"
	StateTopicActive   = "ACTIVE"
	StateTopicInactive = "INACTIVE"
	StateTopicRemoved  = "REMOVED"
)

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
