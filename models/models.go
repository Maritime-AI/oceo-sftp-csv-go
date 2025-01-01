package models

import (
	"errors"
	"strings"
)

// Delimiter is the delimiter used to separate endorsements in a CSV file.
const (
	Delimiter = "*|*"
)

// Crew represents the details of a crew member.
type Crew struct {
	CrewExternalID string  `csv:"External ID"`
	FirstName      string  `csv:"First Name"`
	LastName       string  `csv:"Last Name"`
	MiddleName     *string `csv:"Middle Name"`
	JobTitle       *string `csv:"Job Title"`
	City           *string `csv:"City"`
	State          *string `csv:"State"`
	Country        *string `csv:"Country"`
	Email          *string `csv:"Email"`
	Phone          *string `csv:"Phone"`
}

// Validate checks if the required fields of a Crew are set.
func (c *Crew) Validate() error {
	if len(c.CrewExternalID) == 0 {
		return errors.New("missing crew external ID")
	}

	if len(c.FirstName) == 0 {
		return errors.New("missing first name")
	}

	if len(c.LastName) == 0 {
		return errors.New("missing last name")
	}

	return nil
}

// GetLocation returns the location of the crew member.
func (r *Crew) GetLocation() string {
	if r == nil {
		return ""
	}

	var cityStateParts []string
	if r.City != nil {
		cityStateParts = append(cityStateParts, *r.City)
	}

	if r.State != nil {
		cityStateParts = append(cityStateParts, *r.State)
	}

	loc := strings.Join(cityStateParts, ", ")
	if r.Country != nil {
		loc += ", " + *r.Country
	}

	return loc
}

// CrewCredential represents the credentials of a crew member.
type CrewCredential struct {
	CrewExternalID string  `csv:"External ID"`
	Title          string  `csv:"Title"`
	Type           *string `csv:"Type"`
	//Endorsements is a list of endorsements, separated by *|* Delimiter.
	Endorsements string  `csv:"Endorsements"`
	IssuedAt     *string `csv:"Issued At"`
	ExpiresAt    *string `csv:"Expires At"`
}

// Validate checks if the required fields of a CrewCredential are set.
func (cc *CrewCredential) Validate() error {
	if len(cc.CrewExternalID) == 0 {
		return errors.New("missing crew external id")
	}

	if len(cc.Title) == 0 {
		return errors.New("missing title")
	}

	return nil
}

// Vessel represents the details of a vessel.
type Vessel struct {
	VesselExternalID     string  `csv:"External ID"`
	Name                 string  `csv:"Name"`
	MMSINumber           *string `csv:"MMSI Number"`
	IMONumber            *string `csv:"IMO Number"`
	AdditionalIdentifier *string `csv:"Additional Identifier"`
}

// Validate checks if the required fields of a Vessel are set.
func (v *Vessel) Validate() error {
	if len(v.VesselExternalID) == 0 {
		return errors.New("missing vessel external id")
	}

	if len(v.Name) == 0 {
		return errors.New("missing vessel name")
	}

	return nil
}

// VesselSchedule represents the schedule of a vessel.
type VesselSchedule struct {
	VesselExternalID string  `csv:"Vessel External ID"`
	VesselName       string  `csv:"Vessel Name"`
	VesselIMONumber  *string `csv:"Vessel IMO Number"`
	VesselMMSINumber *string `csv:"Vessel MMSI Number"`
	Client           *string `csv:"Client"`
	Description      *string `csv:"Description"`
	ServiceStartAt   string  `csv:"Service Start At"`
	ServiceEndAt     string  `csv:"Service End At"`
}

// Validate checks if the required fields of a VesselSchedule are set.
func (vs *VesselSchedule) Validate() error {

	if len(vs.VesselName) == 0 {
		return errors.New("missing vessel name")
	}

	if len(vs.VesselExternalID) == 0 {
		return errors.New("missing vessel external id")
	}

	if len(vs.ServiceStartAt) == 0 {
		return errors.New("missing service start at")
	}

	if len(vs.ServiceEndAt) == 0 {
		return errors.New("missing service ended at")
	}

	return nil
}

// VesselSchedulePosition represents the position of a vessel.
type VesselSchedulePosition struct {
	VesselExternalID string `csv:"Vessel External ID"`
	Position         string `csv:"Position"`
	CredentialTitle  string `csv:"Credential Title"`
	//Endorsements is a list of endorsements, separated by *|*.
	Endorsements   *string `csv:"Endorsements"`
	ServiceStartAt *string `csv:"Service Start At"`
	ServiceEndAt   *string `csv:"Service End At"`
}

// Validate checks if the required fields of a VesselSchedulePosition are set.
func (vp *VesselSchedulePosition) Validate() error {
	if vp == nil {
		return nil
	}

	if len(vp.VesselExternalID) == 0 {
		return errors.New("missing vessel external id")
	}

	if len(vp.Position) == 0 {
		return errors.New("missing position")
	}

	if len(vp.CredentialTitle) == 0 {
		return errors.New("missing position credential title")
	}

	return nil
}

// CrewSchedule represents the schedule of a crew member.
type CrewSchedule struct {
	CrewExternalID   string  `csv:"Crew External ID"`
	VesselExternalID string  `csv:"Vessel External ID"`
	VesselName       string  `csv:"Vessel Name"`
	VesselIMONumber  *string `csv:"Vessel IMO Number"`
	VesselMMSINumber *string `csv:"Vessel MMSI Number"`
	ServiceStartAt   string  `csv:"Service Start At"`
	ServiceEndAt     string  `csv:"Service End At"`
}

// Validate checks if the required fields of a CrewSchedule are set.
func (vs *CrewSchedule) Validate() error {
	if vs == nil {
		return errors.New("missing crew schedule")
	}

	if len(vs.VesselExternalID) == 0 {
		return errors.New("missing vessel external id")
	}

	if len(vs.CrewExternalID) == 0 {
		return errors.New("missing crew external id")
	}

	if len(vs.VesselName) == 0 {
		return errors.New("missing vessel name")
	}

	if len(vs.ServiceStartAt) == 0 {
		return errors.New("missing service started at")
	}

	if len(vs.ServiceEndAt) == 0 {
		return errors.New("missing service ended at")
	}

	return nil
}

// CrewSchedulePosition represents the position details of a crew member in a schedule.
type CrewSchedulePosition struct {
	CrewExternalID   string `csv:"Crew External ID"`
	VesselExternalID string `csv:"Vessel External ID"`
	Position         string `csv:"Position"`
	CredentialTitle  string `csv:"Credential Title"`
	//Endorsements is a list of endorsements, separated by *|* Delimiter.
	Endorsements   string  `csv:"Endorsements"`
	ServiceStartAt *string `csv:"Service Start At"`
	ServiceEndAt   *string `csv:"Service End At"`
}

// Validate checks if the required fields of a CrewSchedulePosition are set.
func (vs *CrewSchedulePosition) Validate() error {
	if vs == nil {
		return errors.New("missing crew schedule position")
	}

	if len(vs.VesselExternalID) == 0 {
		return errors.New("missing vessel external id")
	}

	if len(vs.CrewExternalID) == 0 {
		return errors.New("missing crew external id")
	}

	if len(vs.Position) == 0 {
		return errors.New("missing position")
	}

	if len(vs.CredentialTitle) == 0 {
		return errors.New("missing credential")
	}

	return nil
}
