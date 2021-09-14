// Code generated by go-swagger; DO NOT EDIT.

package backup_schedule_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
)

// DeleteBackupScheduleReader is a Reader for the DeleteBackupSchedule structure.
type DeleteBackupScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBackupScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteBackupScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteBackupScheduleOK creates a DeleteBackupScheduleOK with default headers values
func NewDeleteBackupScheduleOK() *DeleteBackupScheduleOK {
	return &DeleteBackupScheduleOK{}
}

/* DeleteBackupScheduleOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteBackupScheduleOK struct {
	Payload models.PlatformResults
}

func (o *DeleteBackupScheduleOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/customers/{cUUID}/schedules/{sUUID}][%d] deleteBackupScheduleOK  %+v", 200, o.Payload)
}
func (o *DeleteBackupScheduleOK) GetPayload() models.PlatformResults {
	return o.Payload
}

func (o *DeleteBackupScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
