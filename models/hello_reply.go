// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HelloReply hello reply
//
// swagger:model hello_reply
type HelloReply struct {

	// reply
	Reply string `json:"reply,omitempty"`
}

// Validate validates this hello reply
func (m *HelloReply) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hello reply based on context it is used
func (m *HelloReply) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HelloReply) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HelloReply) UnmarshalBinary(b []byte) error {
	var res HelloReply
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}