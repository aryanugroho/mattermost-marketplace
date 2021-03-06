package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildURL(t *testing.T) {
	config := &Client{}

	testCases := map[string]struct {
		base     string
		path     string
		expected string
	}{
		"Base url with trailing slash and path with leading slash": {
			base:     "https://api.integrations.mattermost.com/",
			path:     "/api/v1/plugins",
			expected: "https://api.integrations.mattermost.com/api/v1/plugins",
		},
		"Base url without trailing slash and path with leading slash": {
			base:     "https://api.integrations.mattermost.com",
			path:     "/api/v1/plugins",
			expected: "https://api.integrations.mattermost.com/api/v1/plugins",
		},
		"Base url without trailing slash and path without leading slash": {
			base:     "https://api.integrations.mattermost.com",
			path:     "api/v1/plugins",
			expected: "https://api.integrations.mattermost.com/api/v1/plugins",
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			config.Address = testCase.base
			actual := config.buildURL(testCase.path)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
