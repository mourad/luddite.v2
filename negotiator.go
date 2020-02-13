package luddite

import (
	"net/http"
	"strconv"

	"github.com/K-Phoen/negotiation"
)

type negotiator struct {
	acceptedFormats []string
}

func newNegotiatorHandler(acceptedFormats []string) http.Handler {
	return &negotiator{
		acceptedFormats: acceptedFormats,
	}
}

func (n *negotiator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// If no Accept header was included, default to the first accepted format
	accept := req.Header.Get(HeaderAccept)
	if accept == "" {
		accept = n.acceptedFormats[0]
	}

	// Negotiate and set a Content-Type
	//
	// Note: Negotation failures do not return 406 errors here. This allows
	// resource handlers to potentially inspect/handle certain rarely-used
	// content types on their own. If a negotiation failure has occurred and
	// the resource handler doesn't deal with it, then we can expect a 406
	// from WriteResponse.
	if format, err := negotiation.NegotiateAccept(accept, n.acceptedFormats); err == nil {
		rw.Header().Set(HeaderContentType, format.Value)
	}

	// If the X-Spirent-Inhibit-Response header is set and true-ish, then
	// set the same response header. This will cause subsequent calls to
	// WriteResponse to omit the response body for 2xx responses and also
	// makes the behavior obvious to clients (i.e. response header shows
	// intention beyond the 204 status).
	if inhibitResp, _ := strconv.ParseBool(req.Header.Get(HeaderSpirentInhibitResponse)); inhibitResp {
		rw.Header().Set(HeaderSpirentInhibitResponse, "1")
	}
}

// RegisterFormat registers a new format and associated MIME types.
func RegisterFormat(format string, mimeTypes []string) {
	negotiation.RegisterFormat(format, mimeTypes)
}
