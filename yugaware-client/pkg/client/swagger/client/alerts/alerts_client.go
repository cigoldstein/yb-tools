// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new alerts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for alerts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Acknowledge(params *AcknowledgeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AcknowledgeOK, error)

	AcknowledgeByFilter(params *AcknowledgeByFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AcknowledgeByFilterOK, error)

	CreateAlertChannel(params *CreateAlertChannelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAlertChannelOK, error)

	CreateAlertConfiguration(params *CreateAlertConfigurationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAlertConfigurationOK, error)

	CreateAlertDestination(params *CreateAlertDestinationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAlertDestinationOK, error)

	DeleteAlertChannel(params *DeleteAlertChannelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAlertChannelOK, error)

	DeleteAlertConfiguration(params *DeleteAlertConfigurationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAlertConfigurationOK, error)

	DeleteAlertDestination(params *DeleteAlertDestinationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAlertDestinationOK, error)

	Get(params *GetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOK, error)

	GetAlertChannel(params *GetAlertChannelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAlertChannelOK, error)

	GetAlertConfiguration(params *GetAlertConfigurationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAlertConfigurationOK, error)

	GetAlertDestination(params *GetAlertDestinationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAlertDestinationOK, error)

	ListActive(params *ListActiveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListActiveOK, error)

	ListAlertChannels(params *ListAlertChannelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAlertChannelsOK, error)

	ListAlertConfigurations(params *ListAlertConfigurationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAlertConfigurationsOK, error)

	ListAlertDestinations(params *ListAlertDestinationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAlertDestinationsOK, error)

	ListAlertTemplates(params *ListAlertTemplatesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAlertTemplatesOK, error)

	ListOfAlerts(params *ListOfAlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOfAlertsOK, error)

	PageAlertConfigurations(params *PageAlertConfigurationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PageAlertConfigurationsOK, error)

	PageAlerts(params *PageAlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PageAlertsOK, error)

	UpdateAlertChannel(params *UpdateAlertChannelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAlertChannelOK, error)

	UpdateAlertConfiguration(params *UpdateAlertConfigurationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAlertConfigurationOK, error)

	UpdateAlertDestination(params *UpdateAlertDestinationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAlertDestinationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  Acknowledge acknowledges an alert
*/
func (a *Client) Acknowledge(params *AcknowledgeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AcknowledgeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAcknowledgeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "acknowledge",
		Method:             "POST",
		PathPattern:        "/api/v1/customers/{cUUID}/alerts/{alertUUID}/acknowledge",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AcknowledgeReader{formats: a.formats},
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
	success, ok := result.(*AcknowledgeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for acknowledge: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AcknowledgeByFilter acknowledges all alerts
*/
func (a *Client) AcknowledgeByFilter(params *AcknowledgeByFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AcknowledgeByFilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAcknowledgeByFilterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "acknowledgeByFilter",
		Method:             "POST",
		PathPattern:        "/api/v1/customers/{cUUID}/alerts/acknowledge",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AcknowledgeByFilterReader{formats: a.formats},
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
	success, ok := result.(*AcknowledgeByFilterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for acknowledgeByFilter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateAlertChannel creates an alert channel
*/
func (a *Client) CreateAlertChannel(params *CreateAlertChannelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAlertChannelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAlertChannelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createAlertChannel",
		Method:             "POST",
		PathPattern:        "/api/v1/customers/{cUUID}/alert_channels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateAlertChannelReader{formats: a.formats},
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
	success, ok := result.(*CreateAlertChannelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAlertChannel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateAlertConfiguration creates an alert configuration
*/
func (a *Client) CreateAlertConfiguration(params *CreateAlertConfigurationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAlertConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAlertConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createAlertConfiguration",
		Method:             "POST",
		PathPattern:        "/api/v1/customers/{cUUID}/alert_configurations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateAlertConfigurationReader{formats: a.formats},
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
	success, ok := result.(*CreateAlertConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAlertConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateAlertDestination creates an alert destination
*/
func (a *Client) CreateAlertDestination(params *CreateAlertDestinationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAlertDestinationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAlertDestinationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createAlertDestination",
		Method:             "POST",
		PathPattern:        "/api/v1/customers/{cUUID}/alert_destinations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateAlertDestinationReader{formats: a.formats},
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
	success, ok := result.(*CreateAlertDestinationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAlertDestination: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteAlertChannel deletes an alert channel
*/
func (a *Client) DeleteAlertChannel(params *DeleteAlertChannelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAlertChannelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAlertChannelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAlertChannel",
		Method:             "DELETE",
		PathPattern:        "/api/v1/customers/{cUUID}/alert_channels/{acUUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteAlertChannelReader{formats: a.formats},
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
	success, ok := result.(*DeleteAlertChannelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAlertChannel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteAlertConfiguration deletes an alert configuration
*/
func (a *Client) DeleteAlertConfiguration(params *DeleteAlertConfigurationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAlertConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAlertConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAlertConfiguration",
		Method:             "DELETE",
		PathPattern:        "/api/v1/customers/{cUUID}/alert_configurations/{configurationUUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteAlertConfigurationReader{formats: a.formats},
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
	success, ok := result.(*DeleteAlertConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAlertConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteAlertDestination deletes an alert destination
*/
func (a *Client) DeleteAlertDestination(params *DeleteAlertDestinationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAlertDestinationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAlertDestinationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAlertDestination",
		Method:             "DELETE",
		PathPattern:        "/api/v1/customers/{cUUID}/alert_destinations/{adUUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteAlertDestinationReader{formats: a.formats},
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
	success, ok := result.(*DeleteAlertDestinationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAlertDestination: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Get gets details of an alert
*/
func (a *Client) Get(params *GetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get",
		Method:             "GET",
		PathPattern:        "/api/v1/customers/{cUUID}/alerts/{alertUUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetReader{formats: a.formats},
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
	success, ok := result.(*GetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAlertChannel gets an alert channel
*/
func (a *Client) GetAlertChannel(params *GetAlertChannelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAlertChannelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAlertChannelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAlertChannel",
		Method:             "GET",
		PathPattern:        "/api/v1/customers/{cUUID}/alert_channels/{acUUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAlertChannelReader{formats: a.formats},
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
	success, ok := result.(*GetAlertChannelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAlertChannel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAlertConfiguration gets an alert configuration
*/
func (a *Client) GetAlertConfiguration(params *GetAlertConfigurationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAlertConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAlertConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAlertConfiguration",
		Method:             "GET",
		PathPattern:        "/api/v1/customers/{cUUID}/alert_configurations/{configurationUUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAlertConfigurationReader{formats: a.formats},
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
	success, ok := result.(*GetAlertConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAlertConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAlertDestination gets an alert destination
*/
func (a *Client) GetAlertDestination(params *GetAlertDestinationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAlertDestinationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAlertDestinationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAlertDestination",
		Method:             "GET",
		PathPattern:        "/api/v1/customers/{cUUID}/alert_destinations/{adUUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAlertDestinationReader{formats: a.formats},
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
	success, ok := result.(*GetAlertDestinationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAlertDestination: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListActive lists active alerts
*/
func (a *Client) ListActive(params *ListActiveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListActiveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListActiveParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listActive",
		Method:             "GET",
		PathPattern:        "/api/v1/customers/{cUUID}/alerts/active",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListActiveReader{formats: a.formats},
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
	success, ok := result.(*ListActiveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listActive: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAlertChannels lists all alert channels
*/
func (a *Client) ListAlertChannels(params *ListAlertChannelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAlertChannelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAlertChannelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAlertChannels",
		Method:             "GET",
		PathPattern:        "/api/v1/customers/{cUUID}/alert_channels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListAlertChannelsReader{formats: a.formats},
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
	success, ok := result.(*ListAlertChannelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAlertChannels: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAlertConfigurations gets filtered list of alert configurations
*/
func (a *Client) ListAlertConfigurations(params *ListAlertConfigurationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAlertConfigurationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAlertConfigurationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAlertConfigurations",
		Method:             "POST",
		PathPattern:        "/api/v1/customers/{cUUID}/alert_configurations/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListAlertConfigurationsReader{formats: a.formats},
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
	success, ok := result.(*ListAlertConfigurationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAlertConfigurations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAlertDestinations lists alert destinations
*/
func (a *Client) ListAlertDestinations(params *ListAlertDestinationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAlertDestinationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAlertDestinationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAlertDestinations",
		Method:             "GET",
		PathPattern:        "/api/v1/customers/{cUUID}/alert_destinations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListAlertDestinationsReader{formats: a.formats},
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
	success, ok := result.(*ListAlertDestinationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAlertDestinations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAlertTemplates gets filtered list of alert configuration templates
*/
func (a *Client) ListAlertTemplates(params *ListAlertTemplatesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAlertTemplatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAlertTemplatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAlertTemplates",
		Method:             "POST",
		PathPattern:        "/api/v1/customers/{cUUID}/alert_templates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListAlertTemplatesReader{formats: a.formats},
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
	success, ok := result.(*ListAlertTemplatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAlertTemplates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListOfAlerts lists all alerts
*/
func (a *Client) ListOfAlerts(params *ListOfAlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOfAlertsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOfAlertsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOfAlerts",
		Method:             "GET",
		PathPattern:        "/api/v1/customers/{cUUID}/alerts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListOfAlertsReader{formats: a.formats},
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
	success, ok := result.(*ListOfAlertsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listOfAlerts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PageAlertConfigurations lists all alert configurations paginated
*/
func (a *Client) PageAlertConfigurations(params *PageAlertConfigurationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PageAlertConfigurationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPageAlertConfigurationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pageAlertConfigurations",
		Method:             "POST",
		PathPattern:        "/api/v1/customers/{cUUID}/alert_configurations/page",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PageAlertConfigurationsReader{formats: a.formats},
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
	success, ok := result.(*PageAlertConfigurationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pageAlertConfigurations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PageAlerts lists alerts paginated
*/
func (a *Client) PageAlerts(params *PageAlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PageAlertsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPageAlertsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pageAlerts",
		Method:             "POST",
		PathPattern:        "/api/v1/customers/{cUUID}/alerts/page",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PageAlertsReader{formats: a.formats},
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
	success, ok := result.(*PageAlertsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pageAlerts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateAlertChannel updates an alert channel
*/
func (a *Client) UpdateAlertChannel(params *UpdateAlertChannelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAlertChannelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAlertChannelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateAlertChannel",
		Method:             "PUT",
		PathPattern:        "/api/v1/customers/{cUUID}/alert_channels/{acUUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateAlertChannelReader{formats: a.formats},
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
	success, ok := result.(*UpdateAlertChannelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateAlertChannel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateAlertConfiguration updates an alert configuration
*/
func (a *Client) UpdateAlertConfiguration(params *UpdateAlertConfigurationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAlertConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAlertConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateAlertConfiguration",
		Method:             "PUT",
		PathPattern:        "/api/v1/customers/{cUUID}/alert_configurations/{configurationUUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateAlertConfigurationReader{formats: a.formats},
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
	success, ok := result.(*UpdateAlertConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateAlertConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateAlertDestination updates an alert destination
*/
func (a *Client) UpdateAlertDestination(params *UpdateAlertDestinationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAlertDestinationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAlertDestinationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateAlertDestination",
		Method:             "PUT",
		PathPattern:        "/api/v1/customers/{cUUID}/alert_destinations/{adUUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateAlertDestinationReader{formats: a.formats},
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
	success, ok := result.(*UpdateAlertDestinationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateAlertDestination: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
