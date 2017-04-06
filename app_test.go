package main

import(
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestRootHandler(t *testing.T) {
  rootHandler := rootHandler()
  req, _ := http.NewRequest("GET", "", nil)
  w := httptest.NewRecorder()
  rootHandler.ServeHTTP(w, req)

  if w.Code != http.StatusOK {
      t.Errorf("Home page didn't return %v", http.StatusOK)
  }

  body := string(w.Body.Bytes())

  expectedBody := "Hello, World!"
  if body != expectedBody {
    t.Errorf("Body was %s, not %s", body, expectedBody)
  }
}
