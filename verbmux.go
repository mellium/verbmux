// Package verbmux provides functionality for a simple HTTP multiplexer that
// routes requests based on the HTTP verb (GET, POST, CONNECT, etc.). It can
// handle OPTIONS requests automatically.
package verbmux

import (
	"net/http"
	"strings"
)

type verbMux map[string]http.Handler

// New constructs a new multiplexer that responds to requests with the provided
// verb's. OPTIONS requests are automatically handled unless they are
// overridden.
func New(v ...verb) http.Handler {
	vm := make(verbMux)
	for _, vfunc := range v {
		vfunc(vm)
	}

	// If options has not been specified, set a default handler.
	if _, ok := vm["OPTIONS"]; !ok {
		vm["OPTIONS"] = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var verbs []string
			for v, _ := range vm {
				verbs = append(verbs, v)
			}
			w.Header().Add("Allow", strings.Join(verbs, ","))
			// TODO: Log or trace errors?
			w.Write(nil)
		})
	}
	return vm
}

// ServeHTTP dispatches the request to the handler that matches its verb. If no
// handler matches it serves a 405 Method Not Allowed.
func (vm verbMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := vm[r.Method]; ok {
		handler.ServeHTTP(w, r)
		return
	}

	// TODO: i18n or ability to serve a custom handler.
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

type verb func(vm verbMux)

// TODO: HEAD, TRACE, CONNECT, PATCH

// Custom handles custom HTTP verbs, for example, if you wanted to handle the
// WebDAV LOCK or UNLOCK verbs.
func Custom(verb string, handler http.Handler) verb {
	return func(vm verbMux) {
		vm[verb] = handler
	}
}

// Get handles HTTP GET requests.
func Get(handler http.Handler) verb {
	return Custom("GET", handler)
}

// Post handles HTTP POST requests.
func Post(handler http.Handler) verb {
	return Custom("POST", handler)
}

// Put handles HTTP PUT requests.
func Put(handler http.Handler) verb {
	return Custom("PUT", handler)
}

// Delete handles HTTP DELETE requests.
func Delete(handler http.Handler) verb {
	return Custom("DELETE", handler)
}

// Options handles HTTP OPTIONS requests. A new verb muxer handles OPTIONS
// requests by default, so an Options verb only needs to be specified if custom
// behavior is desired, or OPTIONS requests should not be handled (by passing a
// nil handler).
func Options(handler http.Handler) verb {
	return Custom("OPTIONS", handler)
}
