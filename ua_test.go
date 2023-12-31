package useragent_test

import (
	"fmt"
	"testing"

	ua "github.com/medama-io/go-useragent"
	"github.com/stretchr/testify/assert"
)

var testCases = []string{
	// Windows
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; InfoPath.2; GWX:RED)",
	"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)",
	"Mozilla/5.0 (Windows NT 10.0; Trident/7.0; rv:11.0) like Gecko",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322) NS8/0.9.6",
	"Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36 Edge/15.15063",

	// Mac
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/603.3.8 (KHTML, like Gecko) Version/10.1.2 Safari/603.3.8",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:54.0) Gecko/20100101 Firefox/54.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.91 Safari/537.36 Vivaldi/1.92.917.39",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36 Edg/79.0.309.71",

	// Linux
	"Mozilla/5.0 (X11; Linux x86_64; rv:52.0) Gecko/20100101 Firefox/52.0",
	"Mozilla/5.0 (X11; Linux i686; rv:109.0) Gecko/20100101 Firefox/119.0",
	"Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:109.0) Gecko/20100101 Firefox/119.0",
	"Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/119.0",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",

	// iPhone
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.0 Mobile/14F89 Safari/602.1",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) CriOS/60.0.3112.89 Mobile/14F89 Safari/602.1",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 9_3 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) OPiOS/14.0.0.104835 Mobile/13E233 Safari/9537.53",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) FxiOS/8.1.1 Mobile/14F89 Safari/603.2.4",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 13_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0 EdgiOS/44.11.15 Mobile/15E148 Safari/605.1.15",

	// iPad
	"Mozilla/5.0 (iPad; CPU OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.0 Mobile/14F89 Safari/602.1",
	"Mozilla/5.0 (iPad; CPU OS 17_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/119.0.6045.169 Mobile/15E148 Safari/604.1",
	"Mozilla/5.0 (iPad; CPU OS 14_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/119.0 Mobile/15E148 Safari/605.1.15",
}

var resultCases = []ua.UserAgent{
	// Windows
	{
		Browser: ua.Chrome,
		OS:      ua.Windows,
		Desktop: true,
		Version: "118.0.0.0",
	},
	{
		Browser: ua.Chrome,
		OS:      ua.Windows,
		Desktop: true,
		Version: "59.0.3071.115",
	},
	{
		Browser: ua.IE,
		OS:      ua.Windows,
		Desktop: true,
		Version: "8.0",
	},
	{
		Browser: ua.IE,
		OS:      ua.Windows,
		Desktop: true,
		Version: "10.0",
	},
	{
		Browser: ua.IE,
		OS:      ua.Windows,
		Desktop: true,
		Version: "7.0", // Technically should be 11.0, but we need to manually map this after getting it.
	},
	{
		Browser: ua.IE,
		OS:      ua.Windows,
		Desktop: true,
		Version: "6.0",
	},
	{
		Browser: ua.Edge,
		OS:      ua.Windows,
		Desktop: true,
		Version: "15.15063",
	},
	// Mac
	{
		Browser: ua.Safari,
		OS:      ua.MacOS,
		Desktop: true,
		Version: "10.1.2",
	},
	{
		Browser: ua.Chrome,
		OS:      ua.MacOS,
		Desktop: true,
		Version: "60.0.3112.90",
	},
	{
		Browser: ua.Firefox,
		OS:      ua.MacOS,
		Desktop: true,
		Version: "54.0",
	},
	{
		Browser: ua.Vivaldi,
		OS:      ua.MacOS,
		Desktop: true,
		Version: "1.92.917.39",
	},
	{
		Browser: ua.Edge,
		OS:      ua.MacOS,
		Desktop: true,
		Version: "79.0.309.71",
	},
	// Linux
	{
		Browser: ua.Firefox,
		OS:      ua.Linux,
		Desktop: true,
		Version: "52.0",
	},
	{
		Browser: ua.Firefox,
		OS:      ua.Linux,
		Desktop: true,
		Version: "119.0",
	},
	{
		Browser: ua.Firefox,
		OS:      ua.Linux,
		Desktop: true,
		Version: "119.0",
	},
	{
		Browser: ua.Firefox,
		OS:      ua.Linux,
		Desktop: true,
		Version: "119.0",
	},
	{
		Browser: ua.Chrome,
		OS:      ua.Linux,
		Desktop: true,
		Version: "119.0.0.0",
	},
	// iPhone
	{
		Browser: ua.Safari,
		OS:      ua.IOS,
		Mobile:  true,
		Version: "10.0",
	},
	{
		Browser: ua.Chrome,
		OS:      ua.IOS,
		Mobile:  true,
		Version: "60.0.3112.89",
	},
	{
		Browser: ua.Opera,
		OS:      ua.IOS,
		Mobile:  true,
		Version: "14.0.0.104835",
	},
	{
		Browser: ua.Firefox,
		OS:      ua.IOS,
		Mobile:  true,
		Version: "8.1.1",
	},
	{
		Browser: ua.Edge,
		OS:      ua.IOS,
		Mobile:  true,
		Version: "44.11.15",
	},
	// iPad
	{
		Browser: ua.Safari,
		OS:      ua.IOS,
		Tablet:  true,
		Version: "10.0",
	},
	{
		Browser: ua.Chrome,
		OS:      ua.IOS,
		Tablet:  true,
		Version: "119.0.6045.169",
	},
	{
		Browser: ua.Firefox,
		OS:      ua.IOS,
		Tablet:  true,
		Version: "119.0",
	},
}

func TestParse(t *testing.T) {
	parser := ua.NewParser()

	for i, v := range testCases {
		t.Run(fmt.Sprintf("Case:%d", i), func(t *testing.T) {
			result := parser.Parse(v)
			assert.Equal(t, resultCases[i].Browser, result.Browser, "Browser Test Case: %s\nExpected: %s", v, resultCases[i].Browser)
			assert.Equal(t, resultCases[i].OS, result.OS, "OS Test Case: %s\nExpected: %s", v, resultCases[i].OS)
			assert.Equal(t, resultCases[i].Desktop, result.Desktop, "Desktop Test Case: %s\nExpected: %s", v, resultCases[i].Desktop)
			assert.Equal(t, resultCases[i].Version, result.Version, "Version Test Case: %s\nExpected: %s", v, resultCases[i].Version)
			assert.Equal(t, resultCases[i].Mobile, result.Mobile, "Mobile Test Case: %s\nExpected: %s", v, resultCases[i].Mobile)
			assert.Equal(t, resultCases[i].Tablet, result.Tablet, "Tablet Test Case: %s\nExpected: %s", v, resultCases[i].Tablet)
			assert.Equal(t, resultCases[i].TV, result.TV, "TV Test Case: %s\nExpected: %s", v, resultCases[i].TV)
			assert.Equal(t, resultCases[i].Bot, result.Bot, "Bot Test Case: %s\nExpected: %s", v, resultCases[i].Bot)
		})
	}
}
