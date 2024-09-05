package useragent_test

import (
	"fmt"
	"testing"

	ua "github.com/medama-io/go-useragent"
	"github.com/medama-io/go-useragent/testdata"
	"github.com/stretchr/testify/assert"
)

var resultCases = []ua.UserAgent{
	// Windows (7)
	{Browser: ua.Chrome, OS: ua.Windows, Desktop: true, Version: "118.0.0.0"},
	{Browser: ua.Chrome, OS: ua.Windows, Desktop: true, Version: "59.0.3071.115"},
	{Browser: ua.IE, OS: ua.Windows, Desktop: true, Version: "8.0"},
	{Browser: ua.IE, OS: ua.Windows, Desktop: true, Version: "10.0"},
	// Technically should be 11.0, but we need to manually map this after getting it.
	{Browser: ua.IE, OS: ua.Windows, Desktop: true, Version: "7.0"},
	{Browser: ua.IE, OS: ua.Windows, Desktop: true, Version: "6.0"},
	{Browser: ua.Edge, OS: ua.Windows, Desktop: true, Version: "15.15063"},
	// Mac (5) 7
	{Browser: ua.Safari, OS: ua.MacOS, Desktop: true, Version: "10.1.2"},
	{Browser: ua.Chrome, OS: ua.MacOS, Desktop: true, Version: "60.0.3112.90"},
	{Browser: ua.Firefox, OS: ua.MacOS, Desktop: true, Version: "54.0"},
	{Browser: ua.Vivaldi, OS: ua.MacOS, Desktop: true, Version: "1.92.917.39"},
	{Browser: ua.Edge, OS: ua.MacOS, Desktop: true, Version: "79.0.309.71"},
	// Linux (5) 12
	{Browser: ua.Firefox, OS: ua.Linux, Desktop: true, Version: "52.0"},
	{Browser: ua.Firefox, OS: ua.Linux, Desktop: true, Version: "119.0"},
	{Browser: ua.Firefox, OS: ua.Linux, Desktop: true, Version: "119.0"},
	{Browser: ua.Firefox, OS: ua.Linux, Desktop: true, Version: "119.0"},
	{Browser: ua.Chrome, OS: ua.Linux, Desktop: true, Version: "119.0.0.0"},
	// iPhone (5) 17
	{Browser: ua.Safari, OS: ua.IOS, Mobile: true, Version: "10.0"},
	{Browser: ua.Chrome, OS: ua.IOS, Mobile: true, Version: "60.0.3112.89"},
	{Browser: ua.Opera, OS: ua.IOS, Mobile: true, Version: "14.0.0.104835"},
	{Browser: ua.Firefox, OS: ua.IOS, Mobile: true, Version: "8.1.1"},
	{Browser: ua.Edge, OS: ua.IOS, Mobile: true, Version: "44.11.15"},
	// iPad (3) 22
	{Browser: ua.Safari, OS: ua.IOS, Tablet: true, Version: "10.0"},
	{Browser: ua.Chrome, OS: ua.IOS, Tablet: true, Version: "119.0.6045.169"},
	{Browser: ua.Firefox, OS: ua.IOS, Tablet: true, Version: "119.0"},
	// Android (4) 25
	{Browser: ua.Samsung, OS: ua.Android, Mobile: true, Version: "4.0"},
	{Browser: ua.AndroidBrowser, OS: ua.Android, Mobile: true, Version: "4.0"},
	{Browser: ua.AndroidBrowser, OS: ua.Android, Mobile: true, Version: "4.0"},
	{Browser: ua.Chrome, OS: ua.Android, Mobile: true, Version: "38.0.2125.102"},
	// Bots (6) 29
	{Bot: true},
	{Bot: true},
	{Bot: true},
	{Bot: true},
	{Bot: true, Browser: ua.Chrome, Version: "112.0.0.0"},
	{Bot: true, Browser: ua.Chrome, OS: ua.Linux, Version: "125.0.6422.76", Desktop: true},
	// Yandex Browser (1) 35
	{Browser: ua.YandexBrowser, OS: ua.Android, Mobile: true, Version: "24.1.7.27.00"},
}

func TestParse(t *testing.T) {
	parser := ua.NewParser()

	for i, v := range testdata.TestCases {
		t.Run(fmt.Sprintf("Case:%d", i), func(t *testing.T) {
			result := parser.Parse(v)
			assert.Equal(t, resultCases[i].Browser, result.Browser, "Browser\nTest Case: %s\nExpected: %s", v, resultCases[i].Browser)
			assert.Equal(t, resultCases[i].OS, result.OS, "OS\nTest Case: %s\nExpected: %s", v, resultCases[i].OS)
			assert.Equal(t, resultCases[i].Desktop, result.Desktop, "Desktop\nTest Case: %s\nExpected: %s", v, resultCases[i].Desktop)
			assert.Equal(t, resultCases[i].Version, result.Version, "Version\nTest Case: %s\nExpected: %s", v, resultCases[i].Version)
			assert.Equal(t, resultCases[i].Mobile, result.Mobile, "Mobile\nTest Case: %s\nExpected: %s", v, resultCases[i].Mobile)
			assert.Equal(t, resultCases[i].Tablet, result.Tablet, "Tablet\nTest Case: %s\nExpected: %s", v, resultCases[i].Tablet)
			assert.Equal(t, resultCases[i].TV, result.TV, "TV\nTest Case: %s\nExpected: %s", v, resultCases[i].TV)
			assert.Equal(t, resultCases[i].Bot, result.Bot, "Bot\nTest Case: %s\nExpected: %s", v, resultCases[i].Bot)
		})
	}
}
