package main

import (
	"app-htmx/routing"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {
	http.HandleFunc("/", routing.FrontController)
	http.ListenAndServe(":5000", nil)
	open("http://localhost:5000")
}

// https://stackoverflow.com/questions/39320371/how-start-web-server-to-open-page-in-browser-in-golang
// open opens the specified URL in the default browser of the user.
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
