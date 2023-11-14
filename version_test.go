package useragent_test

import (
	"testing"

	ua "github.com/medama-io/go-useragent"
	"github.com/stretchr/testify/assert"
)

// Refer to ua_test.go for main original test cases
var versionResults = []string{
	// Windows
	"Mozilla (Windows NT 10.0; Win64; x64) AppleWebKit (KHTML, like Gecko) Chrome Safari",
	"Mozilla (Windows NT 6.1; WOW64) AppleWebKit (KHTML, like Gecko) Chrome Safari",
	"Mozilla (compatible; MSIE 8.0; Windows NT 6.1; WOW64; Trident SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; InfoPath.2; GWX:RED)",
	"Mozilla (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322) NS8",
	"Mozilla (Windows NT 10.0) AppleWebKit (KHTML, like Gecko) Chrome Safari Edge",
	// Mac
	"Mozilla (Macintosh; Intel Mac OS ) AppleWebKit (KHTML, like Gecko) Version Safari",
	"Mozilla (Macintosh; Intel Mac OS ) AppleWebKit (KHTML, like Gecko) Chrome Safari",
	"Mozilla (Macintosh; Intel Mac OS ) Gecko Firefox",
	"Mozilla (Macintosh; Intel Mac OS ) AppleWebKit (KHTML, like Gecko) Chrome Safari Vivaldi",
	"Mozilla (Macintosh; Intel Mac OS ) AppleWebKit (KHTML, like Gecko) Chrome Safari Edg",
}

func TestRemoveVersions(t *testing.T) {
	assert := assert.New(t)

	for i, v := range versionResults {
		assert.Equal(v, ua.RemoveVersions(testCases[i]), "Test Case: %s\nExpected: %s", testCases[i], v)
	}
}
