package mainDomain

import (
	"GoBaby/internals/models"
	"GoBaby/internals/utils"
	"net/http"
)

func ClockFragment(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.InstanceRoutes.CLOCK)

	files := []string{
		"ui/html/pages/main/clock.tmpl.htmml",
	}
	utils.ParseTemplateFiles(w, "clock", utils.EmptyStruct, utils.EmptyTemplate, files...)
}
