<img src="http://tjholowaychuk.com:6000/svg/title/APEX/GATEWAY">

Package gateway provides a drop-in replacement for net/http's `ListenAndServe` for use in SCF Cloud Function & API Gateway, simply swap it out for `gateway.ListenAndServe`. Extracted from [Up](https://github.com/apex/up) which provides additional middleware features and operational functionality.

```go
package main

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
```

## gin
```go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/diemus/gateway"
	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func welcomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World from Go")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"text": "Welcome to gin lambda server.",
	})
}

func routerEngine() *gin.Engine {
	// set server mode
	gin.SetMode(gin.DebugMode)

	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/welcome", welcomeHandler)
	r.GET("/user/:name", helloHandler)
	r.GET("/", rootHandler)

	return r
}

func main() {
	addr := ":" + os.Getenv("PORT")
	log.Fatal(gateway.ListenAndServe(addr, routerEngine()))
}
```

---

[![GoDoc](https://godoc.org/github.com/apex/up-go?status.svg)](https://godoc.org/github.com/apex/gateway)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-stable-green.svg)

<a href="https://apex.sh"><img src="http://tjholowaychuk.com:6000/svg/sponsor"></a>
