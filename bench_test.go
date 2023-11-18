package useragent_test

import (
	"strings"
	"testing"

	ua "github.com/medama-io/go-useragent"
)

var result ua.UserAgent

func BenchmarkParserAll(b *testing.B) {
	parser := ua.NewParser()

	for _, k := range testCases {
		name := strings.ReplaceAll(k, " ", "")
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = parser.Parse(k)
			}
		})
	}
}

func BenchmarkParserSingle(b *testing.B) {
	parser := ua.NewParser()

	for i := 0; i < b.N; i++ {
		result = parser.Parse(testCases[0])
	}
}
