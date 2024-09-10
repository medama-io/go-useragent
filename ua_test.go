package useragent_test

import (
	"fmt"
	"testing"

	ua "github.com/medama-io/go-useragent"

	"github.com/medama-io/go-useragent/internal"
	"github.com/medama-io/go-useragent/testdata"
	"github.com/stretchr/testify/assert"
)

type ResultCase struct {
	Browser string
	OS      string
	Version string

	Desktop bool
	Mobile  bool
	Tablet  bool
	TV      bool
	Bot     bool
}

var resultCases = []ResultCase{
	// Windows (7)
	{Browser: internal.Chrome, OS: internal.Windows, Desktop: true, Version: "118.0.0.0"},
	{Browser: internal.Chrome, OS: internal.Windows, Desktop: true, Version: "59.0.3071.115"},
	{Browser: internal.IE, OS: internal.Windows, Desktop: true, Version: "8.0"},
	{Browser: internal.IE, OS: internal.Windows, Desktop: true, Version: "10.0"},
	// Technically should be 11.0, but we need to manually map this after getting it.
	{Browser: internal.IE, OS: internal.Windows, Desktop: true, Version: "7.0"},
	{Browser: internal.IE, OS: internal.Windows, Desktop: true, Version: "6.0"},
	{Browser: internal.Edge, OS: internal.Windows, Desktop: true, Version: "15.15063"},
	// Mac (5) 7
	{Browser: internal.Safari, OS: internal.MacOS, Desktop: true, Version: "10.1.2"},
	{Browser: internal.Chrome, OS: internal.MacOS, Desktop: true, Version: "60.0.3112.90"},
	{Browser: internal.Firefox, OS: internal.MacOS, Desktop: true, Version: "54.0"},
	{Browser: internal.Vivaldi, OS: internal.MacOS, Desktop: true, Version: "1.92.917.39"},
	{Browser: internal.Edge, OS: internal.MacOS, Desktop: true, Version: "79.0.309.71"},
	// Linux (5) 12
	{Browser: internal.Firefox, OS: internal.Linux, Desktop: true, Version: "52.0"},
	{Browser: internal.Firefox, OS: internal.Linux, Desktop: true, Version: "119.0"},
	{Browser: internal.Firefox, OS: internal.Linux, Desktop: true, Version: "119.0"},
	{Browser: internal.Firefox, OS: internal.Linux, Desktop: true, Version: "119.0"},
	{Browser: internal.Chrome, OS: internal.Linux, Desktop: true, Version: "119.0.0.0"},
	// iPhone (5) 17
	{Browser: internal.Safari, OS: internal.IOS, Mobile: true, Version: "10.0"},
	{Browser: internal.Chrome, OS: internal.IOS, Mobile: true, Version: "60.0.3112.89"},
	{Browser: internal.Opera, OS: internal.IOS, Mobile: true, Version: "14.0.0.104835"},
	{Browser: internal.Firefox, OS: internal.IOS, Mobile: true, Version: "8.1.1"},
	{Browser: internal.Edge, OS: internal.IOS, Mobile: true, Version: "44.11.15"},
	// iPad (3) 22
	{Browser: internal.Safari, OS: internal.IOS, Tablet: true, Version: "10.0"},
	{Browser: internal.Chrome, OS: internal.IOS, Tablet: true, Version: "119.0.6045.169"},
	{Browser: internal.Firefox, OS: internal.IOS, Tablet: true, Version: "119.0"},
	// Android (4) 25
	{Browser: internal.Samsung, OS: internal.Android, Mobile: true, Version: "4.0"},
	{Browser: internal.AndroidBrowser, OS: internal.Android, Mobile: true, Version: "4.0"},
	{Browser: internal.AndroidBrowser, OS: internal.Android, Mobile: true, Version: "4.0"},
	{Browser: internal.Chrome, OS: internal.Android, Mobile: true, Version: "38.0.2125.102"},
	// Bots (6) 29
	{Bot: true},
	{Bot: true},
	{Bot: true},
	{Bot: true},
	{Bot: true, Browser: internal.Chrome, Version: "112.0.0.0"},
	{Bot: true, Browser: internal.Chrome, OS: internal.Linux, Version: "125.0.6422.76", Desktop: true},
	// Yandex Browser (1) 35
	{Browser: internal.YandexBrowser, OS: internal.Android, Mobile: true, Version: "24.1.7.27.00"},
	// Safari UIWebView (1) 36
	{Browser: internal.Safari, OS: internal.IOS, Mobile: true},
	// Falkon (1) 37
	{Browser: internal.Falkon, OS: internal.Linux, Desktop: true, Version: "24.02.2"},
	// Android Firefox (1) 38
	{Browser: internal.Firefox, OS: internal.Android, Mobile: true, Version: "123.0"},
	// Linux ARM Architecture (1) 39
	{Browser: internal.Chrome, OS: internal.Linux, Desktop: true, Version: "88.0.4324.182"},
}

func TestParse(t *testing.T) {
	parser := ua.NewParser()

	for i, v := range testdata.TestCases {
		t.Run(fmt.Sprintf("Case:%d", i), func(t *testing.T) {
			result := parser.Parse(v)
			assert.Equal(t, resultCases[i].Browser, result.GetBrowser(), "Browser\nTest Case: %s\nExpected: %s", v, resultCases[i].Browser)
			assert.Equal(t, resultCases[i].OS, result.GetOS(), "OS\nTest Case: %s\nExpected: %s", v, resultCases[i].OS)
			assert.Equal(t, resultCases[i].Desktop, result.IsDesktop(), "Desktop\nTest Case: %s\nExpected: %s", v, resultCases[i].Desktop)
			assert.Equal(t, resultCases[i].Version, result.GetVersion(), "Version\nTest Case: %s\nExpected: %s", v, resultCases[i].Version)
			assert.Equal(t, resultCases[i].Mobile, result.IsMobile(), "Mobile\nTest Case: %s\nExpected: %s", v, resultCases[i].Mobile)
			assert.Equal(t, resultCases[i].Tablet, result.IsTablet(), "Tablet\nTest Case: %s\nExpected: %s", v, resultCases[i].Tablet)
			assert.Equal(t, resultCases[i].TV, result.IsTV(), "TV\nTest Case: %s\nExpected: %s", v, resultCases[i].TV)
			assert.Equal(t, resultCases[i].Bot, result.IsBot(), "Bot\nTest Case: %s\nExpected: %s", v, resultCases[i].Bot)
		})
	}
}
