// Code generated by go-swagger; DO NOT EDIT.

package audit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
)

// GetTaskAuditReader is a Reader for the GetTaskAudit structure.
type GetTaskAuditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskAuditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTaskAuditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTaskAuditOK creates a GetTaskAuditOK with default headers values
func NewGetTaskAuditOK() *GetTaskAuditOK {
	return &GetTaskAuditOK{}
}

/* GetTaskAuditOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTaskAuditOK struct {
	Payload *models.Audit
}

func (o *GetTaskAuditOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/customers/{cUUID}/tasks/{tUUID}/audit_info][%d] getTaskAuditOK  %+v", 200, o.Payload)
}
func (o *GetTaskAuditOK) GetPayload() *models.Audit {
	return o.Payload
}

func (o *GetTaskAuditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Audit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
