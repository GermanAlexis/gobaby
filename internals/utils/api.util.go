package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func CheckIfPath(w http.ResponseWriter, r *http.Request, path string) bool {
	if r.URL.Path != path {
		fmt.Println("Error" + r.URL.Path)
		http.NotFound(w, r)
		return false // Indica que la ruta es incorrecta
	}
	return true // Indica que la ruta es correcta
}

func IsValidMethod(method string, acceptMethod string, w http.ResponseWriter) bool {
	if method == acceptMethod {
		w.Header().Set("Allow", acceptMethod)
		return true
	}
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	return false
}

var (
	EmptyStruct   = struct{}{}
	EmptyTemplate = template.FuncMap{}
)
