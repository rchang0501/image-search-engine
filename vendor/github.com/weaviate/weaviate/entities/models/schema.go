//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Schema Definitions of semantic schemas (also see: https://github.com/weaviate/weaviate-semantic-schemas).
//
// swagger:model Schema
type Schema struct {

	// Semantic classes that are available.
	Classes []*Class `json:"classes"`

	// Email of the maintainer.
	// Format: email
	Maintainer strfmt.Email `json:"maintainer,omitempty"`

	// Name of the schema.
	Name string `json:"name,omitempty"`
}

// Validate validates this schema
func (m *Schema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClasses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintainer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Schema) validateClasses(formats strfmt.Registry) error {
	if swag.IsZero(m.Classes) { // not required
		return nil
	}

	for i := 0; i < len(m.Classes); i++ {
		if swag.IsZero(m.Classes[i]) { // not required
			continue
		}

		if m.Classes[i] != nil {
			if err := m.Classes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("classes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("classes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Schema) validateMaintainer(formats strfmt.Registry) error {
	if swag.IsZero(m.Maintainer) { // not required
		return nil
	}

	if err := validate.FormatOf("maintainer", "body", "email", m.Maintainer.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this schema based on the context it is used
func (m *Schema) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClasses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Schema) contextValidateClasses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Classes); i++ {

		if m.Classes[i] != nil {
			if err := m.Classes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("classes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("classes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Schema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Schema) UnmarshalBinary(b []byte) error {
	var res Schema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}