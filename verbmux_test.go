// Copyright 2017 Soquee.
// Use of this source code is governed by the BSD 2-clause license that can be
// found in the LICENSE file.

package verbmux

import (
	"net/http"
	"testing"
)

type mockHandler struct{}

func (mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func TestNew(t *testing.T) {
	// Check that New sets a default OPTIONS handler.
	t.Run("DefaultOptions", func(t *testing.T) {
		vm := New()
		if _, ok := vm.(verbMux)["OPTIONS"]; !ok {
			t.Error("No default OPTIONS handler found")
		}
	})

	// Check that passing in a n OPTIONS handler overrides the default.
	t.Run("OverrideOptions", func(t *testing.T) {
		vm := New(Options(mockHandler{}))
		opts, ok := vm.(verbMux)["OPTIONS"]
		if !ok {
			t.Error("No default OPTIONS handler found")
		}
		if _, ok := opts.(mockHandler); !ok {
			t.Error("Options handler did not override default")
		}
	})
}
