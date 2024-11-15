package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var t = template.Must(template.New("index.html").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
</head>
<body>
<h1>Header</h1>
<ul>
  {{range $key, $value := .Header}}
  <li> {{ $key }}: {{ $value }} </li>
  {{end}}
</ul>

<h1>Hello from Fly</h1>
This is stuff here.
{{ if .Region }}
<h2>I'm running in the {{.Region}} region</h2>
{{end}}

</body>
</html>
`))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"Region": os.Getenv("FLY_REGION"),
			"Header": r.Header,
		}

		t.ExecuteTemplate(w, "index.html", data)
	})
	log.Println("listening")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
