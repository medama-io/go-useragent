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
			assert.Equal(t, tc.ExpectedOSVersion, result.OSVersion(), "os version mismatch")
		})
	}
}

func TestParseNoAllocations(t *testing.T) {
	parser := ua.NewParser()

	// Cover every OS version capture policy.
	testCases := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4 Safari/605.1.15",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (Linux; Android 14; Pixel 8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (X11; CrOS x86_64 14541.0.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64; rv:128.0) Gecko/20100101 Firefox/128.0",
	}

	var sink ua.UserAgent
	for _, tc := range testCases {
		allocs := testing.AllocsPerRun(100, func() {
			sink = parser.Parse(tc)
		})
		assert.Equal(t, 0.0, allocs, "parser should not allocate for: %s", tc)
	}

	_ = sink
}
