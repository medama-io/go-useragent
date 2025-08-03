package useragent_test

import (
	"testing"

	ua "github.com/medama-io/go-useragent"
	"github.com/medama-io/go-useragent/data"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	parser := ua.NewParser()

	for _, tc := range data.AllTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := parser.Parse(tc.UserAgent)

			assert.Equal(t, tc.ExpectedBrowser, result.Browser(), "browser mismatch")
			assert.Equal(t, tc.ExpectedOS, result.OS(), "os mismatch")
			assert.Equal(t, tc.ExpectedDevice, result.Device(), "device mismatch")
			assert.Equal(t, tc.ExpectedVersion, result.BrowserVersion(), "browser version mismatch")
		})
	}
}
