// Code generated by go-swagger; DO NOT EDIT.

package cloud_providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cloud providers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cloud providers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateProviders(params *CreateProvidersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateProvidersOK, error)

	EditProvider(params *EditProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EditProviderOK, error)

	GetListOfProviders(params *GetListOfProvidersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetListOfProvidersOK, error)

	RefreshPricing(params *RefreshPricingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateProviders creates a provider
*/
func (a *Client) CreateProviders(params *CreateProvidersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateProvidersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateProvidersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createProviders",
		Method:             "POST",
		PathPattern:        "/api/v1/customers/{cUUID}/providers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateProvidersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateProvidersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createProviders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EditProvider updates a provider
*/
func (a *Client) EditProvider(params *EditProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EditProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "editProvider",
		Method:             "PUT",
		PathPattern:        "/api/v1/customers/{cUUID}/providers/{pUUID}/edit",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EditProviderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EditProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for editProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetListOfProviders lists cloud providers
*/
func (a *Client) GetListOfProviders(params *GetListOfProvidersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetListOfProvidersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetListOfProvidersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getListOfProviders",
		Method:             "GET",
		PathPattern:        "/api/v1/customers/{cUUID}/providers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetListOfProvidersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetListOfProvidersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getListOfProviders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RefreshPricing refreshes pricing

  Refresh provider pricing info
*/
func (a *Client) RefreshPricing(params *RefreshPricingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRefreshPricingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "refreshPricing",
		Method:             "PUT",
		PathPattern:        "/api/v1/customers/{cUUID}/providers/{pUUID}/refresh_pricing",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RefreshPricingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
