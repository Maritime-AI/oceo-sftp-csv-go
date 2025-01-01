package sftpclient

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/Maritime-AI/oceo-sftp-csv-go/models"
	"github.com/gocarina/gocsv"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

const (
	remoteDir = "./data"
)

const (
	vesselScheduleFileTemplate          = "%s_vesselschedules_%d.csv"
	crewFileTemplate                    = "%s_crew_%d.csv"
	crewCredentialsFileTemplate         = "%s_credentials_%d.csv"
	vesselFileTemplate                  = "%s_vessels_%d.csv"
	vesselSchedulePositionsFileTemplate = "%s_vesselschedulepositions_%d.csv"
	crewScheduleFileTemplate            = "%s_crewschedules_%d.csv"
	crewSchedulePositionsFileTemplate   = "%s_crewschedulepositions_%d.csv"
)

// OCEOSFTPClient manages the connection to an SFTP server and provides methods to upload structured data in CSV format.
type OCEOSFTPClient struct {
	orgName string
	addr    string
	config  ssh.ClientConfig
}

// NewOCEOSFTPCLient initializes a new OCEO SFTPClient with the specified server details.
//
// Parameters:
// - orgName: The name of your organization.
// - host: The SFTP server address.
// - port: The port on which the SFTP server is running.
// - user: The username for authentication.
// - rsaPrivateKeyBytes: RSA private key in byte representation.
//
// Returns:
// - An instance of SFTPClient.
// - An error if there is an issue creating the client.
func NewOCEOSFTPCLient(
	orgName, host, port, user string,
	rsaPrivateKeyBytes []byte) (*OCEOSFTPClient, error) {
	authMethod, err := readPrivateKey(rsaPrivateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key: %w", err)
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	return &OCEOSFTPClient{
		orgName: orgName,
		addr:    addr,
		config: ssh.ClientConfig{
			User:            user,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Auth:            authMethod,
		},
	}, nil
}

// UploadCrew uploads a slice of Crew data to the SFTP server as a CSV file.
//
// Parameters:
// - crew: A slice of Crew structs.
//
// Returns:
// - An error if the upload fails.
func (s *OCEOSFTPClient) UploadCrew(ctx context.Context, crew ...models.Crew) error {
	if len(crew) == 0 {
		fmt.Println("No crew to upload")
		return nil
	}

	nowUnix := time.Now().Unix()
	fn := fmt.Sprintf(crewFileTemplate, s.orgName, nowUnix)
	return s.uploadData(ctx, fn, crew)
}

// UploadCrewCredentials uploads a slice of CrewCredential data to the SFTP server as a CSV file.
//
// Parameters:
// - crewCred: A slice of CrewCredential structs.
//
// Returns:
// - An error if the upload fails.
func (s *OCEOSFTPClient) UploadCrewCredentials(ctx context.Context, credentials ...models.CrewCredential) error {
	if len(credentials) == 0 {
		fmt.Println("No crew to upload")
		return nil
	}

	nowUnix := time.Now().Unix()
	fn := fmt.Sprintf(crewCredentialsFileTemplate, s.orgName, nowUnix)
	return s.uploadData(ctx, fn, credentials)
}

// UploadVessels uploads a slice of Vessel data to the SFTP server as a CSV file.
//
// Parameters:
// - vessels: A slice of Vessel structs.
//
// Returns:
// - An error if the upload fails.
func (s *OCEOSFTPClient) UploadVessels(ctx context.Context, vessels ...models.Vessel) error {
	if len(vessels) == 0 {
		fmt.Println("No vessels to upload")
		return nil
	}

	nowUnix := time.Now().Unix()
	fn := fmt.Sprintf(vesselFileTemplate, s.orgName, nowUnix)
	return s.uploadData(ctx, fn, vessels)
}

// UploadVesselSchedules uploads a slice of VesselSchedule data to the SFTP server as a CSV file.
//
// Parameters:
// - vesselSchedules: A slice of VesselSchedule structs.
//
// Returns:
// - An error if the upload fails.
func (s *OCEOSFTPClient) UploadVesselSchedules(ctx context.Context, vesselSchedules ...models.VesselSchedule) error {
	if len(vesselSchedules) == 0 {
		fmt.Println("No vessel schedules to upload")
		return nil
	}

	nowUnix := time.Now().Unix()
	fn := fmt.Sprintf(vesselScheduleFileTemplate, s.orgName, nowUnix)
	return s.uploadData(ctx, fn, vesselSchedules)
}

// UploadVesselSchedulePositions uploads a slice of VesselSchedulePosition data to the SFTP server as a CSV file.
//
// Parameters:
// - vesselPositions: A slice of VesselSchedulePosition structs.
//
// Returns:
// - An error if the upload fails.
func (s *OCEOSFTPClient) UploadVesselSchedulePositions(ctx context.Context, vesselPositions ...models.VesselSchedulePosition) error {
	if len(vesselPositions) == 0 {
		fmt.Println("No vessel positions to upload")
		return nil
	}

	nowUnix := time.Now().Unix()
	fn := fmt.Sprintf(vesselSchedulePositionsFileTemplate, s.orgName, nowUnix)
	return s.uploadData(ctx, fn, vesselPositions)
}

// UploadCrewSchedules uploads a slice of CrewSchedule data to the SFTP server as a CSV file.
//
// Parameters:
// - crewSchedules: A slice of CrewSchedule structs.
//
// Returns:
// - An error if the upload fails.
func (s *OCEOSFTPClient) UploadCrewSchedules(ctx context.Context, crewSchedules ...models.CrewSchedule) error {
	if len(crewSchedules) == 0 {
		fmt.Println("No crew schedules to upload")
		return nil
	}

	nowUnix := time.Now().Unix()
	fn := fmt.Sprintf(crewScheduleFileTemplate, s.orgName, nowUnix)
	return s.uploadData(ctx, fn, crewSchedules)
}

// UploadCrewSchedulePositions uploads a slice of CrewSchedulePosition data to the SFTP server as a CSV file.
//
// Parameters:
// - crewSchedulePositions: A slice of CrewSchedulePosition structs.
//
// Returns:
// - An error if the upload fails.
func (s *OCEOSFTPClient) UploadCrewSchedulePositions(ctx context.Context, crewSchedulePositions ...models.CrewSchedulePosition) error {
	if len(crewSchedulePositions) == 0 {
		fmt.Println("No crew schedule positions to upload")
		return nil
	}

	nowUnix := time.Now().Unix()
	fn := fmt.Sprintf(crewSchedulePositionsFileTemplate, s.orgName, nowUnix)
	return s.uploadData(ctx, fn, crewSchedulePositions)
}

// uploadData is a helper function to upload data of any type to the SFTP server as a CSV file.
//
// Parameters:
// - data: The data to be uploaded, which must be a slice of structs.
//
// Returns:
// - An error if the upload fails.
func (s *OCEOSFTPClient) uploadData(ctx context.Context, fileName string, data any) error {

	conn, err := ssh.Dial("tcp", s.addr, &s.config)
	if err != nil {
		return fmt.Errorf("failed to dial SFTP server: %w", err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("failed to close SFTP connection: %v", err)
		}
	}()

	sc, err := sftp.NewClient(conn)
	if err != nil {
		return fmt.Errorf("failed to create SFTP client: %w", err)
	}

	defer func() {
		if err := sc.Close(); err != nil {
			log.Printf("failed to close SFTP client: %v", err)
		}
	}()

	// Open the destination file on the remote server
	destFile, err := sc.Create(fmt.Sprintf("./%s/%s", remoteDir, fileName))
	if err != nil {
		return fmt.Errorf("failed to create remote file: %w", err)
	}
	defer func() {
		if err := destFile.Close(); err != nil {
			log.Printf("failed to close remote file: %v", err)
		}
	}()

	// Copy the content to the remote file
	bs, err := gocsv.MarshalBytes(&data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}
	if _, err := io.Copy(destFile, bytes.NewReader(bs)); err != nil {
		return fmt.Errorf("failed to copy data to remote file: %w", err)
	}

	return nil
}

func readPrivateKey(keyBytes []byte) ([]ssh.AuthMethod, error) {
	signer, err := ssh.ParsePrivateKey(keyBytes)
	if err != nil {
		return nil, err
	}
	return []ssh.AuthMethod{
		ssh.PublicKeys(signer),
	}, nil
}
