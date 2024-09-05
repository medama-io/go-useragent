package internal_test

import (
	"fmt"
	"testing"

	"github.com/medama-io/go-useragent/internal"
	"github.com/medama-io/go-useragent/testdata"
	"github.com/stretchr/testify/assert"
)

var matchResults = [][]string{
	// Windows (7)
	{internal.Safari, internal.Chrome, internal.Windows},
	{internal.Safari, internal.Chrome, internal.Windows},
	{internal.Windows, internal.IE},
	{internal.Windows, internal.IE},
	{internal.IE, internal.Windows},
	{internal.Windows, internal.IE},
	{internal.Edge, internal.Safari, internal.Chrome, internal.Windows},

	// Mac (5)
	{internal.Safari, internal.Version, internal.MacOS},
	{internal.Safari, internal.Chrome, internal.MacOS},
	{internal.Firefox, internal.MacOS},
	{internal.Vivaldi, internal.Safari, internal.Chrome, internal.MacOS},
	{internal.Edge, internal.Safari, internal.Chrome, internal.MacOS},

	// Linux (5)
	{internal.Firefox, internal.Linux},
	{internal.Firefox, internal.Linux},
	{internal.Firefox, internal.Linux, internal.Desktop},
	{internal.Firefox, internal.Linux, internal.Desktop},
	{internal.Safari, internal.Chrome, internal.Linux},

	// iPhone (5)
	{internal.Safari, internal.Mobile, internal.Version, internal.MobileDevice, internal.IOS},
	{internal.Safari, internal.Mobile, internal.Chrome, internal.MobileDevice, internal.IOS},
	{internal.Safari, internal.Mobile, internal.Opera, internal.MobileDevice, internal.IOS},
	{internal.Safari, internal.Mobile, internal.Firefox, internal.MobileDevice, internal.IOS},
	{internal.Safari, internal.Mobile, internal.Edge, internal.Version, internal.MobileDevice, internal.IOS},

	// iPad (3)
	{internal.Safari, internal.Mobile, internal.Version, internal.Tablet, internal.IOS},
	{internal.Safari, internal.Mobile, internal.Chrome, internal.Tablet, internal.IOS},
	{internal.Safari, internal.Mobile, internal.Firefox, internal.Tablet, internal.IOS},

	// Android (4)
	{internal.Safari, internal.Mobile, internal.Chrome, internal.Samsung, internal.Android, internal.Linux},
	{internal.Safari, internal.Mobile, internal.Version, internal.Android, internal.Linux},
	{internal.Safari, internal.Mobile, internal.Version, internal.Android, internal.Linux},
	{internal.Safari, internal.Mobile, internal.Chrome, internal.Version, internal.MobileDevice, internal.Android, internal.Linux},

	// Bots (4)
	{internal.Bot},
	{internal.Bot},
	{internal.Bot},
	{internal.Bot},
	{internal.Safari, internal.Chrome, internal.Bot},
	{internal.Safari, internal.Bot, internal.Chrome, internal.Linux},

	// Yandex Browser (1)
	{internal.Safari, internal.Mobile, internal.YandexBrowser, internal.Chrome, internal.Android, internal.Linux},
}

func TestMatchTokenIndexes(t *testing.T) {
	for i, v := range testdata.TestCases {
		t.Run(fmt.Sprintf("Case:%d", i), func(t *testing.T) {
			match := internal.MatchTokenIndexes(v)

			if len(match) != len(matchResults[i]) {
				t.Errorf("Test Case: %s, expected %d matches, got %d\nMatch Index: %d", v, len(match), len(matchResults[i]), i)
				t.FailNow()
			}

			for j, m := range match {
				assert.Equal(t, matchResults[i][j], m.Match, "Test Case: %s\nMatch Number: %d\nExpected: %v\nGot: %v", v, i, matchResults[i], match)
			}
		})
	}
}
