package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func ParseTemplateFiles(w http.ResponseWriter, templateName string, context any, funcToTemplate template.FuncMap, files ...string) {
	ts, er := template.ParseFiles(files...)

	if er != nil {
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		fmt.Println("Error parsing template", er)
		return
	}

	if funcToTemplate != nil {
		ts.Funcs(funcToTemplate)
	}

	err := ts.ExecuteTemplate(w, templateName, context)
	if err != nil {
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		fmt.Println("Error Executing template", err)
		return
	}
}
func TransformationToTemplateFuncMap(key string, f interface{}) template.FuncMap {
	return template.FuncMap{
		key: f,
	}
}
