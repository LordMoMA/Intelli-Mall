// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/LordMoMA/Intelli-Mall/customers/customersclient/models"
)

// EnableCustomerReader is a Reader for the EnableCustomer structure.
type EnableCustomerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableCustomerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnableCustomerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEnableCustomerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEnableCustomerOK creates a EnableCustomerOK with default headers values
func NewEnableCustomerOK() *EnableCustomerOK {
	return &EnableCustomerOK{}
}

/* EnableCustomerOK describes a response with status code 200, with default header values.

A successful response.
*/
type EnableCustomerOK struct {
	Payload models.CustomerspbEnableCustomerResponse
}

// IsSuccess returns true when this enable customer o k response has a 2xx status code
func (o *EnableCustomerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enable customer o k response has a 3xx status code
func (o *EnableCustomerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable customer o k response has a 4xx status code
func (o *EnableCustomerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this enable customer o k response has a 5xx status code
func (o *EnableCustomerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this enable customer o k response a status code equal to that given
func (o *EnableCustomerOK) IsCode(code int) bool {
	return code == 200
}

func (o *EnableCustomerOK) Error() string {
	return fmt.Sprintf("[PUT /api/customers/{id}/enable][%d] enableCustomerOK  %+v", 200, o.Payload)
}

func (o *EnableCustomerOK) String() string {
	return fmt.Sprintf("[PUT /api/customers/{id}/enable][%d] enableCustomerOK  %+v", 200, o.Payload)
}

func (o *EnableCustomerOK) GetPayload() models.CustomerspbEnableCustomerResponse {
	return o.Payload
}

func (o *EnableCustomerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableCustomerDefault creates a EnableCustomerDefault with default headers values
func NewEnableCustomerDefault(code int) *EnableCustomerDefault {
	return &EnableCustomerDefault{
		_statusCode: code,
	}
}

/* EnableCustomerDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EnableCustomerDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the enable customer default response
func (o *EnableCustomerDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this enable customer default response has a 2xx status code
func (o *EnableCustomerDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this enable customer default response has a 3xx status code
func (o *EnableCustomerDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this enable customer default response has a 4xx status code
func (o *EnableCustomerDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this enable customer default response has a 5xx status code
func (o *EnableCustomerDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this enable customer default response a status code equal to that given
func (o *EnableCustomerDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EnableCustomerDefault) Error() string {
	return fmt.Sprintf("[PUT /api/customers/{id}/enable][%d] enableCustomer default  %+v", o._statusCode, o.Payload)
}

func (o *EnableCustomerDefault) String() string {
	return fmt.Sprintf("[PUT /api/customers/{id}/enable][%d] enableCustomer default  %+v", o._statusCode, o.Payload)
}

func (o *EnableCustomerDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *EnableCustomerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
