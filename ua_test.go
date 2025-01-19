package useragent_test

import (
	"fmt"
	"testing"

	ua "github.com/medama-io/go-useragent"

	"github.com/medama-io/go-useragent/agents"
	"github.com/medama-io/go-useragent/testdata"
	"github.com/stretchr/testify/assert"
)

type ResultCase struct {
	Browser agents.Browser
	OS      agents.OS
	Device  agents.Device
	Version string
}

var resultCases = []ResultCase{
	// Windows (7)
	{Browser: agents.BrowserChrome, OS: agents.OSWindows, Device: agents.DeviceDesktop, Version: "118.0.0.0"},
	{Browser: agents.BrowserChrome, OS: agents.OSWindows, Device: agents.DeviceDesktop, Version: "59.0.3071.115"},
	{Browser: agents.BrowserIE, OS: agents.OSWindows, Device: agents.DeviceDesktop, Version: "8.0"},
	{Browser: agents.BrowserIE, OS: agents.OSWindows, Device: agents.DeviceDesktop, Version: "10.0"},
	// Technically should be 11.0, but we need to manually map this after getting it.
	{Browser: agents.BrowserIE, OS: agents.OSWindows, Device: agents.DeviceDesktop, Version: "7.0"},
	{Browser: agents.BrowserIE, OS: agents.OSWindows, Device: agents.DeviceDesktop, Version: "6.0"},
	{Browser: agents.BrowserEdge, OS: agents.OSWindows, Device: agents.DeviceDesktop, Version: "15.15063"},
	// Mac (5) 7
	{Browser: agents.BrowserSafari, OS: agents.OSMacOS, Device: agents.DeviceDesktop, Version: "10.1.2"},
	{Browser: agents.BrowserChrome, OS: agents.OSMacOS, Device: agents.DeviceDesktop, Version: "60.0.3112.90"},
	{Browser: agents.BrowserFirefox, OS: agents.OSMacOS, Device: agents.DeviceDesktop, Version: "54.0"},
	{Browser: agents.BrowserVivaldi, OS: agents.OSMacOS, Device: agents.DeviceDesktop, Version: "1.92.917.39"},
	{Browser: agents.BrowserEdge, OS: agents.OSMacOS, Device: agents.DeviceDesktop, Version: "79.0.309.71"},
	// Linux (5) 12
	{Browser: agents.BrowserFirefox, OS: agents.OSLinux, Device: agents.DeviceDesktop, Version: "52.0"},
	{Browser: agents.BrowserFirefox, OS: agents.OSLinux, Device: agents.DeviceDesktop, Version: "119.0"},
	{Browser: agents.BrowserFirefox, OS: agents.OSLinux, Device: agents.DeviceDesktop, Version: "119.0"},
	{Browser: agents.BrowserFirefox, OS: agents.OSLinux, Device: agents.DeviceDesktop, Version: "119.0"},
	{Browser: agents.BrowserChrome, OS: agents.OSLinux, Device: agents.DeviceDesktop, Version: "119.0.0.0"},
	// iPhone (5) 17
	{Browser: agents.BrowserSafari, OS: agents.OSIOS, Device: agents.DeviceMobile, Version: "10.0"},
	{Browser: agents.BrowserChrome, OS: agents.OSIOS, Device: agents.DeviceMobile, Version: "60.0.3112.89"},
	{Browser: agents.BrowserOpera, OS: agents.OSIOS, Device: agents.DeviceMobile, Version: "14.0.0.104835"},
	{Browser: agents.BrowserFirefox, OS: agents.OSIOS, Device: agents.DeviceMobile, Version: "8.1.1"},
	{Browser: agents.BrowserEdge, OS: agents.OSIOS, Device: agents.DeviceMobile, Version: "44.11.15"},
	// iPad (3) 22
	{Browser: agents.BrowserSafari, OS: agents.OSIOS, Device: agents.DeviceTablet, Version: "10.0"},
	{Browser: agents.BrowserChrome, OS: agents.OSIOS, Device: agents.DeviceTablet, Version: "119.0.6045.169"},
	{Browser: agents.BrowserFirefox, OS: agents.OSIOS, Device: agents.DeviceTablet, Version: "119.0"},
	// Android (4) 25
	{Browser: agents.BrowserSamsung, OS: agents.OSAndroid, Device: agents.DeviceMobile, Version: "4.0"},
	{Browser: agents.BrowserAndroid, OS: agents.OSAndroid, Device: agents.DeviceMobile, Version: "4.0"},
	{Browser: agents.BrowserAndroid, OS: agents.OSAndroid, Device: agents.DeviceMobile, Version: "4.0"},
	{Browser: agents.BrowserChrome, OS: agents.OSAndroid, Device: agents.DeviceMobile, Version: "38.0.2125.102"},
	{Browser: agents.BrowserChrome, OS: agents.OSAndroid, Device: agents.DeviceMobile, Version: "112.0.0.0"},
	// Bots (6) 29
	{Device: agents.DeviceBot},
	{Device: agents.DeviceBot},
	{Device: agents.DeviceBot},
	{Device: agents.DeviceBot},
	{Device: agents.DeviceBot, Browser: agents.BrowserChrome, Version: "112.0.0.0"},
	{Device: agents.DeviceBot, Browser: agents.BrowserChrome, OS: agents.OSLinux, Version: "125.0.6422.76"},
	// Yandex Browser (1) 35
	{Browser: agents.BrowserYandex, OS: agents.OSAndroid, Device: agents.DeviceMobile, Version: "24.1.7.27.00"},
	// Safari UIWebView (1) 36
	{Browser: agents.BrowserSafari, OS: agents.OSIOS, Device: agents.DeviceMobile},
	// Falkon (1) 37
	{Browser: agents.BrowserFalkon, OS: agents.OSLinux, Device: agents.DeviceDesktop, Version: "24.02.2"},
	// Android Firefox (1) 38
	{Browser: agents.BrowserFirefox, OS: agents.OSAndroid, Device: agents.DeviceMobile, Version: "123.0"},
	// Linux ARM Architecture (1) 39
	{Browser: agents.BrowserChrome, OS: agents.OSLinux, Device: agents.DeviceDesktop, Version: "88.0.4324.182"},
	// Samsung
	{Browser: agents.BrowserSamsung, OS: agents.OSLinux, Device: agents.DeviceTV, Version: "2.1"},
	{Browser: agents.BrowserSamsung, OS: agents.OSLinux, Device: agents.DeviceDesktop, Version: "26.0"},
	// OpenBSD (1)
	{Browser: agents.BrowserFirefox, OS: agents.OSOpenBSD, Device: agents.DeviceDesktop, Version: "57.0"},
}

func TestParse(t *testing.T) {
	parser := ua.NewParser()

	for i, v := range testdata.TestCases {
		t.Run(fmt.Sprintf("Case:%d", i), func(t *testing.T) {
			result := parser.Parse(v)
			assert.Equal(t, resultCases[i].Browser, result.Browser(), "Browser\nTest Case: %s\nExpected: %s", v, resultCases[i].Browser)
			assert.Equal(t, resultCases[i].OS, result.OS(), "OS\nTest Case: %s\nExpected: %s", v, resultCases[i].OS)
			assert.Equal(t, resultCases[i].Device, result.Device(), "Device\nTest Case: %s\nExpected: %s", v, resultCases[i].Device)
			assert.Equal(t, resultCases[i].Version, result.BrowserVersion(), "Browser Version\nTest Case: %s\nExpected: %s", v, resultCases[i].Version)
		})
	}
}
