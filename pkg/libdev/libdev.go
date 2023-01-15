package libdev

import (
	"html/template"
	"net/http"

	"github.com/webmachine-dev/lib-dev/pkg/web"
)

type App struct{}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := web.ParsePath(r.URL.Path)
	if len(path) == 0 {
		app.home(w, r)
		return
	}
	if len(path) == 1 {
		tmpl := `<!DOCTYPE html>
		<html>
			<head>
				<meta name="go-import" content="lib.dev/{{ .ID }} git https://github.com/library-development/go-{{ .ID }}" />
			</head>
		</html>`
		t := template.Must(template.New("index").Parse(tmpl))
		t.Execute(w, struct{ ID string }{path[0]})
		return
	}
	http.NotFound(w, r)
}

func (app *App) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("lib.dev!"))
}
