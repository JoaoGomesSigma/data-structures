package prefixtrie

import "testing"

var words = []string{
	"hello", "hi", "test", "world", "trie", "banana", "apple", "programming", "language", "golang",
}

func BenchmarkAdd_V1(b *testing.B) {
	trie := new()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			trie.add([]rune(word))
		}
	}
}

func BenchmarkAdd_V2(b *testing.B) {
	trie := newMapTrie()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			trie.add([]rune(word))
		}
	}
}

func BenchmarkSearch_V1(b *testing.B) {
	trie := new()

	for _, word := range words {
		trie.add([]rune(word))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			trie.search([]rune(word))
		}
	}
}

func BenchmarkSearch_V2(b *testing.B) {
	trie := newMapTrie()

	for _, word := range words {
		trie.add([]rune(word))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			trie.search([]rune(word))
		}
	}
}

func BenchmarkStartsWith_V1(b *testing.B) {
	trie := new()

	for _, word := range words {
		trie.add([]rune(word))
	}

	b.ResetTimer()

	prefixes := []string{"h", "t", "prog", "g", "b"}
	for i := 0; i < b.N; i++ {
		for _, prefix := range prefixes {
			trie.startsWith([]rune(prefix))
		}
	}
}

func BenchmarkStartsWith_V2(b *testing.B) {
	trie := newMapTrie()

	for _, word := range words {
		trie.add([]rune(word))
	}

	b.ResetTimer()

	prefixes := []string{"h", "t", "prog", "g", "b"}
	for i := 0; i < b.N; i++ {
		for _, prefix := range prefixes {
			trie.startsWith([]rune(prefix))
		}
	}
}
