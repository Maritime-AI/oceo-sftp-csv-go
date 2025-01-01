package sftpclient

import (
	"fmt"

	"github.com/Maritime-AI/oceo-sftp-csv-go/models"
	"github.com/gocarina/gocsv"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

const (
	remotePath = "./data"
)

// SFTPClient manages the connection to an SFTP server and provides methods to upload structured data in CSV format.
type SFTPClient struct {
	host     string
	port     int
	username string
	password string
	client   *sftp.Client
}

// NewSFTPClient initializes a new SFTPClient with the specified server details.
//
// Parameters:
// - host: The SFTP server address.
// - port: The port on which the SFTP server is running.
// - username: The username for authentication.
// - password: The password for authentication.
//
// Returns:
// - An instance of SFTPClient.
// - An error if there is an issue creating the client.
func NewSFTPClient(host string, port int, username string, password string) (*SFTPClient, error) {
	config := &ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	connection, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		return nil, fmt.Errorf("failed to dial SSH: %w", err)
	}

	sftpClient, err := sftp.NewClient(connection)
	if err != nil {
		return nil, fmt.Errorf("failed to create SFTP client: %w", err)
	}

	return &SFTPClient{
		host:     host,
		port:     port,
		username: username,
		password: password,
		client:   sftpClient,
	}, nil
}

// Close terminates the SFTP connection.
//
// Returns:
// - An error if closing the connection fails.
func (s *SFTPClient) Close() error {
	if s.client != nil {
		if err := s.client.Close(); err != nil {
			return err
		}
	}

	return nil
}

// UploadCrew uploads a slice of Crew data to the SFTP server as a CSV file.
//
// Parameters:
// - crew: A slice of Crew structs.
//
// Returns:
// - An error if the upload fails.
func (s *SFTPClient) UploadCrew(crew []models.Crew) error {
	return s.uploadData(crew)
}

// UploadCrewCredentials uploads a slice of CrewCredential data to the SFTP server as a CSV file.
//
// Parameters:
// - crewCred: A slice of CrewCredential structs.
//
// Returns:
// - An error if the upload fails.
func (s *SFTPClient) UploadCrewCredentials(crewCred []models.CrewCredential) error {
	return s.uploadData(crewCred)
}

// UploadVessels uploads a slice of Vessel data to the SFTP server as a CSV file.
//
// Parameters:
// - vessels: A slice of Vessel structs.
//
// Returns:
// - An error if the upload fails.
func (s *SFTPClient) UploadVessels(vessels []models.Vessel) error {
	return s.uploadData(vessels)
}

// UploadVesselSchedules uploads a slice of VesselSchedule data to the SFTP server as a CSV file.
//
// Parameters:
// - vesselSchedules: A slice of VesselSchedule structs.
//
// Returns:
// - An error if the upload fails.
func (s *SFTPClient) UploadVesselSchedules(vesselSchedules []models.VesselSchedule) error {
	return s.uploadData(vesselSchedules)
}

// UploadVesselPositions uploads a slice of VesselPosition data to the SFTP server as a CSV file.
//
// Parameters:
// - vesselPositions: A slice of VesselPosition structs.
//
// Returns:
// - An error if the upload fails.
func (s *SFTPClient) UploadVesselPositions(vesselPositions []models.VesselPosition) error {
	return s.uploadData(vesselPositions)
}

// UploadCrewSchedules uploads a slice of CrewSchedule data to the SFTP server as a CSV file.
//
// Parameters:
// - crewSchedules: A slice of CrewSchedule structs.
//
// Returns:
// - An error if the upload fails.
func (s *SFTPClient) UploadCrewSchedules(crewSchedules []models.CrewSchedule) error {
	return s.uploadData(crewSchedules)
}

// UploadCrewSchedulePositions uploads a slice of CrewSchedulePosition data to the SFTP server as a CSV file.
//
// Parameters:
// - crewSchedulePositions: A slice of CrewSchedulePosition structs.
//
// Returns:
// - An error if the upload fails.
func (s *SFTPClient) UploadCrewSchedulePositions(crewSchedulePositions []models.CrewSchedulePosition) error {
	return s.uploadData(crewSchedulePositions)
}

// uploadData is a helper function to upload data of any type to the SFTP server as a CSV file.
//
// Parameters:
// - data: The data to be uploaded, which must be a slice of structs.
//
// Returns:
// - An error if the upload fails.
func (s *SFTPClient) uploadData(data any) error {
	bs, err := gocsv.MarshalBytes(&data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}
	file, err := s.client.Create(remotePath)
	if err != nil {
		return fmt.Errorf("failed to create remote file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(bs)
	if err != nil {
		return fmt.Errorf("failed to write to remote file: %w", err)
	}

	return nil
}
