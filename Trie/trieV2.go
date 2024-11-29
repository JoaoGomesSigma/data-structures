package prefixtrie

type MapTrie struct {
	nodes map[rune]*MapTrie
	end   bool
}

func newMapTrie() *MapTrie {
	return &MapTrie{nodes: make(map[rune]*MapTrie)}
}

func (t *MapTrie) add(s []rune) {
	currentNode := t
	for _, char := range s {
		next, ok := currentNode.nodes[char]
		if !ok || next == nil {
			next = newMapTrie()
			currentNode.nodes[char] = next
		}

		currentNode = next
	}
	currentNode.end = true
}

func (t *MapTrie) search(word []rune) bool {
	currentNode := t
	for _, char := range word {
		next, ok := currentNode.nodes[char]
		if !ok {
			return false
		}
		currentNode = next
	}
	return currentNode.end
}

func (t *MapTrie) startsWith(prefix []rune) []string {
	result := []string{}
	currentNode := t
	for _, char := range prefix {
		next, ok := currentNode.nodes[char]
		if !ok {
			return result
		}
		currentNode = next
	}

	if currentNode.end {
		result = append(result, string(prefix))
	}

	for key, v := range currentNode.nodes {
		if key == 0 {
			continue
		}
		prefixPluKey := append(prefix, key)
		result = append(result, v.mountWords(string(prefixPluKey))...)
	}

	return result
}

func (t *MapTrie) mountWords(indent string) (result []string) {
	for key, v := range t.nodes {
		if key == 0 {
			continue
		}
		indent += string(key)
		result = append(result, v.mountWords(indent)...)
	}
	if t.end {
		result = append(result, indent)
	}
	return result
}
