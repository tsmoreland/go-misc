package middleware

import "net/http"

// OwaspRecommendedApiHeaders adds recommended headers to response
func OwaspRecommendedApiHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)

		addHeaderIfNotPresent(w, "Cache-Control", "no-store")
		addHeaderIfNotPresent(w, "X-Content-Type-Options", "nosniff")
		addHeaderIfNotPresent(w, "X-Frame-Options", "DENY")
		addHeaderIfNotPresent(w, "x-xss-protection", "1; mode=block")
		addHeaderIfNotPresent(w, "Expect-CT", "max-age=0, enforce")
		addHeaderIfNotPresent(w, "referrer-policy", "strict-origin-when-cross-origin")
		addHeaderIfNotPresent(w, "X-Permitted-Cross-Domain-Policies", "none")
		addHeaderIfNotPresent(w, "Content-Security-Policy",
			"default-src: 'none'; FeaturePolicy: 'none'; Referrer-Policy: no-referrer")
	})
}

func addHeaderIfNotPresent(w http.ResponseWriter, header string, value string) {
	if existingValue := w.Header().Get(header); existingValue == "" {
		w.Header().Set(header, value)
	}
}
