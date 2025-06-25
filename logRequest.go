package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Middleware to log method and body
func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log method
		fmt.Println("Request Method:", r.Method)

		// Read body (and allow it to be read again later)
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading body:", err)
		} else {
			bodyStr := string(bodyBytes)
			fmt.Println("Request Body:", bodyStr)

			// Restore the body for downstream handlers
			r.Body = io.NopCloser(strings.NewReader(bodyStr))
		}

		next.ServeHTTP(w, r)
	})
}