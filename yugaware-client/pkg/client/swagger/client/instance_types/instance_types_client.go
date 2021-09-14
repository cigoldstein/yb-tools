// Code generated by go-swagger; DO NOT EDIT.

package instance_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new instance types API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for instance types API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateInstanceType(params *CreateInstanceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInstanceTypeOK, error)

	DeleteInstanceType(params *DeleteInstanceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInstanceTypeOK, error)

	GetAZUTypes(params *GetAZUTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAZUTypesOK, error)

	GetEBSTypes(params *GetEBSTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEBSTypesOK, error)

	GetGCPTypes(params *GetGCPTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGCPTypesOK, error)

	InstanceTypeDetail(params *InstanceTypeDetailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InstanceTypeDetailOK, error)

	ListOfInstanceType(params *ListOfInstanceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOfInstanceTypeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateInstanceType creates an instance type
*/
func (a *Client) CreateInstanceType(params *CreateInstanceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInstanceTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateInstanceTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createInstanceType",
		Method:             "POST",
		PathPattern:        "/api/v1/customers/{cUUID}/providers/{pUUID}/instance_types",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateInstanceTypeReader{formats: a.formats},
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
	success, ok := result.(*CreateInstanceTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createInstanceType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteInstanceType deletes an instance type
*/
func (a *Client) DeleteInstanceType(params *DeleteInstanceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInstanceTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInstanceTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteInstanceType",
		Method:             "DELETE",
		PathPattern:        "/api/v1/customers/{cUUID}/providers/{pUUID}/instance_types/{code}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteInstanceTypeReader{formats: a.formats},
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
	success, ok := result.(*DeleteInstanceTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteInstanceType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAZUTypes lists supported azure disk types
*/
func (a *Client) GetAZUTypes(params *GetAZUTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAZUTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAZUTypesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAZUTypes",
		Method:             "GET",
		PathPattern:        "/api/v1/metadata/azu_types",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAZUTypesReader{formats: a.formats},
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
	success, ok := result.(*GetAZUTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAZUTypes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEBSTypes lists supported e b s volume types
*/
func (a *Client) GetEBSTypes(params *GetEBSTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEBSTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEBSTypesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEBSTypes",
		Method:             "GET",
		PathPattern:        "/api/v1/metadata/ebs_types",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEBSTypesReader{formats: a.formats},
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
	success, ok := result.(*GetEBSTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEBSTypes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetGCPTypes lists supported g c p disk types
*/
func (a *Client) GetGCPTypes(params *GetGCPTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGCPTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGCPTypesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGCPTypes",
		Method:             "GET",
		PathPattern:        "/api/v1/metadata/gcp_types",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGCPTypesReader{formats: a.formats},
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
	success, ok := result.(*GetGCPTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGCPTypes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  InstanceTypeDetail gets details of an instance type
*/
func (a *Client) InstanceTypeDetail(params *InstanceTypeDetailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InstanceTypeDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInstanceTypeDetailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "instanceTypeDetail",
		Method:             "GET",
		PathPattern:        "/api/v1/customers/{cUUID}/providers/{pUUID}/instance_types/{code}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InstanceTypeDetailReader{formats: a.formats},
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
	success, ok := result.(*InstanceTypeDetailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for instanceTypeDetail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListOfInstanceType lists a provider s instance types
*/
func (a *Client) ListOfInstanceType(params *ListOfInstanceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOfInstanceTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOfInstanceTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOfInstanceType",
		Method:             "GET",
		PathPattern:        "/api/v1/customers/{cUUID}/providers/{pUUID}/instance_types",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListOfInstanceTypeReader{formats: a.formats},
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
	success, ok := result.(*ListOfInstanceTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listOfInstanceType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
