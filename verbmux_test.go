// Copyright 2017 The Mellium Contributors.
// Use of this source code is governed by the BSD 2-clause
// license that can be found in the LICENSE file.

package verbmux_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"mellium.im/verbmux"
)

type testCase struct {
	mux  http.Handler
	verb string
	resp int
	body string
}

const helloWorldBody = "<html><body>Hello World!</body></html>"

var helloWorld = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, helloWorldBody)
})

var testCases = [...]testCase{
	0: {verbmux.New(), "GET", 405, "Method not allowed"},
	1: {verbmux.New(), "OPTIONS", 200, ""},
	2: {verbmux.New(verbmux.Options(nil)), "OPTIONS", 405, "Method not allowed"},
	3: {verbmux.New(verbmux.Options(helloWorld)), "OPTIONS", 200, helloWorldBody},
	4: {verbmux.New(verbmux.Get(helloWorld)), "GET", 200, helloWorldBody},
	5: {verbmux.New(verbmux.Post(helloWorld)), "POST", 200, helloWorldBody},
	6: {verbmux.New(verbmux.Put(helloWorld)), "PUT", 200, helloWorldBody},
	7: {verbmux.New(verbmux.Delete(helloWorld)), "DELETE", 200, helloWorldBody},
}

func TestVerbMux(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			req := httptest.NewRequest(tc.verb, "http://example.com/foo", nil)
			w := httptest.NewRecorder()
			tc.mux.ServeHTTP(w, req)

			resp := w.Result()
			body, _ := ioutil.ReadAll(resp.Body)

			if resp.StatusCode != tc.resp {
				t.Errorf("Got invalid status code: want=%d, got=%d", tc.resp, resp.StatusCode)
			}
			if strings.TrimSpace(string(body)) != tc.body {
				t.Errorf("Got invalid body: want=`%s`, got=`%s`", tc.body, body)
			}
		})
	}
}
