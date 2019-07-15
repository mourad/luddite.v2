package luddite

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dimfeld/httptreemux"
	"golang.org/x/tools/godoc/vfs/httpfs"
	"golang.org/x/tools/godoc/vfs/mapfs"
)

const (
	sampleYamlBody = "-id: \"1234\"\nname:\"dave\"\nflag:\"true\"\ndata:\"SGVsbG8gd29ybGQ=\"\ntimestamp:\"2015-03-18T14:30:00Z\""
)

func TestSchemaHandlerDefaultContentType(t *testing.T) {
	fakeFS := httpfs.New(mapfs.New(map[string]string{
		"v1/schema.json": sampleJsonBody,
	}))
	v := make(map[string]string)
	v["version"] = "v1"
	v["filepath"] = "schema.json"

	ctx := httptreemux.AddParamsToContext(context.Background(), v)
	req, _ := http.NewRequest("GET", "/", nil)
	req = req.WithContext(ctx)
	rw := httptest.NewRecorder()
	rw.Header().Set(HeaderContentType, ContentTypeJson)

	s := newSchemaHandler(fakeFS)
	s.ServeHTTP(rw, req)
	if ct := rw.Header().Get(HeaderContentType); ct != ContentTypeJson {
		t.Errorf("incorrrect content type negotiated: %s", ct)
	}

	if body := string(rw.Body.String()); body != sampleJsonBody {
		t.Errorf("JSON serialization failed, got: %s, expected: %s\n", body, sampleJsonBody)
	}
}

func TestSchemaHandlerOctetStreamContentType(t *testing.T) {
	fakeFS := httpfs.New(mapfs.New(map[string]string{
		"v1/schema.yml": sampleYamlBody,
	}))
	v := make(map[string]string)
	v["version"] = "v1"
	v["filepath"] = "schema.yml"

	ctx := httptreemux.AddParamsToContext(context.Background(), v)
	req, _ := http.NewRequest("GET", "/", nil)
	req = req.WithContext(ctx)
	rw := httptest.NewRecorder()
	rw.Header().Set(HeaderContentType, ContentTypeJson)

	s := newSchemaHandler(fakeFS)
	s.ServeHTTP(rw, req)

	if ct := rw.Header().Get(HeaderContentType); ct != ContentTypeOctetStream {
		t.Errorf("incorrrect content type negotiated: %s", ct)
	}

	if body := string(rw.Body.String()); body != sampleYamlBody {
		t.Errorf("YAML serialization failed, got: %s, expected: %s\n", body, sampleYamlBody)
	}
}

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
