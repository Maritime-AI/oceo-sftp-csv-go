package models

import (
	"errors"
	"strings"
	"time"
)

// Delimiter is the delimiter used to separate endorsements in a CSV file.
const (
	Delimiter = "*|*"
)

// Crew represents the details of a crew member.
type Crew struct {
	ContextID      string  `csv:"Context ID" json:"context_id"`
	CrewExternalID string  `csv:"Crew External ID" json:"crew_external_id"`
	FirstName      string  `csv:"First Name" json:"first_name"`
	LastName       string  `csv:"Last Name" json:"last_name"`
	MiddleName     *string `csv:"Middle Name" json:"middle_name"`
	JobTitle       *string `csv:"Job Title" json:"job_title"`
	City           *string `csv:"City" json:"city"`
	State          *string `csv:"State" json:"state"`
	Country        *string `csv:"Country" json:"country"`
	Email          *string `csv:"Email" json:"email"`
	Phone          *string `csv:"Phone" json:"phone"`
}

// Validate checks if the required fields of a Crew are set.
func (c *Crew) Validate() error {
	if c == nil {
		return errors.New("missing crew")
	}

	if len(c.ContextID) == 0 {
		return errors.New("missing context ID")
	}

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
	ContextID      string  `csv:"Context ID" json:"context_id"`
	CrewExternalID string  `csv:"Crew External ID" json:"crew_external_id"`
	Number         *string `csv:"Number" json:"number"`
	Title          string  `csv:"Title" json:"title"`
	Type           *string `csv:"Type" json:"type"`
	//Endorsements is a list of endorsements, separated by *|* Delimiter.
	Endorsements string  `csv:"Endorsements" json:"endorsements"`
	IssuedAt     *string `csv:"Issued At" json:"issued_at"`
	ExpiresAt    *string `csv:"Expires At" json:"expires_at"`
}

// Validate checks if the required fields of a CrewCredential are set.
func (cc *CrewCredential) Validate() error {
	if cc == nil {
		return errors.New("missing crew credential")
	}

	if len(cc.ContextID) == 0 {
		return errors.New("missing context id")
	}

	if len(cc.CrewExternalID) == 0 {
		return errors.New("missing crew external id")
	}

	if len(cc.Title) == 0 {
		return errors.New("missing title")
	}

	return nil
}

type CrewSeatime struct {
	ContextID             string     `csv:"Context ID" json:"context_id"`
	CrewExternalID        string     `csv:"Crew External ID" json:"crew_external_id"`
	CrewedOn              *time.Time `csv:"Crew On" json:"crew_on"`
	CrewedOff             *time.Time `csv:"Crew Off" json:"crew_off"`
	NumDays               *float64   `csv:"Num Days" json:"num_days"`
	Position              *string    `csv:"Position" json:"position"`
	ShiftInHours          *int64     `csv:"Shift In Hours" json:"shift_in_hours"`
	VesselName            string     `csv:"Vessel Name" json:"vessel_name"`
	VesselFlag            *string    `csv:"Vessel Flag" json:"vessel_flag"`
	VesselType            *string    `csv:"Vessel Type" json:"vessel_type"`
	VesselCapacityGT      *int64     `csv:"Vessel Capacity GT" json:"vessel_capacity_gt"`
	VesselHorsePower      *float64   `csv:"Vessel Horse Power" json:"vessel_horse_power"`
	VesselPropulsionType  *string    `csv:"Vessel Propulsion Type" json:"vessel_propulsion_type"`
	VesselIMONumber       *int64     `csv:"Vessel IMO Number" json:"vessel_imo_number"`
	VesselMMSINumber      *int64     `csv:"Vessel MMSI Number" json:"vessel_mmsi_number"`
	Compensation          *string    `csv:"Compensation" json:"compensation"`
	CompensationFrequency *string    `csv:"Compensation Frequency" json:"compensation_frequency"`
	CompanyName           *string    `csv:"Company Name" json:"company_name"`
}

// Validate checks if the required fields of a CrewCredential are set.
func (st *CrewSeatime) Validate() error {
	if st == nil {
		return errors.New("missing crew credential")
	}

	if len(st.ContextID) == 0 {
		return errors.New("missing context id")
	}

	if len(st.CrewExternalID) == 0 {
		return errors.New("missing crew external id")
	}

	if len(st.VesselName) == 0 {
		return errors.New("missing vessel name")
	}

	isCrewOnAndOff := st.CrewedOn != nil && st.CrewedOff != nil
	isNumDays := st.NumDays != nil && *st.NumDays > 0
	if !isCrewOnAndOff && !isNumDays {
		return errors.New("must provide either crew on/off or num days")
	}

	return nil
}

// NumDaysWorked returns the number of days the mariner was at sea
func (st *CrewSeatime) NumDaysWorked() float64 {

	// either start date or num days needs
	// to be defined in order to calculate
	// the sea time
	switch {
	case st.NumDays != nil:
		nd := float64(*st.NumDays)

		if st.ShiftInHours != nil {
			switch *st.ShiftInHours {
			case 8:
				return nd

			case 12:
				return 1.5 * nd
			default:
				return 0
			}
		}

		return nd
	case st.CrewedOn != nil:
		endAt := time.Now()
		if st.CrewedOff != nil {
			endAt = *st.CrewedOff
		}

		endAt = endAt.Add(time.Hour * 24)
		d := endAt.Sub(*st.CrewedOn)
		days := float64(int64(d.Hours() / 24))

		switch *st.ShiftInHours {
		case 8: // 8hr counts for 1 day
			return days
		default: // 12hr counts for 1.5 days
			return days * 1.5
		}
	default:
		return 0
	}
}

// Vessel represents the details of a vessel.
type Vessel struct {
	ContextID            string  `csv:"Context ID"`
	ExternalID           string  `csv:"External ID"`
	VesselExternalID     string  `csv:"Vessel External ID"`
	Name                 string  `csv:"Name"`
	MMSINumber           *string `csv:"MMSI Number"`
	IMONumber            *string `csv:"IMO Number"`
	AdditionalIdentifier *string `csv:"Additional Identifier"`
}

// Validate checks if the required fields of a Vessel are set.
func (v *Vessel) Validate() error {
	if v == nil {
		return errors.New("missing vessel")
	}

	if len(v.ContextID) == 0 {
		return errors.New("missing context id")
	}

	if len(v.ExternalID) == 0 {
		return errors.New("missing external id")
	}

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
	ContextID        string  `csv:"Context ID"`
	ExternalID       string  `csv:"External ID"`
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
	if vs == nil {
		return errors.New("missing vessel schedule")
	}

	if len(vs.ContextID) == 0 {
		return errors.New("missing context id")
	}

	if len(vs.ExternalID) == 0 {
		return errors.New("missing external id")
	}

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
	ContextID        string `csv:"Context ID"`
	ExternalID       string `csv:"External ID"`
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
		return errors.New("missing vessel schedule position")
	}

	if len(vp.ExternalID) == 0 {
		return errors.New("missing external id")
	}

	if len(vp.ContextID) == 0 {
		return errors.New("missing context id")
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
	ContextID        string  `csv:"Context ID"`
	ExternalID       string  `csv:"External ID"`
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

	if len(vs.ContextID) == 0 {
		return errors.New("missing context id")
	}

	if len(vs.ExternalID) == 0 {
		return errors.New("missing external id")
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
	ContextID        string `csv:"Context ID"`
	ExternalID       string `csv:"External ID"`
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

	if len(vs.ContextID) == 0 {
		return errors.New("missing context id")
	}

	if len(vs.ExternalID) == 0 {
		return errors.New("missing external id")
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
