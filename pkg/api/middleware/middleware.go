package middleware

import "net/http"

// SetJSONHeader sets the Header of the response to json type
func SetJSONHeader(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}
