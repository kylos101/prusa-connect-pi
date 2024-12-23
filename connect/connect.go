package connect

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// BaseURL is the default base URL. You should override this.
const BaseURL = "http://localhost:8000"

// Client represents a client for the Connect API.
type Client struct {
	httpClient  *http.Client
	baseURL     *url.URL
	Token       string
	Fingerprint string
	Cookie      string
}

// NewClient creates a new Client instance.
func NewClient(baseURL string) (*Client, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		httpClient: &http.Client{},
		baseURL:    parsedURL,
	}, nil
}

func (c *Client) SetAuth(token, fingerprint, cookie string) {
	c.Token = token
	c.Fingerprint = fingerprint
	c.Cookie = cookie
}

// doRequest sends an HTTP request and returns the response body.
func (c *Client) doRequest(method, path string, body interface{}) ([]byte, error) {
	u := *c.baseURL
	u.Path = path

	var reqBody []byte
	if body != nil {
		var err error
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(context.Background(), method, u.String(), bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	if c.Token != "" {
		req.Header.Set("Token", c.Token)
	}
	if c.Fingerprint != "" {
		req.Header.Set("Fingerprint", c.Fingerprint)
	}
	if c.Cookie != "" {
		req.AddCookie(&http.Cookie{Name: "SESSID", Value: c.Cookie})
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 && resp.StatusCode != 400 && resp.StatusCode != 401 && resp.StatusCode != 403 && resp.StatusCode != 404 && resp.StatusCode != 409 && resp.StatusCode != 503 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

// RegisterCamera registers a camera to Connect.
func (c *Client) RegisterCamera(printerUUID string, origin string) (*CameraResponse, error) {
	if origin != "WEB" && origin != "OTHER" {
		return nil, errors.New("invalid origin")
	}

	path := fmt.Sprintf("/app/printers/%s/camera", printerUUID)
	body := map[string]string{"origin": origin}
	responseBody, err := c.doRequest(http.MethodPost, path, body)
	if err != nil {
		return nil, err
	}

	var cameraResponse CameraResponse
	if err := json.Unmarshal(responseBody, &cameraResponse); err != nil {
		return nil, err
	}

	return &cameraResponse, nil
}

// UploadSnapshot uploads a snapshot to Connect.
func (c *Client) UploadSnapshot(data []byte) error {
	path := "/c/snapshot"
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPut, c.baseURL.String()+path, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "image/jpg")
	if c.Token != "" {
		req.Header.Set("Token", c.Token)
	}
	if c.Fingerprint != "" {
		req.Header.Set("Fingerprint", c.Fingerprint)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

// UpdateCamera updates camera attributes.
func (c *Client) UpdateCamera(cameraRequest CameraRequest) (*CameraResponse, error) {
	path := "/c/info"
	responseBody, err := c.doRequest(http.MethodPut, path, cameraRequest)
	if err != nil {
		return nil, err
	}

	var cameraResponse CameraResponse
	if err := json.Unmarshal(responseBody, &cameraResponse); err != nil {
		return nil, err
	}

	return &cameraResponse, nil
}

// Define the necessary structs based on the OpenAPI spec.
type CameraResponse struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Config       *Config  `json:"config"`
	Options      *Options `json:"options"`
	Capabilities []string `json:"capabilities"`
	TeamID       int      `json:"team_id"`
	PrinterUUID  string   `json:"printer_uuid"`
	Token        string   `json:"token"`
	Origin       string   `json:"origin"`
	Deleted      int      `json:"deleted"`
	Registered   bool     `json:"registered"`
	SortOrder    int      `json:"sort_order"`
}

type Config struct {
	CameraID      string       `json:"camera_id"`
	Path          string       `json:"path"`
	Name          string       `json:"name"`
	Driver        string       `json:"driver"`
	TriggerScheme string       `json:"trigger_scheme"`
	Resolution    *Resolution  `json:"resolution"`
	NetworkInfo   *NetworkInfo `json:"network_info"`
	Firmware      string       `json:"firmware"`
	Manufacturer  string       `json:"manufacturer"`
	Model         string       `json:"model"`
}

type Options struct {
	AvailableResolutions []Resolution `json:"available_resolutions"`
}

type Resolution struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type NetworkInfo struct {
	LanMAC   string `json:"lan_mac"`
	LanIPv4  string `json:"lan_ipv4"`
	LanIPv6  string `json:"lan_ipv6"`
	WifiMAC  string `json:"wifi_mac"`
	WifiIPv4 string `json:"wifi_ipv4"`
	WifiIPv6 string `json:"wifi_ipv6"`
	WifiSSID string `json:"wifi_ssid"`
}

type CameraRequest struct {
	Config       *Config  `json:"config"`
	Options      *Options `json:"options"`
	Capabilities []string `json:"capabilities"`
}

// ... other structs (Status, etc.) as needed based on your usage
