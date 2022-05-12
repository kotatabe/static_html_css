package main

// should use templete
import (
	"net/http"
)

func main() {

	fs_handler := http.FileServer(http.Dir("static"))

	http.Handle("/", fs_handler)

	http.ListenAndServe(":8001", nil)
}
