package internal_test

import (
	"fmt"
	"testing"

	"github.com/medama-io/go-useragent/internal"
	"github.com/medama-io/go-useragent/testdata"
	"github.com/stretchr/testify/assert"
)

var matchResults = [][]internal.Match{
	// Windows (7)
	{internal.BrowserSafari, internal.BrowserChrome, internal.OSWindows},
	{internal.BrowserSafari, internal.BrowserChrome, internal.OSWindows},
	{internal.OSWindows, internal.BrowserIE},
	{internal.OSWindows, internal.BrowserIE},
	{internal.BrowserIE, internal.OSWindows},
	{internal.OSWindows, internal.BrowserIE},
	{internal.BrowserEdge, internal.BrowserSafari, internal.BrowserChrome, internal.OSWindows},

	// Mac (5)
	{internal.BrowserSafari, internal.TokenVersion, internal.OSMacOS},
	{internal.BrowserSafari, internal.BrowserChrome, internal.OSMacOS},
	{internal.BrowserFirefox, internal.OSMacOS},
	{internal.BrowserVivaldi, internal.BrowserSafari, internal.BrowserChrome, internal.OSMacOS},
	{internal.BrowserEdge, internal.BrowserSafari, internal.BrowserChrome, internal.OSMacOS},

	// Linux (5)
	{internal.BrowserFirefox, internal.OSLinux},
	{internal.BrowserFirefox, internal.OSLinux},
	{internal.BrowserFirefox, internal.OSLinux, internal.DeviceDesktop},
	{internal.BrowserFirefox, internal.OSLinux, internal.DeviceDesktop},
	{internal.BrowserSafari, internal.BrowserChrome, internal.OSLinux},

	// iPhone (5)
	{internal.BrowserSafari, internal.DeviceMobile, internal.TokenVersion, internal.OSIOS, internal.TokenMobileDevice},
	{internal.BrowserSafari, internal.DeviceMobile, internal.BrowserChrome, internal.OSIOS, internal.TokenMobileDevice},
	{internal.BrowserSafari, internal.DeviceMobile, internal.BrowserOpera, internal.OSIOS, internal.TokenMobileDevice},
	{internal.BrowserSafari, internal.DeviceMobile, internal.BrowserFirefox, internal.OSIOS, internal.TokenMobileDevice},
	{internal.BrowserSafari, internal.DeviceMobile, internal.BrowserEdge, internal.TokenVersion, internal.OSIOS, internal.TokenMobileDevice},

	// iPad (3)
	{internal.BrowserSafari, internal.DeviceMobile, internal.TokenVersion, internal.OSIOS, internal.DeviceTablet},
	{internal.BrowserSafari, internal.DeviceMobile, internal.BrowserChrome, internal.OSIOS, internal.DeviceTablet},
	{internal.BrowserSafari, internal.DeviceMobile, internal.BrowserFirefox, internal.OSIOS, internal.DeviceTablet},

	// Android (4)
	{internal.BrowserSafari, internal.DeviceMobile, internal.BrowserChrome, internal.BrowserSamsung, internal.OSAndroid, internal.OSLinux},
	{internal.BrowserSafari, internal.DeviceMobile, internal.TokenVersion, internal.OSAndroid, internal.OSLinux},
	{internal.BrowserSafari, internal.DeviceMobile, internal.TokenVersion, internal.OSAndroid, internal.OSLinux},
	{internal.BrowserSafari, internal.DeviceMobile, internal.BrowserChrome, internal.TokenVersion, internal.TokenMobileDevice, internal.OSAndroid, internal.OSLinux},
	{internal.BrowserSafari, internal.DeviceMobile, internal.BrowserChrome, internal.OSAndroid, internal.OSLinux},

	// Bots (4)
	{internal.DeviceBot},
	{internal.DeviceBot},
	{internal.DeviceBot},
	{internal.DeviceBot},
	{internal.BrowserSafari, internal.BrowserChrome, internal.DeviceBot},
	{internal.BrowserSafari, internal.BrowserChrome, internal.DeviceBot, internal.OSLinux},

	// Yandex Browser (1)
	{internal.BrowserSafari, internal.DeviceMobile, internal.BrowserYandex, internal.BrowserChrome, internal.OSAndroid, internal.OSLinux},

	// Safari UIWebView (1)
	{internal.DeviceMobile, internal.BrowserSafari, internal.OSIOS, internal.TokenMobileDevice},

	// Falkon (1)
	{internal.BrowserSafari, internal.BrowserChrome, internal.BrowserFalkon, internal.OSLinux},

	// Android Firefox (1)
	{internal.BrowserFirefox, internal.DeviceMobile, internal.OSAndroid},

	// Linux ARM Architecture (1)
	{internal.BrowserSafari, internal.BrowserChrome, internal.OSLinux},

	// Samsung Browser
	{internal.BrowserSafari, internal.DeviceTV, internal.BrowserChrome, internal.BrowserSamsung, internal.OSLinux},
	{internal.BrowserSafari, internal.BrowserChrome, internal.BrowserSamsung, internal.OSLinux},

	// OpenBSD
	{internal.BrowserFirefox, internal.OSOpenBSD},
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
				expected := []string{}
				got := []string{}

				for _, e := range matchResults[i] {
					expected = append(expected, e.GetMatchName())
				}

				for _, g := range match {
					got = append(got, g.Match.GetMatchName())
				}

				assert.Equal(t, matchResults[i][j], m.Match, "Test Case: %s\nMatch Number: %d\nExpected: %v\nGot: %v", v, i, expected, got)
			}
		})
	}
}
