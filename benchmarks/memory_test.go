package useragent_test

import (
	"testing"

	medama "github.com/medama-io/go-useragent"
	"github.com/medama-io/go-useragent/testdata"
)

// TestMemoryUsage prints memory stats once without benchmarking
func TestMemoryUsage(t *testing.T) {
	// Testcases trie
	trie := medama.NewRuneTrie()
	for _, ua := range testdata.TestCases {
		trie.Put(ua)
	}
	stats := trie.GetTotalMemoryStats()

	t.Logf("\n=== Test Data Memory Usage ===")
	t.Logf("Entries: %d", len(testdata.TestCases))
	t.Logf("Total memory: %.2f MB", float64(stats.TotalMemoryBytes)/1024/1024)
	t.Logf("Number of nodes: %d", stats.NodeCount)
	t.Logf("Average bytes per node: %.1f", stats.AvgBytesPerNode)
	t.Logf("Largest node: %d bytes", stats.LargestNode)
	t.Logf("Smallest node: %d bytes", stats.SmallestNode)
	t.Logf("Array nodes: %d, Map nodes: %d", stats.ArrayNodes, stats.MapNodes)

	// Production trie
	parser := medama.NewParser()
	prodStats := parser.Trie.GetTotalMemoryStats()

	t.Logf("\n=== Production Trie Memory Usage ===")
	t.Logf("Total memory: %.2f MB", float64(prodStats.TotalMemoryBytes)/1024/1024)
	t.Logf("Number of nodes: %d", prodStats.NodeCount)
	t.Logf("Average bytes per node: %.1f", prodStats.AvgBytesPerNode)
	t.Logf("Largest node: %d bytes", prodStats.LargestNode)
	t.Logf("Smallest node: %d bytes", prodStats.SmallestNode)
	t.Logf("Array nodes: %d, Map nodes: %d", prodStats.ArrayNodes, prodStats.MapNodes)
}
