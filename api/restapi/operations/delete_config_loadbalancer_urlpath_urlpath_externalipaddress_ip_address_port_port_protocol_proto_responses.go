// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNoContentCode is the HTTP code returned for type DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNoContent
const DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNoContentCode int = 204

/*
DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNoContent OK

swagger:response deleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIpAddressPortPortProtocolProtoNoContent
*/
type DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNoContent struct {
}

// NewDeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNoContent creates DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNoContent with default headers values
func NewDeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNoContent() *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNoContent {

	return &DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNoContent{}
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoBadRequestCode is the HTTP code returned for type DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoBadRequest
const DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoBadRequestCode int = 400

/*
DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoBadRequest Malformed arguments for API call

swagger:response deleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIpAddressPortPortProtocolProtoBadRequest
*/
type DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoBadRequest creates DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoBadRequest with default headers values
func NewDeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoBadRequest() *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoBadRequest {

	return &DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoBadRequest{}
}

// WithPayload adds the payload to the delete config loadbalancer urlpath urlpath externalipaddress Ip address port port protocol proto bad request response
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoBadRequest) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer urlpath urlpath externalipaddress Ip address port port protocol proto bad request response
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoUnauthorizedCode is the HTTP code returned for type DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoUnauthorized
const DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoUnauthorizedCode int = 401

/*
DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoUnauthorized Invalid authentication credentials

swagger:response deleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIpAddressPortPortProtocolProtoUnauthorized
*/
type DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoUnauthorized creates DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoUnauthorized with default headers values
func NewDeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoUnauthorized() *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoUnauthorized {

	return &DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoUnauthorized{}
}

// WithPayload adds the payload to the delete config loadbalancer urlpath urlpath externalipaddress Ip address port port protocol proto unauthorized response
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoUnauthorized) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer urlpath urlpath externalipaddress Ip address port port protocol proto unauthorized response
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoForbiddenCode is the HTTP code returned for type DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoForbidden
const DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoForbiddenCode int = 403

/*
DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoForbidden Capacity insufficient

swagger:response deleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIpAddressPortPortProtocolProtoForbidden
*/
type DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoForbidden creates DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoForbidden with default headers values
func NewDeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoForbidden() *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoForbidden {

	return &DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoForbidden{}
}

// WithPayload adds the payload to the delete config loadbalancer urlpath urlpath externalipaddress Ip address port port protocol proto forbidden response
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoForbidden) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer urlpath urlpath externalipaddress Ip address port port protocol proto forbidden response
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNotFoundCode is the HTTP code returned for type DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNotFound
const DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNotFoundCode int = 404

/*
DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNotFound Resource not found

swagger:response deleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIpAddressPortPortProtocolProtoNotFound
*/
type DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNotFound creates DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNotFound with default headers values
func NewDeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNotFound() *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNotFound {

	return &DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNotFound{}
}

// WithPayload adds the payload to the delete config loadbalancer urlpath urlpath externalipaddress Ip address port port protocol proto not found response
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNotFound) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer urlpath urlpath externalipaddress Ip address port port protocol proto not found response
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoConflictCode is the HTTP code returned for type DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoConflict
const DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoConflictCode int = 409

/*
DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoConflict Resource Conflict. VLAN already exists OR dependency VRF/VNET not found

swagger:response deleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIpAddressPortPortProtocolProtoConflict
*/
type DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoConflict creates DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoConflict with default headers values
func NewDeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoConflict() *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoConflict {

	return &DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoConflict{}
}

// WithPayload adds the payload to the delete config loadbalancer urlpath urlpath externalipaddress Ip address port port protocol proto conflict response
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoConflict) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer urlpath urlpath externalipaddress Ip address port port protocol proto conflict response
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoInternalServerErrorCode is the HTTP code returned for type DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoInternalServerError
const DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoInternalServerErrorCode int = 500

/*
DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoInternalServerError Internal service error

swagger:response deleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIpAddressPortPortProtocolProtoInternalServerError
*/
type DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoInternalServerError creates DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoInternalServerError with default headers values
func NewDeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoInternalServerError() *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoInternalServerError {

	return &DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoInternalServerError{}
}

// WithPayload adds the payload to the delete config loadbalancer urlpath urlpath externalipaddress Ip address port port protocol proto internal server error response
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoInternalServerError) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer urlpath urlpath externalipaddress Ip address port port protocol proto internal server error response
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailableCode is the HTTP code returned for type DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable
const DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailableCode int = 503

/*
DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable Maintanence mode

swagger:response deleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIpAddressPortPortProtocolProtoServiceUnavailable
*/
type DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable creates DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable with default headers values
func NewDeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable() *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable {

	return &DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable{}
}

// WithPayload adds the payload to the delete config loadbalancer urlpath urlpath externalipaddress Ip address port port protocol proto service unavailable response
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer urlpath urlpath externalipaddress Ip address port port protocol proto service unavailable response
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerUrlpathUrlpathExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
