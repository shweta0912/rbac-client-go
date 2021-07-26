package rbac

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// TODO: support auto-paginating body
const paginationLimit = "100"

// Client is used for making requests to the RBAC service
type Client struct {
	HTTPClient  *http.Client
	BaseURL     string
	Application string
}

// NewClient returns a Client given an application
func NewClient(baseURL, application string) Client {
	return Client{
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		BaseURL:     baseURL,
		Application: application,
	}
}

// getParsed
func (c *Client) getParsed(r *http.Request, data interface{}) error {
	// Perform request and check status
	resp, err := c.do(r)
	if err != nil {
		return fmt.Errorf("request to RBAC service failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK status code: %d", resp.StatusCode)
	}
	// TODO: consider parsing

	// Unmarshal JSON from good response

	body := PaginatedBody{
		Data: &data,
	}
	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	err = json.Unmarshal(rawBody, &body)
	if err != nil {
		return fmt.Errorf("failed to parse response body: %w", err)
	}
	return nil
}

func (c *Client) do(r *http.Request) (*http.Response, error) {
	if c.HTTPClient == nil {
		return nil, errors.New("HTTPClient cannot be nil")
	}
	return c.HTTPClient.Do(r)
}
