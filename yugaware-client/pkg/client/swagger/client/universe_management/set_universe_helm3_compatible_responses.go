// Code generated by go-swagger; DO NOT EDIT.

package universe_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
)

// SetUniverseHelm3CompatibleReader is a Reader for the SetUniverseHelm3Compatible structure.
type SetUniverseHelm3CompatibleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetUniverseHelm3CompatibleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetUniverseHelm3CompatibleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetUniverseHelm3CompatibleOK creates a SetUniverseHelm3CompatibleOK with default headers values
func NewSetUniverseHelm3CompatibleOK() *SetUniverseHelm3CompatibleOK {
	return &SetUniverseHelm3CompatibleOK{}
}

/* SetUniverseHelm3CompatibleOK describes a response with status code 200, with default header values.

successful operation
*/
type SetUniverseHelm3CompatibleOK struct {
	Payload *models.YBPSuccess
}

func (o *SetUniverseHelm3CompatibleOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/customers/{cUUID}/universes/{uniUUID}/mark_helm3_compatible][%d] setUniverseHelm3CompatibleOK  %+v", 200, o.Payload)
}
func (o *SetUniverseHelm3CompatibleOK) GetPayload() *models.YBPSuccess {
	return o.Payload
}

func (o *SetUniverseHelm3CompatibleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.YBPSuccess)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
