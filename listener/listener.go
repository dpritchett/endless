package listener

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dpritchett/endless/snakeoil"
)

type listener struct {
	RespondWith   string
	ResponseDelay float64
	Port          int
	HTTPS         bool
}

func Listener() *listener {
	return &listener{}
}

// Basic request logging
func logRequest(req *http.Request) {
	remoteHost := strings.Split(req.RemoteAddr, ":")[0]
	log.Printf("[%s]\t%s %s", remoteHost, req.Method, req.URL.String())
}

// Web app logic: listen for some number of seconds, then hang up.
func (srv *listener) ListenQuietly(w http.ResponseWriter, req *http.Request) {
	logRequest(req)
	time.Sleep(time.Duration(srv.ResponseDelay) * time.Second)
	io.WriteString(w, srv.RespondWith)
}

// Start a web server with a bit of console output using ListensQuietly to
// handle any incoming requests.
func (srv *listener) Serve() {
	portStr := fmt.Sprintf(":%d", srv.Port)

	fmt.Printf("Endless nameless is now listening ")
	fmt.Printf("on port %d with a delay of %0.2f second(s)...\n", srv.Port, srv.ResponseDelay)

	http.HandleFunc("/", srv.ListenQuietly)

	if srv.HTTPS {
		cert, key := snakeoil.CertFileName(), snakeoil.KeyFileName()
		log.Fatal(http.ListenAndServeTLS(portStr, cert, key, nil))
	} else {
		log.Fatal(http.ListenAndServe(portStr, nil))
	}
}
