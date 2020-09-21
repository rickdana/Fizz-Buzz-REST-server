package basicauthmiddleware

import (
	"net/http"
	"sort"
)

type BasicAuthMiddleware struct {
	username     string
	password     string
	excludedPath []string
}

func NewBasicAuthMiddleware(username string, password string, excludedPath []string) *BasicAuthMiddleware {
	return &BasicAuthMiddleware{username: username, password: password, excludedPath: excludedPath}
}

func (baw *BasicAuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if contains(baw.excludedPath, r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		user, pass, ok := r.BasicAuth()
		if !ok || !baw.checkUsernameAndPassword(user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
			w.WriteHeader(401)
			http.Error(w, "Your are not authorised to access this resource", http.StatusForbidden)
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func (baw *BasicAuthMiddleware) checkUsernameAndPassword(username, password string) bool {
	return username == baw.username && password == baw.password
}

func contains(s []string, searchTerm string) bool {
	i := sort.SearchStrings(s, searchTerm)
	return i < len(s) && s[i] == searchTerm
}
