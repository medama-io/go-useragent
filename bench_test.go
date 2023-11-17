package useragent_test

import (
	"testing"

	ua "github.com/medama-io/go-useragent"
)

var result ua.UserAgent

func BenchmarkParserAll(b *testing.B) {
	parser := ua.NewParser()

	for i := 0; i < b.N; i++ {
		for _, k := range testCases {
			result = parser.Parse(k)
		}
	}
}

func BenchmarkParserSingle(b *testing.B) {
	parser := ua.NewParser()

	for i := 0; i < b.N; i++ {
		result = parser.Parse(testCases[0])
	}
}
