package useragent_test

import (
	"fmt"
	"testing"

	ua "github.com/medama-io/go-useragent"
	"github.com/stretchr/testify/assert"
)

// Refer to ua_test.go for main original test cases
var versionResults = []string{
	// Windows
	"MozillaWindowsNTWinxAppleWebKitKHTMLlikeGeckoChromeSafari",
	"MozillaWindowsNTWOWAppleWebKitKHTMLlikeGeckoChromeSafari",
	"MozillacompatibleMSIEWindowsNTWOWTridentSLCC.NETCLR.NETCLR.NETCLRMediaCenterPC.NETC.NETEInfoPathGWX:RED",
	"MozillacompatibleMSIEWindowsNTTrident",
	"MozillaWindowsNTTridentrv:likeGecko",
	"MozillacompatibleMSIEWindowsNTSV.NETCLRNS",
	"MozillaWindowsNTAppleWebKitKHTMLlikeGeckoChromeSafariEdge",
	// Mac
	"MozillaMacintoshIntelMacOSAppleWebKitKHTMLlikeGeckoVersionSafari",
	"MozillaMacintoshIntelMacOSAppleWebKitKHTMLlikeGeckoChromeSafari",
	"MozillaMacintoshIntelMacOSGeckoFirefox",
	"MozillaMacintoshIntelMacOSAppleWebKitKHTMLlikeGeckoChromeSafariVivaldi",
	"MozillaMacintoshIntelMacOSAppleWebKitKHTMLlikeGeckoChromeSafariEdg",
	// Linux
	"MozillaXLinuxxrv:GeckoFirefox",
	"MozillaXLinuxirv:GeckoFirefox",
	"MozillaXUbuntuLinuxirv:GeckoFirefox",
	"MozillaXFedoraLinuxxrv:GeckoFirefox",
	"MozillaXLinuxxAppleWebKitKHTMLlikeGeckoChromeSafari",
}

func TestRemoveVersions(t *testing.T) {
	for i, v := range versionResults {
		t.Run(fmt.Sprintf("Case:%d", i), func(t *testing.T) {
			assert.Equal(t, v, ua.RemoveVersions(testCases[i]), "Test Case: %s\nExpected: %s", testCases[i], v)
		})
	}
}
