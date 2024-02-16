package web

import (
	"fmt"
	"go-blueprint-htmx/cmd/web/components"
	"net/http"
)

func UpdateStructureHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	name := r.FormValue("projectName")
	selectedDB := r.FormValue("database")
	fmt.Println("selectedDB", selectedDB)
	components.FolderStructure(name, selectedDB).Render(r.Context(), w)
}
