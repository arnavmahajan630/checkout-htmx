//go:build prod
package main
import (
	"embed"
	"net/http"
)

//go:embed public
var publicFs embed.FS

func public() http.Handler {
	return http.FileServer(http.FS(publicFs))
}
