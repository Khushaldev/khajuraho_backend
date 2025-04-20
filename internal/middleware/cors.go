package middleware

import "net/http"

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")

		// TODO: remove comments later
		// if origin == "" {
		// 	w.Header().Set("Access-Control-Allow-Origin", origin)
		// 	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		// 	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		// }
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
