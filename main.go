package main

import (
	"net/http"
	"os"

	"github.com/webmachine-dev/lib-dev/pkg/libdev"
)

func main() {
	port := os.Getenv("PORT")
	app := libdev.App{}
	http.ListenAndServe(":"+port, &app)
}
