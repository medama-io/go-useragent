package useragent_test

import (
	"testing"

	ua "github.com/medama-io/go-useragent"
	"github.com/stretchr/testify/assert"
)

var matchTests = map[string][]ua.MatchResults{
	// Browsers
	// Chrome
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36": {
		{
			StartIndex: 98,
			EndIndex:   104,
			Match:      ua.Safari,
			Precedence: 1,
		},
		{
			StartIndex: 81,
			EndIndex:   87,
			Match:      ua.Chrome,
			Precedence: 2,
		},
		{
			StartIndex: 13,
			EndIndex:   20,
			Match:      ua.Windows,
			Precedence: 1,
		},
	},
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36": {
		{
			StartIndex: 93,
			EndIndex:   99,
			Match:      ua.Safari,
			Precedence: 1,
		},
		{
			StartIndex: 76,
			EndIndex:   82,
			Match:      ua.Chrome,
			Precedence: 2,
		},
		{
			StartIndex: 13,
			EndIndex:   20,
			Match:      ua.Windows,
			Precedence: 1,
		},
	},
	"Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36": {
		{
			StartIndex: 86,
			EndIndex:   92,
			Match:      ua.Safari,
			Precedence: 1,
		},
		{
			StartIndex: 69,
			EndIndex:   75,
			Match:      ua.Chrome,
			Precedence: 2,
		},
		{
			StartIndex: 13,
			EndIndex:   20,
			Match:      ua.Windows,
			Precedence: 1,
		},
	},
}

func TestMatchTokenIndexes(t *testing.T) {
	assert := assert.New(t)

	for test, result := range matchTests {
		match := ua.MatchTokenIndexes(test)
		assert.Equal(result, match)
	}
}
