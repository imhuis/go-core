package context_pkg

import (
	"net/http"
	"net/http/httptest"
)

func SlowServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//time.Sleep(1 * time.Second)
		w.Write([]byte("slow Response!"))
	}))
	return s
}

func FastServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("error") != "true" {
			w.Write([]byte("error"))
			return
		}
	}))
	return s
}
