package gateway_test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diemus/gateway"
)

func Example() {
	http.HandleFunc("/", hello)
	log.Fatal(gateway.ListenAndServe(":3000", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World from Go")
}
