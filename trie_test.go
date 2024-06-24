package useragent_test

import (
	"testing"

	ua "github.com/medama-io/go-useragent"
)

func testPut(t *testing.T) {
	trie := ua.NewRuneTrie()
	trie.Put("MozillaLinuxAndroidAppleWebKitKHTMLlikeGeckoSamsungBrowserChromeMobileSafari")
}
