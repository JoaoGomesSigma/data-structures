package prefixtrie

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterToNumber(t *testing.T) {
	res := LetterToNumber("A")
	assert.Equal(t, 0, res)

	res = LetterToNumber("B")
	assert.Equal(t, 1, res)

	res = LetterToNumber("j")
	assert.Equal(t, 9, res)
}

func TestTrie(t *testing.T) {
	trie := new()

	trie.add([]rune("hello"))
	trie.add([]rune("hi"))
	trie.add([]rune("test"))

	print(trie)

	assert.Equal(t, true, trie.search([]rune("hello")))
	assert.Equal(t, true, trie.search([]rune("hi")))
	assert.Equal(t, false, trie.search([]rune("hellw")))
	assert.Equal(t, false, trie.search([]rune("hellow")))

	assert.ElementsMatch(t, []string{"hello", "hi"}, trie.startsWith([]rune("h")))
	assert.ElementsMatch(t, []string{"test"}, trie.startsWith([]rune("test")))
}

func print(t Trie) {
	fmt.Println("--- Trie vizualization ---")
	for i := range t.rootNode.nodes {
		if t.rootNode.nodes[i].value != 0 {
			t.printNode(t.rootNode.nodes[i], "")
		}
	}
	fmt.Println("---  ---")
}

func (t *Trie) printNode(node Node, indent string) {
	fmt.Printf("%s %c \n", indent, node.value)
	if node.end {
		fmt.Printf("%s END \n", indent)
	}

	for i := range node.nodes {
		if node.nodes[i].value != 0 {
			t.printNode(node.nodes[i], indent+" - ")
		}
	}
}
