package useragent_test

import (
	"fmt"
	"testing"

	ua "github.com/medama-io/go-useragent"
	"github.com/stretchr/testify/assert"
)

// Refer to ua_test.go for main original test cases.
var versionResults = []string{
	// Windows
	"MozillaWindowsNTWinxAppleWebKitKHTMLlikeGeckoChromeSafari",
	"MozillaWindowsNTWOWAppleWebKitKHTMLlikeGeckoChromeSafari",
	"MozillacompatibleMSIEWindowsNTWOWTridentSLCCNETCLRNETCLRNETCLRMediaCenterPCNETCNETEInfoPathGWXRED",
	"MozillacompatibleMSIEWindowsNTTrident",
	"MozillaWindowsNTTridentrvlikeGecko",
	"MozillacompatibleMSIEWindowsNTSVNETCLRNS",
	"MozillaWindowsNTAppleWebKitKHTMLlikeGeckoChromeSafariEdge",

	// Mac
	"MozillaMacintoshIntelMacOSXAppleWebKitKHTMLlikeGeckoVersionSafari",
	"MozillaMacintoshIntelMacOSXAppleWebKitKHTMLlikeGeckoChromeSafari",
	"MozillaMacintoshIntelMacOSXrvGeckoFirefox",
	"MozillaMacintoshIntelMacOSXAppleWebKitKHTMLlikeGeckoChromeSafariVivaldi",
	"MozillaMacintoshIntelMacOSXAppleWebKitKHTMLlikeGeckoChromeSafariEdg",

	// Linux
	"MozillaXLinuxxrvGeckoFirefox",
	"MozillaXLinuxirvGeckoFirefox",
	"MozillaXUbuntuLinuxirvGeckoFirefox",
	"MozillaXFedoraLinuxxrvGeckoFirefox",
	"MozillaXLinuxxAppleWebKitKHTMLlikeGeckoChromeSafari",

	// iPhone
	"MozillaiPhoneCPUiPhoneOSlikeMacOSXAppleWebKitKHTMLlikeGeckoVersionMobileSafari",
	"MozillaiPhoneCPUiPhoneOSlikeMacOSXAppleWebKitKHTMLlikeGeckoCriOSMobileSafari",
	"MozillaiPhoneCPUiPhoneOSlikeMacOSXAppleWebKitKHTMLlikeGeckoOPiOSMobileSafari",
	"MozillaiPhoneCPUiPhoneOSlikeMacOSXAppleWebKitKHTMLlikeGeckoFxiOSMobileSafari",
	"MozillaiPhoneCPUiPhoneOSlikeMacOSXAppleWebKitKHTMLlikeGeckoVersionEdgiOSMobileSafari",
}

func TestCleanVersions(t *testing.T) {
	for i, v := range versionResults {
		t.Run(fmt.Sprintf("Case:%d", i), func(t *testing.T) {
			id := ua.RemoveDeviceIdentifiers(testCases[i], ua.MatchTokenIndexes(testCases[i]))
			line := ua.RemoveVersions(id)
			assert.Equal(t, v, line, "Test Case: %s\nNoID: %s\nExpected: %s", testCases[i], id, v)
		})
	}
}
