// Code generated by go-swagger; DO NOT EDIT.

package region_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
)

// ListAllRegionsReader is a Reader for the ListAllRegions structure.
type ListAllRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAllRegionsOK creates a ListAllRegionsOK with default headers values
func NewListAllRegionsOK() *ListAllRegionsOK {
	return &ListAllRegionsOK{}
}

/* ListAllRegionsOK describes a response with status code 200, with default header values.

successful operation
*/
type ListAllRegionsOK struct {
	Payload []*models.Region
}

func (o *ListAllRegionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/customers/{cUUID}/regions][%d] listAllRegionsOK  %+v", 200, o.Payload)
}
func (o *ListAllRegionsOK) GetPayload() []*models.Region {
	return o.Payload
}

func (o *ListAllRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}