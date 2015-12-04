/*
  Utility functions to be used in main (trying to keep things clean)
*/
package util

import (
    "net/http"
    "log"
    "time"
)

// HTTP logger
func Log(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    start := time.Now()
    handler.ServeHTTP(res, req)
    end := time.Since(start)
    log.Printf("%s %s %s %s", req.Host, req.URL, req.Method, end)

  })
}
