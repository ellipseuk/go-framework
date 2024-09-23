package user

import (
	"fmt"
	"net/http"
	"regexp"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := extractParam(r, "id")
	fmt.Fprintf(w, "User ID: %s", id)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Creating a new user...")
}

func extractParam(r *http.Request, name string) string {
	re := regexp.MustCompile("^/user/([^/]+)id$")
	match := re.FindStringSubmatch(r.URL.Path)
	for i, group := range re.SubexpNames() {
		if group == name {
			return match[i]
		}
	}
	return ""
}
