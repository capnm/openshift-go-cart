package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

var bind = ""

// Run a simple file server with the root in the 'static' directory.
func main() {
	http.HandleFunc("/info/", infohandler)
	http.Handle("/", http.FileServer(http.Dir("static")))

	bind = fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
	if bind == ":" {
		bind = ":8080"
	}

	fmt.Printf("==== listening on '%s'\n", bind)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		panic(err)
	}
}

// An example of a custom handler.
func infohandler(w http.ResponseWriter, r *http.Request) {
	log.Println("access:", r.RequestURI)

	fmt.Fprintf(w, "go version '%s'\nbind on '%s'\nremote addr ",
		runtime.Version(), bind)
	if len(r.RemoteAddr) > 0 {
		w.Write([]byte(strings.Split(r.RemoteAddr, ":")[0]))
	}
}
