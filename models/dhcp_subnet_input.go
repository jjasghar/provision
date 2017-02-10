package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DhcpSubnetInput DHCP Subnet Input
// swagger:model dhcp-subnet-input
type DhcpSubnetInput struct {

	// active end
	// Required: true
	ActiveEnd *strfmt.IPv4 `json:"ActiveEnd"`

	// active lease time
	// Required: true
	ActiveLeaseTime *int64 `json:"ActiveLeaseTime"`

	// active start
	// Required: true
	ActiveStart *strfmt.IPv4 `json:"ActiveStart"`

	// name
	// Required: true
	Name *string `json:"Name"`

	// next server
	NextServer strfmt.IPv4 `json:"NextServer,omitempty"`

	// only bound leases
	OnlyBoundLeases *bool `json:"OnlyBoundLeases,omitempty"`

	// options
	Options []*Dhcpoption `json:"Options"`

	// reserved lease time
	// Required: true
	ReservedLeaseTime *int64 `json:"ReservedLeaseTime"`

	// subnet
	// Required: true
	// Pattern: ^([0-9]{1-3}\.){,3}[0-9]{,3}/[0-9]{,2}$
	Subnet *string `json:"Subnet"`
}

// Validate validates this dhcp subnet input
func (m *DhcpSubnetInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveEnd(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateActiveLeaseTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateActiveStart(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReservedLeaseTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DhcpSubnetInput) validateActiveEnd(formats strfmt.Registry) error {

	if err := validate.Required("ActiveEnd", "body", m.ActiveEnd); err != nil {
		return err
	}

	return nil
}

func (m *DhcpSubnetInput) validateActiveLeaseTime(formats strfmt.Registry) error {

	if err := validate.Required("ActiveLeaseTime", "body", m.ActiveLeaseTime); err != nil {
		return err
	}

	return nil
}

func (m *DhcpSubnetInput) validateActiveStart(formats strfmt.Registry) error {

	if err := validate.Required("ActiveStart", "body", m.ActiveStart); err != nil {
		return err
	}

	return nil
}

func (m *DhcpSubnetInput) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DhcpSubnetInput) validateOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Options) { // not required
		return nil
	}

	for i := 0; i < len(m.Options); i++ {

		if swag.IsZero(m.Options[i]) { // not required
			continue
		}

		if m.Options[i] != nil {

			if err := m.Options[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *DhcpSubnetInput) validateReservedLeaseTime(formats strfmt.Registry) error {

	if err := validate.Required("ReservedLeaseTime", "body", m.ReservedLeaseTime); err != nil {
		return err
	}

	return nil
}

func (m *DhcpSubnetInput) validateSubnet(formats strfmt.Registry) error {

	if err := validate.Required("Subnet", "body", m.Subnet); err != nil {
		return err
	}

	if err := validate.Pattern("Subnet", "body", string(*m.Subnet), `^([0-9]{1-3}\.){,3}[0-9]{,3}/[0-9]{,2}$`); err != nil {
		return err
	}

	return nil
}