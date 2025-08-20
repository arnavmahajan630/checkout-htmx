//go:build !prod
package  main
import (
	"log/slog"
	"net/http"
)

func public() http.Handler {
	slog.Info("Getting static files Ready")
	return http.StripPrefix("/public/", http.FileServer(http.Dir("public")))
}
