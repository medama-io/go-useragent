package useragent_test

import (
	"testing"

	ua "github.com/medama-io/go-useragent"
	"github.com/stretchr/testify/assert"
)

var matchResults = [][]string{
	// Windows
	{ua.Safari, ua.Chrome, ua.Windows},
	{ua.Safari, ua.Chrome, ua.Windows},
	{ua.Windows, ua.IE},
	{ua.Windows, ua.IE},
	{ua.Edge, ua.Safari, ua.Chrome, ua.Windows},
	// Mac
	{ua.Safari, ua.MacOS},
	{ua.Safari, ua.Chrome, ua.MacOS},
	{ua.Firefox, ua.MacOS},
	{ua.Vivaldi, ua.Safari, ua.Chrome, ua.MacOS},
	{ua.Edge, ua.Safari, ua.Chrome, ua.MacOS},
}

func TestMatchTokenIndexes(t *testing.T) {
	assert := assert.New(t)

	// Refer to version_test.go for versionResults test cases
	for i, v := range versionResults {
		match := ua.MatchTokenIndexes(v)

		if len(match) != len(matchResults[i]) {
			t.Errorf("Test Case: %s, expected %d matches, got %d\nMatch Index: %d", v, len(match), len(matchResults[i]), i)
			t.FailNow()
		}

		for j, m := range match {
			assert.Equal(matchResults[i][j], m.Match, "Test Case: %s\nMatch Number: %d\nExpected: %v", v, i, match)
		}

	}
}
