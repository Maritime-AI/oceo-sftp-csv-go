package models

// Crew represents the details of a crew member.
type Crew struct {
	ExternalID string  `csv:"External ID"`
	FirstName  *string `csv:"First Name"`
	LastName   *string `csv:"Last Name"`
	MiddleName *string `csv:"Middle Name"`
	JobTitle   *string `csv:"Job Title"`
	City       *string `csv:"City"`
	State      *string `csv:"State"`
	Country    *string `csv:"Country"`
	Email      *string `csv:"Email"`
	Phone      *string `csv:"Phone"`
}

// CrewCredential represents the credentials of a crew member.
type CrewCredential struct {
	ExternalID string  `csv:"External ID"`
	Title      string  `csv:"Title"`
	Type       *string `csv:"Type"`
	Details    *string `csv:"Details"`
	IssuedAt   *string `csv:"Issued At"`
	ExpiresAt  *string `csv:"Expires At"`
}

// Vessel represents the details of a vessel.
type Vessel struct {
	ExternalID           *string `csv:"External ID"`
	Name                 *string `csv:"Name"`
	MMSINumber           *string `csv:"MMSI Number"`
	IMONumber            *string `csv:"IMO Number"`
	AdditionalIdentifier *string `csv:"Additional Identifier"`
}

// VesselSchedule represents the schedule of a vessel.
type VesselSchedule struct {
	VesselExternalID *string `csv:"Vessel External ID"`
	VesselName       *string `csv:"Vessel Name"`
	VesselIMONumber  *string `csv:"Vessel IMO Number"`
	VesselMMSINumber *string `csv:"Vessel MMSI Number"`
	Client           *string `csv:"Client"`
	Description      *string `csv:"Description"`
	ServiceStartAt   *string `csv:"Service Start At"`
	ServiceEndAt     *string `csv:"Service End At"`
}

// VesselPosition represents the position of a vessel.
type VesselPosition struct {
	VesselExternalID string   `csv:"Vessel External ID"`
	Position         string   `csv:"Position"`
	CredentialTitle  string   `csv:"Credential Title"`
	Endorsements     []string `csv:"Endorsements"`
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

// CrewSchedulePosition represents the position details of a crew member in a schedule.
type CrewSchedulePosition struct {
	CrewExternalID   string   `csv:"Crew External ID"`
	VesselExternalID string   `csv:"Vessel External ID"`
	Position         string   `csv:"Position"`
	CredentialTitle  string   `csv:"Credential Title"`
	Endorsements     []string `csv:"Endorsements"`
}
