package internal_test

import (
	"fmt"
	"testing"

	"github.com/medama-io/go-useragent/internal"
	"github.com/medama-io/go-useragent/testdata"
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

	// Android (4)
	"MozillaLinuxAndroidAppleWebKitKHTMLlikeGeckoSamsungBrowserChromeMobileSafari",
	"MozillaLinuxAndroidAppleWebKitKHTMLlikeGeckoVersionMobileSafari",
	"MozillaLinuxUAndroidAppleWebKitKHTMLlikeGeckoVersionMobileSafari",
	"MozillaLinuxAndroidAppleWebKitKHTMLlikeGeckoVersionChromeMobileSafari",

	// Bots (4)
	"MozillacompatibleGooglebothttpwwwgooglecombothtml",
	"Mozillacompatiblebingbothttpwwwbingcombingbothtm",
	"MozillacompatibleYahooSlurphttphelpyahoocomhelpusysearchslurp",
	"MozillacompatibleYandexBothttpyandexcombots",
	"MozillaAppleWebKitKHTMLlikeGeckocompatiblebingbothttpwwwbingcombingbothtmChromeSafari",
	"MozillaXLinuxxAppleWebKitKHTMLlikeGeckoHeadlessChromeSafari",
	"MozillaLinuxarmAndroidAppleWebKitKHTMLlikeGeckoChromeYaBrowseralphaSAMobileSafari",

	"MozillaiPhoneCPUiPhoneOSlikeMacOSXAppleWebKitKHTMLlikeGeckoMobile",
	"MozillaXLinuxxAppleWebKitKHTMLlikeGeckoFalkonQtWebEngineChromeSafari",
}

func TestCleanVersions(t *testing.T) {
	for i, v := range testdata.TestCases {
		t.Run(fmt.Sprintf("Case:%d", i), func(t *testing.T) {
			line := internal.RemoveMobileIdentifiers(v)
			line = internal.RemoveAndroidIdentifiers(line)
			line = internal.RemoveVersions(line)
			assert.Equal(t, versionResults[i], line, "Test Case: %s\nNoID: %s\nExpected: %s", v, line, versionResults[i])
		})
	}
}
