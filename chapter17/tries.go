package main

import (
	"fmt"
)

type TrieNode struct {
	children map[string]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func newTrieNode() *TrieNode {
	newMap := make(map[string]*TrieNode)
	node := TrieNode{newMap}
	return &node
}

func (t *Trie) search(word string) *TrieNode {
	currentNode := t.root

	for _, runeChar := range word {
		char := string(runeChar)
		// fmt.Println(currentNode.children)

		if _, ok := currentNode.children[char]; ok {
			currentNode = currentNode.children[char]
		} else {
			return nil
		}
	}
	return currentNode
}

func (t *Trie) insert(word string) {
	currentNode := t.root

	for _, runeChar := range word {
		char := string(runeChar)

		if _, ok := currentNode.children[char]; ok {
			currentNode = currentNode.children[char]
		} else {
			newNode := newTrieNode()
			currentNode.children[char] = newNode
			currentNode = newNode
		}
	}
	currentNode.children["*"] = nil
}

func (t *Trie) collectAllWords(node *TrieNode, word string, words []string) []string {
	currentNode := node

	for key, childNode := range currentNode.children {
		if key == "*" {
			words = append(words, word)
		} else {
			words = t.collectAllWords(childNode, word+key, words)
		}
	}
	return words
}

func (t *Trie) autocomplete(prefix string) []string {
	currentNode := t.search(prefix)

	if currentNode == nil {
		return nil
	}
	return t.collectAllWords(currentNode, prefix, make([]string, 0))
}

func (t *Trie) printAllKeys(node *TrieNode) {
	currentNode := node

	for key, childNode := range currentNode.children {
		fmt.Println(key)
		if key != "*" {
			t.printAllKeys(childNode)
		}
	}
}

func (t *Trie) autocorrection(prefix string) []string {
	currentAutocomplete := t.autocomplete(prefix)

	if currentAutocomplete == nil {
		currentAutocomplete = t.autocorrection(prefix[:len(prefix)-1])
	}
	return currentAutocomplete
}

func main() {
	fmt.Println("Trie Node:")
	root := TrieNode{map[string]*TrieNode{"a": newTrieNode(), "b": newTrieNode(), "c": newTrieNode()}}
	fmt.Println(root)
	fmt.Println("")

	// fmt.Println("Trie Tree:")
	trie := Trie{&root}
	// trie.root.children["a"] = &TrieNode{map[string]*TrieNode{"c": nil}}
	// fmt.Println("")

	fmt.Println("insert method:")
	trie.insert("ace")
	trie.insert("act")
	trie.insert("bad")
	trie.insert("bake")
	trie.insert("bat")
	trie.insert("batter")
	trie.insert("cab")
	trie.insert("cat")
	trie.insert("catnip")
	trie.insert("catnap")
	fmt.Println(trie.root)
	fmt.Println("")

	fmt.Println("search method:")
	trie.search("batter")
	fmt.Println("")

	fmt.Println("collectAllWords method:")
	wordsArray := make([]string, 0)
	allWords := trie.collectAllWords(trie.root, "", wordsArray)
	fmt.Println(allWords)
	fmt.Println("")

	fmt.Println("autocomplete method:")
	fmt.Println("ca:", trie.autocomplete("ca"))
	fmt.Println("ba:", trie.autocomplete("bat"))
	fmt.Println("")

	fmt.Println("printAllKeys method:")
	trie.printAllKeys(trie.root)
	fmt.Println("")

	fmt.Println("autocorrection method:")
	fmt.Println(trie.autocorrection("catnar"))
	fmt.Println(trie.autocorrection("bat"))
	fmt.Println(trie.autocorrection("caxasfdij"))

}
