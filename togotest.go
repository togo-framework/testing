// Package togotest is a PHPUnit-style test harness for togo apps: HTTP request
// helpers, fluent response assertions, and an in-memory SQLite database for fast,
// isolated feature tests. Import as `togotest`.
package togotest

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "modernc.org/sqlite"
)

// Response wraps a recorded HTTP response with fluent assertions.
type Response struct {
	Code int
	Body string
	Raw  *httptest.ResponseRecorder
}

// Do sends a request to handler and records the response. body may be nil, a
// string, []byte, or any JSON-serialisable value.
func Do(t *testing.T, handler http.Handler, method, path string, body any) *Response {
	t.Helper()
	var r io.Reader
	switch b := body.(type) {
	case nil:
	case string:
		r = strings.NewReader(b)
	case []byte:
		r = bytes.NewReader(b)
	default:
		buf, err := json.Marshal(b)
		if err != nil {
			t.Fatalf("togotest: marshal body: %v", err)
		}
		r = bytes.NewReader(buf)
	}
	req := httptest.NewRequest(method, path, r)
	if r != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	return &Response{Code: rec.Code, Body: rec.Body.String(), Raw: rec}
}

// Status asserts the HTTP status code.
func (r *Response) Status(t *testing.T, want int) *Response {
	t.Helper()
	if r.Code != want {
		t.Fatalf("expected status %d, got %d — body: %s", want, r.Code, r.Body)
	}
	return r
}

// JSON decodes the body into v.
func (r *Response) JSON(t *testing.T, v any) *Response {
	t.Helper()
	if err := json.Unmarshal([]byte(r.Body), v); err != nil {
		t.Fatalf("togotest: decode JSON: %v — body: %s", err, r.Body)
	}
	return r
}

// Contains asserts the body contains sub.
func (r *Response) Contains(t *testing.T, sub string) *Response {
	t.Helper()
	if !strings.Contains(r.Body, sub) {
		t.Fatalf("expected body to contain %q — body: %s", sub, r.Body)
	}
	return r
}

// SQLite returns an isolated in-memory database, closed automatically when the
// test ends. Apply your schema with Exec.
func SQLite(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("togotest: open sqlite: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })
	return db
}

// Exec runs SQL statements (e.g. schema) against db, failing the test on error.
func Exec(t *testing.T, db *sql.DB, stmts ...string) {
	t.Helper()
	for _, s := range stmts {
		if _, err := db.ExecContext(context.Background(), s); err != nil {
			t.Fatalf("togotest: exec %q: %v", s, err)
		}
	}
}

// Equal asserts two comparable values are equal.
func Equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}
