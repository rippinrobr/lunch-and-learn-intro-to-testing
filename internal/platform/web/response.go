package web

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/pkg/errors"

	clContext "github.com/rippinrobr/lunch-n-learn/internal/platform/context"
)

// These predefined errors cover most of the common error cases that clients
// might encounter. Custom errors can be created easily with the AppError type.
var (
	// ErrInvalidMethod occurs if you get the wrong method
	ErrInvalidMethod = AppError{Message: "invalid method", Status: http.StatusBadRequest, Code: "BAD_REQUEST"}

	// ErrNotFound is abstracting the mgo not found error.
	ErrNotFound = AppError{Message: "Sorry, we couldn't find what you asked for", Status: http.StatusNotFound, Code: "NOT_FOUND"}

	// ErrValidation occurs when there are validation errors.
	ErrValidation = AppError{Message: "One or more fields contain invalid values", Status: http.StatusUnprocessableEntity, Code: "INVALID_INPUT"}

	// ErrBadRequest occurs when the request has the wrong format.
	ErrBadRequest = AppError{Message: "The request wasn't properly formatted", Status: http.StatusBadRequest, Code: "BAD_REQUEST"}

	// ErrInternal occurs when there is an internal server error.
	ErrInternal = AppError{Message: "Unable to process request", Status: http.StatusInternalServerError, Code: "INTERNAL_ERROR"}
)

// AppError is an application error. Its message and error code should be
// displayed to the client and served with a particular HTTP status code.
type AppError struct {
	Message string `json:"message"` // Sorry, we couldn't find what you asked for
	Code    string `json:"code"`    // NOT_FOUND
	Status  int    `json:"status"`  // 404
}

// Error implements the error interface
func (e AppError) Error() string {
	return e.Message
}

// JSONError is the response for errors that occur within the API.
type JSONError struct {
	Error  string       `json:"error"`
	Fields InvalidError `json:"fields,omitempty"`
}

// Response represents the payload response structure in the API
type Response struct {
	Result interface{} `json:"result,omitempty"`
	Errors []error     `json:"errors,omitempty"`
}

// Error handles all error responses for the API.
func Error(ctx context.Context, w http.ResponseWriter, err error) {
	switch e := errors.Cause(err).(type) {
	case AppError:
		RespondError(ctx, w, e, e.Status)

	case InvalidError:
		v := JSONError{
			Error:  "There were problems with the URL or JSON data you submitted",
			Fields: e,
		}

		RespondRaw(ctx, w, v, http.StatusUnprocessableEntity)

	default:
		RespondError(ctx, w, ErrInternal, http.StatusInternalServerError)
	}
}

// RespondError sends JSON describing the error
func RespondError(ctx context.Context, w http.ResponseWriter, err error, code int) {
	RespondRaw(ctx, w, JSONError{Error: err.Error()}, code)
}

// RespondRaw sends JSON to the client.
func RespondRaw(ctx context.Context, w http.ResponseWriter, data interface{}, code int) {

	// Set the status code for the request logger middleware.
	v := ctx.Value(clContext.KeyValues).(*clContext.Values)
	v.StatusCode = code

	// Just set the status code and we are done.
	if code == http.StatusNoContent {
		w.WriteHeader(code)
		return
	}

	// Set the content type.
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response and context.
	w.WriteHeader(code)

	// Marshal the data into a JSON string.
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("%s : ERROR : marshalling error: %s\n", v.TraceID, err)
		jsonData = []byte("{}")
	}

	// Send the result back to the client.
	_, err = w.Write(jsonData)
	if err != nil {
		log.Printf("%s : ERROR : writing response error: %s\n", v.TraceID, err)
	}
}

// Respond sends JSON to the client wrapping the response data.
func Respond(ctx context.Context, w http.ResponseWriter, data interface{}, code int, errs ...error) {
	payload := Response{
		Result: data,
	}
	if len(errs) > 0 {
		for _, err := range errs {
			payload.Errors = append(payload.Errors, errors.Cause(err))
		}
	}

	RespondRaw(ctx, w, payload, code)
}
