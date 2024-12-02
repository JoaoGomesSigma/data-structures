package suffixtrie

type Trie struct {
	nodes map[rune]*Trie
	index *int
}

func new() Trie {
	return Trie{
		nodes: make(map[rune]*Trie),
		index: nil,
	}
}

func (t *Trie) insertSuffix(s []rune) {
	for i := 0; i < len(s); i++ {
		t.insert(s[i:], i)
	}
}

func (t *Trie) insert(s []rune, index int) {
	node := t
	for _, char := range s {
		if _, exists := node.nodes[char]; !exists {
			node.nodes[char] = &Trie{nodes: make(map[rune]*Trie)}
		}
		node = node.nodes[char]
	}
	node.index = &index
}

func (t *Trie) substring(pattern []rune) bool {
	node := t
	for _, char := range pattern {
		nextNode, exists := node.nodes[char]
		if !exists {
			return false
		}
		node = nextNode
	}
	return true
}

func (t *Trie) occurrences(pattern []rune) int {
	node := t
	for _, char := range pattern {
		nextNode, exists := node.nodes[char]
		if !exists {
			return 0
		}
		node = nextNode
	}

	return node.countLeafs()
}

func (t *Trie) countLeafs() int {
	count := 0
	if t.index != nil {
		count++
	}
	for _, childNode := range t.nodes {
		count += childNode.countLeafs()
	}
	return count
}
