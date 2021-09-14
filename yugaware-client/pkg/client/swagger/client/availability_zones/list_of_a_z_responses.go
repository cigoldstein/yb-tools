// Code generated by go-swagger; DO NOT EDIT.

package availability_zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
)

// ListOfAZReader is a Reader for the ListOfAZ structure.
type ListOfAZReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOfAZReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOfAZOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListOfAZOK creates a ListOfAZOK with default headers values
func NewListOfAZOK() *ListOfAZOK {
	return &ListOfAZOK{}
}

/* ListOfAZOK describes a response with status code 200, with default header values.

successful operation
*/
type ListOfAZOK struct {
	Payload []*models.AvailabilityZone
}

func (o *ListOfAZOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/customers/{cUUID}/providers/{pUUID}/regions/{rUUID}/zones][%d] listOfAZOK  %+v", 200, o.Payload)
}
func (o *ListOfAZOK) GetPayload() []*models.AvailabilityZone {
	return o.Payload
}

func (o *ListOfAZOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
