package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"agnos/internal/model"
)

type HospitalAClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewHospitalAClient(baseURL string) *HospitalAClient {
	return &HospitalAClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *HospitalAClient) FetchPatientByID(id string) (*model.HospitalAPatientResponse, error) {
	url := fmt.Sprintf("%s/patient/search/%s", c.baseURL, id)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call Hospital A API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil // หาไม่เจอ
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Hospital A API returned status: %d", resp.StatusCode)
	}

	var patient model.HospitalAPatientResponse
	if err := json.NewDecoder(resp.Body).Decode(&patient); err != nil {
		return nil, fmt.Errorf("failed to decode Hospital A response: %w", err)
	}

	return &patient, nil
}
