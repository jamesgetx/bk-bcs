// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Info Information
//
// General API, OS and hardware information
// swagger:model info
type Info struct {

	// api
	API *InfoAPI `json:"api,omitempty"`

	// system
	System *InfoSystem `json:"system,omitempty"`
}

// Validate validates this info
func (m *Info) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Info) validateAPI(formats strfmt.Registry) error {

	if swag.IsZero(m.API) { // not required
		return nil
	}

	if m.API != nil {
		if err := m.API.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("api")
			}
			return err
		}
	}

	return nil
}

func (m *Info) validateSystem(formats strfmt.Registry) error {

	if swag.IsZero(m.System) { // not required
		return nil
	}

	if m.System != nil {
		if err := m.System.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Info) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Info) UnmarshalBinary(b []byte) error {
	var res Info
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InfoAPI info API
// swagger:model InfoAPI
type InfoAPI struct {

	// HAProxy Dataplane API build date
	// Format: date-time
	BuildDate strfmt.DateTime `json:"build_date,omitempty"`

	// HAProxy Dataplane API version string
	Version string `json:"version,omitempty"`
}

// Validate validates this info API
func (m *InfoAPI) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfoAPI) validateBuildDate(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildDate) { // not required
		return nil
	}

	if err := validate.FormatOf("api"+"."+"build_date", "body", "date-time", m.BuildDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InfoAPI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfoAPI) UnmarshalBinary(b []byte) error {
	var res InfoAPI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InfoSystem info system
// swagger:model InfoSystem
type InfoSystem struct {

	// cpu info
	CPUInfo *InfoSystemCPUInfo `json:"cpu_info,omitempty"`

	// Hostname where the HAProxy is running
	Hostname string `json:"hostname,omitempty"`

	// mem info
	MemInfo *InfoSystemMemInfo `json:"mem_info,omitempty"`

	// OS string
	OsString string `json:"os_string,omitempty"`

	// Current time in milliseconds since Epoch.
	Time int64 `json:"time,omitempty"`

	// System uptime
	Uptime *int64 `json:"uptime,omitempty"`
}

// Validate validates this info system
func (m *InfoSystem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfoSystem) validateCPUInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.CPUInfo) { // not required
		return nil
	}

	if m.CPUInfo != nil {
		if err := m.CPUInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system" + "." + "cpu_info")
			}
			return err
		}
	}

	return nil
}

func (m *InfoSystem) validateMemInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.MemInfo) { // not required
		return nil
	}

	if m.MemInfo != nil {
		if err := m.MemInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system" + "." + "mem_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InfoSystem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfoSystem) UnmarshalBinary(b []byte) error {
	var res InfoSystem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InfoSystemCPUInfo info system CPU info
// swagger:model InfoSystemCPUInfo
type InfoSystemCPUInfo struct {

	// model
	Model string `json:"model,omitempty"`

	// Number of logical CPUs
	NumCpus int64 `json:"num_cpus,omitempty"`
}

// Validate validates this info system CPU info
func (m *InfoSystemCPUInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InfoSystemCPUInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfoSystemCPUInfo) UnmarshalBinary(b []byte) error {
	var res InfoSystemCPUInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InfoSystemMemInfo info system mem info
// swagger:model InfoSystemMemInfo
type InfoSystemMemInfo struct {

	// dataplaneapi memory
	DataplaneapiMemory int64 `json:"dataplaneapi_memory,omitempty"`

	// free memory
	FreeMemory int64 `json:"free_memory,omitempty"`

	// total memory
	TotalMemory int64 `json:"total_memory,omitempty"`
}

// Validate validates this info system mem info
func (m *InfoSystemMemInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InfoSystemMemInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfoSystemMemInfo) UnmarshalBinary(b []byte) error {
	var res InfoSystemMemInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}