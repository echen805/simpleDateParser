package simpleDateParser

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	GetEndpoint endpoint.Endpoint
	StatusEndpoint	endpoint.Endpoint
	ValidateEndpoint	endpoint.Endpoint
}

// MakeGetEndpoint returns the response from our service "get"
func MakeGetEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){
		_ = request.(getRequest) // What's the purpose of this line?
		d, err := srv.Get(ctx)
		if err != nil {
			return getResponse{d, err.Error()}, nil
		}
		return getResponse{d, ""}, nil
	}
}

// MakeStatusEndpoint returns the response from our service "status"
func MakeStatusEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){
		_ = request.(statusRequest) // What's the purpose of this line?
		status, err := srv.Status(ctx)
		if err != nil {
			return statusResponse{status}, nil
		}
		return statusResponse{status}, nil
	}
}

// MakeValidateEndpoint returns the response from our service "validate"
func MakeValidateEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){
		req := request.(validateRequest)
		isValid, err := srv.Validate(ctx, req.Date)
		if err != nil {
			return validateResponse{isValid, err.Error()}, nil
		}
		return validateResponse{isValid, ""}, nil
	}
}

// Get endpoint mapping
func (e Endpoints) Get(ctx context.Context) (string, error) {
	req := getRequest{}
	resp, err := e.GetEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	getResp := resp.(getResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Date, nil
}