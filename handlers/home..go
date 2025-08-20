package handler

import (
	"net/http"

	"github.com/arnavmahajan630/checkout-htmx/views/home"
)

func Home(w http.ResponseWriter, r * http.Request) error {
	return Render(w, r, home.Index())
}
