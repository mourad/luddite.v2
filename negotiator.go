package luddite

import (
	"net/http"

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
	format, err := negotiation.NegotiateAccept(accept, n.acceptedFormats)
	if err != nil {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	rw.Header().Set(HeaderContentType, format.Value)
}

// RegisterFormat registers a new format and associated MIME types.
func RegisterFormat(format string, mimeTypes []string) {
	negotiation.RegisterFormat(format, mimeTypes)
}
