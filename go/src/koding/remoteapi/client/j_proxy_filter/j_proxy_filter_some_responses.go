package j_proxy_filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JProxyFilterSomeReader is a Reader for the JProxyFilterSome structure.
type JProxyFilterSomeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JProxyFilterSomeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJProxyFilterSomeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewJProxyFilterSomeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJProxyFilterSomeOK creates a JProxyFilterSomeOK with default headers values
func NewJProxyFilterSomeOK() *JProxyFilterSomeOK {
	return &JProxyFilterSomeOK{}
}

/*JProxyFilterSomeOK handles this case with default header values.

Request processed successfully
*/
type JProxyFilterSomeOK struct {
	Payload *models.DefaultResponse
}

func (o *JProxyFilterSomeOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JProxyFilter.some][%d] jProxyFilterSomeOK  %+v", 200, o.Payload)
}

func (o *JProxyFilterSomeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJProxyFilterSomeUnauthorized creates a JProxyFilterSomeUnauthorized with default headers values
func NewJProxyFilterSomeUnauthorized() *JProxyFilterSomeUnauthorized {
	return &JProxyFilterSomeUnauthorized{}
}

/*JProxyFilterSomeUnauthorized handles this case with default header values.

Unauthorized request
*/
type JProxyFilterSomeUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *JProxyFilterSomeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JProxyFilter.some][%d] jProxyFilterSomeUnauthorized  %+v", 401, o.Payload)
}

func (o *JProxyFilterSomeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
