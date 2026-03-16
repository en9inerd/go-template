package server

// Custom middleware example:
//
//	func MyMiddleware(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			// before handler
//			next.ServeHTTP(w, r)
//			// after handler
//		})
//	}
