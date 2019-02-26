package http

import (
	"context"
	"encoding/json"
	"errors"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	endpoint "github.com/marcpar/HextoRGB/hex_to_rgb/pkg/endpoint"
)

func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http1.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http1.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}

// makeTransformHandler creates the handler logic
func makeTransformHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/transform").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(
			endpoints.TransformEndpoint,
			decodeTransformRequest,
			encodeTransformResponse,
			options...)))
}

// decodeTransformRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeTransformRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	// req := endpoint.GetRequest{}
	req := endpoint.TransformRequest{S: r.URL.Query().Get("s")}

	//err := endpoint.TransformRequest{}
	return req, nil
}

// encodeTransformResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeTransformResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
