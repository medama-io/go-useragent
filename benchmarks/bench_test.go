package useragent_test

import (
	"testing"

	medama "github.com/medama-io/go-useragent"
	"github.com/medama-io/go-useragent/data"
	mileusna "github.com/mileusna/useragent"
	uap "github.com/ua-parser/uap-go/uaparser"
)

func BenchmarkMedamaParserGetAll(b *testing.B) {
	parser := medama.NewParser()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, k := range data.AllTestCases {
			_ = parser.Parse(k.UserAgent)
		}
	}
}

func BenchmarkUAPParserGetAll(b *testing.B) {
	parser, _ := uap.New("./uap_regexes.yaml")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, k := range data.AllTestCases {
			_ = parser.Parse(k.UserAgent)
		}
	}
}

func BenchmarkMileusnaParserGetAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, k := range data.AllTestCases {
			_ = mileusna.Parse(k.UserAgent)
		}
	}
}

func BenchmarkMedamaParserGetSingle(b *testing.B) {
	parser := medama.NewParser()
	testCase := data.AllTestCases[0].UserAgent
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parser.Parse(testCase)
	}
}

func BenchmarkUAPParserGetSingle(b *testing.B) {
	parser, _ := uap.New("./uap_regexes.yaml")
	testCase := data.AllTestCases[0].UserAgent
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parser.Parse(testCase)
	}
}

func BenchmarkMileusnaParserGetSingle(b *testing.B) {
	testCase := data.AllTestCases[0].UserAgent
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = mileusna.Parse(testCase)
	}
}

// Extra benchmarks for trie implementations
func BenchmarkMedamaParserPutAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = medama.NewParser()
	}
}

func BenchmarkMedamaParserPutSingle(b *testing.B) {
	trie := medama.NewRuneTrie()
	testCase := data.AllTestCases[0].UserAgent
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		trie.Put(testCase)
	}
}
