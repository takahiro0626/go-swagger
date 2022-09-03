// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"go-swagger/models"
)

// PutUsersUserIDOKCode is the HTTP code returned for type PutUsersUserIDOK
const PutUsersUserIDOKCode int = 200

/*
PutUsersUserIDOK OK

swagger:response putUsersUserIdOK
*/
type PutUsersUserIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.MainUser `json:"body,omitempty"`
}

// NewPutUsersUserIDOK creates PutUsersUserIDOK with default headers values
func NewPutUsersUserIDOK() *PutUsersUserIDOK {

	return &PutUsersUserIDOK{}
}

// WithPayload adds the payload to the put users user Id o k response
func (o *PutUsersUserIDOK) WithPayload(payload *models.MainUser) *PutUsersUserIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put users user Id o k response
func (o *PutUsersUserIDOK) SetPayload(payload *models.MainUser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutUsersUserIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutUsersUserIDBadRequestCode is the HTTP code returned for type PutUsersUserIDBadRequest
const PutUsersUserIDBadRequestCode int = 400

/*
PutUsersUserIDBadRequest Bad Request

swagger:response putUsersUserIdBadRequest
*/
type PutUsersUserIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.MainErrorResponse `json:"body,omitempty"`
}

// NewPutUsersUserIDBadRequest creates PutUsersUserIDBadRequest with default headers values
func NewPutUsersUserIDBadRequest() *PutUsersUserIDBadRequest {

	return &PutUsersUserIDBadRequest{}
}

// WithPayload adds the payload to the put users user Id bad request response
func (o *PutUsersUserIDBadRequest) WithPayload(payload *models.MainErrorResponse) *PutUsersUserIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put users user Id bad request response
func (o *PutUsersUserIDBadRequest) SetPayload(payload *models.MainErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutUsersUserIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutUsersUserIDUnprocessableEntityCode is the HTTP code returned for type PutUsersUserIDUnprocessableEntity
const PutUsersUserIDUnprocessableEntityCode int = 422

/*
PutUsersUserIDUnprocessableEntity Unprocessable Entity

swagger:response putUsersUserIdUnprocessableEntity
*/
type PutUsersUserIDUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.MainErrorResponse `json:"body,omitempty"`
}

// NewPutUsersUserIDUnprocessableEntity creates PutUsersUserIDUnprocessableEntity with default headers values
func NewPutUsersUserIDUnprocessableEntity() *PutUsersUserIDUnprocessableEntity {

	return &PutUsersUserIDUnprocessableEntity{}
}

// WithPayload adds the payload to the put users user Id unprocessable entity response
func (o *PutUsersUserIDUnprocessableEntity) WithPayload(payload *models.MainErrorResponse) *PutUsersUserIDUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put users user Id unprocessable entity response
func (o *PutUsersUserIDUnprocessableEntity) SetPayload(payload *models.MainErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutUsersUserIDUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
