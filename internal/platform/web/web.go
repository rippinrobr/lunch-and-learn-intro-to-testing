package web

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/dimfeld/httptreemux"
	"github.com/pborman/uuid"
	clContext "github.com/rippinrobr/lunch-n-learn/internal/platform/context"
)

// A Handler is a type that handles an http request within our own little mini
// framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error

type App struct {
	*httptreemux.TreeMux
	mw []Middleware
}

// New creates an App value that handle a set of routes for the application.
func New(mw ...Middleware) *App {
	return &App{
		TreeMux: httptreemux.New(),
		mw:      mw,
	}
}

// Handle is our mechanism for mounting Handlers for a given HTTP verb and path
// pair, this makes for really easy, convenient routing.
func (a *App) Handle(verb, path string, handler Handler, mw ...Middleware) {

	// Wrap up the application-wide first, this will call the first function
	// of each middleware which will return a function of type Handler.
	handler = wrapMiddleware(wrapMiddleware(handler, mw), a.mw)

	// The function to execute for each request.
	h := func(w http.ResponseWriter, r *http.Request, params map[string]string) {

		// Set the context with the required values to
		// process the request.
		v := clContext.Values{
			TraceID: generateTraceID(r),
			Now:     time.Now(),
		}
		ctx := context.WithValue(r.Context(), clContext.KeyValues, &v)

		// Set the trace id on the outgoing requests before any other header to
		// ensure that the trace id is ALWAYS added to the request regardless of
		// any error occurring or not.
		w.Header().Set(clContext.TraceIDHeader, v.TraceID)

		// Allow CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Call the wrapped handler functions.
		err := handler(ctx, w, r, params)
		if err != nil {
			log.Printf("%s : ERROR : handler error: %s",
				v.TraceID,
				err,
			)
		}
	}

	// Add this handler for the specified verb and route.
	a.TreeMux.Handle(verb, path, h)
}

// generateTraceID generates a new Trace ID UUID
// if the supplied request doesn't already have one specified.
func generateTraceID(r *http.Request) string {
	if traceHeader := r.Header.Get(clContext.TraceIDHeader); traceHeader != "" {
		return traceHeader
	}
	return uuid.New()
}
