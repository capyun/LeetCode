package Trie

import "unicode/utf8"

/**
@author Shitanford
@create 2021-10-04 11:30
*/

func main() {

}

type Trie struct {
	charSet map[rune]*Trie
	end bool
}

func Constructor() Trie {
	return Trie{
		make(map[rune]*Trie),
		false,
	}
}

func (this *Trie) Insert(word string)  {
	root := this
	for _, char := range word {
		//不存在，就加上
		if _, ok := root.charSet[char]; !ok {
			root.charSet[char] = &Trie{
				charSet: make(map[rune]*Trie),
				end: false,
			}
		}
		root = root.charSet[char]
	}
	root.end = true
}


func (this *Trie) Search(word string, wildcard rune) bool {
	if len(word) == 0 {
		if this.end {
			return true
		} else {
			return false
		}
	}
	char, _ := utf8.DecodeRune([]byte{word[0]})
	if char == wildcard {
		for _, v := range this.charSet {
			if v.Search(word[1: ], wildcard) {
				return true
			}
		}
		return false
	}
	if trie, ok := this.charSet[char]; ok {
		return trie.Search(word[1: ], wildcard)
	} else {
		return false
	}
}

func (this *Trie) Search(word string) bool {
	root := this
	for _, char := range word {
		if _, ok := root.charSet[char]; !ok {
			return false
		}
		root = root.charSet[char]
	}
	if root.end {
		return true
	} else {
		return false
	}
}

func (this *Trie) StartsWith(prefix string) bool {
	root := this
	for _, char := range prefix {
		if _, ok := root.charSet[char]; !ok {
			return false
		}
		root = root.charSet[char]
	}
	return true
}
