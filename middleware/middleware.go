package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/golang-web/model"
)

// TokenMiddleware memeriksa token dalam header Authorization
func TokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("token")
		if authHeader != "12345" {
			badResponse := model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized",
				Data:       nil,
			}
			json.NewEncoder(w).Encode(badResponse)
			return
		}

		// Melanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)
	})
	// return func(w http.ResponseWriter, r *http.Request) {
	// 	const validToken = "j3fr1"

	// 	token := r.Header.Get("Authorization")

	// 	if token != validToken {
	// 		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
	// 		return
	// 	}

	// 	next(w, r)
	// }
}
