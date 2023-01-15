package main

import (
	"net/http"
	"os"

	"lib.dev/libdev"
)

func main() {
	port := os.Getenv("PORT")
	app := libdev.App{}
	http.ListenAndServe(":"+port, &app)
}
