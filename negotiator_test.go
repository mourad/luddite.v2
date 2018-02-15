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
	if rw.Header().Get(HeaderContentType) != ContentTypeJson {
		t.Error("default content type not negotiated")
	}
}
