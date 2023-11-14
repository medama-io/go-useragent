package useragent_test

import (
	"testing"

	ua "github.com/medama-io/go-useragent"
	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	assert := assert.New(t)
	trie := ua.NewRuneTrie()
	str := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36"
	trie.Put(str)
	agent := trie.Get(str)
	assert.NotNil(agent)
	assert.Equal(ua.Chrome, agent.Browser)
	assert.Equal(ua.Windows, agent.OS)
	// assert.Equal(ua.Desktop, agent.Device)
}
