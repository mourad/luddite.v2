package luddite

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDefaultContentType(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	rw := httptest.NewRecorder()

	n := newNegotiatorHandler([]string{ContentTypeJson, ContentTypeXml})
	n.ServeHTTP(rw, req)

	if res := rw.Result(); res != nil && res.StatusCode != http.StatusOK {
		t.Error("result was written")
	}
	if rw.Header().Get(HeaderContentType) != ContentTypeJson {
		t.Error("default content type not negotiated")
	}
}

func TestSupportedContentType(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set(HeaderAccept, ContentTypeJson)
	rw := httptest.NewRecorder()

	n := newNegotiatorHandler([]string{ContentTypeJson, ContentTypeXml})
	n.ServeHTTP(rw, req)

	if res := rw.Result(); res != nil && res.StatusCode != http.StatusOK {
		t.Error("result was written")
	}
	if ct := rw.Header().Get(HeaderContentType); ct != ContentTypeJson {
		t.Errorf("incorrrect content type negotiated: %s", ct)
	}
}

func TestUnsupportedContentType(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set(HeaderAccept, ContentTypeCsv)
	rw := httptest.NewRecorder()

	n := newNegotiatorHandler([]string{ContentTypeJson, ContentTypeXml})
	n.ServeHTTP(rw, req)

	if res := rw.Result(); res != nil && res.StatusCode != http.StatusOK {
		t.Error("result was written")
	}
	if ct, ok := rw.Header()[HeaderContentType]; ok {
		t.Errorf("incorrect content type negotiated: %s", ct)
	}
}
