package useragent_test

import (
	"testing"

	ua "github.com/medama-io/go-useragent"
)

var result ua.UserAgent

func BenchmarkParserAll(b *testing.B) {
	parser := ua.NewParser()

	b.Run("All", func(b *testing.B) {
		for _, k := range testCases {
			for i := 0; i < b.N; i++ {
				result = parser.Parse(k)
			}
		}
	})
}

func BenchmarkParserSingle(b *testing.B) {
	parser := ua.NewParser()

	for i := 0; i < b.N; i++ {
		result = parser.Parse(testCases[0])
	}
}

func BenchmarkParserPutAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ua.NewParser()
	}
}

func BenchmarkParserPutSingle(b *testing.B) {
	trie := ua.NewRuneTrie()

	for i := 0; i < b.N; i++ {
		trie.Put(testCases[0])
	}
}
