package web

import (
	"go-blueprint-htmx/cmd/web/components"
	"net/http"
)

func UpdateStructureHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	advancedOptions, ok := r.Form["advancedOptions"]
	if !ok {
		// Handle the case where no checkbox was checked
		advancedOptions = []string{}
	}

	options := components.OptionsStruct{
		SelectedDB:        r.FormValue("database"),
		SelectedFramework: r.FormValue("framework"),
		ProjectName:       r.FormValue("projectName"),
		AdvancedOptions:   advancedOptions,
	}
	commandStr := components.GetCommandString(options)

	components.FolderStructure(options, commandStr).Render(r.Context(), w)
}
