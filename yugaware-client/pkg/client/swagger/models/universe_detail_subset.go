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

// UniverseDetailSubset A small subset of universe information
//
// swagger:model UniverseDetailSubset
type UniverseDetailSubset struct {

	// creation date
	// Required: true
	CreationDate *int64 `json:"creationDate"`

	// name
	// Required: true
	Name *string `json:"name"`

	// universe paused
	// Required: true
	UniversePaused *bool `json:"universePaused"`

	// update in progress
	// Required: true
	UpdateInProgress *bool `json:"updateInProgress"`

	// update succeeded
	// Required: true
	UpdateSucceeded *bool `json:"updateSucceeded"`

	// uuid
	// Required: true
	// Format: uuid
	UUID *strfmt.UUID `json:"uuid"`
}

// Validate validates this universe detail subset
func (m *UniverseDetailSubset) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUniversePaused(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateInProgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateSucceeded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UniverseDetailSubset) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", m.CreationDate); err != nil {
		return err
	}

	return nil
}

func (m *UniverseDetailSubset) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *UniverseDetailSubset) validateUniversePaused(formats strfmt.Registry) error {

	if err := validate.Required("universePaused", "body", m.UniversePaused); err != nil {
		return err
	}

	return nil
}

func (m *UniverseDetailSubset) validateUpdateInProgress(formats strfmt.Registry) error {

	if err := validate.Required("updateInProgress", "body", m.UpdateInProgress); err != nil {
		return err
	}

	return nil
}

func (m *UniverseDetailSubset) validateUpdateSucceeded(formats strfmt.Registry) error {

	if err := validate.Required("updateSucceeded", "body", m.UpdateSucceeded); err != nil {
		return err
	}

	return nil
}

func (m *UniverseDetailSubset) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("uuid", "body", m.UUID); err != nil {
		return err
	}

	if err := validate.FormatOf("uuid", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this universe detail subset based on context it is used
func (m *UniverseDetailSubset) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UniverseDetailSubset) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UniverseDetailSubset) UnmarshalBinary(b []byte) error {
	var res UniverseDetailSubset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
