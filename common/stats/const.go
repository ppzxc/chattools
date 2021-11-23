package stats

var Writer PromHashing

const TypeNamespaceLink = "link"

const TypeSession = "session"
const TypeSessionLogin = "login"
const TypeSessionWebsocket = "websocket"

const TypeTraffic = "traffic"
const TypeTrafficInboundTotal = "inbound_total"
const TypeTrafficOutboundTotal = "outbound_total"
const TypeTrafficFailOut = "fail_out"
const TypeTrafficSuccessOut = "success_out"

const TypeDurationTs = "duration_ts"
const TypeDurationDbQuery = "duration_db_query"

const TypeQueue = "queue"
const TypeQueueInboundRouter = "inbound_router"
const TypeQueueInboundCtrl = "inbound_ctrl"
const TypeQueueInboundMsg = "inbound_msg"
const TypeQueueInboundEmit = "inbound_emit"
const TypeQueueInboundAuth = "inbound_auth"
const TypeQueueInboundMeta = "inbound_meta"
const TypeQueueInboundNotify = "inbound_notify"

const TypeQueueStatus = "queue_status"
const TypeQueueStatusRouter = "router"
const TypeQueueStatusEmitter = "emitter"
const TypeQueueStatusAuth = "auth"
const TypeQueueStatusCtrl = "ctrl"
const TypeQueueStatusMeta = "meta"
const TypeQueueStatusNotify = "notify"

const TypeQueueOutboundRouter = "outbound_router"
const TypeQueueOutboundCtrl = "outbound_ctrl"
const TypeQueueOutboundMsg = "outbound_msg"
const TypeQueueOutboundEmit = "outbound_emit"
const TypeQueueOutboundAuth = "outbound_auth"
const TypeQueueOutboundMeta = "outbound_meta"
const TypeQueueOutboundNotify = "outbound_notify"

const TypeCache = "cache"
const TypeCacheCount = "cache_count"
