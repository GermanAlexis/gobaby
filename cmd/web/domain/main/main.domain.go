package mainDomain

import (
	"GoBaby/internals/models"
	"GoBaby/internals/utils"
	"net/http"
)

func MainView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.InstanceRoutes.MAIN)

	files := []string{
		"ui/html/pages/main/main.tmpl.htmml",
		"ui/html/base.htmml",
	}
	utils.ParseTemplateFiles(w, "base", utils.EmptyStruct, utils.EmptyTemplate, files...)
}
