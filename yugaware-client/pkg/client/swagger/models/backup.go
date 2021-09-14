// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Backup A single backup. Includes the backup's status, expiration time, and configuration.
//
// swagger:model Backup
type Backup struct {

	// Details of the backup
	BackupInfo *BackupTableParams `json:"backupInfo,omitempty"`

	// Backup UUID
	// Read Only: true
	// Format: uuid
	BackupUUID strfmt.UUID `json:"backupUUID,omitempty"`

	// create time
	// Required: true
	// Format: date-time
	CreateTime *strfmt.DateTime `json:"createTime"`

	// Customer UUID that owns this backup
	// Format: uuid
	CustomerUUID strfmt.UUID `json:"customerUUID,omitempty"`

	// Expiry time (unix timestamp) of the backup
	// Format: date-time
	Expiry strfmt.DateTime `json:"expiry,omitempty"`

	// Schedule UUID, if this backup is part of a schedule
	// Format: uuid
	ScheduleUUID strfmt.UUID `json:"scheduleUUID,omitempty"`

	// State of the backup
	// Example: DELETED
	// Read Only: true
	// Enum: [InProgress Completed Failed Deleted Skipped FailedToDelete Stopped]
	State string `json:"state,omitempty"`

	// Backup UUID
	// Read Only: true
	// Format: uuid
	TaskUUID strfmt.UUID `json:"taskUUID,omitempty"`

	// update time
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"updateTime"`
}

// Validate validates this backup
func (m *Backup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduleUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Backup) validateBackupInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupInfo) { // not required
		return nil
	}

	if m.BackupInfo != nil {
		if err := m.BackupInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backupInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Backup) validateBackupUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("backupUUID", "body", "uuid", m.BackupUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Backup) validateCreateTime(formats strfmt.Registry) error {

	if err := validate.Required("createTime", "body", m.CreateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("createTime", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Backup) validateCustomerUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomerUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("customerUUID", "body", "uuid", m.CustomerUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Backup) validateExpiry(formats strfmt.Registry) error {
	if swag.IsZero(m.Expiry) { // not required
		return nil
	}

	if err := validate.FormatOf("expiry", "body", "date-time", m.Expiry.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Backup) validateScheduleUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.ScheduleUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("scheduleUUID", "body", "uuid", m.ScheduleUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

var backupTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["InProgress","Completed","Failed","Deleted","Skipped","FailedToDelete","Stopped"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backupTypeStatePropEnum = append(backupTypeStatePropEnum, v)
	}
}

const (

	// BackupStateInProgress captures enum value "InProgress"
	BackupStateInProgress string = "InProgress"

	// BackupStateCompleted captures enum value "Completed"
	BackupStateCompleted string = "Completed"

	// BackupStateFailed captures enum value "Failed"
	BackupStateFailed string = "Failed"

	// BackupStateDeleted captures enum value "Deleted"
	BackupStateDeleted string = "Deleted"

	// BackupStateSkipped captures enum value "Skipped"
	BackupStateSkipped string = "Skipped"

	// BackupStateFailedToDelete captures enum value "FailedToDelete"
	BackupStateFailedToDelete string = "FailedToDelete"

	// BackupStateStopped captures enum value "Stopped"
	BackupStateStopped string = "Stopped"
)

// prop value enum
func (m *Backup) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, backupTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Backup) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *Backup) validateTaskUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.TaskUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("taskUUID", "body", "uuid", m.TaskUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Backup) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("updateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("updateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this backup based on the context it is used
func (m *Backup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBackupUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaskUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Backup) contextValidateBackupInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.BackupInfo != nil {
		if err := m.BackupInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backupInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Backup) contextValidateBackupUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "backupUUID", "body", strfmt.UUID(m.BackupUUID)); err != nil {
		return err
	}

	return nil
}

func (m *Backup) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

func (m *Backup) contextValidateTaskUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "taskUUID", "body", strfmt.UUID(m.TaskUUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Backup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Backup) UnmarshalBinary(b []byte) error {
	var res Backup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
