// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteConfigEndpointEpipaddressIPAddressParams creates a new DeleteConfigEndpointEpipaddressIPAddressParams object
//
// There are no default values defined in the spec.
func NewDeleteConfigEndpointEpipaddressIPAddressParams() DeleteConfigEndpointEpipaddressIPAddressParams {

	return DeleteConfigEndpointEpipaddressIPAddressParams{}
}

// DeleteConfigEndpointEpipaddressIPAddressParams contains all the bound params for the delete config endpoint epipaddress IP address operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteConfigEndpointEpipaddressIPAddress
type DeleteConfigEndpointEpipaddressIPAddressParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Attributes of end point
	  Required: true
	  In: path
	*/
	IPAddress string
	/*Endpoint Identifier
	  In: query
	*/
	Name *string
	/*Probe port
	  In: query
	*/
	ProbePort *float64
	/*Probe type
	  In: query
	*/
	ProbeType *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteConfigEndpointEpipaddressIPAddressParams() beforehand.
func (o *DeleteConfigEndpointEpipaddressIPAddressParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rIPAddress, rhkIPAddress, _ := route.Params.GetOK("ip_address")
	if err := o.bindIPAddress(rIPAddress, rhkIPAddress, route.Formats); err != nil {
		res = append(res, err)
	}

	qName, qhkName, _ := qs.GetOK("name")
	if err := o.bindName(qName, qhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	qProbePort, qhkProbePort, _ := qs.GetOK("probe_port")
	if err := o.bindProbePort(qProbePort, qhkProbePort, route.Formats); err != nil {
		res = append(res, err)
	}

	qProbeType, qhkProbeType, _ := qs.GetOK("probe_type")
	if err := o.bindProbeType(qProbeType, qhkProbeType, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIPAddress binds and validates parameter IPAddress from path.
func (o *DeleteConfigEndpointEpipaddressIPAddressParams) bindIPAddress(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.IPAddress = raw

	return nil
}

// bindName binds and validates parameter Name from query.
func (o *DeleteConfigEndpointEpipaddressIPAddressParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Name = &raw

	return nil
}

// bindProbePort binds and validates parameter ProbePort from query.
func (o *DeleteConfigEndpointEpipaddressIPAddressParams) bindProbePort(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("probe_port", "query", "float64", raw)
	}
	o.ProbePort = &value

	return nil
}

// bindProbeType binds and validates parameter ProbeType from query.
func (o *DeleteConfigEndpointEpipaddressIPAddressParams) bindProbeType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.ProbeType = &raw

	return nil
}
