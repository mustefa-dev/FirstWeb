package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir   = "views/layouts/"
	TemplateExt = ".gohtml"
)

func layoutFiles() []string {
	files, err := filepath.Glob("views/layouts/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}

func NewView(layout string, files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml", "views/layouts/navbar.gohtml", "views/layouts/bootstrap.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}

}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

// Render to render the view to predefined layout
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	// so we don't have to write content type every time
	w.Header().Set("Content-Type", "text/html")
	//so
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// 53 lines, before using static.go
