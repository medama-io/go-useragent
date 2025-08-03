package useragent_test

import (
	"fmt"
	"testing"

	medama "github.com/medama-io/go-useragent"
	"github.com/medama-io/go-useragent/data"
)

// TestMemoryUsage prints memory stats once without benchmarking
func TestMemoryUsage(t *testing.T) {
	// Testcases trie
	trie := medama.NewRuneTrie()
	for _, ua := range data.AllTestCases {
		trie.Put(ua.UserAgent)
	}
	stats := trie.GetTotalMemoryStats()

	fmt.Println("=== Test Data Memory Usage ===")
	fmt.Printf("Entries: %d\n", len(data.AllTestCases))
	fmt.Printf("Total memory: %.2f MB\n", float64(stats.TotalMemoryBytes)/1024/1024)
	fmt.Printf("Number of nodes: %d\n", stats.NodeCount)
	fmt.Printf("Average bytes per node: %.1f\n", stats.AvgBytesPerNode)
	fmt.Printf("Largest node: %d bytes\n", stats.LargestNode)
	fmt.Printf("Smallest node: %d bytes\n", stats.SmallestNode)
	fmt.Printf("Array nodes: %d\n\n", stats.ArrayNodes)

	// Production trie
	parser := medama.NewParser()
	prodStats := parser.Trie.GetTotalMemoryStats()

	fmt.Println("\n=== Production Trie Memory Usage ===")
	fmt.Printf("Total memory: %.2f MB\n", float64(prodStats.TotalMemoryBytes)/1024/1024)
	fmt.Printf("Number of nodes: %d\n", prodStats.NodeCount)
	fmt.Printf("Average bytes per node: %.1f\n", prodStats.AvgBytesPerNode)
	fmt.Printf("Largest node: %d bytes\n", prodStats.LargestNode)
	fmt.Printf("Smallest node: %d bytes\n", prodStats.SmallestNode)
	fmt.Printf("Array nodes: %d\n", prodStats.ArrayNodes)
}
