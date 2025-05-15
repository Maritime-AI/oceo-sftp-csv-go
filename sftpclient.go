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

type FileType string

const (
	FileTypeCrew                    FileType = "crew"
	FileTypeCrewCredentials         FileType = "credentials"
	FileTypeVessels                 FileType = "vessels"
	FileTypeVesselSchedules         FileType = "vesselschedules"
	FileTypeVesselSchedulePositions FileType = "vesselschedulepositions"
	FileTypeCrewSchedules           FileType = "crewschedules"
	FileTypeCrewSchedulePositions   FileType = "crewschedulepositions"
)

const (
	fileTemplate = "%s_%s_%d.csv"
)

// OCEOSFTPClient manages the connection to an SFTP server and provides methods to upload structured data in CSV format.
type OCEOSFTPClient struct {
	addr   string
	config ssh.ClientConfig
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
	host, port, user string,
	rsaPrivateKeyBytes []byte) (*OCEOSFTPClient, error) {
	authMethod, err := readPrivateKey(rsaPrivateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key: %w", err)
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	return &OCEOSFTPClient{
		addr: addr,
		config: ssh.ClientConfig{
			User:            user,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Auth:            authMethod,
		},
	}, nil
}

// UploadCrewFile uploads a slice of Crew data to the SFTP server as a CSV file.
//
// Parameters:
// - crew: A slice of Crew structs.
//
// Returns:
// - An error if the upload fails.
func (s *OCEOSFTPClient) UploadCrewFile(ctx context.Context,
	orgName string, crew ...models.Crew) error {
	if len(crew) == 0 {
		fmt.Println("No crew to upload")
		return nil
	}

	for _, c := range crew {
		if err := c.Validate(); err != nil {
			return fmt.Errorf("invalid crew data: %w", err)
		}
	}

	nowUnix := time.Now().Unix()
	fn := fmt.Sprintf(fileTemplate, orgName, FileTypeCrew, nowUnix)

	bs, err := gocsv.MarshalBytes(&crew)
	if err != nil {
		return fmt.Errorf("failed to marshal crew: %w", err)
	}

	return s.uploadData(fn, bs)
}

// UploadCrewCredentialFile uploads a slice of CrewCredential data to the SFTP server as a CSV file.
//
// Parameters:
// - crewCred: A slice of CrewCredential structs.
//
// Returns:
// - An error if the upload fails.
func (s *OCEOSFTPClient) UploadCrewCredentialFile(ctx context.Context,
	orgName string, credentials ...models.CrewCredential) error {
	if len(credentials) == 0 {
		fmt.Println("No crew to upload")
		return nil
	}

	for _, cc := range credentials {
		if err := cc.Validate(); err != nil {
			return fmt.Errorf("invalid crew credential data: %w", err)
		}
	}

	nowUnix := time.Now().Unix()
	fn := fmt.Sprintf(fileTemplate, orgName, FileTypeCrewCredentials, nowUnix)

	bs, err := gocsv.MarshalBytes(&credentials)
	if err != nil {
		return fmt.Errorf("failed to marshal crew credentials: %w", err)
	}

	return s.uploadData(fn, bs)
}

// UploadVesselFile uploads a slice of Vessel data to the SFTP server as a CSV file.
//
// Parameters:
// - vessels: A slice of Vessel structs.
//
// Returns:
// - An error if the upload fails.
func (s *OCEOSFTPClient) UploadVesselFile(ctx context.Context,
	orgName string, vessels ...models.Vessel) error {
	if len(vessels) == 0 {
		fmt.Println("No vessels to upload")
		return nil
	}

	for _, v := range vessels {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("invalid vessel data: %w", err)
		}
	}

	nowUnix := time.Now().Unix()
	fn := fmt.Sprintf(fileTemplate, orgName, FileTypeVessels, nowUnix)

	bs, err := gocsv.MarshalBytes(&vessels)
	if err != nil {
		return fmt.Errorf("failed to marshal vessels: %w", err)
	}

	return s.uploadData(fn, bs)
}

// UploadVesselScheduleFile uploads a slice of VesselSchedule data to the SFTP server as a CSV file.
//
// Parameters:
// - vesselSchedules: A slice of VesselSchedule structs.
//
// Returns:
// - An error if the upload fails.
func (s *OCEOSFTPClient) UploadVesselScheduleFile(ctx context.Context,
	orgName string, vesselSchedules ...models.VesselSchedule) error {
	if len(vesselSchedules) == 0 {
		fmt.Println("No vessel schedules to upload")
		return nil
	}

	for _, vs := range vesselSchedules {
		if err := vs.Validate(); err != nil {
			return fmt.Errorf("invalid vessel schedule data: %w", err)
		}
	}

	nowUnix := time.Now().Unix()
	fn := fmt.Sprintf(fileTemplate, orgName, FileTypeVesselSchedules, nowUnix)

	bs, err := gocsv.MarshalBytes(&vesselSchedules)
	if err != nil {
		return fmt.Errorf("failed to marshal vessel schedules: %w", err)
	}

	return s.uploadData(fn, bs)
}

