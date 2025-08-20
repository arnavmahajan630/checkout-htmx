package handler

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

type HTTPHandler func(w http.ResponseWriter, r * http.Request) error
func Make(g HTTPHandler)http.HandlerFunc {
		return func(w http.ResponseWriter, r * http.Request) {
			if err := g(w, r); err != nil {slog.Error("HTTP Handler Error", "err", err, "path", r.URL.Path)}
		}
}

func Render(w http.ResponseWriter, r * http.Request, c  templ.Component) error{
	return c.Render(r.Context(), w)
}







// type 2 for any returns

/* type HTTPHandler func(r *http.Request) (any, error)

func Make(h HTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := h(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) // or custom error mapping
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(data)
	}
}
 */
