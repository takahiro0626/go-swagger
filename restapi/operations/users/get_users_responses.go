// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"go-swagger/models"
)

// GetUsersOKCode is the HTTP code returned for type GetUsersOK
const GetUsersOKCode int = 200

/*
GetUsersOK OK

swagger:response getUsersOK
*/
type GetUsersOK struct {

	/*
	  In: Body
	*/
	Payload *GetUsersOKBody `json:"body,omitempty"`
}

// NewGetUsersOK creates GetUsersOK with default headers values
func NewGetUsersOK() *GetUsersOK {

	return &GetUsersOK{}
}

// WithPayload adds the payload to the get users o k response
func (o *GetUsersOK) WithPayload(payload *GetUsersOKBody) *GetUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users o k response
func (o *GetUsersOK) SetPayload(payload *GetUsersOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUsersBadRequestCode is the HTTP code returned for type GetUsersBadRequest
const GetUsersBadRequestCode int = 400

/*
GetUsersBadRequest Bad Request

swagger:response getUsersBadRequest
*/
type GetUsersBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.MainErrorResponse `json:"body,omitempty"`
}

// NewGetUsersBadRequest creates GetUsersBadRequest with default headers values
func NewGetUsersBadRequest() *GetUsersBadRequest {

	return &GetUsersBadRequest{}
}

// WithPayload adds the payload to the get users bad request response
func (o *GetUsersBadRequest) WithPayload(payload *models.MainErrorResponse) *GetUsersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users bad request response
func (o *GetUsersBadRequest) SetPayload(payload *models.MainErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUsersNotFoundCode is the HTTP code returned for type GetUsersNotFound
const GetUsersNotFoundCode int = 404

/*
GetUsersNotFound Not Found

swagger:response getUsersNotFound
*/
type GetUsersNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.MainErrorResponse `json:"body,omitempty"`
}

// NewGetUsersNotFound creates GetUsersNotFound with default headers values
func NewGetUsersNotFound() *GetUsersNotFound {

	return &GetUsersNotFound{}
}

// WithPayload adds the payload to the get users not found response
func (o *GetUsersNotFound) WithPayload(payload *models.MainErrorResponse) *GetUsersNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users not found response
func (o *GetUsersNotFound) SetPayload(payload *models.MainErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUsersUnprocessableEntityCode is the HTTP code returned for type GetUsersUnprocessableEntity
const GetUsersUnprocessableEntityCode int = 422

/*
GetUsersUnprocessableEntity Unprocessable Entity

swagger:response getUsersUnprocessableEntity
*/
type GetUsersUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.MainErrorResponse `json:"body,omitempty"`
}

// NewGetUsersUnprocessableEntity creates GetUsersUnprocessableEntity with default headers values
func NewGetUsersUnprocessableEntity() *GetUsersUnprocessableEntity {

	return &GetUsersUnprocessableEntity{}
}

// WithPayload adds the payload to the get users unprocessable entity response
func (o *GetUsersUnprocessableEntity) WithPayload(payload *models.MainErrorResponse) *GetUsersUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users unprocessable entity response
func (o *GetUsersUnprocessableEntity) SetPayload(payload *models.MainErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