// UploadVesselSchedulePositionFile uploads a slice of VesselSchedulePosition data to the SFTP server as a CSV file.
//
// Parameters:
// - vesselPositions: A slice of VesselSchedulePosition structs.
//
// Returns:
// - An error if the upload fails.
func (s *OCEOSFTPClient) UploadVesselSchedulePositionFile(ctx context.Context,
	orgName string, vesselPositions ...models.VesselSchedulePosition) error {
	if len(vesselPositions) == 0 {
		fmt.Println("No vessel positions to upload")
		return nil
	}

	for _, vp := range vesselPositions {
		if err := vp.Validate(); err != nil {
			return fmt.Errorf("invalid vessel position data: %w", err)
		}
	}

	nowUnix := time.Now().Unix()
	fn := fmt.Sprintf(fileTemplate, orgName, FileTypeVesselSchedulePositions, nowUnix)

	bs, err := gocsv.MarshalBytes(&vesselPositions)
	if err != nil {
		return fmt.Errorf("failed to marshal vessel positions: %w", err)
	}

	return s.uploadData(fn, bs)
}

// UploadCrewScheduleFile uploads a slice of CrewSchedule data to the SFTP server as a CSV file.
//
// Parameters:
// - crewSchedules: A slice of CrewSchedule structs.
//
// Returns:
// - An error if the upload fails.
func (s *OCEOSFTPClient) UploadCrewScheduleFile(ctx context.Context, orgName string,
	crewSchedules ...models.CrewSchedule) error {
	if len(crewSchedules) == 0 {
		fmt.Println("No crew schedules to upload")
		return nil
	}

	for _, cs := range crewSchedules {
		if err := cs.Validate(); err != nil {
			return fmt.Errorf("invalid crew schedule data: %w", err)
		}
	}

	nowUnix := time.Now().Unix()
	fn := fmt.Sprintf(fileTemplate, orgName, FileTypeCrewSchedules, nowUnix)

	bs, err := gocsv.MarshalBytes(&crewSchedules)
	if err != nil {
		return fmt.Errorf("failed to marshal crew schedules: %w", err)
	}

	return s.uploadData(fn, bs)
}

// UploadCrewSchedulePositionFile uploads a slice of CrewSchedulePosition data to the SFTP server as a CSV file.
//
// Parameters:
// - crewSchedulePositions: A slice of CrewSchedulePosition structs.
//
// Returns:
// - An error if the upload fails.
func (s *OCEOSFTPClient) UploadCrewSchedulePositionFile(ctx context.Context, orgName string,
	crewSchedulePositions ...models.CrewSchedulePosition) error {
	if len(crewSchedulePositions) == 0 {
		fmt.Println("No crew schedule positions to upload")
		return nil
	}

	for _, csp := range crewSchedulePositions {
		if err := csp.Validate(); err != nil {
			return fmt.Errorf("invalid crew schedule position data: %w", err)
		}
	}

	nowUnix := time.Now().Unix()
	fn := fmt.Sprintf(fileTemplate, orgName, FileTypeCrewSchedulePositions, nowUnix)

	bs, err := gocsv.MarshalBytes(&crewSchedulePositions)
	if err != nil {
		return fmt.Errorf("failed to marshal crew schedule positions: %w", err)
	}

	return s.uploadData(fn, bs)
}

// uploadData is a helper function to upload data of any type to the SFTP server as a CSV file.
//
// Parameters:
// - data: The data to be uploaded, which must be a slice of structs.
//
// Returns:
// - An error if the upload fails.
func (s *OCEOSFTPClient) uploadData(fileName string, data []byte) error {

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
	dest := fmt.Sprintf("./%s/%s", remoteDir, fileName)
	log.Printf("uploading data to %s", dest)
	destFile, err := sc.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create remote file: %w", err)
	}
	defer func() {
		if err := destFile.Close(); err != nil {
			log.Printf("failed to close remote file: %v", err)
		}
	}()

	// Copy the content to the remote file
	if _, err := io.Copy(destFile, bytes.NewReader(data)); err != nil {
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
