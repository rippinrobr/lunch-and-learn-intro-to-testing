package context

import (
	"time"
)

// TraceIDHeader is the header added to outgoing requests which adds the
// traceID to it.
const TraceIDHeader = "X-Trace-ID"

// Key represents the type of value for the context key.
type ctxKey int

// KeyValues is how request values or stored/retrieved.
const KeyValues ctxKey = 1

// Values hold the state for each request.
type Values struct {
	TraceID    string
	Now        time.Time
	StatusCode int
}
