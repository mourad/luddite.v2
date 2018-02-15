package luddite

import (
	"net/http"
	"strconv"
)

type version struct {
	minVersion int
	maxVersion int
}

func newVersionHandler(minVersion, maxVersion int) http.Handler {
	return &version{
		minVersion: minVersion,
		maxVersion: maxVersion,
	}
}

func (v *version) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Parse the client's requested API version
	version := v.maxVersion
	if s := req.Header.Get(HeaderSpirentApiVersion); s != "" {
		i, err := strconv.Atoi(s)
		if err != nil || i < 1 {
			e := NewError(nil, EcodeApiVersionInvalid)
			_ = WriteResponse(rw, http.StatusBadRequest, e)
			return
		}
		version = i
	}

	// Range check the requested API version and reject requests that fall outside supported version numbers
	if version < v.minVersion {
		e := NewError(nil, EcodeApiVersionTooOld, v.minVersion)
		_ = WriteResponse(rw, http.StatusGone, e)
		return
	}
	if version > v.maxVersion {
		e := NewError(nil, EcodeApiVersionTooNew, v.maxVersion)
		_ = WriteResponse(rw, http.StatusNotImplemented, e)
		return
	}

	// Add the requested API version to response headers (useful for clients when a default version was negotiated)
	rw.Header().Add(HeaderSpirentApiVersion, strconv.Itoa(version))

	// Add the requested API version to handler context so that downstream handlers can access
	d := contextHandlerDetails(req.Context())
	d.apiVersion = version
}
