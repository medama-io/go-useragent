package useragent_test

import (
	"fmt"
	"testing"

	ua "github.com/medama-io/go-useragent"
	"github.com/stretchr/testify/assert"
)

var matchResults = [][]string{
	// Windows (7)
	{ua.Safari, ua.Chrome, ua.Windows},
	{ua.Safari, ua.Chrome, ua.Windows},
	{ua.Windows, ua.IE},
	{ua.Windows, ua.IE},
	{ua.IE, ua.Windows},
	{ua.Windows, ua.IE},
	{ua.Edge, ua.Safari, ua.Chrome, ua.Windows},

	// Mac (5)
	{ua.Safari, ua.Version, ua.MacOS},
	{ua.Safari, ua.Chrome, ua.MacOS},
	{ua.Firefox, ua.MacOS},
	{ua.Vivaldi, ua.Safari, ua.Chrome, ua.MacOS},
	{ua.Edge, ua.Safari, ua.Chrome, ua.MacOS},

	// Linux (5)
	{ua.Firefox, ua.Linux},
	{ua.Firefox, ua.Linux},
	{ua.Firefox, ua.Linux, ua.Desktop},
	{ua.Firefox, ua.Linux, ua.Desktop},
	{ua.Safari, ua.Chrome, ua.Linux},

	// iPhone (5)
	{ua.Safari, ua.Mobile, ua.Version, ua.IOS, ua.MobileDevice},
	{ua.Safari, ua.Mobile, ua.Chrome, ua.IOS, ua.MobileDevice},
	{ua.Safari, ua.Mobile, ua.Opera, ua.IOS, ua.MobileDevice},
	{ua.Safari, ua.Mobile, ua.Firefox, ua.IOS, ua.MobileDevice},
	{ua.Safari, ua.Mobile, ua.Edge, ua.Version, ua.IOS, ua.MobileDevice},
}

func TestMatchTokenIndexes(t *testing.T) {
	for i, v := range testCases {
		t.Run(fmt.Sprintf("Case:%d", i), func(t *testing.T) {
			match := ua.MatchTokenIndexes(v)

			if len(match) != len(matchResults[i]) {
				t.Errorf("Test Case: %s, expected %d matches, got %d\nMatch Index: %d", v, len(match), len(matchResults[i]), i)
				t.FailNow()
			}

			for j, m := range match {
				assert.Equal(t, matchResults[i][j], m.Match, "Test Case: %s\nMatch Number: %d\nExpected: %v", v, i, match)
			}
		})
	}
}
