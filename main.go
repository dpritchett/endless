package main

import (
	"flag"

	"github.com/dpritchett/endless/listener"
)

// Read any command line flags and start the listener
func main() {
	var portNum *int = flag.Int("port", 8080, "HTTP(s) listener port number")
	var delay *float64 = flag.Float64("delay", 1, "Wait time in seconds before sending the response")
	var respondWith *string = flag.String("response", "*click*", "String response to send to clients after delay has elapsed.")
	var ssl *bool = flag.Bool("ssl", false, "Listen via HTTPS by passing --ssl.")

	flag.Parse()

	daemon := listener.Listener()

	daemon.RespondWith = *respondWith
	daemon.Port = *portNum
	daemon.ResponseDelay = *delay
	daemon.HTTPS = *ssl

	daemon.Serve()
}
