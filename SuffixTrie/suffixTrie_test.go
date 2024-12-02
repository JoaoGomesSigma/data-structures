package suffixtrie

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuffixTrie(t *testing.T) {
	trie := new()

	trie.insertSuffix([]rune("banana"))

	assert.True(t, trie.substring([]rune("ana")))
	assert.False(t, trie.substring([]rune("anas")))

	occurrences := trie.occurrences([]rune("ana"))
	assert.Equal(t, 2, occurrences)

	occurrences = trie.occurrences([]rune("b"))
	assert.Equal(t, 1, occurrences)

	print(trie, "")
}

func print(t Trie, indent string) {
	for key, value := range t.nodes {
		fmt.Printf("%s %c", indent, key)
		if value.index != nil {
			fmt.Printf(" - %d ", *value.index)
		}
		fmt.Printf("\n")
		if len(value.nodes) > 0 {
			print(*value, indent+" - ")
		}
	}
}
