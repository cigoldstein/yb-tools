// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AlertChannelSlackParams alert channel slack params
//
// swagger:model AlertChannelSlackParams
type AlertChannelSlackParams struct {
	textTemplateField *string

	titleTemplateField *string

	// icon Url
	// Required: true
	IconURL *string `json:"iconUrl"`

	// username
	// Required: true
	Username *string `json:"username"`

	// webhook Url
	// Required: true
	WebhookURL *string `json:"webhookUrl"`
}

// TextTemplate gets the text template of this subtype
func (m *AlertChannelSlackParams) TextTemplate() *string {
	return m.textTemplateField
}

// SetTextTemplate sets the text template of this subtype
func (m *AlertChannelSlackParams) SetTextTemplate(val *string) {
	m.textTemplateField = val
}

// TitleTemplate gets the title template of this subtype
func (m *AlertChannelSlackParams) TitleTemplate() *string {
	return m.titleTemplateField
}

// SetTitleTemplate sets the title template of this subtype
func (m *AlertChannelSlackParams) SetTitleTemplate(val *string) {
	m.titleTemplateField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AlertChannelSlackParams) UnmarshalJSON(raw []byte) error {
	var data struct {

		// icon Url
		// Required: true
		IconURL *string `json:"iconUrl"`

		// username
		// Required: true
		Username *string `json:"username"`

		// webhook Url
		// Required: true
		WebhookURL *string `json:"webhookUrl"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		TextTemplate *string `json:"textTemplate"`

		TitleTemplate *string `json:"titleTemplate"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result AlertChannelSlackParams

	result.textTemplateField = base.TextTemplate

	result.titleTemplateField = base.TitleTemplate

	result.IconURL = data.IconURL
	result.Username = data.Username
	result.WebhookURL = data.WebhookURL

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AlertChannelSlackParams) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// icon Url
		// Required: true
		IconURL *string `json:"iconUrl"`

		// username
		// Required: true
		Username *string `json:"username"`

		// webhook Url
		// Required: true
		WebhookURL *string `json:"webhookUrl"`
	}{

		IconURL: m.IconURL,

		Username: m.Username,

		WebhookURL: m.WebhookURL,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		TextTemplate *string `json:"textTemplate"`

		TitleTemplate *string `json:"titleTemplate"`
	}{

		TextTemplate: m.TextTemplate(),

		TitleTemplate: m.TitleTemplate(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this alert channel slack params
func (m *AlertChannelSlackParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTextTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitleTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIconURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhookURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertChannelSlackParams) validateTextTemplate(formats strfmt.Registry) error {

	if err := validate.Required("textTemplate", "body", m.TextTemplate()); err != nil {
		return err
	}

	return nil
}

func (m *AlertChannelSlackParams) validateTitleTemplate(formats strfmt.Registry) error {

	if err := validate.Required("titleTemplate", "body", m.TitleTemplate()); err != nil {
		return err
	}

	return nil
}

func (m *AlertChannelSlackParams) validateIconURL(formats strfmt.Registry) error {

	if err := validate.Required("iconUrl", "body", m.IconURL); err != nil {
		return err
	}

	return nil
}

func (m *AlertChannelSlackParams) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

func (m *AlertChannelSlackParams) validateWebhookURL(formats strfmt.Registry) error {

	if err := validate.Required("webhookUrl", "body", m.WebhookURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this alert channel slack params based on the context it is used
func (m *AlertChannelSlackParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AlertChannelSlackParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertChannelSlackParams) UnmarshalBinary(b []byte) error {
	var res AlertChannelSlackParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
