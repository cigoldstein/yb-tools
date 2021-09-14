// Code generated by go-swagger; DO NOT EDIT.

package instance_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAZUTypesReader is a Reader for the GetAZUTypes structure.
type GetAZUTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAZUTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAZUTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAZUTypesOK creates a GetAZUTypesOK with default headers values
func NewGetAZUTypesOK() *GetAZUTypesOK {
	return &GetAZUTypesOK{}
}

/* GetAZUTypesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAZUTypesOK struct {
	Payload []string
}

func (o *GetAZUTypesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/metadata/azu_types][%d] getAZUTypesOK  %+v", 200, o.Payload)
}
func (o *GetAZUTypesOK) GetPayload() []string {
	return o.Payload
}

func (o *GetAZUTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
