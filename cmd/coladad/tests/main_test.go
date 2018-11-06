package tests

import (
	"bytes"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rippinrobr/lunch-n-learn/cmd/coladad/config"
	"github.com/rippinrobr/lunch-n-learn/cmd/coladad/handlers"
	lnlContext "github.com/rippinrobr/lunch-n-learn/internal/platform/context"
	"github.com/rippinrobr/lunch-n-learn/internal/platform/web"
)

// FeatureContextFn represents custom FeatureContext funcs that we want
// to be run as part of the test suite.
type FeatureContextFn func(*Feature, *godog.Suite)

var (

	// Values that get bootstrapped once
	app               *web.App
	cfg               config.Config
	ctx               context.Context
	featureContextFns = []FeatureContextFn{}
	godogOpts         = godog.Options{
		Format:    "progress",
		Randomize: -1,
	}
)

func init() {

	// If we're running under normal test mode, let's discard the log output. It's just noise.
	log.SetOutput(ioutil.Discard)

	// Handle various options/flags.
	for _, arg := range os.Args[1:] {
		switch {

		// If `go test -v` was run, switch to use the pretty formatter, and print log output as well.
		case arg == "-test.v=true":
			log.SetOutput(os.Stderr)
			godogOpts.Format = "pretty"

		// If -godog.tags=X was specfied, only run tests with those specified tags.
		case strings.HasPrefix(arg, "-godog.tags"):
			matches := regexp.MustCompile("-godog.tags=(.*)").FindStringSubmatch(arg)
			if len(matches) == 2 {
				godogOpts.Tags = matches[1]
			}
			flag.CommandLine.StringVar(&godogOpts.Tags, "godog.tags", matches[1], "Filter scenarios by tags. Expression can be:\n"+
				"@wip: run all scenarios with wip tag\n"+
				"~@wip: exclude all scenarios with wip tag\n"+
				"@wip && ~@new: run wip scenarios, but exclude new\n"+
				"@wip,@undone: run wip or undone scenarios")
		}

	}

}

func RegisterFeatureContext(fn FeatureContextFn) {
	featureContextFns = append(featureContextFns, fn)
}

// Feature represents a base feature that is under test.
// This is where all of our common tests are bound.
type Feature struct {
	keystore       map[string]string // supplied key : actual entity id
	request        *http.Request
	requestHeaders http.Header
	response       *httptest.ResponseRecorder
	singleResult   map[string]interface{}
}

func (f *Feature) reset() {
	f.keystore = map[string]string{}
	f.requestHeaders = http.Header{}
	f.request = nil
	f.response = nil
}

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	// Reset the logger when we're done.
	defer log.SetOutput(os.Stderr)

	// Setup our test context.
	ctx = context.WithValue(context.Background(), lnlContext.KeyValues, &lnlContext.Values{
		TraceID: uuid.New().String(),
		Now:     time.Now(),
	})

	// Configure our logger format.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	// ============================================================
	// get db connection
	cfg := config.New()
	coladaDB, err := sql.Open(cfg.DBType, cfg.DBConnInfo)
	if err != nil {
		log.Fatalf("startup : Register DB : %v", err)
	}

	// Initialize our app.
	app = handlers.API(coladaDB, cfg).(*web.App)

	// Run GoDog.
	status := godog.RunWithOptions("", FeatureContext, godogOpts)
	if st := m.Run(); st > status {
		status = st
	}
	return status
}

func FeatureContext(s *godog.Suite) {
	f := &Feature{
		keystore:       map[string]string{},
		requestHeaders: http.Header{},
	}

	s.Step(`^a request is made to "([^"]*)" "([^"]*)"$`, f.makeRequest)
	s.Step(`^the response should be a (\d+) status code$`, f.theResponseShouldBeAStatusCode)
	s.Step(`^the response body should contain a result with( at [A-Za-z]*)? (\d+) item(s?)`, f.theResponseBodyShouldContainAResultWithAtLeastItem)
	s.Step(`^the response body should return a single item$`, f.theResponseBodyShouldReturnASingleItem)
	s.Step(`^the response body should contain the field "([^"]*)"$`, f.theResopnseBodyShouldContainTheField)
}

func (f *Feature) makeRequest(method, url string) error {
	body := ""

	r, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed while creating a new request to %s %s: %s", method, url, err.Error())
	}

	r.Header.Set("Content-Type", "application/json")
	f.request = r
	w := httptest.NewRecorder()
	// Make the request against our app.
	app.ServeHTTP(w, r)
	f.response = w

	return nil
}

func (f *Feature) theResponseShouldBeAStatusCode(statusCode int) error {
	// Make sure we have a response to even check.
	if f.response == nil {
		return fmt.Errorf("no response was found- was a request made in a previous step?")
	}

	// Return an error if the response codes don't match.
	if f.response.Code != statusCode {
		return fmt.Errorf("response code should be %d, got %d", statusCode, f.response.Code)
	}

	return nil
}

func (f *Feature) theResponseBodyShouldContainAResultWithAtLeastItem(comparator string, n int) error {
	// Make sure we have a request/response to even check.
	if f.request == nil {
		return fmt.Errorf("no request was found- was one made in a previous step?")
	}
	if f.response == nil {
		return fmt.Errorf("no response was found- was a request made in a previous step?")
	}

	// Figure out the compareFn we need to use for testing results.
	type compareFn func(length, expected int) bool
	var lengthMatches compareFn
	switch strings.Trim(comparator, " ") {
	case "at least":
		lengthMatches = func(length, expected int) bool {
			return length >= expected
		}
	case "at most":
		lengthMatches = func(length, expected int) bool {
			return length <= expected
		}
	default:
		lengthMatches = func(length, expected int) bool {
			return length == expected
		}
	}

	// Read the response body.
	buffer, _ := ioutil.ReadAll(f.response.Body)
	f.response.Body = bytes.NewBuffer(buffer) // Reset the body on the response so any subsequent steps can read it too.
	responseBody := bytes.NewBuffer(buffer)

	result := struct {
		Results []interface{} `json:"result"`
	}{}
	if err := web.Unmarshal(responseBody, &result); err != nil {
		return fmt.Errorf("failed while unmarshalling JSON into list result: %s", err.Error())
	}

	if !lengthMatches(len(result.Results), n) {
		return fmt.Errorf("number of items returned should be >= %d, got %d", n, len(result.Results))
	}

	return nil
}

func (f *Feature) theResponseBodyShouldReturnASingleItem() error {
	// Make sure we have a request/response to even check.
	if f.request == nil {
		return fmt.Errorf("no request was found- was one made in a previous step?")
	}
	if f.response == nil {
		return fmt.Errorf("no response was found- was a request made in a previous step?")
	}

	// Read the response body.
	buffer, _ := ioutil.ReadAll(f.response.Body)
	f.response.Body = bytes.NewBuffer(buffer) // Reset the body on the response so any subsequent steps can read it too.
	responseBody := bytes.NewBuffer(buffer)

	result := struct {
		Results map[string]interface{} `json:"result"`
	}{}
	if err := web.Unmarshal(responseBody, &result); err != nil {
		return fmt.Errorf("failed while unmarshalling JSON into a single object: %s", err.Error())
	}

	f.singleResult = result.Results
	return nil
}

func (f *Feature) theResopnseBodyShouldContainTheField(key string) error {
	if _, exists := f.singleResult[key]; !exists {
		return fmt.Errorf("the key '%s' was not found in the returned item", key)
	}

	return nil
}
