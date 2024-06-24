package useragent_test

import (
	"fmt"
	"testing"

	ua "github.com/medama-io/go-useragent"
	"github.com/stretchr/testify/assert"
)

// Refer to ua_test.go for main original test cases.
var versionResults = []string{
	// Windows (7)
	"MozillaWindowsNTWinxAppleWebKitKHTMLlikeGeckoChromeSafari",
	"MozillaWindowsNTWOWAppleWebKitKHTMLlikeGeckoChromeSafari",
	"MozillacompatibleMSIEWindowsNTWOWTridentSLCCNETCLRNETCLRNETCLRMediaCenterPCNETCNETEInfoPathGWXRED",
	"MozillacompatibleMSIEWindowsNTTrident",
	"MozillaWindowsNTTridentrvlikeGecko",
	"MozillacompatibleMSIEWindowsNTSVNETCLRNS",
	"MozillaWindowsNTAppleWebKitKHTMLlikeGeckoChromeSafariEdge",

	// Mac (5)
	"MozillaMacintoshIntelMacOSXAppleWebKitKHTMLlikeGeckoVersionSafari",
	"MozillaMacintoshIntelMacOSXAppleWebKitKHTMLlikeGeckoChromeSafari",
	"MozillaMacintoshIntelMacOSXrvGeckoFirefox",
	"MozillaMacintoshIntelMacOSXAppleWebKitKHTMLlikeGeckoChromeSafariVivaldi",
	"MozillaMacintoshIntelMacOSXAppleWebKitKHTMLlikeGeckoChromeSafariEdg",

	// Linux (5)
	"MozillaXLinuxxrvGeckoFirefox",
	"MozillaXLinuxirvGeckoFirefox",
	"MozillaXUbuntuLinuxirvGeckoFirefox",
	"MozillaXFedoraLinuxxrvGeckoFirefox",
	"MozillaXLinuxxAppleWebKitKHTMLlikeGeckoChromeSafari",

	// iPhone (5)
	"MozillaiPhoneCPUiPhoneOSlikeMacOSXAppleWebKitKHTMLlikeGeckoVersionMobileSafari",
	"MozillaiPhoneCPUiPhoneOSlikeMacOSXAppleWebKitKHTMLlikeGeckoCriOSMobileSafari",
	"MozillaiPhoneCPUiPhoneOSlikeMacOSXAppleWebKitKHTMLlikeGeckoOPiOSMobileSafari",
	"MozillaiPhoneCPUiPhoneOSlikeMacOSXAppleWebKitKHTMLlikeGeckoFxiOSMobileSafari",
	"MozillaiPhoneCPUiPhoneOSlikeMacOSXAppleWebKitKHTMLlikeGeckoVersionEdgiOSMobileSafari",

	// iPad (3)
	"MozillaiPadCPUOSlikeMacOSXAppleWebKitKHTMLlikeGeckoVersionMobileSafari",
	"MozillaiPadCPUOSlikeMacOSXAppleWebKitKHTMLlikeGeckoCriOSMobileSafari",
	"MozillaiPadCPUOSlikeMacOSXAppleWebKitKHTMLlikeGeckoFxiOSMobileSafari",

	// Android (6)
	"MozillaLinuxAndroidAppleWebKitKHTMLlikeGeckoSamsungBrowserChromeMobileSafari",
	"MozillaLinuxAndroidAppleWebKitKHTMLlikeGeckoVersionMobileSafari",
	"MozillaLinuxUAndroidAppleWebKitKHTMLlikeGeckoVersionMobileSafari",
	"MozillaLinuxAndroidAppleWebKitKHTMLlikeGeckoVersionChromeMobileSafari",
}

func TestCleanVersions(t *testing.T) {
	for i, v := range testCases {
		t.Run(fmt.Sprintf("Case:%d", i), func(t *testing.T) {
			line := ua.RemoveMobileIdentifiers(v)
			line = ua.RemoveAndroidIdentifiers(line)
			line = ua.RemoveVersions(line)
			assert.Equal(t, versionResults[i], line, "Test Case: %s\nNoID: %s\nExpected: %s", v, line, versionResults[i])
		})
	}
}
