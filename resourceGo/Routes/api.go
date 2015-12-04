/*
   API Routes that will be used by the source of the application.
   Creates a ServeMux from the default http package
*/
package routes

import (
    "fmt"
    "net/http"
)

// Routes with /api/ prefix
var ApiMux = http.NewServeMux()

// Setup Routes with Mux
func init() {
  ApiMux.HandleFunc("/api/test", test)
  ApiMux.HandleFunc("/api/other", other)
}


// Handle /api/test
func test(res http.ResponseWriter, req *http.Request) {
  if req.Method != "GET" {
    http.Error(res, http.StatusText(405), 405)
    return
  }
  fmt.Fprintf(res, "Test Passed!")
}

// Handle /api/test
func other(res http.ResponseWriter, req *http.Request) {
  if req.Method != "GET" {
    http.Error(res, http.StatusText(405), 405)
    return
  }
  fmt.Fprintf(res, "Other Test Passed!")
}
