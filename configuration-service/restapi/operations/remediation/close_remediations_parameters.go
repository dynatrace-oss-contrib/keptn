// Code generated by go-swagger; DO NOT EDIT.

package remediation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCloseRemediationsParams creates a new CloseRemediationsParams object
// with the default values initialized.
func NewCloseRemediationsParams() CloseRemediationsParams {

	var (
		// initialize parameters with default values

		pageSizeDefault = int64(20)
	)

	return CloseRemediationsParams{
		PageSize: &pageSizeDefault,
	}
}

// CloseRemediationsParams contains all the bound params for the close remediations operation
// typically these are obtained from a http.Request
//
// swagger:parameters closeRemediations
type CloseRemediationsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Keptn Context
	  Required: true
	  In: path
	*/
	KeptnContext string
	/*Pointer to the next set of items
	  In: query
	*/
	NextPageKey *string
	/*The number of items to return
	  Maximum: 50
	  Minimum: 1
	  In: query
	  Default: 20
	*/
	PageSize *int64
	/*Name of the project
	  Required: true
	  In: path
	*/
	ProjectName string
	/*Name of the service
	  Required: true
	  In: path
	*/
	ServiceName string
	/*Name of the stage
	  Required: true
	  In: path
	*/
	StageName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCloseRemediationsParams() beforehand.
func (o *CloseRemediationsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rKeptnContext, rhkKeptnContext, _ := route.Params.GetOK("keptnContext")
	if err := o.bindKeptnContext(rKeptnContext, rhkKeptnContext, route.Formats); err != nil {
		res = append(res, err)
	}

	qNextPageKey, qhkNextPageKey, _ := qs.GetOK("nextPageKey")
	if err := o.bindNextPageKey(qNextPageKey, qhkNextPageKey, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageSize, qhkPageSize, _ := qs.GetOK("pageSize")
	if err := o.bindPageSize(qPageSize, qhkPageSize, route.Formats); err != nil {
		res = append(res, err)
	}

	rProjectName, rhkProjectName, _ := route.Params.GetOK("projectName")
	if err := o.bindProjectName(rProjectName, rhkProjectName, route.Formats); err != nil {
		res = append(res, err)
	}

	rServiceName, rhkServiceName, _ := route.Params.GetOK("serviceName")
	if err := o.bindServiceName(rServiceName, rhkServiceName, route.Formats); err != nil {
		res = append(res, err)
	}

	rStageName, rhkStageName, _ := route.Params.GetOK("stageName")
	if err := o.bindStageName(rStageName, rhkStageName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindKeptnContext binds and validates parameter KeptnContext from path.
func (o *CloseRemediationsParams) bindKeptnContext(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.KeptnContext = raw

	return nil
}

// bindNextPageKey binds and validates parameter NextPageKey from query.
func (o *CloseRemediationsParams) bindNextPageKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.NextPageKey = &raw

	return nil
}

// bindPageSize binds and validates parameter PageSize from query.
func (o *CloseRemediationsParams) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewCloseRemediationsParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("pageSize", "query", "int64", raw)
	}
	o.PageSize = &value

	if err := o.validatePageSize(formats); err != nil {
		return err
	}

	return nil
}

// validatePageSize carries on validations for parameter PageSize
func (o *CloseRemediationsParams) validatePageSize(formats strfmt.Registry) error {

	if err := validate.MinimumInt("pageSize", "query", int64(*o.PageSize), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("pageSize", "query", int64(*o.PageSize), 50, false); err != nil {
		return err
	}

	return nil
}

// bindProjectName binds and validates parameter ProjectName from path.
func (o *CloseRemediationsParams) bindProjectName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProjectName = raw

	return nil
}

// bindServiceName binds and validates parameter ServiceName from path.
func (o *CloseRemediationsParams) bindServiceName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ServiceName = raw

	return nil
}

// bindStageName binds and validates parameter StageName from path.
func (o *CloseRemediationsParams) bindStageName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.StageName = raw

	return nil
}
