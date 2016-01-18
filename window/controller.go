package window

import (
	"net/http"
	"text/template"
)

type Options struct {
	Title string
	Dev   bool
}

type WindowController struct {
	options *Options
}

func NewController(opts *Options) *WindowController {
	return &WindowController{opts}
}

func (c *WindowController) Render(res http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./window/window.html")
	if err != nil {
		http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	t.Execute(res, c.options)
}
