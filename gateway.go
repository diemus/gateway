// Package gateway provides a drop-in replacement for net/http.ListenAndServe for use in AWS Lambda & API Gateway.
package gateway

import (
	"context"
	"net/http"

	lambda "github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/tencentyun/scf-go-lib/events"
	//"github.com/aws/aws-lambda-go/events"
	//"github.com/aws/aws-lambda-go/lambda"
)

// ListenAndServe is a drop-in replacement for
// http.ListenAndServe for use within AWS Lambda.
//
// ListenAndServe always returns a non-nil error.
func ListenAndServe(addr string, h http.Handler) error {
	if h == nil {
		h = http.DefaultServeMux
	}

	lambda.Start(func(ctx context.Context, e events.APIGatewayRequest) (events.APIGatewayResponse, error) {
		r, err := NewRequest(ctx, e)
		if err != nil {
			return events.APIGatewayResponse{}, err
		}

		w := NewResponse()
		h.ServeHTTP(w, r)
		return w.End(), nil
	})

	return nil
}
