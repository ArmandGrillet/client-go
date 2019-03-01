// Code generated by go-swagger; DO NOT EDIT.

package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteAclsRidGroupsGidActionReader is a Reader for the DeleteAclsRidGroupsGidAction structure.
type DeleteAclsRidGroupsGidActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAclsRidGroupsGidActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteAclsRidGroupsGidActionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAclsRidGroupsGidActionNoContent creates a DeleteAclsRidGroupsGidActionNoContent with default headers values
func NewDeleteAclsRidGroupsGidActionNoContent() *DeleteAclsRidGroupsGidActionNoContent {
	return &DeleteAclsRidGroupsGidActionNoContent{}
}

/*DeleteAclsRidGroupsGidActionNoContent handles this case with default header values.

Success.
*/
type DeleteAclsRidGroupsGidActionNoContent struct {
}

func (o *DeleteAclsRidGroupsGidActionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /acls/{rid}/groups/{gid}/{action}][%d] deleteAclsRidGroupsGidActionNoContent ", 204)
}

func (o *DeleteAclsRidGroupsGidActionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
