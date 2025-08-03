package internal_test

import (
	"testing"

	"github.com/medama-io/go-useragent/data"
	"github.com/medama-io/go-useragent/internal"
	"github.com/stretchr/testify/assert"
)

func TestCleanVersions(t *testing.T) {
	for _, tc := range data.AllTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			line := internal.RemoveMobileIdentifiers(tc.UserAgent)
			line = internal.RemoveAndroidIdentifiers(line)
			line = internal.RemoveVersions(line)

			assert.Equal(t, tc.ExpectedCleanVersion, line, "clean version mismatch for: %s", tc.UserAgent)
		})
	}
}
