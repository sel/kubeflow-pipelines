// Code generated by go-swagger; DO NOT EDIT.

package run_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1beta1ListRunsResponse v1beta1 list runs response
// swagger:model v1beta1ListRunsResponse
type V1beta1ListRunsResponse struct {

	// The token to list the next page of runs.
	NextPageToken string `json:"next_page_token,omitempty"`

	// runs
	Runs []*V1beta1Run `json:"runs"`

	// The total number of runs for the given query.
	TotalSize int32 `json:"total_size,omitempty"`
}

// Validate validates this v1beta1 list runs response
func (m *V1beta1ListRunsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRuns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1ListRunsResponse) validateRuns(formats strfmt.Registry) error {

	if swag.IsZero(m.Runs) { // not required
		return nil
	}

	for i := 0; i < len(m.Runs); i++ {
		if swag.IsZero(m.Runs[i]) { // not required
			continue
		}

		if m.Runs[i] != nil {
			if err := m.Runs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("runs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1ListRunsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1ListRunsResponse) UnmarshalBinary(b []byte) error {
	var res V1beta1ListRunsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
