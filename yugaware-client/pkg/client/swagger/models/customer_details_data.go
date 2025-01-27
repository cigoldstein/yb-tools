// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomerDetailsData Customer details, including their universe UUIDs. Only the customer code and name are modifiable.
//
// swagger:model CustomerDetailsData
type CustomerDetailsData struct {

	// Alerts
	// Read Only: true
	AlertingData *AlertingData `json:"alertingData,omitempty"`

	// Call-home level
	// Example: MEDIUM
	// Read Only: true
	CallhomeLevel string `json:"callhomeLevel,omitempty"`

	// Customer code
	// Example: admin
	// Required: true
	Code *string `json:"code"`

	// Creation timestamp
	// Example: 2021-06-17 15:00:05
	// Read Only: true
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// Customer ID
	// Read Only: true
	CustomerID int32 `json:"customerId,omitempty"`

	// Customer name
	// Example: Sridhar
	// Required: true
	Name *string `json:"name"`

	// SMTP
	// Read Only: true
	SMTPData *SMTPData `json:"smtpData,omitempty"`

	// Associated universe IDs
	// Example: [c3595ca7-68a3-47f0-b1b2-1725886d5ed5, 9e0bb733-556c-4935-83dd-6b742a2c32e6]
	// Read Only: true
	UniverseUUIDs []strfmt.UUID `json:"universeUUIDs"`

	// User UUID
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this customer details data
func (m *CustomerDetailsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlertingData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSMTPData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUniverseUUIDs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomerDetailsData) validateAlertingData(formats strfmt.Registry) error {
	if swag.IsZero(m.AlertingData) { // not required
		return nil
	}

	if m.AlertingData != nil {
		if err := m.AlertingData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alertingData")
			}
			return err
		}
	}

	return nil
}

func (m *CustomerDetailsData) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *CustomerDetailsData) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomerDetailsData) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CustomerDetailsData) validateSMTPData(formats strfmt.Registry) error {
	if swag.IsZero(m.SMTPData) { // not required
		return nil
	}

	if m.SMTPData != nil {
		if err := m.SMTPData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smtpData")
			}
			return err
		}
	}

	return nil
}

func (m *CustomerDetailsData) validateUniverseUUIDs(formats strfmt.Registry) error {
	if swag.IsZero(m.UniverseUUIDs) { // not required
		return nil
	}

	for i := 0; i < len(m.UniverseUUIDs); i++ {

		if err := validate.FormatOf("universeUUIDs"+"."+strconv.Itoa(i), "body", "uuid", m.UniverseUUIDs[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validate this customer details data based on the context it is used
func (m *CustomerDetailsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlertingData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCallhomeLevel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreationDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomerID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSMTPData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUniverseUUIDs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomerDetailsData) contextValidateAlertingData(ctx context.Context, formats strfmt.Registry) error {

	if m.AlertingData != nil {
		if err := m.AlertingData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alertingData")
			}
			return err
		}
	}

	return nil
}

func (m *CustomerDetailsData) contextValidateCallhomeLevel(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "callhomeLevel", "body", string(m.CallhomeLevel)); err != nil {
		return err
	}

	return nil
}

func (m *CustomerDetailsData) contextValidateCreationDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "creationDate", "body", strfmt.DateTime(m.CreationDate)); err != nil {
		return err
	}

	return nil
}

func (m *CustomerDetailsData) contextValidateCustomerID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "customerId", "body", int32(m.CustomerID)); err != nil {
		return err
	}

	return nil
}

func (m *CustomerDetailsData) contextValidateSMTPData(ctx context.Context, formats strfmt.Registry) error {

	if m.SMTPData != nil {
		if err := m.SMTPData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smtpData")
			}
			return err
		}
	}

	return nil
}

func (m *CustomerDetailsData) contextValidateUniverseUUIDs(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "universeUUIDs", "body", []strfmt.UUID(m.UniverseUUIDs)); err != nil {
		return err
	}

	return nil
}

func (m *CustomerDetailsData) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomerDetailsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerDetailsData) UnmarshalBinary(b []byte) error {
	var res CustomerDetailsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
