package types

import "errors"

// for msa
var (
	ErrSessionIsNotRegister = errors.New("session is not register")
)

var (
	ErrAuthSessionAlreadyLogin              = errors.New("ws sessions login fail, sessions is already login")
	ErrAuthNotSupportUsingType              = errors.New("ws sessions login fail, login using type invalid")
	ErrAuthNotSupportHowType                = errors.New("ws sessions login fail, login how type invalid")
	ErrAuthNotSupportWhatType               = errors.New("ws sessions login fail, invalid auth.domain.what == session, jwt")
	ErrAuthSessionIsNotLogin                = errors.New("ws sessions login fail, sessions is not login")
	ErrAuthRequireJwtToken                  = errors.New("ws sessions login fail, using = key is require token")
	ErrAuthAlreadyLogin                     = errors.New("ws sessions login fail, user id is already login")
	ErrAuthNotMatchPassword                 = errors.New("ws sessions login fail, password not match")
	ErrCtrlInvalidInviteMe                  = errors.New("invalid ctrl.invite, invite = me")
	ErrCtrlInvalidNotOwner                  = errors.New("invalid ctrl.invite, request user is not owner in topic")
	ErrCtrlInvalidRequestUserNotSubs        = errors.New("invalid ctrl.invite, request user id not in topic")
	ErrCtrlInvalidInviteUserAlreadyJoin     = errors.New("invalid ctrl.invite, invite user id already in topic")
	ErrCtrlInvalidNotFoundTopic             = errors.New("invalid ctrl.join, not found topic")
	ErrCtrlInvalidPrivateRoom               = errors.New("invalid ctrl.join, private room")
	ErrCtrlInvalidAlreadyTopicJoinUser      = errors.New("invalid ctrl.join, already join")
	ErrDataBaseUpdateRowAffectIsZero        = errors.New("update query result row affected is zero")
	ErrDataBaseCacheSetFailed               = errors.New("database cache set failed")
	ErrDataBaseCacheGetFailed               = errors.New("database cache get failed")
	ErrDataBaseRoomSelectErrorRoomIdInvalid = errors.New("database topic select error, topic id invalid")
	ErrDataBaseRoomSelectErrorSubsIsNil     = errors.New("database topic subscription is nil")
	ErrDataBaseDeleteSubsFail               = errors.New("database delete subs fail")
	ErrInvalidRoot                          = errors.New("invalid type != msg, auth, meta, ctrl, file")
	ErrInvalidRequestJsonFormat             = errors.New("invalid request json format")
	ErrInvalidRequestTopicId                = errors.New("invalid request.topic.id")
	ErrInvalidRequestUser                   = errors.New("invalid request.user")
	ErrInvalidRequestUserDeviceInfo         = errors.New("invalid request.user.device_info")
	ErrInvalidRequestNotify                 = errors.New("invalid request.notify")
	ErrInvalidRequestUsers                  = errors.New("invalid request.users")
	ErrInvalidRequestWhat                   = errors.New("invalid request.what")
	ErrInvalidRequestHow                    = errors.New("invalid request.how")
	ErrInvalidRequestUsing                  = errors.New("invalid request.using")
	ErrInvalidRequestWho                    = errors.New("invalid request.who")
	ErrInvalidRequestUserAuthEmail          = errors.New("invalid request.user.auth.email")
	ErrInvalidRequestUserAuthPassword       = errors.New("invalid request.user.auth.password")
	ErrRequireRoot                          = errors.New("require root object")
	ErrRequireUUID                          = errors.New("require uuid")
	ErrRequireRequest                       = errors.New("require request")
	ErrRequireRequestNotify                 = errors.New("require request.notify")
	ErrRequireRequestWhat                   = errors.New("require request.what")
	ErrRequireRequestHow                    = errors.New("require request.how")
	ErrRequireRequestWho                    = errors.New("require request.who")
	ErrRequireRequestUsing                  = errors.New("require request.using")
	ErrRequireRequestUser                   = errors.New("require request.user")
	ErrRequireRequestUsers                  = errors.New("require request.users")
	ErrRequireRequestUserAuth               = errors.New("require request.user.auth")
	ErrRequireRequestUserDeviceInfo         = errors.New("require request.user.device_info")
	ErrRequireRequestUserAuthUserName       = errors.New("require request.user.auth.username")
	ErrRequireRequestUserAuthEmail          = errors.New("require request.user.auth.email")
	ErrRequireRequestTopic                  = errors.New("require request.topic")
	ErrRequireRequestTopicId                = errors.New("require request.topic_id")
	ErrRequireRequestUserId                 = errors.New("require request.user.id")
	ErrRequireRequestMessage                = errors.New("require request.message")
	ErrRequireRequestFileId                 = errors.New("require request.file_id")
	ErrRequireRequestSequenceId             = errors.New("require request.sequence_id")
	ErrRequireRequestToken                  = errors.New("require request.token")
	ErrRequireRequestProfile                = errors.New("require request.profile")
	ErrMetaNotContainsMessages              = errors.New("invalid meta.message, not contains message")
	ErrMetaTopicCountZero                   = errors.New("invalid meta.topic, topic count zero")
	ErrMsgNotEnteredTopic                   = errors.New("invalid msg.send, not entered topic")
)
