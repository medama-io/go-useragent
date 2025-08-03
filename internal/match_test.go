package internal_test

import (
	"testing"

	"github.com/medama-io/go-useragent/data"
	"github.com/medama-io/go-useragent/internal"
	"github.com/stretchr/testify/assert"
)

func TestMatchTokenIndexes(t *testing.T) {
	for _, tc := range data.AllTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			matchResults := internal.MatchTokenIndexes(tc.UserAgent)

			// Convert expected and actual matches to their string names for better debugging.
			expectedNames := make([]string, len(tc.ExpectedMatches))
			for i, m := range tc.ExpectedMatches {
				expectedNames[i] = m.GetMatchName()
			}

			actualNames := make([]string, len(matchResults))
			for i, m := range matchResults {
				actualNames[i] = m.Match.GetMatchName()
			}

			// Use assert.Equal because the order of matches is important.
			assert.Equal(t, expectedNames, actualNames, "mismatch in token matches for: %s", tc.UserAgent)
		})
	}
}
