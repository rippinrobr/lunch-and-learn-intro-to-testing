package middleware

// RequestLogger writes some information about the request to the logs in
import (
	"context"
	"log"
	"net/http"
	"time"

	clContext "github.com/rippinrobr/lunch-n-learn/internal/platform/context"
	"github.com/rippinrobr/lunch-n-learn/internal/platform/web"
)

// RequestLogger the format: TraceID : (200) GET /foo -> IP ADDR (latency)
func RequestLogger(next web.Handler) web.Handler {

	// Wrap this handler around the next one provided.
	h := func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
		v := ctx.Value(clContext.KeyValues).(*clContext.Values)

		err := next(ctx, w, r, params)

		log.Printf("%s : (%d) : %s %s -> %s (%s)",
			v.TraceID,
			v.StatusCode,
			r.Method, r.URL.Path,
			r.RemoteAddr, time.Since(v.Now),
		)

		return err
	}

	return h
}
