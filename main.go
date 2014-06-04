package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var responseDelayInSeconds = 10.0
var respondWith string

// Basic request logging
func logRequest(req *http.Request) {
	remoteHost := strings.Split(req.RemoteAddr, ":")[0]
	log.Printf("[%s]\t%s %s", remoteHost, req.Method, req.URL.String())
}

// Web app logic: listen for some number of seconds, then hang up.
func ListensQuietly(w http.ResponseWriter, req *http.Request) {
	logRequest(req)
	time.Sleep(time.Duration(responseDelayInSeconds) * time.Second)
	io.WriteString(w, respondWith)
}

// Start a web server with a bit of console output using ListensQuietly to
// handle any incoming requests.
func serve(port int, delay float64) {
	responseDelayInSeconds = delay
	portStr := fmt.Sprintf(":%d", port)

	fmt.Printf("Endless nameless is now listening ")
	fmt.Printf("on port %d with a delay of %0.2f second(s)...\n", port, delay)

	http.HandleFunc("/", ListensQuietly)
	log.Fatal(http.ListenAndServe(portStr, nil))
}

// Read any command line flags and start the listener
func main() {
	var portNum *int = flag.Int("port", 8080, "HTTP listener port number")
	var delay *float64 = flag.Float64("delay", 1, "Wait time before sending response")
	var responseInput *string = flag.String("response", "*click*", "String response to send to clients after delay has elapsed.")

	flag.Parse()

	respondWith = *responseInput

	serve(*portNum, *delay)
}
