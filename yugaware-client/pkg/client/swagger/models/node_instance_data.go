// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NodeInstanceData Details of a node instance. Used by the API to validate data against input constraints.
//
// swagger:model NodeInstanceData
type NodeInstanceData struct {

	// Node instance name
	// Example: Mumbai instance
	// Required: true
	InstanceName *string `json:"instanceName"`

	// Node instance type
	// Example: c5large
	// Required: true
	InstanceType *string `json:"instanceType"`

	// IP address
	// Example: 1.1.1.1
	// Required: true
	IP *string `json:"ip"`

	// Node name
	// Example: India node
	NodeName string `json:"nodeName,omitempty"`

	// Region
	// Example: south-east
	// Required: true
	Region *string `json:"region"`

	// SSH user
	// Example: centos
	// Required: true
	SSHUser *string `json:"sshUser"`

	// Zone
	// Example: south-east
	// Required: true
	Zone *string `json:"zone"`
}

// Validate validates this node instance data
func (m *NodeInstanceData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeInstanceData) validateInstanceName(formats strfmt.Registry) error {

	if err := validate.Required("instanceName", "body", m.InstanceName); err != nil {
		return err
	}

	return nil
}

func (m *NodeInstanceData) validateInstanceType(formats strfmt.Registry) error {

	if err := validate.Required("instanceType", "body", m.InstanceType); err != nil {
		return err
	}

	return nil
}

func (m *NodeInstanceData) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

func (m *NodeInstanceData) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

func (m *NodeInstanceData) validateSSHUser(formats strfmt.Registry) error {

	if err := validate.Required("sshUser", "body", m.SSHUser); err != nil {
		return err
	}

	return nil
}

func (m *NodeInstanceData) validateZone(formats strfmt.Registry) error {

	if err := validate.Required("zone", "body", m.Zone); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this node instance data based on context it is used
func (m *NodeInstanceData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodeInstanceData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeInstanceData) UnmarshalBinary(b []byte) error {
	var res NodeInstanceData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
