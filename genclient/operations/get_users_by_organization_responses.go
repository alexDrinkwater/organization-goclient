// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3dsim/organization-goclient/models"
)

// GetUsersByOrganizationReader is a Reader for the GetUsersByOrganization structure.
type GetUsersByOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersByOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUsersByOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetUsersByOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetUsersByOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetUsersByOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetUsersByOrganizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUsersByOrganizationOK creates a GetUsersByOrganizationOK with default headers values
func NewGetUsersByOrganizationOK() *GetUsersByOrganizationOK {
	return &GetUsersByOrganizationOK{}
}

/*GetUsersByOrganizationOK handles this case with default header values.

Successfully returned the list of items
*/
type GetUsersByOrganizationOK struct {
	Payload []*models.User
}

func (o *GetUsersByOrganizationOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/users][%d] getUsersByOrganizationOK  %+v", 200, o.Payload)
}

func (o *GetUsersByOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersByOrganizationUnauthorized creates a GetUsersByOrganizationUnauthorized with default headers values
func NewGetUsersByOrganizationUnauthorized() *GetUsersByOrganizationUnauthorized {
	return &GetUsersByOrganizationUnauthorized{}
}

/*GetUsersByOrganizationUnauthorized handles this case with default header values.

Not authorized
*/
type GetUsersByOrganizationUnauthorized struct {
}

func (o *GetUsersByOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/users][%d] getUsersByOrganizationUnauthorized ", 401)
}

func (o *GetUsersByOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersByOrganizationForbidden creates a GetUsersByOrganizationForbidden with default headers values
func NewGetUsersByOrganizationForbidden() *GetUsersByOrganizationForbidden {
	return &GetUsersByOrganizationForbidden{}
}

/*GetUsersByOrganizationForbidden handles this case with default header values.

Forbidden
*/
type GetUsersByOrganizationForbidden struct {
}

func (o *GetUsersByOrganizationForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/users][%d] getUsersByOrganizationForbidden ", 403)
}

func (o *GetUsersByOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersByOrganizationNotFound creates a GetUsersByOrganizationNotFound with default headers values
func NewGetUsersByOrganizationNotFound() *GetUsersByOrganizationNotFound {
	return &GetUsersByOrganizationNotFound{}
}

/*GetUsersByOrganizationNotFound handles this case with default header values.

Organization not found
*/
type GetUsersByOrganizationNotFound struct {
}

func (o *GetUsersByOrganizationNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/users][%d] getUsersByOrganizationNotFound ", 404)
}

func (o *GetUsersByOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersByOrganizationDefault creates a GetUsersByOrganizationDefault with default headers values
func NewGetUsersByOrganizationDefault(code int) *GetUsersByOrganizationDefault {
	return &GetUsersByOrganizationDefault{
		_statusCode: code,
	}
}

/*GetUsersByOrganizationDefault handles this case with default header values.

unexpected error
*/
type GetUsersByOrganizationDefault struct {
	_statusCode int
}

// Code gets the status code for the get users by organization default response
func (o *GetUsersByOrganizationDefault) Code() int {
	return o._statusCode
}

func (o *GetUsersByOrganizationDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/users][%d] getUsersByOrganization default ", o._statusCode)
}

func (o *GetUsersByOrganizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}