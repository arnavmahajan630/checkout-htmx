package handler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/arnavmahajan630/checkout-htmx/models"
	"github.com/arnavmahajan630/checkout-htmx/views/user"
)

func GetUsers(w http.ResponseWriter, r * http.Request) error{
	slog.Info("Serving in memory user list")

	time.Sleep(2 * time.Second)
	myUsers := []models.User {
		{Id: 1, Name: "Arnav", Email: "example@gmail.com"},
		{Id: 2, Name: "Jhon", Email: "example@gmail.com"},
		{Id: 3, Name: "Random", Email: "example@gmail.com"},
	}
	return Render(w, r, user.List(myUsers))
}