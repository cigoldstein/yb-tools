// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SupportBundleFormData Support bundle form metadata
//
// swagger:model SupportBundleFormData
type SupportBundleFormData struct {

	// List of components to be included in the support bundle
	// Required: true
	Components []string `json:"components"`

	// End date to filter logs till
	// Example: 2022-01-26
	// Required: true
	// Format: date-time
	EndDate *strfmt.DateTime `json:"endDate"`

	// Start date to filter logs from
	// Example: 2022-01-25
	// Required: true
	// Format: date-time
	StartDate *strfmt.DateTime `json:"startDate"`
}

// Validate validates this support bundle form data
func (m *SupportBundleFormData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var supportBundleFormDataComponentsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UniverseLogs","ApplicationLogs","OutputFiles","ErrorFiles","GFlags","Instance","ConsensusMeta","TabletMeta"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		supportBundleFormDataComponentsItemsEnum = append(supportBundleFormDataComponentsItemsEnum, v)
	}
}

func (m *SupportBundleFormData) validateComponentsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, supportBundleFormDataComponentsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SupportBundleFormData) validateComponents(formats strfmt.Registry) error {

	if err := validate.Required("components", "body", m.Components); err != nil {
		return err
	}

	for i := 0; i < len(m.Components); i++ {

		// value enum
		if err := m.validateComponentsItemsEnum("components"+"."+strconv.Itoa(i), "body", m.Components[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *SupportBundleFormData) validateEndDate(formats strfmt.Registry) error {

	if err := validate.Required("endDate", "body", m.EndDate); err != nil {
		return err
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SupportBundleFormData) validateStartDate(formats strfmt.Registry) error {

	if err := validate.Required("startDate", "body", m.StartDate); err != nil {
		return err
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this support bundle form data based on context it is used
func (m *SupportBundleFormData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SupportBundleFormData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupportBundleFormData) UnmarshalBinary(b []byte) error {
	var res SupportBundleFormData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}