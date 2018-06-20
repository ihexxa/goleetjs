package tree

import (
	"strings"
)

type Node struct {
	Leaf string
	Children map[rune]*Node
}

func NewTrie() map[rune]*Node {
	return make(map[rune]*Node, 0)
}

// complexity could be refined
func insert(word, suffix string, root map[rune]*Node) {
	runes := []rune(suffix)
	char := runes[0]
	if _, ok := root[char]; !ok {
		if len(runes) == 1 {
			root[char] = &Node{ Leaf: word }
		} else {
			root[char] = &Node{ Children: make(map[rune]*Node, 0) }
			insert(word, string(runes[1:]), root[char].Children)
		}
	} else {
		if len(runes) == 1 {
			root[char].Leaf = word 
		} else {
			if root[char].Children == nil { // bug fixed here: children is not created
				root[char].Children = make(map[rune]*Node, 0)
			}
			insert(word, string(runes[1:]), root[char].Children)
		}
	}
}

// complexity could be refined
func find(word string, root map[rune]*Node) (string, bool) {
	runes := []rune(word)
	node, ok := root[runes[0]]
	if ok {
		if node.Leaf != "" {
			return node.Leaf, true // found
		}
		if len(runes) > 1 {
			return find(string(runes[1:]), node.Children) // go to children and search for suffix
		}
		return "", false // no char is left, not found
	}

	return "", false // node not found
}

func replaceWords(dict []string, sentence string) string {
	if len(sentence) == 0 || len(dict) == 0 {
		return sentence
	}

	root := NewTrie()

	for _, word := range dict {
		insert(word, word, root)
	}

	words := strings.Split(sentence, " ")
	newWords := make([]string, 0)
	for _, word := range words {
		prefix, exist := find(word, root)
		if exist {
			newWords = append(newWords, prefix)
		} else {
			newWords = append(newWords, word)
		}
	}

	return strings.Join(newWords, " ")
}
