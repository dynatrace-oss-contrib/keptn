// Code generated by go-swagger; DO NOT EDIT.

package configure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/keptn/keptn/api/models"
)

// PostConfigureBridgeExposeOKCode is the HTTP code returned for type PostConfigureBridgeExposeOK
const PostConfigureBridgeExposeOKCode int = 200

/*PostConfigureBridgeExposeOK Bridge was successfully exposed/disposed

swagger:response postConfigureBridgeExposeOK
*/
type PostConfigureBridgeExposeOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostConfigureBridgeExposeOK creates PostConfigureBridgeExposeOK with default headers values
func NewPostConfigureBridgeExposeOK() *PostConfigureBridgeExposeOK {

	return &PostConfigureBridgeExposeOK{}
}

// WithPayload adds the payload to the post configure bridge expose o k response
func (o *PostConfigureBridgeExposeOK) WithPayload(payload string) *PostConfigureBridgeExposeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post configure bridge expose o k response
func (o *PostConfigureBridgeExposeOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigureBridgeExposeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostConfigureBridgeExposeBadRequestCode is the HTTP code returned for type PostConfigureBridgeExposeBadRequest
const PostConfigureBridgeExposeBadRequestCode int = 400

/*PostConfigureBridgeExposeBadRequest Bridge could not be exposed/disposed

swagger:response postConfigureBridgeExposeBadRequest
*/
type PostConfigureBridgeExposeBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigureBridgeExposeBadRequest creates PostConfigureBridgeExposeBadRequest with default headers values
func NewPostConfigureBridgeExposeBadRequest() *PostConfigureBridgeExposeBadRequest {

	return &PostConfigureBridgeExposeBadRequest{}
}

// WithPayload adds the payload to the post configure bridge expose bad request response
func (o *PostConfigureBridgeExposeBadRequest) WithPayload(payload *models.Error) *PostConfigureBridgeExposeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post configure bridge expose bad request response
func (o *PostConfigureBridgeExposeBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigureBridgeExposeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostConfigureBridgeExposeDefault Error

swagger:response postConfigureBridgeExposeDefault
*/
type PostConfigureBridgeExposeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigureBridgeExposeDefault creates PostConfigureBridgeExposeDefault with default headers values
func NewPostConfigureBridgeExposeDefault(code int) *PostConfigureBridgeExposeDefault {
	if code <= 0 {
		code = 500
	}

	return &PostConfigureBridgeExposeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post configure bridge expose default response
func (o *PostConfigureBridgeExposeDefault) WithStatusCode(code int) *PostConfigureBridgeExposeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post configure bridge expose default response
func (o *PostConfigureBridgeExposeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post configure bridge expose default response
func (o *PostConfigureBridgeExposeDefault) WithPayload(payload *models.Error) *PostConfigureBridgeExposeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post configure bridge expose default response
func (o *PostConfigureBridgeExposeDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigureBridgeExposeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
