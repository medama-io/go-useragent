package useragent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionPart(t *testing.T) {
	tests := []struct {
		version  string
		index    int
		expected string
	}{
		{version: "118.0.0.0", index: 0, expected: "118"},
		{version: "118.0.0.0", index: 1, expected: "0"},
		{version: "118.0.0.0", index: 2, expected: "0.0"},
		{version: "10.12.6", index: 2, expected: "6"},
		{version: "1.2.3b", index: 2, expected: "3b"},
		{version: "17", index: 0, expected: "17"},
		{version: "17", index: 1, expected: ""},
		{version: "17", index: 2, expected: ""},
		{version: "11.2", index: 2, expected: ""},
		{version: "", index: 0, expected: ""},
		{version: "", index: 2, expected: ""},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expected, versionPart(tc.version, tc.index), "version part mismatch for %q", tc.version)
	}
}

func TestOSVersionParts(t *testing.T) {
	ua := UserAgent{}
	copy(ua.osVersion[:], "10.15.7")
	ua.osVersionIndex = len("10.15.7")

	assert.Equal(t, "10.15.7", ua.OSVersion())
	assert.Equal(t, "10", ua.OSVersionMajor())
	assert.Equal(t, "15", ua.OSVersionMinor())
	assert.Equal(t, "7", ua.OSVersionPatch())

	ua = UserAgent{}
	assert.Equal(t, "", ua.OSVersion())
	assert.Equal(t, "", ua.OSVersionMajor())
	assert.Equal(t, "", ua.OSVersionMinor())
	assert.Equal(t, "", ua.OSVersionPatch())
}
