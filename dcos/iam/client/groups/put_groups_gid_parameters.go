// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/dcos/client-go/dcos/iam/models"
)

// NewPutGroupsGidParams creates a new PutGroupsGidParams object
// with the default values initialized.
func NewPutGroupsGidParams() *PutGroupsGidParams {
	var ()
	return &PutGroupsGidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutGroupsGidParamsWithTimeout creates a new PutGroupsGidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutGroupsGidParamsWithTimeout(timeout time.Duration) *PutGroupsGidParams {
	var ()
	return &PutGroupsGidParams{

		timeout: timeout,
	}
}

// NewPutGroupsGidParamsWithContext creates a new PutGroupsGidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutGroupsGidParamsWithContext(ctx context.Context) *PutGroupsGidParams {
	var ()
	return &PutGroupsGidParams{

		Context: ctx,
	}
}

// NewPutGroupsGidParamsWithHTTPClient creates a new PutGroupsGidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutGroupsGidParamsWithHTTPClient(client *http.Client) *PutGroupsGidParams {
	var ()
	return &PutGroupsGidParams{
		HTTPClient: client,
	}
}

/*PutGroupsGidParams contains all the parameters to send to the API endpoint
for the put groups gid operation typically these are written to a http.Request
*/
type PutGroupsGidParams struct {

	/*GroupCreationObject*/
	GroupCreationObject *models.GroupCreate
	/*Gid
	  The ID of the group.

	*/
	Gid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put groups gid params
func (o *PutGroupsGidParams) WithTimeout(timeout time.Duration) *PutGroupsGidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put groups gid params
func (o *PutGroupsGidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put groups gid params
func (o *PutGroupsGidParams) WithContext(ctx context.Context) *PutGroupsGidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put groups gid params
func (o *PutGroupsGidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put groups gid params
func (o *PutGroupsGidParams) WithHTTPClient(client *http.Client) *PutGroupsGidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put groups gid params
func (o *PutGroupsGidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupCreationObject adds the groupCreationObject to the put groups gid params
func (o *PutGroupsGidParams) WithGroupCreationObject(groupCreationObject *models.GroupCreate) *PutGroupsGidParams {
	o.SetGroupCreationObject(groupCreationObject)
	return o
}

// SetGroupCreationObject adds the groupCreationObject to the put groups gid params
func (o *PutGroupsGidParams) SetGroupCreationObject(groupCreationObject *models.GroupCreate) {
	o.GroupCreationObject = groupCreationObject
}

// WithGid adds the gid to the put groups gid params
func (o *PutGroupsGidParams) WithGid(gid string) *PutGroupsGidParams {
	o.SetGid(gid)
	return o
}

// SetGid adds the gid to the put groups gid params
func (o *PutGroupsGidParams) SetGid(gid string) {
	o.Gid = gid
}

// WriteToRequest writes these params to a swagger request
func (o *PutGroupsGidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.GroupCreationObject != nil {
		if err := r.SetBodyParam(o.GroupCreationObject); err != nil {
			return err
		}
	}

	// path param gid
	if err := r.SetPathParam("gid", o.Gid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
