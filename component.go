package main

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"
)

type Component struct {
	State interface{}
	Body  string
}

func (c Component) Render(params interface{}) string {
	var b bytes.Buffer
	tmpl, err := template.New("component").Parse(c.Body)
	if err != nil {
		return fmt.Sprintf("<p> ERROR %s </p>", err)
	}

	err = tmpl.Execute(&b, params)
	if err != nil {
		return fmt.Sprintf("<p> ERROR %s </p>", err)
	}

	// err = tmpl.Execute(&b, c.State)
	// if err != nil {
	// 	return fmt.Sprintf("<p> ERROR %s </p>", err)
	// }

	return b.String()
}

var counterViewer = Component{
	Body: `<h1>{{.Counter}}</h1>`,
}

var pagecounter = Component{
	State: struct{ CounterViewer string }{
		CounterViewer: counterViewer.Render(struct{ Counter int }{Counter: 43}),
	},
	Body: `<h1>Counter:</h1>
			{{.CounterViewer}}`,
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render := pagecounter.Render(pagecounter.State)
		fmt.Fprint(w, render)
	})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
	}
}
