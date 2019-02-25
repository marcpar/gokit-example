package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/marcpar/HextoRGB/hex_to_rgb/pkg/service"
)

// TrasformRequest collects the request parameters for the Trasform method.
type TrasformRequest struct {
	S string `json:"s"`
}

// TrasformResponse collects the response parameters for the Trasform method.
type TrasformResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// TransformRequest collects the request parameters for the Transform method.
type TransformRequest struct {
	S string `json:"s"`
}

// TransformResponse collects the response parameters for the Transform method.
type TransformResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeTransformEndpoint returns an endpoint that invokes Transform on the service.
func MakeTransformEndpoint(s service.HexToRgbService) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(TransformRequest)
		rs, err := s.Transform(ctx, req.S)
		return TransformResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r TransformResponse) Failed() error {
	return r.Err
}

// Transform implements Service. Primarily useful in a client.
func (e Endpoints) Transform(ctx context.Context, s string) (rs string, err error) {
	request := TransformRequest{S: s}
	response, err := e.TransformEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(TransformResponse).Rs, response.(TransformResponse).Err
}
