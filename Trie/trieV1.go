package prefixtrie

import (
	"strings"
)

type Trie struct {
	rootNode Node
}

type Node struct {
	value rune
	nodes []Node
	end   bool
}

func new() Trie {
	return Trie{
		rootNode: Node{
			nodes: make([]Node, 26),
		},
	}
}

func (n *Node) new(value rune) {
	n.value = value
	n.nodes = make([]Node, 26)
}

func (t *Trie) add(s []rune) {
	currentNode := &t.rootNode
	var position int
	for _, char := range s {
		position = LetterToNumber(string(char))

		if currentNode.nodes[position].value == 0 {
			currentNode.nodes[position].new(char)
		}

		currentNode = &currentNode.nodes[position]
	}
	currentNode.end = true
}

func (t *Trie) search(word []rune) bool {
	currentNode := t.rootNode
	for _, char := range word {
		position := LetterToNumber(string(char))
		if currentNode.nodes[position].value == 0 {
			return false
		}
		currentNode = currentNode.nodes[position]
	}
	return currentNode.end
}

func (t *Trie) startsWith(prefix []rune) []string {
	result := []string{}
	currentNode := t.rootNode
	for _, char := range prefix {
		position := LetterToNumber(string(char))
		if currentNode.nodes[position].value == 0 {
			return result
		}
		currentNode = currentNode.nodes[position]
	}

	if currentNode.end {
		result = append(result, string(prefix))
	}

	for _, v := range currentNode.nodes {
		if v.value == 0 {
			continue
		}
		result = append(result, v.mountWords(string(prefix))...)
	}

	return result
}

func (node *Node) mountWords(indent string) (result []string) {
	indent += string(node.value)
	for _, v := range node.nodes {
		if v.value == 0 {
			continue
		}
		result = append(result, v.mountWords(indent)...)
	}
	if node.end {
		result = append(result, indent)
	}
	return result
}

func (t *Trie) remove() {

}

func LetterToNumber(letter string) int {
	letter = strings.ToUpper(letter)

	return int(letter[0] - 'A')
}
