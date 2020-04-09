package gateway

import (
	"context"
	"github.com/tencentyun/scf-go-lib/events"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// NewRequest returns a new http.Request from the given Lambda event.
func NewRequest(ctx context.Context, e events.APIGatewayRequest) (*http.Request, error) {
	// path
	u, err := url.Parse(e.Path)
	if err != nil {
		return nil, errors.Wrap(err, "parsing path")
	}

	// querystring
	q := u.Query()
	for k, v := range e.QueryString {
		q[k] = v
	}

	u.RawQuery = q.Encode()

	// base64 encoded body
	body := e.Body
	//if e.IsBase64Encoded {
	//	b, err := base64.StdEncoding.DecodeString(body)
	//	if err != nil {
	//		return nil, errors.Wrap(err, "decoding base64 body")
	//	}
	//	body = string(b)
	//}

	// new request
	req, err := http.NewRequest(e.Method, u.String(), strings.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}

	// manually set RequestURI because NewRequest is for clients and req.RequestURI is for servers
	req.RequestURI = e.Path

	// remote addr
	req.RemoteAddr = e.Context.SourceIP

	// header fields
	for k, v := range e.Headers {
		req.Header.Set(k, v)
	}

	// content-length
	if req.Header.Get("Content-Length") == "" && body != "" {
		req.Header.Set("Content-Length", strconv.Itoa(len(body)))
	}

	// custom fields
	req.Header.Set("X-Request-Id", e.Context.RequestID)
	req.Header.Set("X-Stage", e.Context.Stage)

	// custom context values
	req = req.WithContext(newContext(ctx, e))

	// host
	req.URL.Host = req.Header.Get("Host")
	req.Host = req.URL.Host

	return req, nil
}
