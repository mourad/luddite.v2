package luddite

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dimfeld/httptreemux"
)

func TestSchemaHandlerGivenInvalidVersionStringLength(t *testing.T) {
	v := make(map[string]string)
	v["version"] = "v"
	ctx := httptreemux.AddParamsToContext(context.Background(), v)
	req, _ := http.NewRequest("GET", "/", nil)
	req = req.WithContext(ctx)
	rw := httptest.NewRecorder()
	rw.Header().Set(HeaderContentType, ContentTypeJson)

	s := newSchemaHandler(nil)
	s.ServeHTTP(rw, req)

	if rw.Code != http.StatusNotFound {
		t.Error("expected 400/Not found")
	}
}

func TestSchemaHandlerGivenInvalidVersionValue(t *testing.T) {
	v := make(map[string]string)
	v["version"] = "w1"
	ctx := httptreemux.AddParamsToContext(context.Background(), v)
	req, _ := http.NewRequest("GET", "/", nil)
	req = req.WithContext(ctx)
	rw := httptest.NewRecorder()
	rw.Header().Set(HeaderContentType, ContentTypeJson)

	s := newSchemaHandler(nil)
	s.ServeHTTP(rw, req)

	if rw.Code != http.StatusNotFound {
		t.Error("expected 400/Not found")
	}
}

func TestSchemaHandlerGivenInvalidVersionNumber(t *testing.T) {
	v := make(map[string]string)
	v["version"] = "v0"
	ctx := httptreemux.AddParamsToContext(context.Background(), v)
	req, _ := http.NewRequest("GET", "/", nil)
	req = req.WithContext(ctx)
	rw := httptest.NewRecorder()
	rw.Header().Set(HeaderContentType, ContentTypeJson)

	s := newSchemaHandler(nil)
	s.ServeHTTP(rw, req)

	if rw.Code != http.StatusNotFound {
		t.Error("expected 400/Not found")
	}
}
