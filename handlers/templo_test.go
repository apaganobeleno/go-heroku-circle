package handlers

import (
	"net/http"
	"testing"

	"github.com/bmizerany/pat"
)

func TestTemplo(t *testing.T) {
	// router := muxFor("/hi/:name", Templo)
	// request, _ := http.NewRequest("GET", "/hi/Tony", nil)
	// response := httptest.NewRecorder()
	//
	// router.ServeHTTP(response, request)
	// assert.Contains(t, string(response.Body.Bytes()), "Hello Tony")
}

func muxFor(path string, handler http.HandlerFunc) *pat.PatternServeMux {
	router := pat.New()
	router.Get(path, http.HandlerFunc(handler))
	return router
}
