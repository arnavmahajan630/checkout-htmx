package handler

import "net/http"

func GetUsers(w http.ResponseWriter, r * http.Request) error{
	htmlContent := "<h2>Content Loaded Successfully!</h2>"
	w.Header().Set("Content-Type", "text/html")
	_ , err := w.Write([]byte(htmlContent))
	return err
}