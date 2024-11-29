package prefixtrie

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapTrie(t *testing.T) {
	trie := newMapTrie()

	trie.add([]rune("hello"))
	trie.add([]rune("hi"))
	trie.add([]rune("test"))

	fmt.Println(trie)

	assert.Equal(t, true, trie.search([]rune("hello")))
	assert.Equal(t, true, trie.search([]rune("hi")))
	assert.Equal(t, false, trie.search([]rune("hellw")))
	assert.Equal(t, false, trie.search([]rune("hellow")))

	assert.ElementsMatch(t, []string{"hello", "hi"}, trie.startsWith([]rune("h")))
	assert.ElementsMatch(t, []string{"test"}, trie.startsWith([]rune("test")))
}
