package zlog

const (
	ContextKeyLogID     = "logId"
	HeaderKeyLogID      = "X-Log-Id"
	HeaderKeyLowerLogID = "x-log-id"
	ContextKeyIp        = "ip"
	ContextKeyUri       = "uri"
	ContextKeyUid       = "uid"
)

const (
	LogFileLevelStdout    = "stdout"
	LogFileLevelNormal    = "normal"
	LogFileLevelWarnFatal = "warnfatal"
)
