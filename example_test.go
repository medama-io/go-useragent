package useragent_test

import (
	"fmt"

	"github.com/medama-io/go-useragent"
)

func Example() {
	// Create a new parser. Initialize only once during application startup.
	ua := useragent.NewParser()

	// Example user-agent string.
	str := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36"

	// Parse the user-agent string.
	agent := ua.Parse(str)

	// Access parsed information using agent fields.
	fmt.Println(agent.Browser())        // agents.BrowserChrome
	fmt.Println(agent.BrowserVersion()) // 118.0.0.0
	fmt.Println(agent.OS())             // agents.OSWindows
	fmt.Println(agent.Device())         // agents.DeviceDesktop

	// Boolean helper functions.
	fmt.Println(agent.IsChrome())  // true
	fmt.Println(agent.IsFirefox()) // false

	fmt.Println(agent.IsWindows()) // true
	fmt.Println(agent.IsLinux())   // false

	fmt.Println(agent.IsDesktop()) // true
	fmt.Println(agent.IsTV())      // false
	fmt.Println(agent.IsBot())     // false
	// and many more...

	// Version helper functions.
	fmt.Println(agent.BrowserVersionMajor()) // 118
	fmt.Println(agent.BrowserVersionMinor()) // 0
	fmt.Println(agent.BrowserVersionPatch()) // 0.0

	// Output:
	// Chrome
	// 118.0.0.0
	// Windows
	// Desktop
	// true
	// false
	// true
	// false
	// true
	// false
	// false
	// 118
	// 0
	// 0.0
}
