package global

type Bound string

//func (b Bound) ToString() string {
//
//}

const (
	BoundMessage        Bound = "msg"
	BoundControl        Bound = "ctrl"
	BoundMeta           Bound = "meta"
	BoundAuthentication Bound = "auth"
	BoundNotify         Bound = "notify"
	BoundFile           Bound = "file"
	BoundPing           Bound = "ping"
	BoundPong           Bound = "pong"
	BoundEtc            Bound = "etc"
)

type Authentication string

const (
	AuthLogin    Authentication = "login"
	AuthLogout   Authentication = "logout"
	AuthToken    Authentication = "token"
	AuthRegister Authentication = "register"
)

type Control string

const (
	CtrlCreate Control = "create"
	CtrlJoin   Control = "join"
	CtrlInvite Control = "invite"
	CtrlLeave  Control = "leave"
)

type Meta string

const (
	MetaTopic   Meta = "topic"
	MetaUser    Meta = "user"
	MetaMessage Meta = "message"
	MetaNotify  Meta = "notify"
	MetaProfile Meta = "profile"
)
