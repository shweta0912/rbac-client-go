package rbac

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const mockApp = "chipotle"

const mockSimpleAccess = `{
	"data": [
	  {
		"resourceDefinitions": [],
		"permission": "chipotle:burrito:order"
	  },
	  {
		"resourceDefinitions": [
		  {
			"attributeFilter": {
			  "key": "beanType",
			  "value": "black",
			  "operation": "equal"
			}
		  }
		],
		"permission": "chipotle:burrito:eat"
	  }
	]
  }`

const mockEmptyAccess = `{
	"data": []
  }`

func TestGetAccess(t *testing.T) {
	// Set up mock service
	var mockBody *[]byte
	var mockStatus *int
	mockService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(*mockStatus)
		w.Write([]byte(*mockBody))
	}))
	defer mockService.Close()

	// Build a client using the mock service
	c := NewClient(mockService.URL, mockApp)

	tests := map[string]struct {
		respBody   []byte
		respStatus int
		expected   AccessList
		ok         bool
	}{
		"simple": {respBody: []byte(mockSimpleAccess), respStatus: 200, ok: true, expected: AccessList{
			Access{Permission: "chipotle:burrito:order", ResourceDefinitions: []ResourceDefinition{}},
			Access{Permission: "chipotle:burrito:eat", ResourceDefinitions: []ResourceDefinition{
				{Filter: ResourceDefinitionFilter{
					Key:       "beanType",
					Value:     "black",
					Operation: "equal",
				}},
			}},
		}},
		"empty": {respBody: []byte(mockEmptyAccess), respStatus: 200, ok: true, expected: AccessList{}},
		"error": {respBody: []byte{}, respStatus: 500, ok: false, expected: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			mockBody = &tc.respBody
			mockStatus = &tc.respStatus
			got, err := c.GetAccess(context.Background(), "", "")

			if tc.ok {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestGetAccess_RequestParams(t *testing.T) {
	identity := "aHR0cHM6Ly93d3cueW91dHViZS5jb20vd2F0Y2g/dj1kUXc0dzlXZ1hjUQo="
	username := "rick"
	handlerFired := false

	// Set up mock service
	mockService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, identity, r.Header.Get(identityHeader))
		assert.Equal(t, username, r.URL.Query().Get("username"))
		handlerFired = true
	}))
	defer mockService.Close()

	// Build a client and get access using mock service
	c := NewClient(mockService.URL, mockApp)
	c.GetAccess(context.Background(), identity, username)

	// Safety check to ensure the mock service handler was executed
	assert.True(t, handlerFired)
}
