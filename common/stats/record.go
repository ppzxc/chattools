package stats

import (
	"strings"
	"time"
)

const (
	INSERT = "insert"
	SELECT = "select"
	UPDATE = "update"
	CREATE = "create"
	DELETE = "delete"
)

func QueryRecord(query string, table string, caller string, start time.Time) {
	if strings.Contains(query, INSERT) {
		Writer.IncHistogramQuery(TypeDurationDbQuery, "insert", table, caller, start)
	} else if strings.Contains(query, CREATE) {
		Writer.IncHistogramQuery(TypeDurationDbQuery, "create", table, caller, start)
	} else if strings.Contains(query, DELETE) {
		Writer.IncHistogramQuery(TypeDurationDbQuery, "delete", table, caller, start)
	} else if strings.Contains(query, UPDATE) {
		Writer.IncHistogramQuery(TypeDurationDbQuery, "update", table, caller, start)
	} else if strings.Contains(query, SELECT) {
		Writer.IncHistogramQuery(TypeDurationDbQuery, "select", table, caller, start)
	} else {
		Writer.IncHistogramQuery(TypeDurationDbQuery, "etc", table, caller, start)
	}
}
