package rbac

import (
	"context"
	"fmt"
	"net/http"
)

// GetAccess returns an AccessList for a principal.
// When username is empty, the authenticated principal is used.
func (c *Client) GetAccess(ctx context.Context, identity string, username string) (AccessList, error) {
	// Build request to RBAC service
	url := c.BaseURL + "/access"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}

	q := req.URL.Query()
	q.Add("application", c.Application)
	if len(username) > 0 {
		q.Add("username", username)
	}
	q.Add("limit", paginationLimit)
	req.URL.RawQuery = q.Encode()

	// Add X-RH-Identity header for authenticating the current principal
	req.Header.Set(identityHeader, identity)

	var access AccessList
	err = c.getParsed(req, &access)
	if err != nil {
		return nil, fmt.Errorf("failed to get permitted access: %w", err)
	}

	return access, nil
}
