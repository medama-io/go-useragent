package useragent_test

import (
	"testing"

	ua "github.com/medama-io/go-useragent"
	"github.com/medama-io/go-useragent/testdata"
)

var result ua.UserAgent

func BenchmarkParserGetAll(b *testing.B) {
	parser := ua.NewParser()

	b.Run("All", func(b *testing.B) {
		for _, k := range testdata.TestCases {
			for i := 0; i < b.N; i++ {
				result = parser.Parse(k)
			}
		}
	})
}

func BenchmarkParserGetSingle(b *testing.B) {
	parser := ua.NewParser()

	for i := 0; i < b.N; i++ {
		result = parser.Parse(testdata.TestCases[0])
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
		trie.Put(testdata.TestCases[0])
	}
}
