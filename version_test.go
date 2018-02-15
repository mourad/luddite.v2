package luddite

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMinApiVersionConstraint(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add(HeaderSpirentApiVersion, "1")
	rw := httptest.NewRecorder()
	rw.Header().Set(HeaderContentType, ContentTypeJson)

	v := newVersionHandler(2, 42)
	v.ServeHTTP(rw, req)
	if rw.Code != http.StatusGone {
		t.Error("expected 410/Gone response for outdated version")
	}
}

func TestMaxApiVersionConstraint(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add(HeaderSpirentApiVersion, "43")
	rw := httptest.NewRecorder()
	rw.Header().Set(HeaderContentType, ContentTypeJson)

	v := newVersionHandler(2, 42)
	v.ServeHTTP(rw, req)
	if rw.Code != http.StatusNotImplemented {
		t.Error("expected 501/Not Implemented response for future version")
	}
}

func TestApiVersionContext(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add(HeaderSpirentApiVersion, "1")
	req = req.WithContext(withHandlerDetails(req.Context(), &handlerDetails{}))
	rw := httptest.NewRecorder()

	v := newVersionHandler(1, 1)
	v.ServeHTTP(rw, req)
	if ContextApiVersion(req.Context()) != 1 {
		t.Error("missing API version in request context")
	}
	if _, ok := rw.HeaderMap[HeaderSpirentApiVersion]; !ok {
		t.Errorf("missing %s header in response", HeaderSpirentApiVersion)
	}
}
