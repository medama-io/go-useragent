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
	"Mozilla (Windows NT ; Win; x) AppleWebKit (KHTML, like Gecko) Chrome Safari",
	"Mozilla (Windows NT ; WOW) AppleWebKit (KHTML, like Gecko) Chrome Safari",
	"Mozilla (compatible; MSIE ; Windows NT ; WOW; Trident SLCC; .NET CLR ; .NET CLR ; .NET CLR ; Media Center PC ; .NETC; .NETE; InfoPath; GWX:RED)",
	"Mozilla (compatible; MSIE ; Windows NT ; Trident",
	"Mozilla (Windows NT ; Trident rv:) like Gecko",
	"Mozilla (compatible; MSIE ; Windows NT ; SV; .NET CLR ) NS",
	"Mozilla (Windows NT ) AppleWebKit (KHTML, like Gecko) Chrome Safari Edge",
	// Mac
	"Mozilla (Macintosh; Intel Mac OS ) AppleWebKit (KHTML, like Gecko) Version Safari",
	"Mozilla (Macintosh; Intel Mac OS ) AppleWebKit (KHTML, like Gecko) Chrome Safari",
	"Mozilla (Macintosh; Intel Mac OS ) Gecko Firefox",
	"Mozilla (Macintosh; Intel Mac OS ) AppleWebKit (KHTML, like Gecko) Chrome Safari Vivaldi",
	"Mozilla (Macintosh; Intel Mac OS ) AppleWebKit (KHTML, like Gecko) Chrome Safari Edg",
	// Linux
	"Mozilla (X; Linux x_; rv:) Gecko Firefox",
	"Mozilla (X; Linux i; rv:) Gecko Firefox",
	"Mozilla (X; Ubuntu; Linux i; rv:) Gecko Firefox",
	"Mozilla (X; Fedora; Linux x_; rv:) Gecko Firefox",
	"Mozilla (X; Linux x_) AppleWebKit (KHTML, like Gecko) Chrome Safari",
}

func TestRemoveVersions(t *testing.T) {
	for i, v := range versionResults {
		t.Run(fmt.Sprintf("Case:%d", i), func(t *testing.T) {
			assert.Equal(t, v, ua.RemoveVersions(testCases[i]), "Test Case: %s\nExpected: %s", testCases[i], v)
		})
	}
}
